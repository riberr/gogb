package bus

import "gogb/gameboy/utils"

func (b *Bus) handleBanking(address uint16, value uint8) {

	if address < 0x2000 {
		// do RAM enabling
		// todo prettify
		if b.cart.header.cartType == 1 || b.cart.header.cartType == 2 || b.cart.header.cartType == 3 ||
			b.cart.header.cartType == 5 || b.cart.header.cartType == 6 {
			b.doRamBankEnable(address, value)
		}
	} else if address < 0x4000 {
		// do ROM bank change
		if b.cart.header.cartType == 1 || b.cart.header.cartType == 2 || b.cart.header.cartType == 3 ||
			b.cart.header.cartType == 5 || b.cart.header.cartType == 6 {
			b.doChangeLoROMBank(value)
		}
	} else if address < 0x6000 {
		// do ROM or RAM bank change
		// there is no rambank in mbc2 so always use rambank 0
		if b.cart.header.cartType == 1 || b.cart.header.cartType == 2 || b.cart.header.cartType == 3 {
			if b.romBanking {
				b.doChangeHiRomBank(value)
			} else {
				b.doRamBankChange(value)
			}
		}
	} else if address < 0x8000 {
		// this will change whether we are doing ROM banking
		// or RAM banking with the above if statement
		if b.cart.header.cartType == 1 || b.cart.header.cartType == 2 || b.cart.header.cartType == 3 {
			b.doChangeRomRamMode(value)
		}
	}
}

func (b *Bus) doRamBankEnable(address uint16, value uint8) {
	if b.cart.header.cartType == 5 || b.cart.header.cartType == 6 {
		if utils.TestBit16(address, 4) {
			return
		}
	}
	testValue := value & 0xF
	if testValue == 0xA {
		b.enableRam = true
	} else if testValue == 0x0 {
		b.enableRam = false
	}
}

func (b *Bus) doChangeLoROMBank(value uint8) {
	if b.cart.header.cartType == 5 || b.cart.header.cartType == 6 {
		b.currentRomBank = uint16(value) & 0xF
		if b.currentRomBank == 0 {
			b.currentRomBank++
		}
		return
	}

	lower5 := uint16(value) & 31
	b.currentRomBank &= 224 // turn off the lower 5
	b.currentRomBank |= lower5
	if b.currentRomBank == 0 {
		b.currentRomBank++
	}
}

func (b *Bus) doChangeHiRomBank(value uint8) {
	// turn off the upper 3 bits of the current rom
	b.currentRomBank &= 31

	// turn off the lower 5 bits of the data
	value = value & 224
	b.currentRomBank |= uint16(value)
	if b.currentRomBank == 0 {
		b.currentRomBank++
	}
}

func (b *Bus) doRamBankChange(value uint8) {
	b.currentRamBank = uint16(value) & 0x3
}

func (b *Bus) doChangeRomRamMode(value uint8) {
	newValue := value & 0x1
	if newValue == 0 {
		b.romBanking = true
		b.currentRamBank = 0
	} else {
		b.romBanking = false
	}

}
