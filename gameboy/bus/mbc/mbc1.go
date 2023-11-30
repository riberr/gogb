package mbc

type MBC1 struct {
	rom         []uint8
	currRomBank uint16

	ram         []uint8
	currRamBank uint16
	ramEnabled  bool

	romBanking bool
}

func NewMBC1(rom []uint8) MBC {
	return &MBC1{
		rom:         rom,
		currRomBank: 1,
		ram:         make([]uint8, 0x8000), //32KB
	}
}

func (m *MBC1) Read(address uint16) uint8 {
	switch {
	case address < 0x4000:
		return m.rom[address]
	case address < 0x8000:
		newAddr := address - 0x4000
		return m.rom[newAddr+(m.currRomBank*0x4000)]
	case 0xA000 <= address && address < 0xC000:
		newAddr := address - 0xA000
		return m.ram[newAddr+(m.currRamBank*0x2000)]
	default:
		panic("mbc error?")
	}
}

// WriteRom attempts to switch the ROM or RAM bank.
func (m *MBC1) WriteRom(address uint16, value uint8) {
	switch {
	case address < 0x2000:
		// RAM enable
		if value&0xF == 0xA {
			m.ramEnabled = true
		} else if value&0xF == 0x0 {
			m.ramEnabled = false
		}
	case address < 0x4000:
		// ROM bank number (lower 5)
		m.currRomBank = (m.currRomBank & 0xe0) | uint16(value&0x1f)
		m.updateRomBankIfZero()
	case address < 0x6000:
		// ROM/RAM banking
		if m.romBanking {
			m.currRomBank = (m.currRomBank & 0x1F) | uint16(value&0xe0)
			m.updateRomBankIfZero()
		} else {
			m.currRamBank = uint16(value & 0x3)
		}
	case address < 0x8000:
		// ROM/RAM select mode
		m.romBanking = value&0x1 == 0x00
		if m.romBanking {
			m.currRamBank = 0
		} else {
			m.currRomBank = m.currRomBank & 0x1F
		}
	}
}

// Update the currRomBank if it is on a value which cannot be used.
func (m *MBC1) updateRomBankIfZero() {
	if m.currRomBank == 0x00 || m.currRomBank == 0x20 || m.currRomBank == 0x40 || m.currRomBank == 0x60 {
		m.currRomBank++
	}
}

// WriteRam writes data to the ram if it is enabled.
func (m *MBC1) WriteRam(address uint16, value uint8) {
	if m.ramEnabled {
		m.ram[(0x2000*m.currRamBank)+address-0xA000] = value
	}
}
