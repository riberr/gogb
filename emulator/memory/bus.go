package memory

import (
	"fmt"
	"log"
)

// MBC1
// The majority of games for the original Game Boy use the MBC1 chip
// 0x0000 - 0x3FFF : ROM Bank 0
// 0x4000 - 0x7FFF : ROM Bank 1 - Switchable
// 0x8000 - 0x97FF : CHR RAM
// 0x9800 - 0x9BFF : BG Map 1
// 0x9C00 - 0x9FFF : BG Map 2
// 0xA000 - 0xBFFF : Cartridge RAM
// 0xC000 - 0xCFFF : RAM Bank 0
// 0xD000 - 0xDFFF : RAM Bank 1-7 - switchable - Color only
// 0xE000 - 0xFDFF : Reserved - Echo RAM
// 0xFE00 - 0xFE9F : Object Attribute Memory
// 0xFEA0 - 0xFEFF : Reserved - Unusable
// 0xFF00 - 0xFF7F : I/O Registers
// 0xFF80 - 0xFFFE : Zero Page

var wram = NewMemory(0xC000, 0xCFFF)

func BusRead(address uint16) uint8 {
	if address < 0x8000 {
		//ROM Data
		return cartRead(address)
	}

	if wram.from <= address && address <= wram.to {
		return wram.read(address)
	}

	if wram.has(address) {
		return wram.read(address)
	}

	log.Panicf("READ NO IMPL (%02x)", address)
	return 0
}

func BusWrite(address uint16, value uint8) {
	// for testing with blargg's cpu_instrs roms
	if address == 0xFF02 && value == 0x81 {
		fmt.Printf("!!! %v\n", BusRead(0xFF01))
	}

	if address < 0x8000 {
		//ROM Data
		cartWrite(address, value)
		return
	}

	if wram.has(address) {
		wram.write(address, value)
		return
	}

	log.Panicf("WRITE NO IMPL (%02x)", address)
}
