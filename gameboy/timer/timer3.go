/*
* Implements the Timer functionality of the Gameboy
 */

package timer

import "gogb/gameboy/interrupts"

type Timer3 struct {
	interrupts  *interrupts.Interrupts
	DivCounter  int    // Divider counter
	TimaCounter int    // Timer counter
	DIV         uint32 // Divider register (0xFF04)
	TIMA        uint32 // Timer counter (0xFF05)
	TMA         uint32 // Timer modulo (0xFF06)
	TAC         uint32 // Timer control (0xFF07)
}

func NewTimer3(interrupts *interrupts.Interrupts) *Timer3 {
	return &Timer3{
		interrupts: interrupts,
		DivCounter: 0,
		DIV:        0x18,
		TIMA:       0x00,
		TMA:        0x00,
		TAC:        0xF8,
	}
}

func (t *Timer3) Reset() {
	t.DivCounter = 0
	t.TimaCounter = 0
	t.DIV = 0x00
	t.TIMA = 0x00
	t.TMA = 0x00
	t.TAC = 0x00
}

func (t *Timer3) Enabled() bool {
	return (t.TAC>>2)&1 == 1
}

func (t *Timer3) getClockFreqCount() int {
	switch t.TAC & 0x03 {
	case 0x00:
		return 1024
	case 0x01:
		return 16
	case 0x02:
		return 64
	default:
		return 256
	}
}

func (t *Timer3) updateDividerRegister(cycles int, doubleSpeedMode bool) {
	ds := 1
	if doubleSpeedMode {
		ds = 2
	}
	maxDivCycles := 4194304 / 16384 * ds // (TODO: or 2 if in double speed mode)

	t.DivCounter += cycles

	if t.DivCounter >= maxDivCycles {
		t.DivCounter -= maxDivCycles
		t.DIV++

		if t.DIV > 0xff {
			t.DIV = 0
		}
	}
}

func (t *Timer3) Tick(cycles int) {

	t.updateDividerRegister(cycles, false)

	if t.Enabled() {
		t.TimaCounter += cycles
		freq := t.getClockFreqCount()
		for t.TimaCounter >= freq {
			t.TimaCounter -= freq
			if t.TIMA == 0xFF {
				t.TIMA = t.TMA
				t.interrupts.SetIF(interrupts.TIMER)
				break
			} else {
				t.TIMA++
			}
		}
	}

}

func (t *Timer3) Read(address uint16) uint8 {
	switch address {
	case 0xFF04:
		return uint8(t.DIV)
	case 0xFF05:
		return uint8(t.TIMA)
	case 0xFF06:
		return uint8(t.TMA)
	case 0xFF07:
		return uint8(t.TAC)
	default:
		panic("timer does not have this address")
	}
}

func (t *Timer3) Write(address uint16, value uint8) {
	switch address {
	case 0xFF04: // DIV, which is upper 8 bits of SYSCLK. Writing to it resets it
		t.TimaCounter = 0
		t.DivCounter = 0
		t.DIV = 0
	case 0xFF05: // TIMA, the timer counter
		t.TIMA = uint32(value)
	case 0xFF06: // TMA, the timer modulo
		t.TMA = uint32(value)
	case 0xFF07: // TAC, the timer control
		currentFreq := t.TAC & 0x03
		t.TAC = uint32(value) | 0xF8
		newFreq := t.TAC & 0x03
		if currentFreq != newFreq {
			t.TimaCounter = 0
		}
		return
	}
}
