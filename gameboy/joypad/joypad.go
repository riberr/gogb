package joypad

import (
	"fmt"
	"gogb/gameboy/interrupts"
	"gogb/gameboy/utils"
)

type Button int

const (
	PadLeft   Button = 0
	PadUp     Button = 1
	PadRight  Button = 2
	PadDown   Button = 3
	PadStart  Button = 4
	PadSelect Button = 5
	PadB      Button = 6
	PadA      Button = 7
)

type JoyPad struct {
	interrupts  *interrupts.Interrupts
	joyp        uint8 // 0xFF00
	buttonState uint8 // holds all 8 buttons
}

func New(interrupts *interrupts.Interrupts) *JoyPad {
	return &JoyPad{
		interrupts:  interrupts,
		joyp:        0,
		buttonState: 0xff,
	}
}

func (j *JoyPad) ButtonPressed(button Button) {
	previouslyUnset := false
	fmt.Printf("press button: %v\n", button)

	// if setting from 1 to 0 we may have to request an interupt
	if utils.TestBit(j.buttonState, int(button)) {
		previouslyUnset = true
	}

	// remember if a keypressed its bit is 0 not 1
	j.buttonState = utils.ClearBit(j.buttonState, int(button))

	// is this a standard button or a directional button?
	var isButton bool
	if button > 3 {
		isButton = true
	} else { // directional button pressed
		isButton = false
	}

	requestInterrupt := false

	// only request interrupt if the button just pressed is the style of button the game is interested in
	if isButton && !utils.TestBit(j.joyp, 5) {
		requestInterrupt = true
	} else if !isButton && !utils.TestBit(j.joyp, 4) {
		requestInterrupt = true
	}

	if requestInterrupt && !previouslyUnset {
		fmt.Printf("req int: keys: %08b\n", j.buttonState)
		j.interrupts.SetIF(interrupts.JOYPAD)
	}
}

func (j *JoyPad) ButtonReleased(button Button) {
	fmt.Printf("release button: %v\n", button)
	j.buttonState = utils.SetBit(j.buttonState, int(button))
}

func (j *JoyPad) GetJoyPadState() uint8 {
	res := j.joyp
	// flip all bits
	res ^= 0xFF

	// are we interested in the standard buttons?
	if !utils.TestBit(res, 4) {
		topJoyPad := j.buttonState >> 4
		topJoyPad |= 0xF0 // turn the top 4 bits on
		res &= topJoyPad  // show what buttons are pressed
	} else if !utils.TestBit(res, 5) { //directional buttons
		bottomJoyPad := j.buttonState & 0xF
		bottomJoyPad |= 0xF0
		res &= bottomJoyPad
	}
	return res
}

func (j *JoyPad) Write(address uint16, value uint8) {
	if address == 0xFF00 {
		j.joyp = value
	}
}
