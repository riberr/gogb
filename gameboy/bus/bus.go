package bus

import (
	"gogb/gameboy/interrupts"
	"gogb/gameboy/joypad"
	"gogb/gameboy/ppu"
	"gogb/gameboy/seriallink"
	"gogb/gameboy/timer"
	"gogb/gameboy/utils"
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
	timer2     *timer.Timer2
	sl         *seriallink.SerialLink
	ppu        *ppu.PPU
	joypad     *joypad.JoyPad
	//vram       Space
	eram    utils.Space
	wramC   utils.Space
	wramD   utils.Space
	echoRam utils.Space
	//oam        Space
	notUsable utils.Space
	ioRegs    utils.Space
	hram      utils.Space

	currentRomBank uint16
	currentRamBank uint16
	enableRam      bool
	romBanking     bool
}

//var ieReg uint8 = 0 //Interrupt Enable register

func New(interrupts *interrupts.Interrupts, timer *timer.Timer, timer2 *timer.Timer2, sl *seriallink.SerialLink, ppu *ppu.PPU, joypad *joypad.JoyPad) *Bus {
	return &Bus{
		cart:       &Cart{},
		interrupts: interrupts,
		timer:      timer,
		timer2:     timer2,
		sl:         sl,
		ppu:        ppu,
		joypad:     joypad,
		//vram:       NewSpace(0x8000, 0x9FFF), // Video RAM
		eram:    utils.NewSpace(0xA000, 0xBFFF), // External RAM
		wramC:   utils.NewSpace(0xC000, 0xCFFF),
		wramD:   utils.NewSpace(0xD000, 0xDFFF),
		echoRam: utils.NewSpace(0xE000, 0xFDFF),
		//oam:        NewSpace(0xFE00, 0xFE9F), // Object attribute bus
		notUsable: utils.NewSpace(0xFEA0, 0xFEFF),
		ioRegs:    utils.NewSpace(0xFF00, 0xFF7F), // I/O Registers
		hram:      utils.NewSpace(0xFF80, 0xFFFE),

		currentRomBank: 1,
		currentRamBank: 0,
		enableRam:      false,
		romBanking:     false,
	}
}

func (b *Bus) LoadCart(romPath string, romName string) bool {
	return b.cart.Load(romPath, romName)
}

func (b *Bus) Read(address uint16) uint8 {

	if address < 0x4000 {
		//ROM Data
		return b.cart.read(address)
	}

	// are we reading from the rom memory bank?
	if address < 0x8000 {
		newAddress := address - 0x4000
		return b.cart.read(newAddress + (b.currentRomBank * 0x4000))
	}

	if b.ppu.Vram.Has(address) {
		//println("reading from vram")
		return b.ppu.Vram.Read(address)
	}

	if b.eram.Has(address) {
		if b.enableRam {
			bank := b.currentRamBank % b.cart.header.ramBanks
			return b.eram.Read(address + bank)
		} else {
			return 0xFF
		}
	}

	if b.wramC.Has(address) {
		return b.wramC.Read(address)
	}

	if b.wramD.Has(address) {
		return b.wramD.Read(address)
	}

	if b.echoRam.Has(address) {
		return b.echoRam.Read(address)
	}

	if b.ppu.Oam.Has(address) {
		//println("reading from vram")
		return b.ppu.Oam.Read(address)
	}

	if b.notUsable.Has(address) {
		return b.notUsable.Read(address)
	}

	if b.hram.Has(address) {
		return b.hram.Read(address)
	}

	switch address {
	case 0xFF00:
		/*
			if b.joypad.GetJoyPadState() != 0xFF {
				fmt.Printf("joypad: %08b\n", b.joypad.GetJoyPadState())
			}
		*/
		return b.joypad.GetJoyPadState()
	case 0xFF01:
		return b.sl.GetSB()
	case 0xFF02:
		return b.sl.GetSC()
	case 0xFF04, 0xFF05, 0xFF06, 0xFF07:
		return b.timer.Read(address)
	case 0xFF0F:
		return b.interrupts.GetIF()
	case 0xFF40, 0xFF41, 0xFF44, 0xFF45, 0xFF47, 0xFF48, 0xFF49:
		return b.ppu.Read(address)
	case 0xFFFF:
		return b.interrupts.GetIE()
	default:
		if b.ioRegs.Has(address) {
			return b.ioRegs.Read(address)
		}
	}

	log.Panicf("READ NO IMPL (%02x)", address)
	return 0
}

func (b *Bus) Write(address uint16, value uint8) {
	if address < 0x8000 {
		//ROM Data
		//b.cart.write(address, value)
		b.handleBanking(address, value)
		return
	}

	if b.ppu.Vram.Has(address) {
		b.ppu.Vram.Write(address, value)
		return
	}

	if b.eram.Has(address) {
		b.eram.Write(address, value)
		return
	}

	if b.wramC.Has(address) {
		b.wramC.Write(address, value)
		return
	}

	if b.wramD.Has(address) {
		b.wramD.Write(address, value)
		return
	}

	if b.echoRam.Has(address) {
		b.echoRam.Write(address, value)
		return
	}

	if b.ppu.Oam.Has(address) {
		b.ppu.Oam.Write(address, value)
		return
	}

	if b.notUsable.Has(address) {
		b.notUsable.Write(address, value)
		return
	}

	if b.hram.Has(address) {
		b.hram.Write(address, value)
		return
	}

	switch address {
	case 0xFF00:
		// do nothing?
		b.joypad.Write(address, value)
		return
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
	case 0xFF40, 0xFF41, 0xFF44, 0xFF45, 0xFF47, 0xFF48, 0xFF49:
		b.ppu.Write(address, value)
		return
	case 0xFF46:
		b.doDMATransfer(value)
		return
	case 0xFFFF:
		b.interrupts.SetAllIE(value)
		return
	default:
		if b.ioRegs.Has(address) {
			b.ioRegs.Write(address, value)
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
