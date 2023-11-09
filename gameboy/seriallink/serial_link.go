package seriallink

// https://gbdev.io/pandocs/Serial_Data_Transfer_(Link_Cable).html
// https://gbdev.gg8.se/wiki/articles/Serial_Data_Transfer_(Link_Cable)
type SerialLink struct {
	sb  uint8  // 0xFF01: Serial transfer data
	sc  uint8  // 0xFF02: Serial transfer control
	log []byte // log the output of the serial link
}

func New() *SerialLink {
	return &SerialLink{
		sb: 0,
		sc: 0,
	}
}

const (
	cpuFreq   = 4194304 // Hz
	clockFreq = 8192    // Hz
)

func (sl *SerialLink) GetSB() uint8 {
	return sl.sb
}

func (sl *SerialLink) SetSB(value uint8) {
	sl.sb = value
}

func (sl *SerialLink) GetSC() uint8 {
	return sl.sc
}

func (sl *SerialLink) SetSC(value uint8) {
	sl.sc = value

	if value == 0x81 {
		sl.log = append(sl.log, sl.sb)
	}
}

func (sl *SerialLink) GetLog() string {
	return string(sl.log)
}

func (sl *SerialLink) Tick() {

}
