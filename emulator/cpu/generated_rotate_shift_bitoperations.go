package cpu

var OpCodesRotateShiftBitoperations = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Rotates, Shifts, Bit operations
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// RLCA
	0x07: NewOpCode(0x07, "RLCA /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x17: NewOpCode(0x17, "RLA /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x0f: NewOpCode(0x0f, "RRCA /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x1f: NewOpCode(0x1f, "RRA /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),

	// Increment (register), Increment SP
	0xcb: NewOpCode(0xcb, "PREFIX CB /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
}
