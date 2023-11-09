package cpu

var OpCodesRotateShiftBitoperations = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Rotates, Shifts, Bit operations
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Rotate A left. Old bit 7 to Carry flag
	0x07: NewOpCode(0x07, "RLCA", 1, 4, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.a = RLC(cpu, cpu.regs.a); cpu.regs.setFlag(FLAG_ZERO_Z_BIT, false) },
	}),
	// Rotate A left through Carry flag
	0x17: NewOpCode(0x17, "RLA", 1, 4, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.a = RL(cpu, cpu.regs.a); cpu.regs.setFlag(FLAG_ZERO_Z_BIT, false) },
	}),
	// Rotate A right. Old bit 0 to Carry flag
	0x0f: NewOpCode(0x0f, "RRCA", 1, 4, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.a = RRC(cpu, cpu.regs.a); cpu.regs.setFlag(FLAG_ZERO_Z_BIT, false) },
	}),
	// Rotate A right through Carry flag
	0x1f: NewOpCode(0x1f, "RRA", 1, 4, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.a = RR(cpu, cpu.regs.a); cpu.regs.setFlag(FLAG_ZERO_Z_BIT, false) },
	}),

	// Increment (register), Increment SP
	0xcb: NewOpCode(0xcb, "PREFIX CB", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.cb = true }}),
}
