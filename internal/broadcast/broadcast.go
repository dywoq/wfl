// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

// Package broadcast represents a base interface to log messages universally.
package broadcast

// MsgType identifies the type of a message.
type MsgType int

// Messenger defines a set of methods to log messages without depending
// on an implementation.
type Messenger interface {
	Msg(t MsgType, v any)
	Msgf(t MsgType, format string, v ...any)
}

const (
	MsgTypeInfo MsgType = iota
	MsgTypeWarn
	MsgTypeError
)

// MsgOpt invokes [Messenger.Msg] if m is not nil.
func MsgOpt(m Messenger, t MsgType, v any) {
	if m != nil {
		m.Msg(t, v)
	}
}

// MsgfOpt invokes [Messenger.Msgf] if m is not nil.
func MsgfOpt(m Messenger, t MsgType, format string, v ...any) {
	if m != nil {
		m.Msgf(t, format, v...)
	}
}
