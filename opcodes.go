package main

import "gogb/utils"

type OpCode struct {
	value   uint8            // for instance 0xC3
	label   string           // JP {0:x4}
	length  int              // in bytes
	tCycles int              // clock cycles
	mCycles int              // machine cycles
	steps   []func(cpu *CPU) // function array
}

func NewOpCode(value uint8, label string, length int, tCycles int, steps []func(cpu *CPU)) OpCode {
	return OpCode{
		value:   value,
		label:   label,
		length:  length,
		tCycles: tCycles,
		mCycles: tCycles / 4,
		steps:   steps,
	}
}

var lsb, msb uint8 = 0, 0
var OpCodes = map[uint8]OpCode{
	0x00: NewOpCode(0x00, "NOP", 1, 4, []func(cpu *CPU){}),
	0x03: NewOpCode(0x03, "INC BC", 1, 8, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),
	0x0B: NewOpCode(0x0B, "DEC BC", 1, 8, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),

	0x11: NewOpCode(0x11, "LD DE,u16", 3, 12, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.e = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.regs.d = busRead(cpu.pc); cpu.pc++ },
	}),

	0x21: NewOpCode(0x21, "LD HL,u16", 3, 12, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.l = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.regs.h = busRead(cpu.pc); cpu.pc++ },
	}),

	0x41: NewOpCode(0x41, "LD B,C", 1, 4, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.b = cpu.regs.c },
	}),
	0x47: NewOpCode(0x47, "LD B,A", 1, 4, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.b = cpu.regs.a },
	}),

	0x66: NewOpCode(0x66, "LD H,(HL)", 1, 8, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),

	0x73: NewOpCode(0x73, "LD (HL),E", 1, 8, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),

	0x83: NewOpCode(0x83, "ADD A,E", 1, 4, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),

	0xC3: NewOpCode(0xC3, "JP u16", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xCC: NewOpCode(0xCC, "CALL Z,u16", 3, 12 /*12-24*/, []func(cpu *CPU){
		func(cpu *CPU) { cpu.pc = cpu.pc + 2 },
	}),
	0xCE: NewOpCode(0xCE, "ADC A,u8", 2, 8, []func(cpu *CPU){
		func(cpu *CPU) { cpu.pc++ },
	}),
}
