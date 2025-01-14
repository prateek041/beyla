// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package gpuevent

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_debugGpuKernelLaunchT struct {
	Flags   uint8
	PidInfo struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	KernFuncOff uint64
	GridX       int32
	GridY       int32
	GridZ       int32
	BlockX      int32
	BlockY      int32
	BlockZ      int32
	Stream      uint64
	Args        [16]uint64
	UstackSz    uint64
	Ustack      [128]uint64
}

type bpf_debugGpuMallocT struct {
	Flags   uint8
	Size    uint64
	PidInfo struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
}

// loadBpf_debug returns the embedded CollectionSpec for bpf_debug.
func loadBpf_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_debug: %w", err)
	}

	return spec, err
}

// loadBpf_debugObjects loads bpf_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_debugObjects
//	*bpf_debugPrograms
//	*bpf_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugSpecs struct {
	bpf_debugProgramSpecs
	bpf_debugMapSpecs
}

// bpf_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugProgramSpecs struct {
	HandleCudaLaunch *ebpf.ProgramSpec `ebpf:"handle_cuda_launch"`
	HandleCudaMalloc *ebpf.ProgramSpec `ebpf:"handle_cuda_malloc"`
}

// bpf_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugMapSpecs struct {
	DebugEvents *ebpf.MapSpec `ebpf:"debug_events"`
	PidCache    *ebpf.MapSpec `ebpf:"pid_cache"`
	Rb          *ebpf.MapSpec `ebpf:"rb"`
	ValidPids   *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpf_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugObjects struct {
	bpf_debugPrograms
	bpf_debugMaps
}

func (o *bpf_debugObjects) Close() error {
	return _Bpf_debugClose(
		&o.bpf_debugPrograms,
		&o.bpf_debugMaps,
	)
}

// bpf_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugMaps struct {
	DebugEvents *ebpf.Map `ebpf:"debug_events"`
	PidCache    *ebpf.Map `ebpf:"pid_cache"`
	Rb          *ebpf.Map `ebpf:"rb"`
	ValidPids   *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_debugMaps) Close() error {
	return _Bpf_debugClose(
		m.DebugEvents,
		m.PidCache,
		m.Rb,
		m.ValidPids,
	)
}

// bpf_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugPrograms struct {
	HandleCudaLaunch *ebpf.Program `ebpf:"handle_cuda_launch"`
	HandleCudaMalloc *ebpf.Program `ebpf:"handle_cuda_malloc"`
}

func (p *bpf_debugPrograms) Close() error {
	return _Bpf_debugClose(
		p.HandleCudaLaunch,
		p.HandleCudaMalloc,
	)
}

func _Bpf_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_debug_arm64_bpfel.o
var _Bpf_debugBytes []byte
