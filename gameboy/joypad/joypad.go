package joypad

import (
	"gogb/gameboy/interrupts"
	"gogb/gameboy/utils"
)

type Button int

const (
	PadRight Button = iota
	PadLeft
	PadUp
	PadDown

	PadA
	PadB
	PadSelect
	PadStart
)

const (
	joyp0             = 0
	joyp1             = 1
	joyp2             = 2
	joyp3             = 3
	joypSelectPad     = 4
	joypSelectButtons = 5
)

type JoyPad struct {
	interrupts *interrupts.Interrupts2
	joyp       uint8 // 0xFF00
	pad        uint8
	buttons    uint8
}

func New(interrupts *interrupts.Interrupts2) *JoyPad {
	return &JoyPad{
		interrupts: interrupts,
		joyp:       0,
		pad:        0xF,
		buttons:    0xF,
	}
}

func (j *JoyPad) KeyEvent(button Button, isPress bool) {
	prevPad := j.pad
	prevButtons := j.buttons

	switch button {
	case PadRight:
		j.pad = setOrClearBit(j.pad, joyp0, isPress)
	case PadLeft:
		j.pad = setOrClearBit(j.pad, joyp1, isPress)
	case PadUp:
		j.pad = setOrClearBit(j.pad, joyp2, isPress)
	case PadDown:
		j.pad = setOrClearBit(j.pad, joyp3, isPress)
	case PadA:
		j.buttons = setOrClearBit(j.buttons, joyp0, isPress)
	case PadB:
		j.buttons = setOrClearBit(j.buttons, joyp1, isPress)
	case PadSelect:
		j.buttons = setOrClearBit(j.buttons, joyp2, isPress)
	case PadStart:
		j.buttons = setOrClearBit(j.buttons, joyp3, isPress)
	}

	res := ((prevPad ^ j.pad) & j.pad) | ((prevButtons ^ j.buttons) & j.buttons)
	if res != 0 {
		j.interrupts.SetInterruptFlag(interrupts.INTR_JOYPAD)
	}
}

func setOrClearBit(n uint8, pos int, isPress bool) uint8 {
	if isPress {
		return utils.ClearBit(n, pos)
	} else {
		return utils.SetBit(n, pos)
	}
}

func (j *JoyPad) GetJoyPadState() uint8 {
	P14 := (j.joyp >> joypSelectPad) & 0x01
	P15 := (j.joyp >> joypSelectButtons) & 0x01

	joyPadState := 0xFF & (j.joyp | 0b11001111)
	if P14 == 0 {
		joyPadState &= j.pad
	}

	if P15 == 0 {
		joyPadState &= j.buttons
	}

	return joyPadState
}

func (j *JoyPad) Write(address uint16, value uint8) {
	if address == 0xFF00 {
		j.joyp = value
	}
}
