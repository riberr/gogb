package cpu

var OpCodesMisc = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Miscellaneous
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Halt system clock
	0x76: NewOpCode(0x76, "HALT /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),

	// Stop system and main clocks
	0x10: NewOpCode(0x10, "STOP /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),

	// Disable interrupts
	0xf3: NewOpCode(0xf3, "DI /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),

	// Enable interrupts
	0xfb: NewOpCode(0xfb, "EI /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),

	// No operation
	0x00: NewOpCode(0x00, "NOP", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
}
