// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

package memory

import "github.com/unicorn-engine/unicorn/bindings/go/unicorn"

// RegionFlag identifies the guest program's memory access.
type RegionFlag int

// RegionType identifies the type of a memory region.
type RegionType int

type Region struct {
	PhysAddr uint64
	Data     []byte
	Size     uint64
	Flags    RegionFlag
	Type     RegionType
}

// NewRegion pre-allocates the region's data slice with size aligned to 4KiB.
func NewRegion(physAddr uint64, size uint64, ttype RegionType, flags RegionFlag) *Region {
	return &Region{
		PhysAddr: physAddr,
		Data:     make([]byte, 0, (size+4096-1)&^(4096-1)),
		Size:     size,
		Flags:    flags,
		Type:     ttype,
	}
}

const (
	RegionFlagRead  RegionFlag = unicorn.PROT_READ
	RegionFlagWrite RegionFlag = unicorn.PROT_WRITE
	RegionFlagExec  RegionFlag = unicorn.PROT_EXEC
)

const (
	RegionTypeConventional RegionType = iota
)
