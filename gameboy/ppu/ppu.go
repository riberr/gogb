package ppu

import (
	"gogb/gameboy/utils"
	"image"
)

type PPU struct {
	Vram utils.Space
	Oam  utils.Space

	Fb     *image.RGBA // main frame buffer
	FbVram *image.RGBA // frame buffer of vram tiles

	lcdc uint8 // 0xFF40
	ly   uint8 // 0xFF44
	lyc  uint8 // 0xFF45
	stat uint8 // 0xFF41

	scanlineCounter int
}

type StatBitValue int

const (
	LycIntSelect   StatBitValue = 0b0100_0000
	Mode2IntSelect              = 0b0010_0000
	Mode1IntSelect              = 0b0001_0000
	Mode0IntSelect              = 0b0000_1000
	LycEqualsLy                 = 0b0000_0100
	PpuMode                     = 0b0000_0011
)

const (
	Mode2bounds = 456 - 80
	Mode3bounds = Mode2bounds - 172
)

func New() *PPU {
	return &PPU{
		Vram:   utils.NewSpace(0x8000, 0x9FFF), // Video RAM
		Oam:    utils.NewSpace(0xFE00, 0xFE9F), // Object attribute bus
		FbVram: image.NewRGBA(image.Rect(0, 0, vramWidth, vramHeight)),
	}
}

func (ppu *PPU) GenerateGraphics() {
	ppu.GenerateDebugVram()
}

// Update from http://www.codeslinger.co.uk/pages/projects/gameboy/lcd.html
func (ppu *PPU) Update(cycles int) {
	ppu.setLcdStatus()

	if ppu.isLcdEnabled() {
		ppu.scanlineCounter -= cycles
	} else {
		return
	}

	if ppu.scanlineCounter <= 0 {
		// time to move onto next scanline
		ppu.ly++
		ppu.scanlineCounter = 456

		// we have entered vertical blank period
		if ppu.ly == 144 {
			//requestInterrupt(0)

		} else if ppu.ly > 153 {
			ppu.ly = 0
		} else if ppu.ly < 144 {
			//ppu.drawScanLine()
		}
	}
}

func (ppu *PPU) isLcdEnabled() bool {
	return utils.TestBit(ppu.lcdc, 7)
}

func (ppu *PPU) setLcdStatus() {
	if !ppu.isLcdEnabled() {
		// set the mode to 1 during lcd disabled and reset scanline
		ppu.scanlineCounter = 456
		ppu.ly = 0
		// todo: make stat write/read nicer
		ppu.stat &= 0b1111_1100              // clear PpuMode
		ppu.stat = utils.SetBit(ppu.stat, 0) // set ppumode = 1
		return
	}

	currentMode := ppu.stat & 0x3
	mode := uint8(0)
	reqInt := false

	switch {
	case ppu.ly >= 144:
		// in vblank so set mode to 1
		mode = 1
		// todo: make stat write/read nicer
		ppu.stat = utils.ClearBit(ppu.stat, 1)
		ppu.stat = utils.SetBit(ppu.stat, 0)
		reqInt = utils.TestBit(ppu.stat, 4)
	case ppu.scanlineCounter >= Mode2bounds:
		mode = 2
		ppu.stat = utils.SetBit(ppu.stat, 1)
		ppu.stat = utils.ClearBit(ppu.stat, 0)
		reqInt = utils.TestBit(ppu.stat, 5)
	case ppu.scanlineCounter >= Mode3bounds:
		mode = 3
		ppu.stat = utils.SetBit(ppu.stat, 1)
		ppu.stat = utils.SetBit(ppu.stat, 0)
	default:
		mode = 0
		ppu.stat = utils.ClearBit(ppu.stat, 1)
		ppu.stat = utils.ClearBit(ppu.stat, 0)
		reqInt = utils.TestBit(ppu.stat, 3)
	}

	// just entered a new mode so request interupt
	if reqInt && (mode != currentMode) {
		//RequestInterupt(1)
	}

	// check the conincidence flag
	if ppu.ly == ppu.lyc {
		ppu.stat = utils.SetBit(ppu.stat, 2)
		if utils.TestBit(ppu.stat, 6) {
			// RequestInterupt(1)
		}
	} else {
		ppu.stat = utils.ClearBit(ppu.stat, 2)
	}
}

func (ppu *PPU) Read(address uint16) uint8 {
	switch address {
	case 0xFF40:
		return ppu.lcdc
	case 0xFF41:
		return ppu.stat
	case 0xFF44:
		return ppu.ly
	case 0xFF45:
		return ppu.lyc
	default:
		panic("not handled ppu read")
	}
}

func (ppu *PPU) Write(address uint16, value uint8) {
	switch address {
	case 0xFF40:
		println("trying to write to 0xFF40")
	case 0xFF41:
		ppu.stat = value | 0x80
	case 0xFF44:
		ppu.ly = 0
	case 0xFF45:
		ppu.lyc = value
	default:
		panic("not handled ppu write")
	}
}
