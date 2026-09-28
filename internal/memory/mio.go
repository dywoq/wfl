// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

package memory

import (
	"fmt"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

// MioFunc represents a function that is activated when a guest program wrties to
// BaseAddr + offset.
type MioFunc func(uc unicorn.Unicorn, r *MioRange, offset, val uint64) error

type MioRange struct {
	BaseAddr uint64
	Funcs    map[uint64]MioFunc
}

func NewMioRange(baseAddr uint64) *MioRange {
	return &MioRange{
		BaseAddr: baseAddr,
	}
}

// Invoke retrieves a function at m.Funcs[offset] and invokes it, providing it
// with the parameters. Returns an error if m.Funcs[offset] does not exist.
func (m *MioRange) Invoke(uc unicorn.Unicorn, offset uint64, val uint64) error {
	f, ok := m.Funcs[offset]
	if !ok {
		return fmt.Errorf("unable to find a function for 0x%X offset", offset)
	}
	return f(uc, m, offset, val)
}
