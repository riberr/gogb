package mbc

import "gogb/gameboy/utils"

type MBC3 struct {
	rom             []uint8
	selectedRomBank uint16
	romOffset       int

	ram       []uint8
	ramOffset int

	mapEnable bool
	mapSelect uint8

	latchClockReg uint8
	clockLatched  bool

	rtc *RTC
}

func NewMBC3(rom []uint8, romBanks uint16, ramBanks uint16) MBC {
	ram := make([]uint8, 0x0200) //512
	for i := range ram {
		ram[i] = 0xFF
	}

	return &MBC3{
		rom:             rom,
		romOffset:       0x4000,
		ram:             ram,
		selectedRomBank: 1,
		mapEnable:       false,
		mapSelect:       0,
		latchClockReg:   0xFF,
		clockLatched:    false,
		rtc:             NewRTC(),
	}
}

func (m *MBC3) WriteRom(address uint16, value uint8) {
	/*
		if address < 0x4000 {
			if !utils.TestBit16(address, 8) {
				m.ramWriteEnabled = (value & 0x0F) == 0x0A
			} else {
				if value&0x0F == 0x00 {
					m.selectedRomBank = 0x01
				} else {
					m.selectedRomBank = uint16(value) & 0x0F
				}
				m.romOffset = 0x4000 * int(m.selectedRomBank)
			}
		}
	*/

	switch {
	case address < 0x2000:
		m.mapEnable = (value & 0x0F) == 0x0A
	case address < 0x4000:
		if value == 0 {
			m.selectedRomBank = 1
		} else {
			m.selectedRomBank = uint16(value)
			m.romOffset = 0x4000 * int(m.selectedRomBank)
		}
	case address < 0x6000:
		m.mapSelect = value & 0xF
		/*
			if mbc30 {
			            self.ram_offset = RAM_BANK_SIZE * (state.map_select & 0b111) as usize;
			          } else {
			            self.ram_offset = RAM_BANK_SIZE * (state.map_select & 0b011) as usize;
			          }
		*/
		m.ramOffset = 0x2000 * int(m.mapSelect&0b011)
		//fmt.Printf("ramoffset %04x\n", m.ramOffset)
		//fmt.Printf("mapSelect %04x\n", m.mapSelect)
	case address < 0x8000:
		if value == 1 && m.latchClockReg == 0 {
			if m.clockLatched {
				m.rtc.Unlatch()
				m.clockLatched = false
			} else {
				m.rtc.Latch()
				m.clockLatched = true
			}
		}
		m.latchClockReg = value
	}

}

func (m *MBC3) WriteRam(address uint16, value uint8) {
	if m.mapEnable {
		switch m.mapSelect {
		case 0x0, 0x1, 0x2, 0x3:
			ramAddress := m.getRamAddress(address)
			if ramAddress < uint16(len(m.ram)) {
				m.ram[ramAddress] = value & 0x0F
			}
		case 0x8, 0x9, 0xA, 0xB, 0xC:
			m.setTimer(value)
		default:
			panic("OHNO")
		}
	}
}

func (m *MBC3) getRamAddress(address uint16) uint16 {
	addr := address - uint16(m.ramOffset) //address - 0xA000
	return addr & 0b0000_0001_1111_1111
}

func (m *MBC3) Read(address uint16) uint8 {
	switch {
	case address < 0x4000:
		return m.rom[(0x00|(address&0x3fff))&(uint16(len(m.rom))-1)]
	case address < 0x8000:
		return m.rom[(m.romOffset|int(address&0x3fff))&(len(m.rom)-1)]
	case 0xA000 <= address && address < 0xC000 && m.mapSelect < 4:
		if m.mapEnable {
			ramAddress := m.getRamAddress(address)
			if ramAddress < uint16(len(m.ram)) {
				return m.ram[ramAddress] | 0xF0
			} else {
				return 0xF0
			}
		} else {
			return 0xF0
		}
	case 0xA000 <= address && address < 0xC000:
		if m.mapEnable {
			// if mbc30 => self.read_ram(addr, default_value),
			return m.getTimer()
		} else {
			return 0xF0
		}
	default:
		return 0xFF
	}
}

func (m *MBC3) getTimer() uint8 {
	switch m.mapSelect {
	case 0x08:
		return uint8(m.rtc.Seconds())
	case 0x09:
		return uint8(m.rtc.Minutes())
	case 0x0A:
		return uint8(m.rtc.Hours())
	case 0x0B:
		return uint8(m.rtc.Days() & 0xFF)
	case 0x0C:
		result := (m.rtc.Days() & 0x100) >> 8
		if m.rtc.isHalt() {
			result |= 1 << 6
		}
		if m.rtc.isCounterOverflow() {
			result |= 1 << 7
		}
		return uint8(result)
	default:
		return 0xFF
	}
}

func (m *MBC3) setTimer(value uint8) {
	switch m.mapSelect {
	case 0x08:
		m.rtc.SetSeconds(int64(value))
	case 0x09:
		m.rtc.SetMinutes(int64(value))
	case 0x0A:
		m.rtc.SetHours(int64(value))
	case 0x0B:
		m.rtc.SetDays((m.rtc.Days() & 0x100) | (int64(value) & 0xFF))
	case 0x0C:
		m.rtc.SetDays((m.rtc.Days() & 0xff) | (int64(value)&1)<<8)
		//m.rtc.SetHalt((value & (1 << 6)) != 0)
		m.rtc.SetHalt(utils.TestBit(value, 6))
		if utils.TestBit(value, 7) {
			m.rtc.ClearCounterOverflow()
		}
		//if (value & (1 << 7)) == 0 {}
	}
}
