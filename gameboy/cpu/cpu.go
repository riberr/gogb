package cpu

import (
	"fmt"
	"gogb/gameboy/bus"
	"gogb/gameboy/interrupts"
	"gogb/gameboy/utils"
)

type CPU struct {
	bus        *bus.Bus
	interrupts *interrupts.Interrupts2
	debug      bool // debug print

	regs    Registers
	sp      uint16 // stack pointer
	pc      uint16 // program counter
	halted  bool
	haltbug bool

	// cpu emulation variables
	jumpStop    bool   // used for conditional jumps
	lsb, msb, e uint8  // temp values when executing opcodes
	ee          uint16 // temp value when executing opcodes

	thisCpuTicks int
}

func New(bus *bus.Bus, interrupts *interrupts.Interrupts2, debug bool) *CPU {
	return &CPU{
		bus:        bus,
		interrupts: interrupts,
		regs:       NewRegisters(),
		sp:         0xFFFE, // post boot rom
		pc:         0x100,  // post boot rom
		halted:     false,
		haltbug:    false,
		debug:      debug,
		jumpStop:   false,
		lsb:        0,
		msb:        0,
		e:          0,
		ee:         0,
	}
}

func (cpu *CPU) Step() int {
	opcode := OpCodes[cpu.bus.Read(cpu.pc)]
	cpu.printDebug(opcode)

	if cpu.halted {
		cpu.thisCpuTicks = 4
		return cpu.thisCpuTicks
	}

	if cpu.haltbug {
		cpu.haltbug = false
	} else {
		cpu.pc++
	}

	cpu.jumpStop = false
	if opcode.tCycles == 4 {
		cpu.thisCpuTicks = 0 // there are some ALU operations that can be completed in the same cycle as fetch (fetch / overlap)
	} else {
		cpu.thisCpuTicks = 4 // fetch counts for 4 tCycles
	}

	for _, step := range opcode.steps {
		if cpu.jumpStop {
			break
		}

		step(cpu)
		cpu.thisCpuTicks += 4

	}

	/*
		var temp = opcode.tCycles
		if cpu.jumpStop {
			temp -= 4
		}
		if opcode.value != 0xc4 && opcode.value != 0xcb {
			if cpu.thisCpuTicks != temp {
				println(cpu.jumpStop)
				fmt.Printf("%v: got %v, want %v ", opcode.label, cpu.thisCpuTicks, temp)
				panic("not equal!")
			}
		}
	*/

	//fmt.Printf("ticks %v jumpstop %v\n", cpu.thisCpuTicks, cpu.jumpStop)
	return cpu.thisCpuTicks
}

func (cpu *CPU) DoInterrupts() (cycles int) {
	if cpu.interrupts.InterruptsEnabling {
		cpu.interrupts.InterruptsOn = true
		cpu.interrupts.InterruptsEnabling = false
		return 0
	}
	if !cpu.interrupts.InterruptsOn && !cpu.halted {
		return 0
	}

	req := cpu.bus.Read(0xFF0F)
	enabled := cpu.bus.Read(0xFFFF)

	if req > 0 {
		var i byte
		for i = 0; i < 5; i++ {
			if utils.TestBit(req, int(i)) && utils.TestBit(enabled, int(i)) {
				cpu.serviceInterrupt(i)
				return 20
			}
		}
	}
	return 0
}

// Called if an interrupt has been raised. Will check if interrupts are
// enabled and will jump to the interrupt address.
func (cpu *CPU) serviceInterrupt(interrupt byte) {
	// If was halted without interrupts, do not jump or reset IF
	if !cpu.interrupts.InterruptsOn && cpu.halted {
		cpu.halted = false
		return
	}
	cpu.interrupts.InterruptsOn = false
	cpu.halted = false

	req := cpu.bus.Read(0xFF0F)
	req = utils.ClearBit(req, int(interrupt))
	cpu.bus.Write(0xFF0F, req)

	cpu.sp--
	cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc))
	cpu.sp--
	cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc))
	cpu.pc = interrupts.ISR_address[interrupt]
}

/*
func (cpu *CPU) DoInterrupts() int {
	if cpu.interrupts.GetIMEEnabling() {
		cpu.interrupts.SetIMEEnabling(false)
		cpu.interrupts.EnableIME()
		return 0
	}
	if !cpu.interrupts.IsIME() && !cpu.halted {
		return 0
	}

	interruptFlag := cpu.interrupts.GetEnabledFlaggedInterrupt()
	if interruptFlag == -1 {
		return 0
	}

	cpu.serviceInterrupts(interruptFlag)
	return 20
}

func (cpu *CPU) serviceInterrupts(interruptFlag interrupts.Flag) {

	// If was halted without interrupts, do not jump or reset IF
	if !cpu.interrupts.IsIME() && cpu.halted {
		cpu.halted = false
		return
	}

	cpu.halted = false
	cpu.interrupts.DisableIME()
	cpu.interrupts.ClearIF(interruptFlag)

	cpu.sp--
	cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc))
	cpu.sp--
	cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc))
	cpu.pc = interrupts.ISR_address[interruptFlag]
}

*/

// GetInternalString returns a string representing the internal state of the cpu
func (cpu *CPU) GetInternalString() string {
	return fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X\n",
		cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
		cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3))
}

func (cpu *CPU) printDebug(opcode OpCode) {
	if cpu.debug {
		// A: 01 F: B0 B: 00 C: 13 D: 00 E: D8 H: 01 L: 4D SP: FFFE PC: 00:0101 (C3 13 02 CE)
		fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X) [Z: %t, N: %t, H: %t, C: %t] last ticks: %v, %v\n",
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
			cpu.bus.Read(cpu.pc), cpu.bus.Read(cpu.pc+1), cpu.bus.Read(cpu.pc+2), cpu.bus.Read(cpu.pc+3),
			cpu.regs.getFlag(FLAG_ZERO_Z_BIT), cpu.regs.getFlag(FLAG_SUBTRACTION_N_BIT), cpu.regs.getFlag(FLAG_HALF_CARRY_H_BIT), cpu.regs.getFlag(FLAG_CARRY_C_BIT),
			cpu.thisCpuTicks, opcode.label,
		)
	}
}

func (cpu *CPU) GetPC() uint16 {
	return cpu.pc
}

func (cpu *CPU) GetHalted() bool {
	return cpu.halted
}
