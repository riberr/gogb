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

	state      state
	currOpcode OpCode
	currStep   int
	Cycle      int
}

type state int

const (
	FetchOpCode state = 1 << iota
	FetchExtendedOpcode
	Execute
	Halted
	Stopped
	InterruptWait0
	InterruptWait1
	InterruptPushPCHigh
	InterruptPushPCLow
	InterruptCall

	// Useful combinations
	Interruptable     = FetchOpCode | Halted | Stopped
	HandlingInterrupt = InterruptWait0 | InterruptWait1 | InterruptPushPCHigh | InterruptPushPCLow | InterruptCall
)

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
		state:      FetchOpCode,
		Cycle:      0,
	}
}

func (cpu *CPU) Step() {
	cpu.Cycle++
	if cpu.Cycle < 4 {
		return
	}
	cpu.Cycle = 0

	// handle interrupts
	if cpu.state == Interruptable {
		// todo
	}

	// exit HALT even if IME is not set
	// todo

	switch cpu.state {
	case FetchOpCode:
		//println("FetchOpCode")
		cpu.currOpcode = OpCodes[cpu.bus.Read(cpu.pc)]

		if cpu.debug {
			// A: 01 F: B0 B: 00 C: 13 D: 00 E: D8 H: 01 L: 4D SP: FFFE PC: 00:0101 (C3 13 02 CE)
			fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X) [Z: %t, N: %t, H: %t, C: %t] %v\n",
				cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
				cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3),
				cpu.regs.getFlag(FLAG_ZERO_Z_BIT), cpu.regs.getFlag(FLAG_SUBTRACTION_N_BIT), cpu.regs.getFlag(FLAG_HALF_CARRY_H_BIT), cpu.regs.getFlag(FLAG_CARRY_C_BIT),
				cpu.currOpcode.label,
			)
		}

		cpu.pc++
		if cpu.currOpcode.mCycles == 1 {
			cpu.currOpcode.steps[0](cpu)
		} else {
			cpu.currStep = 0
			cpu.state = Execute
		}
	case Execute:
		//println("Execute")
		cpu.jumpStop = false
		cpu.currOpcode.steps[cpu.currStep](cpu)
		cpu.currStep++

		if cpu.jumpStop || cpu.currStep == len(cpu.currOpcode.steps) { // fetch takes one m_cycle
			cpu.state = FetchOpCode
		}

		/*
			for _, step := range cpu.currOpcode.steps {
				step(cpu)

				if cpu.jumpStop {
					break
				}
			}
		*/
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
	if cpu.state == FetchOpCode {
		return fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n",
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
			cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3),
		)
	} else {
		return ""
	}
}

func (cpu *CPU) GetState() state {
	return cpu.state
}
