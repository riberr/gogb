package timer

import "gogb/gameboy/interrupts"

// Timer is kinda ported from https://github.com/raddad772/jsmoo/blob/main/system/gb/gb_cpu.js
type Timer struct {
	interrupts interrupts.Interrupts

	sysclk uint16
	tima   uint8
	tma    uint8
	tac    uint8

	lastBit          uint8
	cyclesTilTimaIrq uint8
	timaReloadCycle  bool
	tmaReloadCycle   bool
}

func New(interrupts *interrupts.Interrupts) *Timer {
	return &Timer{
		interrupts: *interrupts,
	}
}

func (t *Timer) Tick() {
	t.timaReloadCycle = false
	t.updateSysClk((t.sysclk + 1) & 0xFFFF)

	if t.cyclesTilTimaIrq > 0 {
		t.cyclesTilTimaIrq--
		if t.cyclesTilTimaIrq == 0 {
			t.interrupts.SetIF(interrupts.TIMER)
			t.tima = t.tma
			t.timaReloadCycle = true
		}
	}
}

func (t *Timer) updateSysClk(newValue uint16) {
	t.sysclk = newValue

	var thisBit uint8
	switch t.tac & 0b11 {
	case 0b00: // 4096Hz, CPU Clock/1024
		//fmt.Printf("%02x \n", (t.sysclk>>6)&0xFF)
		thisBit = uint8(t.sysclk>>7) & 1
	case 0b01: // 262144Hz CPU Clock/16
		thisBit = uint8(t.sysclk>>1) & 1
	case 0b10: // 65536Hz, CPU Clock/64
		thisBit = uint8(t.sysclk>>3) & 1
	case 0b11: // 16384Hz, CPU Clock/256
		thisBit = uint8(t.sysclk>>5) & 1
	default:
		panic("illegal clockSelect")
	}
	thisBit &= (t.tac & 4) >> 2 // thisBit = clock enable

	t.detectEdge(t.lastBit, thisBit)
	t.lastBit = thisBit
}

// detectEdge detects falling edge
func (t *Timer) detectEdge(before uint8, after uint8) {
	if (before == 1) && (after == 0) {
		t.tima = (t.tima + 1) & 0xFF // Increment TIMA
		if t.tima == 0 {             // If we overflow, schedule IRQ
			t.cyclesTilTimaIrq = 1
		}
	}
}

func (t *Timer) Read(address uint16) uint8 {
	switch address {
	case 0xFF04: // DIV, bits 6-13 of SYSCLK	// todo: should be upper 8 bits of sysclk?
		return uint8((t.sysclk >> 6) & 0xFF)
	case 0xFF05:
		return t.tima
	case 0xFF06:
		return t.tma
	case 0xFF07:
		return t.tac
	default:
		panic("timer does not have this address")
	}
}

func (t *Timer) Write(address uint16, value uint8) {
	switch address {
	case 0xFF04: // DIV, which is upper 8 bits of SYSCLK. Writing to it resets it
		t.updateSysClk(0)
	case 0xFF05: // TIMA, the timer counter
		if !t.timaReloadCycle {
			t.tima = value
		}
		// "During the strange cycle [A] you can prevent the IF flag from being set and prevent the TIMA from
		// being reloaded from TMA by writing a value to TIMA.
		// That new value will be the one that stays in the TIMA register after the instruction.
		// Writing to DIV, TAC or other registers wont prevent the IF flag from being set or TIMA from being reloaded."
		if t.cyclesTilTimaIrq == 1 {
			t.cyclesTilTimaIrq = 0
		}
	case 0xFF06: // TMA, the timer modulo
		// "If TMA is written the same cycle it is loaded to TIMA [B], TIMA is also loaded with that value."
		if t.timaReloadCycle {
			t.tima = value
		}
		t.tma = value
	case 0xFF07: // TAC, the timer control
		lastBit := t.lastBit
		t.lastBit &= (value & 4) >> 2
		t.detectEdge(lastBit, t.lastBit)
		t.tac = value
	}
}
