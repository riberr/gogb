package interrupts

const (
	INTR_VBLANK_ADDR    = 0x0040
	INTR_LCDSTAT_ADDR   = 0x0048
	INTR_TIMER_ADDR     = 0x0050
	INTR_SERIAL_ADDR    = 0x0058
	INTR_HIGHTOLOW_ADDR = 0x0060
)

var InterruptAddresses = map[byte]uint16{
	0: INTR_VBLANK_ADDR,    // V-Blank
	1: INTR_LCDSTAT_ADDR,   // LCDC Status
	2: INTR_TIMER_ADDR,     // Timer Overflow
	3: INTR_SERIAL_ADDR,    // Serial Transfer
	4: INTR_HIGHTOLOW_ADDR, // Hi-Lo P10-P13
}

const (
	INTR_VBLANK      = 0
	INTR_LCD         = 1
	INTR_TIMER       = 2
	INTR_SERIAL_LINK = 3
	INTR_JOYPAD      = 4
)

type Interrupts2 struct {
	InterruptsEnabling bool  // Interrupts are being enabled
	InterruptsOn       bool  // Interrupts are on
	IE                 uint8 // Interrupt enable register
	IF                 uint8 // Interrupt flag register

}

func NewInterrupts2() *Interrupts2 {
	return &Interrupts2{}
}

func (i *Interrupts2) CheckValidInterrupts() uint8 {
	valid_interrupts := i.IE & i.IF

	// Find flags that are set in IF but not enabled in IE
	// disabled_interrupts := i.IF &^ i.IE

	// if disabled_interrupts != 0 {
	// 	internal.Logger.Warningf("Warning: Interrupt flags are set but not enabled. IE: %08b, IF: %08b", i.IE, i.IF)
	// }

	return valid_interrupts
}

func (i *Interrupts2) SetInterruptFlag(f uint8) {
	req := i.IF | 0xE0
	i.IF = req | (1 << f)
}
