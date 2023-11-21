package timer

import (
	interrupts2 "gogb/gameboy/interrupts"
	"gogb/gameboy/utils"
)

// Not used. Implementation taken from GoBoy
type Timer2 struct {
	interrupts   *interrupts2.Interrupts
	timerCounter int
	Divider      int

	div  uint8
	tima uint8
	tma  uint8
	tac  uint8
}

func NewTimer2(interrupts *interrupts2.Interrupts) *Timer2 {
	return &Timer2{
		interrupts: interrupts,
		div:        0x1E,
		tima:       0x00,
		tma:        0x00,
		tac:        0xF8,
	}
}

func (t *Timer2) UpdateTimers(cycles int) {
	t.dividerRegister(cycles)
	if t.isClockEnabled() {
		t.timerCounter += cycles

		freq := t.getClockFreqCount()
		for t.timerCounter >= freq {
			t.timerCounter -= freq
			if t.tima == 0xFF {
				t.tima = t.tma
				t.interrupts.SetIF(interrupts2.TIMER)
			} else {
				t.tima++
			}
		}
	}
}

func (t *Timer2) isClockEnabled() bool {
	return utils.HasBit(t.tac, 2)
}

func (t *Timer2) GetClockFreq() byte {
	return t.tac & 0x3
}

func (t *Timer2) getClockFreqCount() int {
	switch t.GetClockFreq() {
	case 0:
		return 1024
	case 1:
		return 16
	case 2:
		return 64
	default:
		return 256
	}
}

func (t *Timer2) SetClockFreq() {
	t.timerCounter = 0
}

func (t *Timer2) dividerRegister(cycles int) {
	t.Divider += cycles
	if t.Divider >= 255 {
		t.Divider -= 255
		t.div++
	}
}

func (t *Timer2) Read(address uint16) uint8 {
	switch address {
	case 0xFF04:
		return t.div
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

func (t *Timer2) Write(address uint16, value uint8) {
	switch address {
	case 0xFF04:
		t.SetClockFreq()
		t.Divider = 0
		t.div = 0
	case 0xFF05:
		t.tima = value
	case 0xFF06:
		t.tma = value
	case 0xFF07:
		// Timer control
		currentFreq := t.GetClockFreq()
		t.tac = value | 0xF8
		newFreq := t.GetClockFreq()

		if currentFreq != newFreq {
			t.SetClockFreq()
		}
	}

}
