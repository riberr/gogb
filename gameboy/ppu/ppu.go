package ppu

import (
	"gogb/gameboy/interrupts"
	"gogb/gameboy/utils"
	"image"
)

type PPU struct {
	interrupts *interrupts.Interrupts

	Vram utils.Space
	Oam  utils.Space

	Fb     *image.RGBA // main frame buffer
	FbVram *image.RGBA // frame buffer of vram tiles

	lcdc            uint8 // 0xFF40
	ly              uint8 // 0xFF44
	lyc             uint8 // 0xFF45
	stat            uint8 // 0xFF41
	scy             uint8 // 0xFF42
	scx             uint8 //0xFF43
	wy              uint8 // 0xFF4A
	wx              uint8 // 0xFF4B
	scanlineCounter int
}

// STAT
const (
	LycIntSelect   = 6
	Mode2IntSelect = 5
	Mode1IntSelect = 4
	Mode0IntSelect = 3
	LycEqualsLy    = 2
	PpuModeHi      = 1
	PpuModeLo      = 0
)

// LCDC
const (
	LcdAndPpuEnable   = 7
	WindowTileMap     = 6
	WindowEnable      = 5
	BgAndWindowTiles  = 4
	BGTileMap         = 3
	ObjSize           = 2
	ObjEnable         = 1
	BgAndWindowEnable = 0
)

const (
	Mode2bounds = 456 - 80
	Mode3bounds = Mode2bounds - 172
)

func New(interrupts *interrupts.Interrupts) *PPU {
	return &PPU{
		interrupts: interrupts,
		Vram:       utils.NewSpace(0x8000, 0x9FFF), // Video RAM
		Oam:        utils.NewSpace(0xFE00, 0xFE9F), // Object attribute bus
		Fb:         image.NewRGBA(image.Rect(0, 0, vramWidth, vramHeight)),
		FbVram:     image.NewRGBA(image.Rect(0, 0, vramWidth, vramHeight)),
	}
}

func (ppu *PPU) GenerateDebugGraphics() {
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
			ppu.interrupts.SetIF(interrupts.VBLANK)
		} else if ppu.ly > 153 {
			ppu.ly = 0
		} else if ppu.ly < 144 {
			ppu.drawScanLine()
		}
	}
}

func (ppu *PPU) isLcdEnabled() bool {
	return utils.TestBit(ppu.lcdc, LcdAndPpuEnable)
}

func (ppu *PPU) setLcdStatus() {
	if !ppu.isLcdEnabled() {
		// set the mode to 1 during lcd disabled and reset scanline
		ppu.scanlineCounter = 456
		ppu.ly = 0
		ppu.stat = utils.ClearBit(ppu.stat, PpuModeHi)
		ppu.stat = utils.SetBit(ppu.stat, PpuModeLo) // set ppumode = 1
		return
	}

	currentMode := ppu.stat & 0x3
	mode := uint8(0)
	reqInt := false

	switch {
	case ppu.ly >= 144:
		// in vblank so set mode to 1
		mode = 1
		ppu.stat = utils.ClearBit(ppu.stat, PpuModeHi)
		ppu.stat = utils.SetBit(ppu.stat, PpuModeLo)
		reqInt = utils.TestBit(ppu.stat, Mode1IntSelect)
	case ppu.scanlineCounter >= Mode2bounds:
		mode = 2
		ppu.stat = utils.SetBit(ppu.stat, PpuModeHi)
		ppu.stat = utils.ClearBit(ppu.stat, PpuModeLo)
		reqInt = utils.TestBit(ppu.stat, Mode2IntSelect)
	case ppu.scanlineCounter >= Mode3bounds:
		mode = 3
		ppu.stat = utils.SetBit(ppu.stat, PpuModeHi)
		ppu.stat = utils.SetBit(ppu.stat, PpuModeLo)
	default:
		mode = 0
		ppu.stat = utils.ClearBit(ppu.stat, PpuModeHi)
		ppu.stat = utils.ClearBit(ppu.stat, PpuModeLo)
		reqInt = utils.TestBit(ppu.stat, Mode0IntSelect)
	}

	// just entered a new mode so request interupt
	if reqInt && (mode != currentMode) {
		ppu.interrupts.SetIF(interrupts.LCD)
	}

	// check the coincidence flag
	if ppu.ly == ppu.lyc {
		ppu.stat = utils.SetBit(ppu.stat, LycEqualsLy)
		if utils.TestBit(ppu.stat, LycIntSelect) {
			ppu.interrupts.SetIF(interrupts.LCD)
		}
	} else {
		ppu.stat = utils.ClearBit(ppu.stat, LycEqualsLy)
	}
}

func (ppu *PPU) drawScanLine() {
	if utils.TestBit(ppu.lcdc, BgAndWindowEnable) {
		ppu.renderTiles()
	}
	if utils.TestBit(ppu.lcdc, ObjEnable) {
		//ppu.renderSprites()
	}
}

func (ppu *PPU) renderTiles() {
	tileData := uint16(0)
	backgroundMemory := uint16(0)
	unsigned := true
	usingWindow := false

	// is the window enabled?
	if utils.TestBit(ppu.lcdc, WindowEnable) {
		// is the current scanline we're drawing within the windows Y pos?,
		if ppu.wy <= ppu.ly {
			usingWindow = true
		}
	}

	// which tile data are we using?
	if utils.TestBit(ppu.lcdc, BgAndWindowTiles) {
		tileData = 0x8000
	} else {
		// IMPORTANT: This memory region uses signed bytes as tile identifiers
		tileData = 0x8800
		unsigned = false
	}

	// which background mem?
	if !usingWindow {
		if utils.TestBit(ppu.lcdc, BGTileMap) {
			backgroundMemory = 0x9C00
		} else {
			backgroundMemory = 0x9800
		}
	} else {
		// which window memory?
		if utils.TestBit(ppu.lcdc, WindowTileMap) {
			backgroundMemory = 0x9C00
		} else {
			backgroundMemory = 0x9800
		}
	}

	yPos := uint8(0)

	// yPos is used to calculate which of 32 vertical tiles the current scanline is drawing
	if !usingWindow {
		yPos = ppu.scy + ppu.ly
	} else {
		yPos = ppu.ly - ppu.wy
	}

	// which of the 8 vertical pixels of the current tile is the scanline on?
	tileRow := uint16(yPos/8) * 32

	// time to start drawing the 160 horizontal pixels for this scanline
	for pixel := uint8(0); pixel < 160; pixel++ {
		xPos := pixel + ppu.scx

		// translate the current x pos to window space if necessary
		if usingWindow {
			if pixel >= (ppu.wx - 7) {
				xPos = pixel - (ppu.wx - 7)
			}
		}

		// which of the 32 horizontal tiles does this xPos fall within?
		tileCol := uint16(xPos / 8)

		// deduce where this tile identifier is in memory.
		tileLocation := tileData

		// get the tile identity number. Remember it can be signed or unsigned
		tileAddress := backgroundMemory + tileRow + tileCol
		if unsigned {
			tileNum := int16(ppu.Vram.Read(tileAddress))
			tileLocation = tileLocation + uint16(tileNum*16)
		} else {
			tileNum := int16(int8(ppu.Vram.Read(tileAddress)))
			tileLocation = uint16(int32(tileLocation) + int32((tileNum+128)*16))
		}

		// find the correct vertical line we're on of the tile to get the tile data from in memory
		line := yPos % 8
		line *= 2 // each vertical line takes up two bytes of memory
		b1 := ppu.Vram.Read(tileLocation + uint16(line))
		b2 := ppu.Vram.Read(tileLocation + uint16(line) + 1)

		/*
			// pixel 0 in the tile is it 7 of data 1 and data2. Pixel 1 is bit 6 etc..
			colorBit := int(xPos) % 8
			colorBit -= 7
			colorBit *= -1

			// combine data 2 and data 1 to get the colour id for this pixel in the tile
			colorNum := utils.ToInt(utils.TestBit(data2, colorBit))
			colorNum <<= 1
			colorNum |= utils.ToInt(utils.TestBit(data1, colorBit))

			// now we have the colour id get the actual colour from palette 0xFF47
		*/

		for bit := 7; bit >= 0; bit-- {
			hi := uint8(b1 & (1 << bit))
			if hi > 0 {
				hi = 2
			} else {
				hi = 0
			}

			lo := uint8(b2 & (1 << bit))
			if lo > 0 {
				lo = 1
			} else {
				lo = 0
			}

			c := hi | lo // color

			xx := int(xPos) + ((7 - bit) * scale)
			//yy := int(yPos) + (int(tileY) / 2 * scale)

			drawSquare(ppu.Fb, &coloredRects[c], 4, xx, int(ppu.ly))
			//drawSquare2(img /*colors[c]*/, colors[c], 4, xx, yy)

		}
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
		ppu.lcdc = value
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
