package mbc

import (
	"gogb/gameboy/utils"
)

type MBC2 struct {
	rom             []uint8
	selectedRomBank uint16
	romOffset       int

	ram             []uint8
	ramWriteEnabled bool
}

func NewMBC2(rom []uint8) MBC {
	ram := make([]uint8, 0x0200) //512
	for i := range ram {
		ram[i] = 0xFF
	}

	return &MBC2{
		rom:             rom,
		romOffset:       0x4000,
		ram:             ram,
		selectedRomBank: 1,
		ramWriteEnabled: false,
	}
}

func (m *MBC2) WriteRom(address uint16, value uint8) {
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

}

func (m *MBC2) WriteRam(address uint16, value uint8) {
	if m.ramWriteEnabled {
		ramAddress := m.getRamAddress(address)
		if ramAddress < uint16(len(m.ram)) {
			m.ram[ramAddress] = value & 0x0F
		}
	}
}

func (m *MBC2) getRamAddress(address uint16) uint16 {
	addr := address - 0xA000
	return addr & 0b0000_0001_1111_1111
}

func (m *MBC2) Read(address uint16) uint8 {
	switch {
	case address < 0x4000:
		return m.rom[(0x00|(address&0x3fff))&(uint16(len(m.rom))-1)]
	case address < 0x8000:
		return m.rom[(m.romOffset|int(address&0x3fff))&(len(m.rom)-1)]
	case 0xA000 <= address && address < 0xC000:
		if m.ramWriteEnabled {
			ramAddress := m.getRamAddress(address)
			if ramAddress < uint16(len(m.ram)) {
				return m.ram[ramAddress] | 0xF0
			}
		}
		return 0xFF
	default:
		return 0xFF
	}
}
