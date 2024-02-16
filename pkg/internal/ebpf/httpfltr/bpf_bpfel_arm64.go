// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package httpfltr

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpfHttpConnectionMetadataT struct {
	Pid struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Type uint8
}

type bpfHttpInfoT struct {
	Flags           uint8
	_               [1]byte
	ConnInfo        bpfConnectionInfoT
	_               [2]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [160]uint8
	Len             uint32
	RespLen         uint32
	Status          uint16
	Type            uint8
	Ssl             uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpfPidConnectionInfoT struct {
	Conn bpfConnectionInfoT
	Pid  uint32
}

type bpfPidKeyT struct {
	Pid uint32
	Ns  uint32
}

type bpfRecvArgsT struct {
	SockPtr  uint64
	IovecPtr uint64
}

type bpfSockArgsT struct {
	Addr       uint64
	AcceptTime uint64
}

type bpfSslArgsT struct {
	Ssl    uint64
	Buf    uint64
	LenPtr uint64
}

type bpfTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid   uint32
	Valid uint8
	_     [3]byte
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	KprobeSysExit           *ebpf.ProgramSpec `ebpf:"kprobe_sys_exit"`
	KprobeTcpConnect        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.ProgramSpec `ebpf:"kprobe_tcp_rcv_established"`
	KprobeTcpRecvmsg        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_recvmsg"`
	KprobeTcpSendmsg        *ebpf.ProgramSpec `ebpf:"kprobe_tcp_sendmsg"`
	KretprobeSockAlloc      *ebpf.ProgramSpec `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysClone       *ebpf.ProgramSpec `ebpf:"kretprobe_sys_clone"`
	KretprobeSysConnect     *ebpf.ProgramSpec `ebpf:"kretprobe_sys_connect"`
	KretprobeTcpRecvmsg     *ebpf.ProgramSpec `ebpf:"kretprobe_tcp_recvmsg"`
	SocketHttpFilter        *ebpf.ProgramSpec `ebpf:"socket__http_filter"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	ActiveAcceptArgs    *ebpf.MapSpec `ebpf:"active_accept_args"`
	ActiveConnectArgs   *ebpf.MapSpec `ebpf:"active_connect_args"`
	ActiveRecvArgs      *ebpf.MapSpec `ebpf:"active_recv_args"`
	ActiveSslHandshakes *ebpf.MapSpec `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs   *ebpf.MapSpec `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs  *ebpf.MapSpec `ebpf:"active_ssl_write_args"`
	CloneMap            *ebpf.MapSpec `ebpf:"clone_map"`
	Events              *ebpf.MapSpec `ebpf:"events"`
	FilteredConnections *ebpf.MapSpec `ebpf:"filtered_connections"`
	HttpInfoMem         *ebpf.MapSpec `ebpf:"http_info_mem"`
	OngoingHttp         *ebpf.MapSpec `ebpf:"ongoing_http"`
	OngoingHttpFallback *ebpf.MapSpec `ebpf:"ongoing_http_fallback"`
	PidCache            *ebpf.MapSpec `ebpf:"pid_cache"`
	PidTidToConn        *ebpf.MapSpec `ebpf:"pid_tid_to_conn"`
	ServerTraces        *ebpf.MapSpec `ebpf:"server_traces"`
	SslToConn           *ebpf.MapSpec `ebpf:"ssl_to_conn"`
	SslToPidTid         *ebpf.MapSpec `ebpf:"ssl_to_pid_tid"`
	TpCharBufMem        *ebpf.MapSpec `ebpf:"tp_char_buf_mem"`
	TpInfoMem           *ebpf.MapSpec `ebpf:"tp_info_mem"`
	TraceMap            *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids           *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	ActiveAcceptArgs    *ebpf.Map `ebpf:"active_accept_args"`
	ActiveConnectArgs   *ebpf.Map `ebpf:"active_connect_args"`
	ActiveRecvArgs      *ebpf.Map `ebpf:"active_recv_args"`
	ActiveSslHandshakes *ebpf.Map `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs   *ebpf.Map `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs  *ebpf.Map `ebpf:"active_ssl_write_args"`
	CloneMap            *ebpf.Map `ebpf:"clone_map"`
	Events              *ebpf.Map `ebpf:"events"`
	FilteredConnections *ebpf.Map `ebpf:"filtered_connections"`
	HttpInfoMem         *ebpf.Map `ebpf:"http_info_mem"`
	OngoingHttp         *ebpf.Map `ebpf:"ongoing_http"`
	OngoingHttpFallback *ebpf.Map `ebpf:"ongoing_http_fallback"`
	PidCache            *ebpf.Map `ebpf:"pid_cache"`
	PidTidToConn        *ebpf.Map `ebpf:"pid_tid_to_conn"`
	ServerTraces        *ebpf.Map `ebpf:"server_traces"`
	SslToConn           *ebpf.Map `ebpf:"ssl_to_conn"`
	SslToPidTid         *ebpf.Map `ebpf:"ssl_to_pid_tid"`
	TpCharBufMem        *ebpf.Map `ebpf:"tp_char_buf_mem"`
	TpInfoMem           *ebpf.Map `ebpf:"tp_info_mem"`
	TraceMap            *ebpf.Map `ebpf:"trace_map"`
	ValidPids           *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.ActiveAcceptArgs,
		m.ActiveConnectArgs,
		m.ActiveRecvArgs,
		m.ActiveSslHandshakes,
		m.ActiveSslReadArgs,
		m.ActiveSslWriteArgs,
		m.CloneMap,
		m.Events,
		m.FilteredConnections,
		m.HttpInfoMem,
		m.OngoingHttp,
		m.OngoingHttpFallback,
		m.PidCache,
		m.PidTidToConn,
		m.ServerTraces,
		m.SslToConn,
		m.SslToPidTid,
		m.TpCharBufMem,
		m.TpInfoMem,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	KprobeSysExit           *ebpf.Program `ebpf:"kprobe_sys_exit"`
	KprobeTcpConnect        *ebpf.Program `ebpf:"kprobe_tcp_connect"`
	KprobeTcpRcvEstablished *ebpf.Program `ebpf:"kprobe_tcp_rcv_established"`
	KprobeTcpRecvmsg        *ebpf.Program `ebpf:"kprobe_tcp_recvmsg"`
	KprobeTcpSendmsg        *ebpf.Program `ebpf:"kprobe_tcp_sendmsg"`
	KretprobeSockAlloc      *ebpf.Program `ebpf:"kretprobe_sock_alloc"`
	KretprobeSysAccept4     *ebpf.Program `ebpf:"kretprobe_sys_accept4"`
	KretprobeSysClone       *ebpf.Program `ebpf:"kretprobe_sys_clone"`
	KretprobeSysConnect     *ebpf.Program `ebpf:"kretprobe_sys_connect"`
	KretprobeTcpRecvmsg     *ebpf.Program `ebpf:"kretprobe_tcp_recvmsg"`
	SocketHttpFilter        *ebpf.Program `ebpf:"socket__http_filter"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.KprobeSysExit,
		p.KprobeTcpConnect,
		p.KprobeTcpRcvEstablished,
		p.KprobeTcpRecvmsg,
		p.KprobeTcpSendmsg,
		p.KretprobeSockAlloc,
		p.KretprobeSysAccept4,
		p.KretprobeSysClone,
		p.KretprobeSysConnect,
		p.KretprobeTcpRecvmsg,
		p.SocketHttpFilter,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_arm64.o
var _BpfBytes []byte
