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

	regs Registers
	sp   uint16 // stack pointer
	pc   uint16 // program counter

	// cpu emulation variables
	jumpStop    bool   // used for conditional jumps
	lsb, msb, e uint8  // temp values when executing opcodes
	ee          uint16 // temp value when executing opcodes
}

func New(bus *bus.Bus, interrupts *interrupts.Interrupts, debug bool) *CPU {
	return &CPU{
		bus:        bus,
		interrupts: interrupts,
		regs:       NewRegisters(),
		sp:         0xFFFE, // post boot rom
		pc:         0x100,  // post boot rom
		debug:      debug,
		jumpStop:   false,
		lsb:        0,
		msb:        0,
		e:          0,
		ee:         0,
	}
}

func (cpu *CPU) Step() {
	curOpCode := OpCodes[cpu.bus.Read(cpu.pc)]

	if cpu.debug {
		// A: 01 F: B0 B: 00 C: 13 D: 00 E: D8 H: 01 L: 4D SP: FFFE PC: 00:0101 (C3 13 02 CE)
		fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X) [Z: %t, N: %t, H: %t, C: %t] %v\n",
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
			cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3),
			cpu.regs.getFlag(FLAG_ZERO_Z_BIT), cpu.regs.getFlag(FLAG_SUBTRACTION_N_BIT), cpu.regs.getFlag(FLAG_HALF_CARRY_H_BIT), cpu.regs.getFlag(FLAG_CARRY_C_BIT),
			curOpCode.label,
		)
	}

	// represents the 'fetch' step
	cpu.pc++
	cpu.jumpStop = false

	for _, step := range curOpCode.steps {
		step(cpu)

		if cpu.jumpStop {
			break
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
