// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package generictracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_tpCallProtocolArgsT struct {
	PidConn    bpf_tpPidConnectionInfoT
	SmallBuf   [24]uint8
	U_buf      uint64
	BytesLen   int32
	Ssl        uint8
	Direction  uint8
	OrigDport  uint16
	PacketType uint8
	_          [7]byte
}

type bpf_tpConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_tpEgressKeyT struct {
	S_port uint16
	D_port uint16
}

type bpf_tpGrpcFramesCtxT struct {
	PrevInfo        bpf_tpHttp2GrpcRequestT
	HasPrevInfo     uint8
	_               [3]byte
	Pos             int32
	SavedBufPos     int32
	SavedStreamId   uint32
	FoundDataFrame  uint8
	Iterations      uint8
	TerminateSearch uint8
	_               [1]byte
	Stream          bpf_tpHttp2ConnStreamT
	Args            bpf_tpCallProtocolArgsT
}

type bpf_tpHttp2ConnStreamT struct {
	PidConn  bpf_tpPidConnectionInfoT
	StreamId uint32
}

type bpf_tpHttp2GrpcRequestT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_tpConnectionInfoT
	Data            [256]uint8
	RetData         [64]uint8
	Type            uint8
	_               [3]byte
	Len             int32
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Ssl     uint8
	NewConn uint8
	_       [2]byte
	Tp      struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_tpHttpConnectionMetadataT struct {
	Pid struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Type uint8
}

type bpf_tpHttpInfoT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_tpConnectionInfoT
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [192]uint8
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
	ExtraId uint64
	TaskTid uint32
	_       [4]byte
}

type bpf_tpMsgBufferT struct {
	Buf [256]uint8
	Pos uint16
}

type bpf_tpPartialConnectionInfoT struct {
	S_addr [16]uint8
	S_port uint16
	D_port uint16
	TcpSeq uint32
}

type bpf_tpPidConnectionInfoT struct {
	Conn bpf_tpConnectionInfoT
	Pid  uint32
}

type bpf_tpPidKeyT struct {
	Pid uint32
	Ns  uint32
}

type bpf_tpRecvArgsT struct {
	SockPtr  uint64
	IovecCtx [40]uint8
}

type bpf_tpSendArgsT struct {
	P_conn  bpf_tpPidConnectionInfoT
	Size    uint64
	SockPtr uint64
}

type bpf_tpSockArgsT struct {
	Addr       uint64
	AcceptTime uint64
}

type bpf_tpSslArgsT struct {
	Ssl    uint64
	Buf    uint64
	LenPtr uint64
}

type bpf_tpSslPidConnectionInfoT struct {
	P_conn    bpf_tpPidConnectionInfoT
	OrigDport uint16
	_         [2]byte
}

type bpf_tpTcpReqT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_tpConnectionInfoT
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	Rbuf            [128]uint8
	Len             uint32
	RespLen         uint32
	Ssl             uint8
	Direction       uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	_  [2]byte
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	ExtraId uint64
}

type bpf_tpTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid     uint32
	Valid   uint8
	ReqType uint8
	_       [2]byte
}

type bpf_tpTraceKeyT struct {
	P_key   bpf_tpPidKeyT
	ExtraId uint64
}

type bpf_tpTraceMapKeyT struct {
	Conn bpf_tpConnectionInfoT
	Type uint32
}

// loadBpf_tp returns the embedded CollectionSpec for bpf_tp.
func loadBpf_tp() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_tpBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_tp: %w", err)
	}

	return spec, err
}

// loadBpf_tpObjects loads bpf_tp and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_tpObjects
//	*bpf_tpPrograms
//	*bpf_tpMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_tpObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_tp()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_tpSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpSpecs struct {
	bpf_tpProgramSpecs
	bpf_tpMapSpecs
}

// bpf_tpSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpProgramSpecs struct {
	BeylaAsyncReset                        *ebpf.ProgramSpec `ebpf:"beyla_async_reset"`
	BeylaEmitAsyncInit                     *ebpf.ProgramSpec `ebpf:"beyla_emit_async_init"`
	BeylaKprobeSysExit                     *ebpf.ProgramSpec `ebpf:"beyla_kprobe_sys_exit"`
	BeylaKprobeTcpCleanupRbuf              *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_cleanup_rbuf"`
	BeylaKprobeTcpClose                    *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_close"`
	BeylaKprobeTcpConnect                  *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_connect"`
	BeylaKprobeTcpRateCheckAppLimited      *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_rate_check_app_limited"`
	BeylaKprobeTcpRcvEstablished           *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_rcv_established"`
	BeylaKprobeTcpRecvmsg                  *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_recvmsg"`
	BeylaKprobeTcpSendmsg                  *ebpf.ProgramSpec `ebpf:"beyla_kprobe_tcp_sendmsg"`
	BeylaKprobeUnixStreamRecvmsg           *ebpf.ProgramSpec `ebpf:"beyla_kprobe_unix_stream_recvmsg"`
	BeylaKprobeUnixStreamSendmsg           *ebpf.ProgramSpec `ebpf:"beyla_kprobe_unix_stream_sendmsg"`
	BeylaKretprobeSockAlloc                *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sock_alloc"`
	BeylaKretprobeSysAccept4               *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sys_accept4"`
	BeylaKretprobeSysClone                 *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sys_clone"`
	BeylaKretprobeSysConnect               *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_sys_connect"`
	BeylaKretprobeTcpRecvmsg               *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_tcp_recvmsg"`
	BeylaKretprobeTcpSendmsg               *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_tcp_sendmsg"`
	BeylaKretprobeUnixStreamRecvmsg        *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_unix_stream_recvmsg"`
	BeylaKretprobeUnixStreamSendmsg        *ebpf.ProgramSpec `ebpf:"beyla_kretprobe_unix_stream_sendmsg"`
	BeylaProtocolHttp                      *ebpf.ProgramSpec `ebpf:"beyla_protocol_http"`
	BeylaProtocolHttp2                     *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2"`
	BeylaProtocolHttp2GrpcFrames           *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_frames"`
	BeylaProtocolHttp2GrpcHandleEndFrame   *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_handle_end_frame"`
	BeylaProtocolHttp2GrpcHandleStartFrame *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_handle_start_frame"`
	BeylaProtocolTcp                       *ebpf.ProgramSpec `ebpf:"beyla_protocol_tcp"`
	BeylaSocketHttpFilter                  *ebpf.ProgramSpec `ebpf:"beyla_socket__http_filter"`
	BeylaUprobeSslDoHandshake              *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_do_handshake"`
	BeylaUprobeSslRead                     *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_read"`
	BeylaUprobeSslReadEx                   *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_read_ex"`
	BeylaUprobeSslShutdown                 *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_shutdown"`
	BeylaUprobeSslWrite                    *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_write"`
	BeylaUprobeSslWriteEx                  *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ssl_write_ex"`
	BeylaUretprobeSslDoHandshake           *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_do_handshake"`
	BeylaUretprobeSslRead                  *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_read"`
	BeylaUretprobeSslReadEx                *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_read_ex"`
	BeylaUretprobeSslWrite                 *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_write"`
	BeylaUretprobeSslWriteEx               *ebpf.ProgramSpec `ebpf:"beyla_uretprobe_ssl_write_ex"`
}

// bpf_tpMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tpMapSpecs struct {
	ActiveAcceptArgs        *ebpf.MapSpec `ebpf:"active_accept_args"`
	ActiveConnectArgs       *ebpf.MapSpec `ebpf:"active_connect_args"`
	ActiveNodejsIds         *ebpf.MapSpec `ebpf:"active_nodejs_ids"`
	ActiveRecvArgs          *ebpf.MapSpec `ebpf:"active_recv_args"`
	ActiveSendArgs          *ebpf.MapSpec `ebpf:"active_send_args"`
	ActiveSendSockArgs      *ebpf.MapSpec `ebpf:"active_send_sock_args"`
	ActiveSslConnections    *ebpf.MapSpec `ebpf:"active_ssl_connections"`
	ActiveSslHandshakes     *ebpf.MapSpec `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs       *ebpf.MapSpec `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs      *ebpf.MapSpec `ebpf:"active_ssl_write_args"`
	ActiveUnixSocks         *ebpf.MapSpec `ebpf:"active_unix_socks"`
	AsyncResetArgs          *ebpf.MapSpec `ebpf:"async_reset_args"`
	ClientConnectInfo       *ebpf.MapSpec `ebpf:"client_connect_info"`
	CloneMap                *ebpf.MapSpec `ebpf:"clone_map"`
	ConnectionMetaMem       *ebpf.MapSpec `ebpf:"connection_meta_mem"`
	Events                  *ebpf.MapSpec `ebpf:"events"`
	GrpcFramesCtxMem        *ebpf.MapSpec `ebpf:"grpc_frames_ctx_mem"`
	Http2InfoMem            *ebpf.MapSpec `ebpf:"http2_info_mem"`
	HttpInfoMem             *ebpf.MapSpec `ebpf:"http_info_mem"`
	IncomingTraceMap        *ebpf.MapSpec `ebpf:"incoming_trace_map"`
	IovecMem                *ebpf.MapSpec `ebpf:"iovec_mem"`
	JumpTable               *ebpf.MapSpec `ebpf:"jump_table"`
	MsgBuffers              *ebpf.MapSpec `ebpf:"msg_buffers"`
	NodejsParentMap         *ebpf.MapSpec `ebpf:"nodejs_parent_map"`
	OngoingHttp             *ebpf.MapSpec `ebpf:"ongoing_http"`
	OngoingHttp2Connections *ebpf.MapSpec `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc        *ebpf.MapSpec `ebpf:"ongoing_http2_grpc"`
	OngoingHttpFallback     *ebpf.MapSpec `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq           *ebpf.MapSpec `ebpf:"ongoing_tcp_req"`
	OutgoingTraceMap        *ebpf.MapSpec `ebpf:"outgoing_trace_map"`
	PidCache                *ebpf.MapSpec `ebpf:"pid_cache"`
	PidTidToConn            *ebpf.MapSpec `ebpf:"pid_tid_to_conn"`
	ProtocolArgsMem         *ebpf.MapSpec `ebpf:"protocol_args_mem"`
	ServerTraces            *ebpf.MapSpec `ebpf:"server_traces"`
	SslToConn               *ebpf.MapSpec `ebpf:"ssl_to_conn"`
	SslToPidTid             *ebpf.MapSpec `ebpf:"ssl_to_pid_tid"`
	TcpConnectionMap        *ebpf.MapSpec `ebpf:"tcp_connection_map"`
	TcpReqMem               *ebpf.MapSpec `ebpf:"tcp_req_mem"`
	TpCharBufMem            *ebpf.MapSpec `ebpf:"tp_char_buf_mem"`
	TpInfoMem               *ebpf.MapSpec `ebpf:"tp_info_mem"`
	TraceMap                *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids               *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpf_tpObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpObjects struct {
	bpf_tpPrograms
	bpf_tpMaps
}

func (o *bpf_tpObjects) Close() error {
	return _Bpf_tpClose(
		&o.bpf_tpPrograms,
		&o.bpf_tpMaps,
	)
}

// bpf_tpMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpMaps struct {
	ActiveAcceptArgs        *ebpf.Map `ebpf:"active_accept_args"`
	ActiveConnectArgs       *ebpf.Map `ebpf:"active_connect_args"`
	ActiveNodejsIds         *ebpf.Map `ebpf:"active_nodejs_ids"`
	ActiveRecvArgs          *ebpf.Map `ebpf:"active_recv_args"`
	ActiveSendArgs          *ebpf.Map `ebpf:"active_send_args"`
	ActiveSendSockArgs      *ebpf.Map `ebpf:"active_send_sock_args"`
	ActiveSslConnections    *ebpf.Map `ebpf:"active_ssl_connections"`
	ActiveSslHandshakes     *ebpf.Map `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs       *ebpf.Map `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs      *ebpf.Map `ebpf:"active_ssl_write_args"`
	ActiveUnixSocks         *ebpf.Map `ebpf:"active_unix_socks"`
	AsyncResetArgs          *ebpf.Map `ebpf:"async_reset_args"`
	ClientConnectInfo       *ebpf.Map `ebpf:"client_connect_info"`
	CloneMap                *ebpf.Map `ebpf:"clone_map"`
	ConnectionMetaMem       *ebpf.Map `ebpf:"connection_meta_mem"`
	Events                  *ebpf.Map `ebpf:"events"`
	GrpcFramesCtxMem        *ebpf.Map `ebpf:"grpc_frames_ctx_mem"`
	Http2InfoMem            *ebpf.Map `ebpf:"http2_info_mem"`
	HttpInfoMem             *ebpf.Map `ebpf:"http_info_mem"`
	IncomingTraceMap        *ebpf.Map `ebpf:"incoming_trace_map"`
	IovecMem                *ebpf.Map `ebpf:"iovec_mem"`
	JumpTable               *ebpf.Map `ebpf:"jump_table"`
	MsgBuffers              *ebpf.Map `ebpf:"msg_buffers"`
	NodejsParentMap         *ebpf.Map `ebpf:"nodejs_parent_map"`
	OngoingHttp             *ebpf.Map `ebpf:"ongoing_http"`
	OngoingHttp2Connections *ebpf.Map `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc        *ebpf.Map `ebpf:"ongoing_http2_grpc"`
	OngoingHttpFallback     *ebpf.Map `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq           *ebpf.Map `ebpf:"ongoing_tcp_req"`
	OutgoingTraceMap        *ebpf.Map `ebpf:"outgoing_trace_map"`
	PidCache                *ebpf.Map `ebpf:"pid_cache"`
	PidTidToConn            *ebpf.Map `ebpf:"pid_tid_to_conn"`
	ProtocolArgsMem         *ebpf.Map `ebpf:"protocol_args_mem"`
	ServerTraces            *ebpf.Map `ebpf:"server_traces"`
	SslToConn               *ebpf.Map `ebpf:"ssl_to_conn"`
	SslToPidTid             *ebpf.Map `ebpf:"ssl_to_pid_tid"`
	TcpConnectionMap        *ebpf.Map `ebpf:"tcp_connection_map"`
	TcpReqMem               *ebpf.Map `ebpf:"tcp_req_mem"`
	TpCharBufMem            *ebpf.Map `ebpf:"tp_char_buf_mem"`
	TpInfoMem               *ebpf.Map `ebpf:"tp_info_mem"`
	TraceMap                *ebpf.Map `ebpf:"trace_map"`
	ValidPids               *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_tpMaps) Close() error {
	return _Bpf_tpClose(
		m.ActiveAcceptArgs,
		m.ActiveConnectArgs,
		m.ActiveNodejsIds,
		m.ActiveRecvArgs,
		m.ActiveSendArgs,
		m.ActiveSendSockArgs,
		m.ActiveSslConnections,
		m.ActiveSslHandshakes,
		m.ActiveSslReadArgs,
		m.ActiveSslWriteArgs,
		m.ActiveUnixSocks,
		m.AsyncResetArgs,
		m.ClientConnectInfo,
		m.CloneMap,
		m.ConnectionMetaMem,
		m.Events,
		m.GrpcFramesCtxMem,
		m.Http2InfoMem,
		m.HttpInfoMem,
		m.IncomingTraceMap,
		m.IovecMem,
		m.JumpTable,
		m.MsgBuffers,
		m.NodejsParentMap,
		m.OngoingHttp,
		m.OngoingHttp2Connections,
		m.OngoingHttp2Grpc,
		m.OngoingHttpFallback,
		m.OngoingTcpReq,
		m.OutgoingTraceMap,
		m.PidCache,
		m.PidTidToConn,
		m.ProtocolArgsMem,
		m.ServerTraces,
		m.SslToConn,
		m.SslToPidTid,
		m.TcpConnectionMap,
		m.TcpReqMem,
		m.TpCharBufMem,
		m.TpInfoMem,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpf_tpPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tpObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tpPrograms struct {
	BeylaAsyncReset                        *ebpf.Program `ebpf:"beyla_async_reset"`
	BeylaEmitAsyncInit                     *ebpf.Program `ebpf:"beyla_emit_async_init"`
	BeylaKprobeSysExit                     *ebpf.Program `ebpf:"beyla_kprobe_sys_exit"`
	BeylaKprobeTcpCleanupRbuf              *ebpf.Program `ebpf:"beyla_kprobe_tcp_cleanup_rbuf"`
	BeylaKprobeTcpClose                    *ebpf.Program `ebpf:"beyla_kprobe_tcp_close"`
	BeylaKprobeTcpConnect                  *ebpf.Program `ebpf:"beyla_kprobe_tcp_connect"`
	BeylaKprobeTcpRateCheckAppLimited      *ebpf.Program `ebpf:"beyla_kprobe_tcp_rate_check_app_limited"`
	BeylaKprobeTcpRcvEstablished           *ebpf.Program `ebpf:"beyla_kprobe_tcp_rcv_established"`
	BeylaKprobeTcpRecvmsg                  *ebpf.Program `ebpf:"beyla_kprobe_tcp_recvmsg"`
	BeylaKprobeTcpSendmsg                  *ebpf.Program `ebpf:"beyla_kprobe_tcp_sendmsg"`
	BeylaKprobeUnixStreamRecvmsg           *ebpf.Program `ebpf:"beyla_kprobe_unix_stream_recvmsg"`
	BeylaKprobeUnixStreamSendmsg           *ebpf.Program `ebpf:"beyla_kprobe_unix_stream_sendmsg"`
	BeylaKretprobeSockAlloc                *ebpf.Program `ebpf:"beyla_kretprobe_sock_alloc"`
	BeylaKretprobeSysAccept4               *ebpf.Program `ebpf:"beyla_kretprobe_sys_accept4"`
	BeylaKretprobeSysClone                 *ebpf.Program `ebpf:"beyla_kretprobe_sys_clone"`
	BeylaKretprobeSysConnect               *ebpf.Program `ebpf:"beyla_kretprobe_sys_connect"`
	BeylaKretprobeTcpRecvmsg               *ebpf.Program `ebpf:"beyla_kretprobe_tcp_recvmsg"`
	BeylaKretprobeTcpSendmsg               *ebpf.Program `ebpf:"beyla_kretprobe_tcp_sendmsg"`
	BeylaKretprobeUnixStreamRecvmsg        *ebpf.Program `ebpf:"beyla_kretprobe_unix_stream_recvmsg"`
	BeylaKretprobeUnixStreamSendmsg        *ebpf.Program `ebpf:"beyla_kretprobe_unix_stream_sendmsg"`
	BeylaProtocolHttp                      *ebpf.Program `ebpf:"beyla_protocol_http"`
	BeylaProtocolHttp2                     *ebpf.Program `ebpf:"beyla_protocol_http2"`
	BeylaProtocolHttp2GrpcFrames           *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_frames"`
	BeylaProtocolHttp2GrpcHandleEndFrame   *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_handle_end_frame"`
	BeylaProtocolHttp2GrpcHandleStartFrame *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_handle_start_frame"`
	BeylaProtocolTcp                       *ebpf.Program `ebpf:"beyla_protocol_tcp"`
	BeylaSocketHttpFilter                  *ebpf.Program `ebpf:"beyla_socket__http_filter"`
	BeylaUprobeSslDoHandshake              *ebpf.Program `ebpf:"beyla_uprobe_ssl_do_handshake"`
	BeylaUprobeSslRead                     *ebpf.Program `ebpf:"beyla_uprobe_ssl_read"`
	BeylaUprobeSslReadEx                   *ebpf.Program `ebpf:"beyla_uprobe_ssl_read_ex"`
	BeylaUprobeSslShutdown                 *ebpf.Program `ebpf:"beyla_uprobe_ssl_shutdown"`
	BeylaUprobeSslWrite                    *ebpf.Program `ebpf:"beyla_uprobe_ssl_write"`
	BeylaUprobeSslWriteEx                  *ebpf.Program `ebpf:"beyla_uprobe_ssl_write_ex"`
	BeylaUretprobeSslDoHandshake           *ebpf.Program `ebpf:"beyla_uretprobe_ssl_do_handshake"`
	BeylaUretprobeSslRead                  *ebpf.Program `ebpf:"beyla_uretprobe_ssl_read"`
	BeylaUretprobeSslReadEx                *ebpf.Program `ebpf:"beyla_uretprobe_ssl_read_ex"`
	BeylaUretprobeSslWrite                 *ebpf.Program `ebpf:"beyla_uretprobe_ssl_write"`
	BeylaUretprobeSslWriteEx               *ebpf.Program `ebpf:"beyla_uretprobe_ssl_write_ex"`
}

func (p *bpf_tpPrograms) Close() error {
	return _Bpf_tpClose(
		p.BeylaAsyncReset,
		p.BeylaEmitAsyncInit,
		p.BeylaKprobeSysExit,
		p.BeylaKprobeTcpCleanupRbuf,
		p.BeylaKprobeTcpClose,
		p.BeylaKprobeTcpConnect,
		p.BeylaKprobeTcpRateCheckAppLimited,
		p.BeylaKprobeTcpRcvEstablished,
		p.BeylaKprobeTcpRecvmsg,
		p.BeylaKprobeTcpSendmsg,
		p.BeylaKprobeUnixStreamRecvmsg,
		p.BeylaKprobeUnixStreamSendmsg,
		p.BeylaKretprobeSockAlloc,
		p.BeylaKretprobeSysAccept4,
		p.BeylaKretprobeSysClone,
		p.BeylaKretprobeSysConnect,
		p.BeylaKretprobeTcpRecvmsg,
		p.BeylaKretprobeTcpSendmsg,
		p.BeylaKretprobeUnixStreamRecvmsg,
		p.BeylaKretprobeUnixStreamSendmsg,
		p.BeylaProtocolHttp,
		p.BeylaProtocolHttp2,
		p.BeylaProtocolHttp2GrpcFrames,
		p.BeylaProtocolHttp2GrpcHandleEndFrame,
		p.BeylaProtocolHttp2GrpcHandleStartFrame,
		p.BeylaProtocolTcp,
		p.BeylaSocketHttpFilter,
		p.BeylaUprobeSslDoHandshake,
		p.BeylaUprobeSslRead,
		p.BeylaUprobeSslReadEx,
		p.BeylaUprobeSslShutdown,
		p.BeylaUprobeSslWrite,
		p.BeylaUprobeSslWriteEx,
		p.BeylaUretprobeSslDoHandshake,
		p.BeylaUretprobeSslRead,
		p.BeylaUretprobeSslReadEx,
		p.BeylaUretprobeSslWrite,
		p.BeylaUretprobeSslWriteEx,
	)
}

func _Bpf_tpClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_tp_x86_bpfel.o
var _Bpf_tpBytes []byte
