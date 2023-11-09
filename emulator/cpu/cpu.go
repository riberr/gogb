package cpu

import (
	"fmt"
	"gogb/emulator/memory"
)

type CPU struct {
	regs      Registers
	curOpCode OpCode
	sp        uint16 // stack pointer
	pc        uint16 // program counter
	cb        bool   // cb prefix
	debug     bool   // debug print
}

func NewCPU(debug bool) CPU {
	cpu := CPU{
		regs:  NewRegisters(),
		sp:    0xFFFE, //sp:   0x01,
		pc:    0x100,
		debug: debug,
	}
	return cpu
}

func (cpu *CPU) Step() {
	//if !cpu.cb {
	cpu.curOpCode = OpCodes[memory.BusRead(cpu.pc)]
	//} else {
	//	cpu.curOpCode = OpCodesCB[memory.BusRead(cpu.pc)]
	//	cpu.cb = false
	//}

	/*
		fmt.Printf("%04X: (%02X %02X %02X) AF: %02X%02X BC: %02X%02X DE: %02X%02X HL: %02X%02X SP: %02X op-length: %v op: %v\n",
			pc, memory.BusRead(pc), memory.BusRead(pc+1), memory.BusRead(pc+2),
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.curOpCode.length, cpu.curOpCode.label)
	*/
	// represents the 'fetch' step
	cpu.pc++
	stop = false

	for _, step := range cpu.curOpCode.steps {
		step(cpu)

		if stop {
			break
		}
	}

	if cpu.cb {
		cpu.curOpCode = OpCodesCB[memory.BusRead(cpu.pc)]
		cpu.cb = false
		cpu.pc++
		for _, step := range cpu.curOpCode.steps {
			step(cpu)
		}
	}

	pc := cpu.pc
	// A: 01 F: B0 B: 00 C: 13 D: 00 E: D8 H: 01 L: 4D SP: FFFE PC: 00:0101 (C3 13 02 CE)
	if cpu.debug {
		fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X) [Z: %t, N: %t, H: %t, C: %t] %v\n",
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, pc,
			memory.BusRead(pc), memory.BusRead(pc+1), memory.BusRead(pc+2), memory.BusRead(pc+3),
			cpu.regs.getFlag(FLAG_ZERO_Z_BIT), cpu.regs.getFlag(FLAG_SUBTRACTION_N_BIT), cpu.regs.getFlag(FLAG_HALF_CARRY_H_BIT), cpu.regs.getFlag(FLAG_CARRY_C_BIT),
			cpu.curOpCode.label,
		)
	}
}
