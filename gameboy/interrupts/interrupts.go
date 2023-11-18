package interrupts

import (
	"gogb/utils"
)

// Bit position
// 0   Vblank
// 1   LCD
// 2   Timer
// 3   Serial Link
// 4   Joypad

type Interrupts struct {
	ime         bool // IME: Interrupt master enable flag [write only]
	imeEnabling bool
	_ie         uint8 // FFFF — IE: Interrupt enable
	_if         uint8 // FF0F — IF: Interrupt flag
}

var ISR_address = []uint16{
	0x0040, // Vblank
	0x0048, // LCD Status
	0x0050, // TimerOverflow
	0x0058, // SerialLink
	0x0060, // JoypadPress,
}

type Flag int

const (
	VBLANK      Flag = 0
	LCD         Flag = 1
	TIMER       Flag = 2
	SERIAL_LINK Flag = 3
	JOYPAD      Flag = 4
)

func New() *Interrupts {
	return &Interrupts{
		ime: false,
		_ie: 0,
		_if: 0xE1, //0xE1??
	}
}

func (i *Interrupts) IsInterruptsRequested() bool {
	return i._if&i._ie != 0
}

func (i *Interrupts) IsHaltBug() bool {
	println("is haltbug!")
	return (i._ie&i._if&0x1F) != 0 && !i.ime
}

func (i *Interrupts) GetIF() uint8 {
	return i._if
}

func (i *Interrupts) SetAllIF(value uint8) {
	i._if = value
}

func (i *Interrupts) SetIF(flag Flag) {
	i._if = utils.SetBit(i._if, int(flag))
}

func (i *Interrupts) ClearIF(flag Flag) {
	i._if = utils.ClearBit(i._if, int(flag))
}

func (i *Interrupts) IsIF(flag Flag) bool {
	return utils.HasBit(i._if, int(flag))
}

func (i *Interrupts) GetIE() uint8 {
	return i._ie
}

func (i *Interrupts) SetAllIE(value uint8) {
	i._ie = value
}

func (i *Interrupts) SetIE(flag Flag) {
	i._ie = utils.SetBit(i._ie, int(flag))
}

func (i *Interrupts) IsIE(flag Flag) bool {
	return utils.HasBit(i._ie, int(flag))
}

func (i *Interrupts) GetEnabledFlaggedInterrupt() Flag {
	switch {
	case i.IsIE(VBLANK) && i.IsIF(VBLANK):
		return VBLANK
	case i.IsIE(LCD) && i.IsIF(LCD):
		return LCD
	case i.IsIE(TIMER) && i.IsIF(TIMER):
		return TIMER
	case i.IsIE(SERIAL_LINK) && i.IsIF(SERIAL_LINK):
		return SERIAL_LINK
	case i.IsIE(JOYPAD) && i.IsIF(JOYPAD):
		return JOYPAD
	default:
		return -1
	}
}

func (i *Interrupts) GetIMEEnabling() bool {
	return i.imeEnabling
}

func (i *Interrupts) SetIMEEnabling(value bool) {
	i.imeEnabling = value
}

func (i *Interrupts) DisableIME() {
	i.ime = false
}

func (i *Interrupts) EnableIME() {
	i.ime = true
}

func (i *Interrupts) IsIME() bool {
	return i.ime
}
