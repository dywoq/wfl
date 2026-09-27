// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

package entry

import (
	"errors"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type Arch int

// Entry represents the WFL's entry.
type Entry struct {
	Arch         Arch
	Code         []byte
	StartingAddr uint64
	uc           unicorn.Unicorn
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

	if err := e.uc.MemMapProt(e.StartingAddr, 16*4096, unicorn.PROT_ALL); err != nil {
		return err
	}
	if err := e.uc.MemWrite(e.StartingAddr, e.Code); err != nil {
		return err
	}

	if err := e.uc.Start(e.StartingAddr, uint64(len(e.Code))); err != nil {
		return err
	}

	return nil
}
