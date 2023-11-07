package cpu

var OpCodes16bitArithmeticsGenerated = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 16BIT ARITHMETICS & LOGICAL
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Add (register?), Add SP?
	0x09: NewOpCode(0x09, "ADD HL,BC /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x19: NewOpCode(0x19, "ADD HL,DE /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x29: NewOpCode(0x29, "ADD HL,HL /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x39: NewOpCode(0x39, "ADD HL,SP /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0xe8: NewOpCode(0xe8, "ADD SP,i8 /*todo*/", 2, 16, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),

	// Increment (register), Increment SP
	0x03: NewOpCode(0x03, "INC BC /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x13: NewOpCode(0x13, "INC DE /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x23: NewOpCode(0x23, "INC HL /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x33: NewOpCode(0x33, "INC SP /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),

	// Decrement (register), Decrement SP
	0x0b: NewOpCode(0x0b, "DEC BC /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x1b: NewOpCode(0x1b, "DEC DE /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x2b: NewOpCode(0x2b, "DEC HL /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x3b: NewOpCode(0x3b, "DEC SP /*todo*/", 1, 8, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
}
