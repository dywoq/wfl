// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

package entry

import (
	"errors"
	"fmt"

	"github.com/dywoq/wfl/internal/memory"
	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type Arch int

// Entry represents the WFL's entry.
type Entry struct {
	Arch          Arch
	Code          []byte
	memoryRegions []*memory.Region
	uc            unicorn.Unicorn
}

const (
	ArchX86_32 Arch = iota
	ArchX86_64
	ArchMips
	ArchPowerPc
)

var (
	ErrUnknownArch = errors.New("unknown arch")
)

func New(arch Arch, code []byte) *Entry {
	return &Entry{
		Arch:          arch,
		Code:          code,
		memoryRegions: nil,
		uc:            nil,
	}
}

// Start relies on e.Arch to find a matching processor's architecture and mode, required by Unicorn Engine.
// The function maps e.Code at e.StartingAddr with all protection flags (Read/Write/Execute).
// After these steps, Start runs e.Code.
//
// Returns [ErrUnknownArch] if e.Arch is unknown.
//
// Returns any error from the Unicorn Engine's API.
func (e *Entry) Start() error {
	mode := 0
	arch := 0
	switch e.Arch {
	case ArchX86_32:
		arch = unicorn.ARCH_X86
		mode = unicorn.MODE_32
	case ArchX86_64:
		arch = unicorn.ARCH_X86
		mode = unicorn.MODE_64
	case ArchMips:
		arch = unicorn.ARCH_MIPS
		mode = unicorn.CPU_MIPS64_R4000
	case ArchPowerPc:
		arch = unicorn.ARCH_PPC
		mode = unicorn.CPU_PPC32_604
	default:
		return ErrUnknownArch
	}
	got, err := unicorn.NewUnicorn(arch, mode)
	if err != nil {
		return err
	}
	e.uc = got
	defer e.uc.Close()

	e.initMemoryRegions()

	// Retrieve the conventional memory's physical address and size.
	// Check if length of e.Code overflows conventionalSize
	ok := false
	conventionalPhysAddr := uint64(0)
	conventionalSize := uint64(0)
	for _, r := range e.memoryRegions {
		if r.Type == memory.RegionTypeConventional {
			conventionalPhysAddr = r.PhysAddr
			conventionalSize = r.Size
			ok = true
			break
		}
	}
	if !ok {
		return errors.New("unable to find the conventional region")
	}
	if len(e.Code) >= int(conventionalSize) {
		return fmt.Errorf("code's length overflows the conventional region's size (%d)", conventionalSize)
	}

	if err := e.uc.MemMapProt(conventionalPhysAddr, conventionalSize, unicorn.PROT_ALL); err != nil {
		return fmt.Errorf("failed to map conventional region: %v", err)
	}
	if err := e.uc.MemWrite(conventionalPhysAddr, e.Code); err != nil {
		return err
	}

	// Map other memory regions skipping the conventional one
	for _, r := range e.memoryRegions {
		if r.Type == memory.RegionTypeConventional {
			continue
		}
		if err := e.uc.MemMapProt(r.PhysAddr, r.Size, int(r.Flags)); err != nil {
			return fmt.Errorf("failed to map a memory region (phys addr: %d, size: %d): %v", r.PhysAddr, r.Size, err)
		}
	}

	if err := e.uc.Start(conventionalPhysAddr, conventionalPhysAddr+conventionalSize); err != nil {
		return fmt.Errorf("executing conventional code failed: %v", err)
	}

	return nil
}

func (e *Entry) initMemoryRegions() {
	e.memoryRegions = []*memory.Region{
		memory.NewRegion(0x0, 0x10000, memory.RegionTypeConventional, memory.RegionFlagRead|memory.RegionFlagWrite|memory.RegionFlagExec),
	}
}
