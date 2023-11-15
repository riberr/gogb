package cpu

var OpCodesMisc = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Miscellaneous
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Halt system clock
	0x76: NewOpCode(0x76, "HALT", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /* handled in cpu loop*/ }}),

	// Stop system and main clocks
	0x10: NewOpCode(0x10, "STOP /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { panic("STOP") }}),

	// Disable interrupts
	0xf3: NewOpCode(0xf3, "DI", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.interrupts.DisableIME() }}),

	// Enable interrupts
	0xfb: NewOpCode(0xfb, "EI", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.interrupts.EnableIME() }}),

	// No operation
	0x00: NewOpCode(0x00, "NOP", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
}
