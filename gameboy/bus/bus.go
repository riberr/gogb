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
	vram       Space
	eram       Space
	wramC      Space
	wramD      Space
	echoRam    Space
	oam        Space
	notUsable  Space
	ioRegs     Space
	hram       Space
}

//var ieReg uint8 = 0 //Interrupt Enable register

func New(interrupts *interrupts.Interrupts, timer *timer.Timer, sl *seriallink.SerialLink) *Bus {
	return &Bus{
		cart:       &Cart{},
		interrupts: interrupts,
		timer:      timer,
		sl:         sl,
		vram:       NewSpace(0x8000, 0x9FFF), // Video RAM
		eram:       NewSpace(0xA000, 0xBFFF), // External RAM
		wramC:      NewSpace(0xC000, 0xCFFF),
		wramD:      NewSpace(0xD000, 0xDFFF),
		echoRam:    NewSpace(0xE000, 0xFDFF),
		oam:        NewSpace(0xFE00, 0xFE9F), // Object attribute bus
		notUsable:  NewSpace(0xFEA0, 0xFEFF),
		ioRegs:     NewSpace(0xFF00, 0xFF7F), // I/O Registers
		hram:       NewSpace(0xFF80, 0xFFFE),
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

	if b.vram.has(address) {
		return b.vram.read(address)
	}

	if b.eram.has(address) {
		return b.eram.read(address)
	}

	if b.wramC.has(address) {
		return b.wramC.read(address)
	}

	if b.wramD.has(address) {
		return b.wramD.read(address)
	}

	if b.echoRam.has(address) {
		return b.echoRam.read(address)
	}

	if b.oam.has(address) {
		return b.oam.read(address)
	}

	if b.notUsable.has(address) {
		return b.notUsable.read(address)
	}

	if b.hram.has(address) {
		return b.hram.read(address)
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
		if b.ioRegs.has(address) {
			return b.ioRegs.read(address)
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

	if b.vram.has(address) {
		b.vram.write(address, value)
		return
	}

	if b.eram.has(address) {
		b.eram.write(address, value)
		return
	}

	if b.wramC.has(address) {
		b.wramC.write(address, value)
		return
	}

	if b.wramD.has(address) {
		b.wramD.write(address, value)
		return
	}

	if b.echoRam.has(address) {
		b.echoRam.write(address, value)
		return
	}

	if b.oam.has(address) {
		b.oam.write(address, value)
		return
	}

	if b.notUsable.has(address) {
		b.notUsable.write(address, value)
		return
	}

	if b.hram.has(address) {
		b.hram.write(address, value)
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
	case 0xFF46:
		b.doDMATransfer(value)
		return
	case 0xFFFF:

		b.interrupts.SetAllIE(value)
		return
	default:
		if b.ioRegs.has(address) {
			b.ioRegs.write(address, value)
			return
		}
	}

	log.Panicf("WRITE NO IMPL (%02x)", address)
}

func (b *Bus) doDMATransfer(value uint8) {
	address := uint16(value) << 8 // source address is data * 100
	for i := uint16(0); i < 0xA0; i++ {
		b.Write(uint16(0xFE00)+i, b.Read(address+i))
	}
}
