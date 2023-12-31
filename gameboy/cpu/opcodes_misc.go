package cpu

var OpCodesMisc = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Miscellaneous
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Halt system clock
	0x76: NewOpCode(0x76, "HALT", 1, 4, []func(cpu *CPU){func(cpu *CPU) {
		if cpu.interrupts.IsHaltBug() {
			cpu.haltbug = true
		} else {
			cpu.halted = true
		}

	}}),

	// Stop system and main clocks
	0x10: NewOpCode(0x10, "STOP /*todo*/", 1, 4, []func(cpu *CPU){func(cpu *CPU) { println("warning: STOP!") }}),

	// Disable interrupts
	0xf3: NewOpCode(0xf3, "DI", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*cpu.interrupts.DisableIME()*/ cpu.interrupts.InterruptsOn = false }}),

	// Enable interrupts
	0xfb: NewOpCode(0xfb, "EI", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*cpu.interrupts.EnableIME()*/ cpu.interrupts.InterruptsEnabling = true }}),

	// No operation
	0x00: NewOpCode(0x00, "NOP", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
}
