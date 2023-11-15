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

	state         state
	currOpcode    OpCode
	currStep      int
	Cycle         int
	interruptFlag interrupts.Flag
	Log           string
	NewLog        bool
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

	Interruptable = FetchOpCode | Halted | Stopped
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

	if (cpu.state&Interruptable != 0) && cpu.interrupts.IsIME() && cpu.interrupts.GetEnabledFlaggedInterrupt() != -1 {
		cpu.state = InterruptWait0
	}

	if (cpu.state == Halted) && cpu.interrupts.GetEnabledFlaggedInterrupt() != -1 {
		cpu.state = FetchOpCode
	}

	switch cpu.state {
	case Halted, Stopped:
		// do nothing
		cpu.Log = fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X",
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
			cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3))
		cpu.NewLog = true
		return
	case FetchOpCode:
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
		cpu.Log = fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X",
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
			cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3))
		cpu.NewLog = true

		cpu.pc++
		if cpu.currOpcode.mCycles == 1 {
			if cpu.currOpcode.value == 0x76 {
				cpu.state = Halted
				return
			}
			cpu.currOpcode.steps[0](cpu)
		} else {
			cpu.currStep = 0
			cpu.state = Execute
		}
	case Execute:
		cpu.jumpStop = false
		cpu.currOpcode.steps[cpu.currStep](cpu)
		cpu.currStep++

		if cpu.jumpStop || cpu.currStep == len(cpu.currOpcode.steps) { // fetch takes one m_cycle
			cpu.state = FetchOpCode
		}
	case InterruptWait0:
		// [TCAGBD:4.9] mentions a 2-cycle idle upon handling interrupt request.
		cpu.state = InterruptWait1
	case InterruptWait1:
		cpu.interruptFlag = cpu.interrupts.GetEnabledFlaggedInterrupt()
		if cpu.interruptFlag == -1 {
			cpu.state = FetchOpCode
		} else {
			cpu.interrupts.DisableIME()
			cpu.interrupts.ClearIF(cpu.interruptFlag)
			cpu.state = InterruptPushPCHigh
		}
	case InterruptPushPCHigh:
		cpu.sp--
		cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc))
		cpu.state = InterruptPushPCLow
	case InterruptPushPCLow:
		cpu.sp--
		cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc))
		cpu.state = InterruptCall
	case InterruptCall:
		cpu.pc = interrupts.ISR_address[cpu.interruptFlag]
		cpu.state = FetchOpCode
	}
}

// GetInternalString returns a string representing the internal state of the cpu
func (cpu *CPU) GetInternalString() string {
	return fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n",
		cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
		cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3))
}

func (cpu *CPU) GetState() state {
	return cpu.state
}
