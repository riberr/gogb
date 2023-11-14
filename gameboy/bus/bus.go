package bus

import (
	"gogb/gameboy/interrupts"
	"gogb/gameboy/seriallink"
	"gogb/gameboy/timer"
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
	cart       *Cart
	interrupts *interrupts.Interrupts
	timer      *timer.Timer
	sl         *seriallink.SerialLink
}

var vram = NewMemory(0x8000, 0x9FFF) // Video RAM
var eram = NewMemory(0xA000, 0xBFFF) // External RAM
var wramC = NewMemory(0xC000, 0xCFFF)
var wramD = NewMemory(0xD000, 0xDFFF)
var echoRam = NewMemory(0xE000, 0xFDFF)
var oam = NewMemory(0xFE00, 0xFE9F) // Object attribute bus
var notUsable = NewMemory(0xFEA0, 0xFEFF)
var ioRegs = NewMemory(0xFF00, 0xFF7F) // I/O Registers
var hram = NewMemory(0xFF80, 0xFFFE)

//var ieReg uint8 = 0 //Interrupt Enable register

func New(interrupts *interrupts.Interrupts, timer *timer.Timer, sl *seriallink.SerialLink) *Bus {
	return &Bus{
		cart:       &Cart{},
		interrupts: interrupts,
		timer:      timer,
		sl:         sl,
	}
}

func (b *Bus) LoadCart(romPath string, romName string) bool {
	return b.cart.Load(romPath, romName)
}

func (b *Bus) Read(address uint16) uint8 {
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

	if eram.has(address) {
		return eram.read(address)
	}

	if wramC.has(address) {
		return wramC.read(address)
	}

	if wramD.has(address) {
		return wramD.read(address)
	}

	if echoRam.has(address) {
		return echoRam.read(address)
	}

	if oam.has(address) {
		return oam.read(address)
	}

	if notUsable.has(address) {
		return notUsable.read(address)
	}

	if hram.has(address) {
		return hram.read(address)
	}

	switch address {
	case 0xFF01:
		return b.sl.GetSB()
	case 0xFF02:
		return 0xFF // TODO REMOVE!!
		//return b.sl.GetSC()
	case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
		return b.timer.Read(address)
	case 0xFF0F:
		return b.interrupts.GetIF()
	case 0xFFFF:
		return b.interrupts.GetIE()
	default:
		if ioRegs.has(address) {
			return ioRegs.read(address)
		}
	}

	log.Panicf("READ NO IMPL (%02x)", address)
	return 0
}

func (b *Bus) Write(address uint16, value uint8) {
	if address < 0x8000 {
		//ROM Data
		b.cart.write(address, value)
		return
	}

	if vram.has(address) {
		vram.write(address, value)
		return
	}

	if eram.has(address) {
		eram.write(address, value)
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

	if echoRam.has(address) {
		echoRam.write(address, value)
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

	if hram.has(address) {
		hram.write(address, value)
		return
	}

	switch address {
	case 0xFF01:
		b.sl.SetSB(value)
		return
	case 0xFF02:
		b.sl.SetSC(value)
		return
	case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
		b.timer.Write(address, value)
		return
	case 0xFF0F:
		b.interrupts.SetAllIF(value)
		return
	case 0xFFFF:

		b.interrupts.SetAllIE(value)
		return
	default:
		if ioRegs.has(address) {
			ioRegs.write(address, value)
			return
		}
	}

	log.Panicf("WRITE NO IMPL (%02x)", address)
}
