// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

// Package broadcast represents a base interface to log messages universally.
package broadcast

// MsgType identifies the type of a message.
type MsgType int

// Message defines a set of methods to log messages without depending
// on an implementation.
type Messager interface {
	Msg(t MsgType, v any)
	Msgf(t MsgType, format string, v ...any)
}

const (
	MsgTypeInfo MsgType = iota
	MsgTypeWarn
	MsgTypeError
)
