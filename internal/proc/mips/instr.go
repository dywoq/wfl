// Copyright 2026 dywoq - Apache License 2.0
// https://github.com/dywoq/wlf

package mips

type InstrFormat int

type InstrRegister uint8

type InstrRFunction uint8

type InstrIOpcode uint8

type InstrJOpcode uint8

type Instr struct {
	Format InstrFormat
	RInfo  *InstrRInfo
	IInfo  *InstrIInfo
	JInfo  *InstrJInfo
}

type InstrRInfo struct {
	SrcRegister1 InstrRegister
	SrcRegister2 InstrRegister
	DestRegister InstrRegister
	Function     InstrRFunction
	ShiftAmount  uint8
}

type InstrIInfo struct {
	SrcRegister  InstrRegister
	DestRegister InstrRegister
	Value        uint16
}

type InstrJInfo struct {
	Target uint32
}

const (
	InstrFormatR InstrFormat = iota
	InstrFormatI
	InstrFormatJ
)

const (
	InstrRegisterZero InstrRegister = iota
	InstrRegisterAt
	InstrRegisterValue0
	InstrRegisterValue1
	InstrRegisterArgument0
	InstrRegisterArgument1
	InstrRegisterArgument2
	InstrRegisterArgument3
	InstrRegisterTemporary0
	InstrRegisterTemporary1
	InstrRegisterTemporary2
	InstrRegisterTemporary3
	InstrRegisterTemporary4
	InstrRegisterTemporary5
	InstrRegisterTemporary6
	InstrRegisterTemporary7
	InstrRegisterSaved0
	InstrRegisterSaved1
	InstrRegisterSaved2
	InstrRegisterSaved3
	InstrRegisterSaved4
	InstrRegisterSaved5
	InstrRegisterSaved6
	InstrRegisterSaved7
	InstrRegisterTemporary8
	InstrRegisterTemporary9
	InstrRegisterKernel0
	InstrRegisterKernel1
	InstrRegisterGlobalPointer
	InstrRegisterStackPointer
	InstrRegisterFramePointer
	InstrRegisterReturnAddress
)

const (
	InstrRFunctionSll  InstrRFunction = 0x00
	InstrRFunctionSrl  InstrRFunction = 0x02
	InstrRFunctionSra  InstrRFunction = 0x03
	InstrRFunctionSllv InstrRFunction = 0x04
	InstrRFunctionSrlv InstrRFunction = 0x06
	InstrRFunctionSrav InstrRFunction = 0x07

	InstrRFunctionJr      InstrRFunction = 0x08
	InstrRFunctionJalr    InstrRFunction = 0x09
	InstrRFunctionSyscall InstrRFunction = 0x0C
	InstrRFunctionBreak   InstrRFunction = 0x0D
	InstrRFunctionSync    InstrRFunction = 0x0F

	InstrRFunctionMfhi InstrRFunction = 0x10
	InstrRFunctionMthi InstrRFunction = 0x11
	InstrRFunctionMflo InstrRFunction = 0x12
	InstrRFunctionMtlo InstrRFunction = 0x13

	InstrRFunctionDsllv InstrRFunction = 0x14
	InstrRFunctionDsrlv InstrRFunction = 0x16
	InstrRFunctionDsrav InstrRFunction = 0x17

	InstrRFunctionMult  InstrRFunction = 0x18
	InstrRFunctionMultu InstrRFunction = 0x19
	InstrRFunctionDiv   InstrRFunction = 0x1A
	InstrRFunctionDivu  InstrRFunction = 0x1B

	InstrRFunctionDmult  InstrRFunction = 0x1C
	InstrRFunctionDmultu InstrRFunction = 0x1D
	InstrRFunctionDdiv   InstrRFunction = 0x1E
	InstrRFunctionDdivu  InstrRFunction = 0x1F

	InstrRFunctionAdd  InstrRFunction = 0x20
	InstrRFunctionAddU InstrRFunction = 0x21
	InstrRFunctionSub  InstrRFunction = 0x22
	InstrRFunctionSubU InstrRFunction = 0x23
	InstrRFunctionAnd  InstrRFunction = 0x24
	InstrRFunctionOr   InstrRFunction = 0x25
	InstrRFunctionXor  InstrRFunction = 0x26
	InstrRFunctionNor  InstrRFunction = 0x27
	InstrRFunctionSlt  InstrRFunction = 0x2A
	InstrRFunctionSLtU InstrRFunction = 0x2B

	InstrRFunctionDAdd  InstrRFunction = 0x2C
	InstrRFunctionDAddU InstrRFunction = 0x2D
	InstrRFunctionDSub  InstrRFunction = 0x2E
	InstrRFunctionDSubU InstrRFunction = 0x2F

	InstrRFunctionTge  InstrRFunction = 0x30
	InstrRFunctionTgeu InstrRFunction = 0x31
	InstrRFunctionTlt  InstrRFunction = 0x32
	InstrRFunctionTltu InstrRFunction = 0x33
	InstrRFunctionTeq  InstrRFunction = 0x34
	InstrRFunctionTne  InstrRFunction = 0x36

	InstrRFunctionDsll   InstrRFunction = 0x38
	InstrRFunctionDsrl   InstrRFunction = 0x3A
	InstrRFunctionDsra   InstrRFunction = 0x3B
	InstrRFunctionDsll32 InstrRFunction = 0x3C
	InstrRFunctionDsrl32 InstrRFunction = 0x3E
	InstrRFunctionDsra32 InstrRFunction = 0x3F
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
	InstrJOpcodeJ   InstrIOpcode = 0x02
	InstrJOpcodeJal InstrJOpcode = 0x03
)
