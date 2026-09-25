// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

package mips

import "fmt"

type InstrFormat int

type InstrReg uint8

type InstrRFunc uint8

type InstrIOpcode uint8

type InstrJOpcode uint8

type Instr struct {
	Format InstrFormat
	RInfo  *InstrRInfo
	IInfo  *InstrIInfo
	JInfo  *InstrJInfo
}

type InstrRInfo struct {
	SrcReg1     InstrReg
	SrcReg2     InstrReg
	DestReg     InstrReg
	Func        InstrRFunc
	ShiftAmount uint8
}

type InstrIInfo struct {
	SrcReg    InstrReg
	TargetReg InstrReg
	Val       uint16
	Opcode    InstrIOpcode
}

type InstrJInfo struct {
	Target uint32
	Opcode InstrJOpcode
}

const (
	InstrFormatR InstrFormat = iota
	InstrFormatI
	InstrFormatJ
)

const (
	InstrRegZero InstrReg = iota
	InstrRegAt
	InstrRegVal0
	InstrRegVal1
	InstrRegArgument0
	InstrRegArgument1
	InstrRegArgument2
	InstrRegArgument3
	InstrRegTemporary0
	InstrRegTemporary1
	InstrRegTemporary2
	InstrRegTemporary3
	InstrRegTemporary4
	InstrRegTemporary5
	InstrRegTemporary6
	InstrRegTemporary7
	InstrRegSaved0
	InstrRegSaved1
	InstrRegSaved2
	InstrRegSaved3
	InstrRegSaved4
	InstrRegSaved5
	InstrRegSaved6
	InstrRegSaved7
	InstrRegTemporary8
	InstrRegTemporary9
	InstrRegKernel0
	InstrRegKernel1
	InstrRegGlobalPointer
	InstrRegStackPointer
	InstrRegFramePointer
	InstrRegReturnAddress
)

const (
	InstrRFuncSll  InstrRFunc = 0x00
	InstrRFuncSrl  InstrRFunc = 0x02
	InstrRFuncSra  InstrRFunc = 0x03
	InstrRFuncSllv InstrRFunc = 0x04
	InstrRFuncSrlv InstrRFunc = 0x06
	InstrRFuncSrav InstrRFunc = 0x07

	InstrRFuncJr      InstrRFunc = 0x08
	InstrRFuncJalr    InstrRFunc = 0x09
	InstrRFuncSyscall InstrRFunc = 0x0C
	InstrRFuncBreak   InstrRFunc = 0x0D
	InstrRFuncSync    InstrRFunc = 0x0F

	InstrRFuncMfhi InstrRFunc = 0x10
	InstrRFuncMthi InstrRFunc = 0x11
	InstrRFuncMflo InstrRFunc = 0x12
	InstrRFuncMtlo InstrRFunc = 0x13

	InstrRFuncDsllv InstrRFunc = 0x14
	InstrRFuncDsrlv InstrRFunc = 0x16
	InstrRFuncDsrav InstrRFunc = 0x17

	InstrRFuncMult  InstrRFunc = 0x18
	InstrRFuncMultu InstrRFunc = 0x19
	InstrRFuncDiv   InstrRFunc = 0x1A
	InstrRFuncDivu  InstrRFunc = 0x1B

	InstrRFuncDmult  InstrRFunc = 0x1C
	InstrRFuncDmultu InstrRFunc = 0x1D
	InstrRFuncDdiv   InstrRFunc = 0x1E
	InstrRFuncDdivu  InstrRFunc = 0x1F

	InstrRFuncAdd  InstrRFunc = 0x20
	InstrRFuncAddU InstrRFunc = 0x21
	InstrRFuncSub  InstrRFunc = 0x22
	InstrRFuncSubU InstrRFunc = 0x23
	InstrRFuncAnd  InstrRFunc = 0x24
	InstrRFuncOr   InstrRFunc = 0x25
	InstrRFuncXor  InstrRFunc = 0x26
	InstrRFuncNor  InstrRFunc = 0x27
	InstrRFuncSlt  InstrRFunc = 0x2A
	InstrRFuncSLtU InstrRFunc = 0x2B

	InstrRFuncDAdd  InstrRFunc = 0x2C
	InstrRFuncDAddU InstrRFunc = 0x2D
	InstrRFuncDSub  InstrRFunc = 0x2E
	InstrRFuncDSubU InstrRFunc = 0x2F

	InstrRFuncTge  InstrRFunc = 0x30
	InstrRFuncTgeu InstrRFunc = 0x31
	InstrRFuncTlt  InstrRFunc = 0x32
	InstrRFuncTltu InstrRFunc = 0x33
	InstrRFuncTeq  InstrRFunc = 0x34
	InstrRFuncTne  InstrRFunc = 0x36

	InstrRFuncDsll   InstrRFunc = 0x38
	InstrRFuncDsrl   InstrRFunc = 0x3A
	InstrRFuncDsra   InstrRFunc = 0x3B
	InstrRFuncDsll32 InstrRFunc = 0x3C
	InstrRFuncDsrl32 InstrRFunc = 0x3E
	InstrRFuncDsra32 InstrRFunc = 0x3F
)

const (
	InstrIOpcodeBeq  InstrIOpcode = 0x04
	InstrIOpcodeBne  InstrIOpcode = 0x05
	InstrIOpcodeBlez InstrIOpcode = 0x06
	InstrIOpcodeBgtz InstrIOpcode = 0x07

	InstrIOpcodeAddi  InstrIOpcode = 0x08
	InstrIOpcodeAddiu InstrIOpcode = 0x09
	InstrIOpcodeSlti  InstrIOpcode = 0x0A
	InstrIOpcodeSltiu InstrIOpcode = 0x0B
	InstrIOpcodeAndi  InstrIOpcode = 0x0C
	InstrIOpcodeOri   InstrIOpcode = 0x0D
	InstrIOpcodeXori  InstrIOpcode = 0x0E
	InstrIOpcodeLui   InstrIOpcode = 0x0F

	InstrIOpcodeBeql  InstrIOpcode = 0x14
	InstrIOpcodeBnel  InstrIOpcode = 0x15
	InstrIOpcodeBlezl InstrIOpcode = 0x16
	InstrIOpcodeBgtzl InstrIOpcode = 0x17

	InstrIOpcodeDaddi  InstrIOpcode = 0x18
	InstrIOpcodeDaddiu InstrIOpcode = 0x19

	InstrIOpcodeLdl InstrIOpcode = 0x1A
	InstrIOpcodeLdr InstrIOpcode = 0x1B

	InstrIOpcodeLb  InstrIOpcode = 0x20
	InstrIOpcodeLh  InstrIOpcode = 0x21
	InstrIOpcodeLwl InstrIOpcode = 0x22
	InstrIOpcodeLw  InstrIOpcode = 0x23
	InstrIOpcodeLbu InstrIOpcode = 0x24
	InstrIOpcodeLhu InstrIOpcode = 0x25
	InstrIOpcodeLwr InstrIOpcode = 0x26
	InstrIOpcodeLwu InstrIOpcode = 0x27

	InstrIOpcodeSb    InstrIOpcode = 0x28
	InstrIOpcodeSh    InstrIOpcode = 0x29
	InstrIOpcodeSwl   InstrIOpcode = 0x2A
	InstrIOpcodeSw    InstrIOpcode = 0x2B
	InstrIOpcodeSdl   InstrIOpcode = 0x2C
	InstrIOpcodeSdr   InstrIOpcode = 0x2D
	InstrIOpcodeSwr   InstrIOpcode = 0x2E
	InstrIOpcodeCache InstrIOpcode = 0x2F

	InstrIOpcodeLl  InstrIOpcode = 0x30
	InstrIOpcodeLld InstrIOpcode = 0x34
	InstrIOpcodeSc  InstrIOpcode = 0x38
	InstrIOpcodeScd InstrIOpcode = 0x39

	InstrIOpcodeLwc1 InstrIOpcode = 0x31
	InstrIOpcodeLwc2 InstrIOpcode = 0x32
	InstrIOpcodeLdc1 InstrIOpcode = 0x35
	InstrIOpcodeLdc2 InstrIOpcode = 0x36
	InstrIOpcodeSwc1 InstrIOpcode = 0x39
	InstrIOpcodeSwc2 InstrIOpcode = 0x3A
	InstrIOpcodeSdc1 InstrIOpcode = 0x3D
	InstrIOpcodeSdc2 InstrIOpcode = 0x3E

	InstrIOpcodeLd InstrIOpcode = 0x37
	InstrIOpcodeSd InstrIOpcode = 0x3F
)

const (
	InstrJOpcodeJ   InstrJOpcode = 0x02
	InstrJOpcodeJal InstrJOpcode = 0x03
)

func DecodeInstr(binary uint32) (*Instr, error) {
	// Retrieve the operation code to identify the instruction's format.
	// By default it is I.
	format := InstrFormatI
	opcode := (binary >> 26) & 0x3F
	if opcode == 0 {
		format = InstrFormatR
	}
	if opcode == uint32(InstrJOpcodeJ) || opcode == uint32(InstrJOpcodeJal) {
		format = InstrFormatJ
	}

	switch format {
	case InstrFormatR:
		functionCode := binary & 0x3F
		shiftAmount := (binary >> 6) & 0x1F
		destReg := (binary >> 11) & 0x1F
		srcReg2 := (binary >> 16) & 0x1F
		srcReg1 := (binary >> 21) & 0x1F
		return &Instr{
			Format: format,
			RInfo: &InstrRInfo{
				SrcReg1:     InstrReg(srcReg1),
				SrcReg2:     InstrReg(srcReg2),
				DestReg:     InstrReg(destReg),
				Func:        InstrRFunc(functionCode),
				ShiftAmount: uint8(shiftAmount),
			},
		}, nil
	case InstrFormatI:
		Val := binary & 0xFFFF
		destReg := (binary >> 16) & 0x1F
		srcReg := (binary >> 21) & 0x1F
		return &Instr{
			Format: format,
			IInfo: &InstrIInfo{
				SrcReg:    InstrReg(srcReg),
				TargetReg: InstrReg(destReg),
				Val:       uint16(Val),
				Opcode:    InstrIOpcode(opcode),
			},
		}, nil
	case InstrFormatJ:
		target := binary & 0x3FFFFFF
		return &Instr{
			Format: format,
			JInfo: &InstrJInfo{
				Target: target,
				Opcode: InstrJOpcode(opcode),
			},
		}, nil
	}

	return nil, fmt.Errorf("could not resolve an instruction's format: 0x%X", binary)
}
