package ebpfcommon

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log/slog"
	"net"
	"os"
	"strings"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"

	"github.com/grafana/beyla/v2/pkg/config"
	"github.com/grafana/beyla/v2/pkg/internal/ebpf/ringbuf"
	"github.com/grafana/beyla/v2/pkg/internal/request"
)

//go:generate $BPF2GO -cc $BPF_CLANG -cflags $BPF_CFLAGS -target amd64,arm64 -type http_request_trace -type sql_request_trace -type http_info_t -type connection_info_t -type http2_grpc_request_t -type tcp_req_t -type kafka_client_req_t -type kafka_go_req_t  -type redis_client_req_t bpf ../../../../bpf/tracer_common.c -- -I../../../../bpf/headers

// HTTPRequestTrace contains information from an HTTP request as directly received from the
// eBPF layer. This contains low-level C structures for accurate binary read from ring buffer.
type HTTPRequestTrace bpfHttpRequestTrace
type SQLRequestTrace bpfSqlRequestTrace
type BPFHTTPInfo bpfHttpInfoT
type BPFConnInfo bpfConnectionInfoT
type TCPRequestInfo bpfTcpReqT
type GoSaramaClientInfo bpfKafkaClientReqT
type GoRedisClientInfo bpfRedisClientReqT
type GoKafkaGoClientInfo bpfKafkaGoReqT

const EventTypeSQL = 5        // EVENT_SQL_CLIENT
const EventTypeKHTTP = 6      // HTTP Events generated by kprobes
const EventTypeKHTTP2 = 7     // HTTP2/gRPC Events generated by kprobes
const EventTypeTCP = 8        // Unknown TCP protocol to be classified by user space
const EventTypeGoSarama = 9   // Kafka client for Go (Shopify/IBM Sarama)
const EventTypeGoRedis = 10   // Redis client for Go
const EventTypeGoKafkaGo = 11 // Kafka-Go client from Segment-io

var IntegrityModeOverride = false

var ActiveNamespaces = make(map[uint32]uint32)

// ProbeDesc holds the information of the instrumentation points of a given
// function/symbol
type ProbeDesc struct {
	// Required, if true, will cancel the execution of the eBPF Tracer
	// if the function has not been found in the executable
	Required bool

	// The eBPF program to attach to the symbol as a uprobe (either to the
	// symbol name or to StartOffset)
	Start *ebpf.Program

	// The eBPF program to attach to the symbol either as a uretprobe or as a
	// uprobe to ReturnOffsets
	End *ebpf.Program

	// Optional offset to the start of the symbol
	StartOffset uint64

	// Optional list of the offsets of every RET instruction in the symbol
	ReturnOffsets []uint64
}

type Filter struct {
	io.Closer
	Fd int
}

type SockOps struct {
	io.Closer
	Program       *ebpf.Program
	AttachAs      ebpf.AttachType
	SockopsCgroup link.Link
}

type SockMsg struct {
	io.Closer
	Program  *ebpf.Program
	MapFD    int
	AttachAs ebpf.AttachType
}

type MisclassifiedEvent struct {
	EventType int
	TCPInfo   *TCPRequestInfo
}

var MisclassifiedEvents = make(chan MisclassifiedEvent)

func ptlog() *slog.Logger { return slog.With("component", "ebpf.ProcessTracer") }

func ReadBPFTraceAsSpan(cfg *config.EBPFTracer, record *ringbuf.Record, filter ServiceFilter) (request.Span, bool, error) {
	var eventType uint8

	// we read the type first, depending on the type we decide what kind of record we have
	err := binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &eventType)
	if err != nil {
		return request.Span{}, true, err
	}

	switch eventType {
	case EventTypeSQL:
		return ReadSQLRequestTraceAsSpan(record)
	case EventTypeKHTTP:
		return ReadHTTPInfoIntoSpan(record, filter)
	case EventTypeKHTTP2:
		return ReadHTTP2InfoIntoSpan(record, filter)
	case EventTypeTCP:
		return ReadTCPRequestIntoSpan(cfg, record, filter)
	case EventTypeGoSarama:
		return ReadGoSaramaRequestIntoSpan(record)
	case EventTypeGoRedis:
		return ReadGoRedisRequestIntoSpan(record)
	case EventTypeGoKafkaGo:
		return ReadGoKafkaGoRequestIntoSpan(record)
	}

	var event HTTPRequestTrace

	err = binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &event)
	if err != nil {
		return request.Span{}, true, err
	}

	return HTTPRequestTraceToSpan(&event), false, nil
}

func ReadSQLRequestTraceAsSpan(record *ringbuf.Record) (request.Span, bool, error) {
	var event SQLRequestTrace
	if err := binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &event); err != nil {
		return request.Span{}, true, err
	}

	return SQLRequestTraceToSpan(&event), false, nil
}

type KernelLockdown uint8

const (
	KernelLockdownNone KernelLockdown = iota + 1
	KernelLockdownIntegrity
	KernelLockdownConfidentiality
	KernelLockdownOther
)

func SupportsContextPropagationWithProbe(log *slog.Logger) bool {
	kernelMajor, kernelMinor := KernelVersion()
	log.Debug("Linux kernel version", "major", kernelMajor, "minor", kernelMinor)

	if kernelMajor < 5 || (kernelMajor == 5 && kernelMinor < 10) {
		log.Debug("Found Linux kernel earlier than 5.10, Go trace context propagation at library level is supported", "major", kernelMajor, "minor", kernelMinor)
		return true
	}

	// bpf_probe_write_user(), used to inject the context, requires CAP_SYS_ADMIN

	if !hasCapSysAdmin() {
		log.Info("Go context propagation at library level disabled due to missing capability CAP_SYS_ADMIN")
		return false
	}

	lockdown := KernelLockdownMode()

	if lockdown == KernelLockdownNone {
		log.Debug("Kernel not in lockdown mode, Go trace context propagation at library level is supported.")
		return true
	}

	return false
}

func SupportsEBPFLoops(log *slog.Logger, overrideKernelVersion bool) bool {
	if overrideKernelVersion {
		log.Debug("Skipping kernel version check for bpf_loop functionality: user supplied confirmation of support")
		return true
	}
	kernelMajor, kernelMinor := KernelVersion()
	return kernelMajor > 5 || (kernelMajor == 5 && kernelMinor >= 17)
}

// Injectable for tests
var lockdownPath = "/sys/kernel/security/lockdown"

func KernelLockdownMode() KernelLockdown {
	plog := ptlog()
	plog.Debug("checking kernel lockdown mode, [none] allows us to propagate trace context")
	// If we can't find the file, assume no lockdown
	if _, err := os.Stat(lockdownPath); err == nil {
		f, err := os.Open(lockdownPath)

		if err != nil {
			plog.Warn("failed to open /sys/kernel/security/lockdown, assuming lockdown [integrity]", "error", err)
			return KernelLockdownIntegrity
		}

		defer f.Close()
		scanner := bufio.NewScanner(f)
		if scanner.Scan() {
			lockdown := scanner.Text()
			switch {
			case strings.Contains(lockdown, "[none]"):
				return KernelLockdownNone
			case strings.Contains(lockdown, "[integrity]"):
				return KernelLockdownIntegrity
			case strings.Contains(lockdown, "[confidentiality]"):
				return KernelLockdownConfidentiality
			default:
				return KernelLockdownOther
			}
		}

		plog.Warn("file /sys/kernel/security/lockdown is empty, assuming lockdown [integrity]")
		return KernelLockdownIntegrity
	}

	plog.Debug("can't find /sys/kernel/security/lockdown, assuming no lockdown")
	return KernelLockdownNone
}

func cstr(chars []uint8) string {
	addrLen := bytes.IndexByte(chars, 0)
	if addrLen < 0 {
		addrLen = len(chars)
	}

	return string(chars[:addrLen])
}

func (connInfo *BPFConnInfo) reqHostInfo() (source, target string) {
	src := make(net.IP, net.IPv6len)
	dst := make(net.IP, net.IPv6len)
	copy(src, connInfo.S_addr[:])
	copy(dst, connInfo.D_addr[:])

	srcStr := src.String()
	dstStr := dst.String()

	if src.IsUnspecified() {
		srcStr = ""
	}

	if dst.IsUnspecified() {
		dstStr = ""
	}

	return srcStr, dstStr
}
