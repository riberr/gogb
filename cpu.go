package main

import "fmt"

type CPU struct {
	regs      Registers
	curOpCode OpCode
	sp        uint16
	pc        uint16
}

func NewCPU() CPU {
	cpu := CPU{
		regs: NewRegisters(),
		sp:   0x01,
		pc:   0x100,
	}
	return cpu
}

func (cpu *CPU) Step() {
	cpu.curOpCode = OpCodes[busRead(cpu.pc)]

	pc := cpu.pc
	fmt.Printf("%04X: (%02X %02X %02X) A: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X. opcode: %v\n",
		pc, busRead(pc), busRead(pc+1), busRead(pc+2),
		cpu.regs.a, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.curOpCode.label)

	// represents the 'fetch' step
	cpu.pc++

	for _, step := range cpu.curOpCode.steps {
		step(cpu)
	}
}
