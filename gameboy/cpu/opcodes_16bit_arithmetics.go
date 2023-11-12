package cpu

var OpCodes16bitArithmeticsGenerated = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 16BIT ARITHMETICS & LOGICAL
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Add (register), Add SP
	0x09: NewOpCode(0x09, "ADD HL,BC", 1, 8, []func(cpu *CPU){
		func(cpu *CPU) { addHL(cpu, cpu.regs.getBC()) },
	}),
	0x19: NewOpCode(0x19, "ADD HL,DE", 1, 8, []func(cpu *CPU){
		func(cpu *CPU) { addHL(cpu, cpu.regs.getDE()) },
	}),
	0x29: NewOpCode(0x29, "ADD HL,HL", 1, 8, []func(cpu *CPU){
		func(cpu *CPU) { addHL(cpu, cpu.regs.getHL()) },
	}),
	0x39: NewOpCode(0x39, "ADD HL,SP", 1, 8, []func(cpu *CPU){
		func(cpu *CPU) { addHL(cpu, cpu.sp) },
	}),
	0xe8: NewOpCode(0xe8, "ADD SP,i8", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.e = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.ee = addSigned8(cpu, cpu.sp, cpu.e) },
		func(cpu *CPU) { cpu.sp = cpu.ee },
	}),

	// Increment (register), Increment SP
	0x03: NewOpCode(0x03, "INC BC", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.incBC() }}),
	0x13: NewOpCode(0x13, "INC DE", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.incDE() }}),
	0x23: NewOpCode(0x23, "INC HL", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.incHL() }}),
	0x33: NewOpCode(0x33, "INC SP", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.sp++ }}),

	// Decrement (register), Decrement SP
	0x0b: NewOpCode(0x0b, "DEC BC", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.decBC() }}),
	0x1b: NewOpCode(0x1b, "DEC DE", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.decDE() }}),
	0x2b: NewOpCode(0x2b, "DEC HL", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.decHL() }}),
	0x3b: NewOpCode(0x3b, "DEC SP", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.sp-- }}),
}
