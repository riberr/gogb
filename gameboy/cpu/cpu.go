package cpu

import (
	"fmt"
	"gogb/gameboy/bus"
	"gogb/gameboy/interrupts"
	"gogb/utils"
)

type CPU struct {
	bus        *bus.Bus
	interrupts *interrupts.Interrupts
	debug      bool // debug print

	regs      Registers
	curOpCode OpCode
	sp        uint16 // stack pointer
	pc        uint16 // program counter
	cb        bool   // cb prefix
}

func New(bus *bus.Bus, interrupts *interrupts.Interrupts, debug bool) *CPU {
	return &CPU{
		bus:        bus,
		interrupts: interrupts,
		regs:       NewRegisters(),
		sp:         0xFFFE, // post boot rom
		pc:         0x100,  // post boot rom
		debug:      debug,
	}
}

func (cpu *CPU) Step() {
	cpu.curOpCode = OpCodes[cpu.bus.Read(cpu.pc)]

	/*
		fmt.Printf("%04X: (%02X %02X %02X) AF: %02X%02X BC: %02X%02X DE: %02X%02X HL: %02X%02X SP: %02X op-length: %v op: %v\n",
			pc, bus.Read(pc), bus.Read(pc+1), bus.Read(pc+2),
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.curOpCode.length, cpu.curOpCode.label)
	*/

	pc := cpu.pc
	// A: 01 F: B0 B: 00 C: 13 D: 00 E: D8 H: 01 L: 4D SP: FFFE PC: 00:0101 (C3 13 02 CE)
	if cpu.debug {
		fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X) [Z: %t, N: %t, H: %t, C: %t] %v\n",
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, pc,
			cpu.bus.Read(pc), cpu.bus.Read(pc+1), cpu.bus.Read(pc+2), cpu.bus.Read(pc+3),
			cpu.regs.getFlag(FLAG_ZERO_Z_BIT), cpu.regs.getFlag(FLAG_SUBTRACTION_N_BIT), cpu.regs.getFlag(FLAG_HALF_CARRY_H_BIT), cpu.regs.getFlag(FLAG_CARRY_C_BIT),
			cpu.curOpCode.label,
		)
	}

	// represents the 'fetch' step
	cpu.pc++
	stop = false

	for _, step := range cpu.curOpCode.steps {
		step(cpu)

		if stop {
			break
		}
	}

	// if prefix CB
	if cpu.cb {
		cpu.curOpCode = OpCodesCB[cpu.bus.Read(cpu.pc)]
		cpu.cb = false
		cpu.pc++
		for _, step := range cpu.curOpCode.steps {
			step(cpu)
		}
	}

	// interrupts
	if cpu.interrupts.IsIME() {
		flag := cpu.interrupts.GetEnabledFlaggedInterrupt()
		if flag == -1 {

		} else {
			cpu.interrupts.DisableIME()
			cpu.interrupts.ClearIF(flag)
			cpu.sp--
			cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc))
			cpu.sp--
			cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc))
			cpu.pc = interrupts.ISR_address[flag]
		}
	}
}

// GetInternalState returns a string representing the internal state of the cpu
func (cpu *CPU) GetInternalState() string {
	return fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n",
		cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
		cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3),
	)
}
