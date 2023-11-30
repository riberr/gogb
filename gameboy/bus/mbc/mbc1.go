package mbc

type MBC1 struct {
	romBanks        uint16
	rom             []uint8
	selectedRomBank uint16

	ramBanks        uint16
	ram             []uint8
	selectedRamBank uint16
	ramWriteEnabled bool

	memoryModel            uint8
	multicart              bool
	cachedRomBankFor0x0000 int
	cachedRomBankFor0x4000 int
}

func NewMBC1(rom []uint8, romBanks uint16, ramBanks uint16) MBC {
	return &MBC1{
		rom:                    rom,
		romBanks:               romBanks,
		ramBanks:               ramBanks,
		selectedRomBank:        1,
		cachedRomBankFor0x0000: -1,
		cachedRomBankFor0x4000: -1,
		multicart:              false,                          //todo not impl
		ram:                    make([]uint8, 0x2000*ramBanks), //make([]uint8, 0x8000), //32KB
	}
}

// WriteRom attempts to switch the ROM or RAM bank.
func (m *MBC1) WriteRom(address uint16, value uint8) {
	switch {
	case address < 0x2000:
		m.ramWriteEnabled = (value & 0b1111) == 0b1010
		if !m.ramWriteEnabled {
			//battery.saveRam(ram)
		}
	case address < 0x4000:
		bank := m.selectedRomBank & 0b01100000
		bank = bank | (uint16(value) & 0b00011111)
		m.selectedRomBank = bank
		m.cachedRomBankFor0x0000 = -1
		m.cachedRomBankFor0x4000 = -1
	case address < 0x6000 && m.memoryModel == 0:
		bank := m.selectedRomBank & 0b00011111
		bank = bank | ((uint16(value) & 0b11) << 5)
		m.selectedRomBank = bank
		m.cachedRomBankFor0x0000 = -1
		m.cachedRomBankFor0x4000 = -1
	case address < 0x6000 && m.memoryModel == 1:
		m.selectedRamBank = uint16(value) & 0b11
		m.cachedRomBankFor0x0000 = -1
		m.cachedRomBankFor0x4000 = -1
	case address < 0x8000:
		m.memoryModel = value & 1
		m.cachedRomBankFor0x0000 = -1
		m.cachedRomBankFor0x4000 = -1
	}
}

// WriteRam writes data to the ram if it is enabled.
func (m *MBC1) WriteRam(address uint16, value uint8) {
	if m.ramWriteEnabled {
		ramAddress := m.getRamAddress(address)
		if ramAddress < uint16(len(m.ram)) {
			m.ram[ramAddress] = value
		}
	}
}

func (m *MBC1) getRamAddress(address uint16) uint16 {
	if m.memoryModel == 0 {
		return address - 0xA000
	} else {
		return (m.selectedRamBank%m.ramBanks)*0x2000 + (address - 0xA000)
	}
}

func (m *MBC1) Read(address uint16) uint8 {
	switch {
	case address < 0x4000:
		return m.getRomByte(m.getRomBankFor0x0000(), address)
	case address < 0x8000:
		return m.getRomByte(m.getRomBankFor0x4000(), address-0x4000)
	case 0xA000 <= address && address < 0xC000:
		if m.ramWriteEnabled {
			ramAddress := m.getRamAddress(address)
			if ramAddress < uint16(len(m.ram)) {
				return m.ram[ramAddress]
			} else {
				return 0xFF
			}
		} else {
			return 0xFF
		}
	default:
		panic("mbc error?")

	}
}

func (m *MBC1) getRomBankFor0x0000() int {
	if m.cachedRomBankFor0x0000 == -1 {
		if m.memoryModel == 0 {
			m.cachedRomBankFor0x0000 = 0
		} else {
			bank := m.selectedRamBank << 5
			if m.multicart {
				bank >>= 1
			}
			bank %= m.romBanks
			m.cachedRomBankFor0x0000 = int(bank)
		}
	}
	return m.cachedRomBankFor0x0000
}

func (m *MBC1) getRomBankFor0x4000() int {
	if m.cachedRomBankFor0x4000 == -1 {
		bank := m.selectedRomBank
		if bank%0x20 == 0 {
			bank++
		}
		if m.memoryModel == 1 {
			bank &= 0b00011111
			bank |= m.selectedRamBank << 5
		}
		if m.multicart {
			bank = ((bank >> 1) & 0x30) | (bank & 0x0f)
		}
		bank %= m.romBanks
		m.cachedRomBankFor0x4000 = int(bank)
	}
	return m.cachedRomBankFor0x4000
}

func (m *MBC1) getRomByte(bank int, address uint16) uint8 {
	cartOffset := bank*0x4000 + int(address)
	if cartOffset < len(m.rom) {
		return m.rom[cartOffset]
	} else {
		return 0xff
	}
}
