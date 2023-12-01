package mbc

import "gogb/gameboy/utils"

type MBC2 struct {
	//romBanks        uint16
	rom             []uint8
	selectedRomBank uint16
	romOffset       uint16

	//ramBanks        uint16
	ram []uint8
	//selectedRamBank uint16
	ramWriteEnabled bool
}

func NewMBC2(rom []uint8, romBanks uint16, ramBanks uint16) MBC {
	ram := make([]uint8, 0x0200) //512
	for i := range ram {
		ram[i] = 0xFF
	}

	return &MBC2{
		rom:       rom,
		romOffset: 0x4000,
		//romBanks:        romBanks,
		ram: ram,
		//ramBanks:        ramBanks,
		selectedRomBank: 1,
		ramWriteEnabled: false,
	}
}

func (m *MBC2) WriteRom(address uint16, value uint8) {

	/*
		if address < 0x2000 {
			if !utils.TestBit16(address, 8) {
				if (value & 0x0F) == 0x0A {
					println("ramwrite enabled")
					m.ramWriteEnabled = true
				} else {
					println("ramwrite disabled")
					m.ramWriteEnabled = false
				}

				//if (!ramWriteEnabled) {
				//	battery.saveRam(ram);
				//}

			}

		} else if address < 0x4000 {
			if utils.TestBit16(address, 8) {
				bank := value & 0x0F
				if bank == 0 {
					bank = 1
				}
				m.selectedRomBank = uint16(bank)
				fmt.Printf("selected rom bank: %v\n", m.selectedRomBank)
			}
		}
	*/

	if address < 0x4000 {
		if !utils.TestBit16(address, 8) {
			m.ramWriteEnabled = (value & 0x0F) == 0x0A
		} else {
			if value&0x0F == 0x00 {
				m.selectedRomBank = 0x01
			} else {
				m.selectedRomBank = uint16(value) & 0x0F
			}
			m.romOffset = 0x4000 * m.selectedRomBank
		}
		/*
			if utils.TestBit16(address, 8) {
				if value&0x0f == 0x0 {
					m.selectedRomBank = 1
				} else {
					m.selectedRomBank = uint16(value) & 0xF
				}
				fmt.Printf("selected rom bank: %v\n", m.selectedRomBank)
			} else {
				//m.ramWriteEnabled = (value & 0x0A) != 0

				if value == 0x0A {
					println("ramwrite enabled")
					m.ramWriteEnabled = true
				} else {
					println("ramwrite disabled")
					m.ramWriteEnabled = false
				}
			}
		*/
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
		return m.getRomByte(0, address)
	case address < 0x8000:
		//todo skillnad här vilket test som passar
		//return m.rom[(m.romOffset|(address&0x3fff))&(uint16(len(m.rom))-1)]
		return m.getRomByte(int(m.selectedRomBank), address-0x4000)
	case 0xA000 <= address && address < 0xC000:
		//ramAddress := m.getRamAddress(address)
		//return (m.ram[ramAddress] & 0x0F) | 0xF0
		ramAddress := m.getRamAddress(address)
		if ramAddress < uint16(len(m.ram)) {
			return m.ram[ramAddress] | 0xF0
		} else {
			return 0xF0
		}
	default:
		return 0xFF
	}
}

func (m *MBC2) getRomByte(bank int, address uint16) uint8 {
	cartOffset := bank*0x4000 + int(address)
	if cartOffset < len(m.rom) {
		return m.rom[cartOffset]
	} else {
		return 0x00
	}
}
