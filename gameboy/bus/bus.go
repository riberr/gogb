package bus

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
// 0xFE00 - 0xFE9F : Object Attribute Space
// 0xFEA0 - 0xFEFF : Reserved - Unusable
// 0xFF00 - 0xFF7F : I/O Registers
// 0xFF80 - 0xFFFE : Zero Page

type Bus struct {
	cart *Cart
}

var vram = NewMemory(0x8000, 0x9FFF) // Video RAM
var wramC = NewMemory(0xC000, 0xCFFF)
var wramD = NewMemory(0xD000, 0xDFFF)
var oam = NewMemory(0xFE00, 0xFE9F) // Object attribute bus
var notUsable = NewMemory(0xFEA0, 0xFEFF)
var ioRegs = NewMemory(0xFF00, 0xFF7F) // I/O Registers
var hram = NewMemory(0xFF80, 0xFFFE)
var ieReg uint8 = 0 //Interrupt Enable register

func New() *Bus {
	return &Bus{
		cart: &Cart{},
	}
}

func (b *Bus) LoadCart(romPath string, romName string) bool {
	return b.cart.Load(romPath, romName)
}

func (b *Bus) BusRead(address uint16) uint8 {
	if address < 0x8000 {
		//ROM Data
		return b.cart.read(address)
	}

	// todo remove when implementing PPU
	if address == 0xFF44 {
		return 0x90
	}

	if vram.has(address) {
		return vram.read(address)
	}

	if wramC.has(address) {
		return wramC.read(address)
	}

	if wramD.has(address) {
		return wramD.read(address)
	}

	if oam.has(address) {
		return oam.read(address)
	}

	if notUsable.has(address) {
		return notUsable.read(address)
	}

	if ioRegs.has(address) {
		return ioRegs.read(address)
	}

	if hram.has(address) {
		return hram.read(address)
	}

	if address == 0xFFFF {
		return ieReg
	}

	log.Panicf("READ NO IMPL (%02x)", address)
	return 0
}

func (b *Bus) BusWrite(address uint16, value uint8) {
	// for testing with blargg's cpu_instrs roms
	if address == 0xFF02 && value == 0x81 {
		fmt.Printf("%02x ", b.BusRead(0xFF01))
	}

	if address < 0x8000 {
		//ROM Data
		println("warning: writing to ROM")
		b.cart.write(address, value)
		return
	}

	if vram.has(address) {
		vram.write(address, value)
		return
	}

	if wramC.has(address) {
		wramC.write(address, value)
		return
	}

	if wramD.has(address) {
		wramD.write(address, value)
		return
	}

	if oam.has(address) {
		oam.write(address, value)
		return
	}

	if notUsable.has(address) {
		notUsable.write(address, value)
		return
	}

	if ioRegs.has(address) {
		ioRegs.write(address, value)
		return
	}

	if hram.has(address) {
		hram.write(address, value)
		return
	}

	if address == 0xFFFF {
		ieReg = value
		return
	}

	log.Panicf("WRITE NO IMPL (%02x)", address)
}
