package timer

import (
	"fmt"
	"gogb/gameboy/interrupts"
)

// Timer is kinda ported from https://github.com/raddad772/jsmoo/blob/main/system/gb/gb_cpu.js
// and https://github.com/rvaccarim/FrozenBoy/blob/master/FrozenBoyCore/Processor/Timer.cs
type Timer struct {
	interrupts *interrupts.Interrupts2

	sysclk int
	tima   uint8
	tma    uint8
	tac    uint8

	lastBit            bool
	ticksSinceOverflow uint8
	overflow           bool
}

var FreqToBit = []int{9, 3, 5, 7}

func New(interrupts *interrupts.Interrupts2) *Timer {
	return &Timer{
		interrupts: interrupts,
	}
}

func (t *Timer) UpdateTimers(cycles int) {
	fmt.Printf("%04x\n", t.sysclk)
	for i := 0; i < cycles; i++ {
		t.Tick()
	}
}

func (t *Timer) Tick() {
	t.updateSysClk((t.sysclk + 1) & 0xFFFF)
	if !t.overflow {
		return
	}
	t.ticksSinceOverflow++
	if t.ticksSinceOverflow == 4 {
		t.interrupts.SetInterruptFlag(interrupts.INTR_TIMER)
	}
	if t.ticksSinceOverflow == 5 {
		t.tima = t.tma
	}
	if t.ticksSinceOverflow == 6 {
		t.tima = t.tma
		t.overflow = false
		t.ticksSinceOverflow = 0
	}
}

func (t *Timer) updateSysClk(newValue int) {
	t.sysclk = newValue

	bitPos := FreqToBit[t.tac&0b11]
	//bitPos <<= _speedMode.GetSpeedMode() - 1;
	//bitPos <<= 0

	/*
		bit := utils.ToInt((t.sysclk & (1 << bitPos)) != 0)
		bit &= utils.ToInt((t.tac & (1 << 2)) != 0)
		if bit == 0 && t.lastBit {
			t.UpdateTima()
		}
	*/
	bit := (t.sysclk & (1 << bitPos)) != 0
	bitTemp := (t.tac & (1 << 2)) != 0
	bit = bit && bitTemp
	if !bit && t.lastBit {
		t.UpdateTima()
	}

	//bitBool := bit != 0
	t.lastBit = bit
}

func (t *Timer) UpdateTima() {
	//t.tima = (t.tima + 1) & 0xFF // Increment TIMA

	t.tima++
	println("tock")
	fmt.Printf("tock tima: %02x\n", t.tima)
	if t.tima == 0 { // or == 0xFF?
		t.overflow = true
		t.ticksSinceOverflow = 0
	}
}

func (t *Timer) Read(address uint16) uint8 {
	switch address {
	case 0xFF04:
		return uint8((t.sysclk >> 8) & 0xFF)
	case 0xFF05:
		fmt.Printf("read tima: %02x\n", t.tima)
		return t.tima
	case 0xFF06:
		return t.tma
	case 0xFF07:
		return t.tac | 0b11111000
	default:
		panic("timer does not have this address")
	}
}

func (t *Timer) Write(address uint16, value uint8) {
	switch address {
	case 0xFF04: // DIV, which is upper 8 bits of SYSCLK. Writing to it resets it
		t.updateSysClk(0)
	case 0xFF05: // TIMA, the timer counter
		fmt.Printf("write tima, before: %02x\n", t.tima)
		if t.ticksSinceOverflow < 5 {
			t.tima = value
			t.overflow = false
			t.ticksSinceOverflow = 0
		}
		fmt.Printf("write tima, after: %02x\n", t.tima)
	case 0xFF06: // TMA, the timer modulo
		t.tma = value
	case 0xFF07: // TAC, the timer control
		t.tac = value
	}
}
