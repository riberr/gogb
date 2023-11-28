package ppu

import (
	"gogb/gameboy/interrupts"
	"gogb/gameboy/utils"
	"image"
)

type PPU struct {
	interrupts *interrupts.Interrupts2

	Vram utils.Space
	Oam  utils.Space

	Fb     *image.RGBA // main frame buffer
	FbVram *image.RGBA // frame buffer of vram tiles

	scanlineCounter int
	tileScanline    [160]uint8

	// registers
	lcdc uint8 // 0xFF40
	ly   uint8 // 0xFF44
	lyc  uint8 // 0xFF45
	stat uint8 // 0xFF41
	scy  uint8 // 0xFF42
	scx  uint8 //0xFF43
	wy   uint8 // 0xFF4A
	wx   uint8 // 0xFF4B
	bgp  uint8 // 0xFF47 BG palette data
	obp0 uint8 // 0xFF48 OBJ palette data 0
	obp1 uint8 // 0xFF49 OBJ palette data 1
}

const (
	ScreenWidth  = 160 * Scale
	ScreenHeight = 144 * Scale

	spritePriorityOffset = uint8(100)
)

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

func New(interrupts *interrupts.Interrupts2) *PPU {
	return &PPU{
		interrupts: interrupts,
		Vram:       utils.NewSpace(0x8000, 0x9FFF), // Video RAM
		Oam:        utils.NewSpace(0xFE00, 0xFE9F), // Object attribute bus
		Fb:         image.NewRGBA(image.Rect(0, 0, ScreenWidth, ScreenHeight)),
		FbVram:     image.NewRGBA(image.Rect(0, 0, vramWidth, vramHeight)),
	}
}

func (ppu *PPU) GenerateDebugGraphics() {
	ppu.GenerateDebugVram()
}

// Update from http://www.codeslinger.co.uk/pages/projects/gameboy/lcd.html
func (ppu *PPU) Update(cycles int) {
	/*
		ppu.setLcdStatus()

		if ppu.isLcdEnabled() {
			ppu.scanlineCounter -= cycles
		} else {
			return
		}
	*/
	if !ppu.isLcdEnabled() {
		return
	}
	ppu.scanlineCounter -= cycles
	ppu.setLcdStatus()
	/**/
	if ppu.scanlineCounter <= 0 {
		// time to move onto next scanline
		ppu.ly++
		ppu.scanlineCounter = 456

		// we have entered vertical blank period
		if ppu.ly == 144 {
			ppu.interrupts.SetInterruptFlag(interrupts.INTR_VBLANK)
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

	// just entered a new mode so request interrupt
	if reqInt && (mode != currentMode) {
		ppu.interrupts.SetInterruptFlag(interrupts.INTR_LCD)
	}

	// check the coincidence flag
	if ppu.ly == ppu.lyc {
		ppu.stat = utils.SetBit(ppu.stat, LycEqualsLy)
		if utils.TestBit(ppu.stat, LycIntSelect) {
			ppu.interrupts.SetInterruptFlag(interrupts.INTR_LCD)
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
		ppu.renderSprites()
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
	ppu.tileScanline = [160]uint8{}
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
			//gameboy color
			if l.Mb.Cgb && internal.IsBitSet(tileAttr, 5) {
				// horizontal flip
				xPos = (7 - (pixel+(scrollX&0b111))%8)
			}
			if l.Mb.Cgb && !internal.IsBitSet(lcdControl, LCDC_BGEN) {
				priority = false
			}
		*/

		colorBit := uint8(int8((xPos%8)-7) * -1)
		colorNum := (utils.BitValue(b1, colorBit) << 1) | utils.BitValue(b2, colorBit)
		drawSquare(ppu.Fb, &coloredRects[colorNum], 4, int(pixel)*Scale, int(ppu.ly)*Scale)
		ppu.tileScanline[pixel] = colorNum
	}
}

func (ppu *PPU) renderSprites() {
	var ySize uint8 = 8
	if utils.TestBit(ppu.lcdc, 2) {
		ySize = 16
	}

	palette0 := ppu.obp0
	palette1 := ppu.obp1
	var minx [ScreenWidth]uint8
	var lineSprites = 0
	for sprite := 0; sprite < 40; sprite++ {
		// sprite occupies 4 bytes in the sprite attributes table
		index := uint16(sprite) * 4

		// If this is true the scanline is out of the area we care about
		yPos := ppu.Oam.Read(0xFE00+index) - 16
		if ppu.ly < yPos || ppu.ly >= (yPos+ySize) {
			continue
		}

		// Only 10 sprites are allowed to be displayed on each line
		if lineSprites >= 10 {
			break
		}
		lineSprites++

		xPos := ppu.Oam.Read(0xFE00+index+1) - 8
		tileLocation := ppu.Oam.Read(0xFE00 + index + 2)
		if ySize == 16 {
			tileLocation &= 0b11111110
		}
		attributes := ppu.Oam.Read(0xFE00 + index + 3)

		yFlip := utils.TestBit(attributes, 6)
		xFlip := utils.TestBit(attributes, 5)
		priority := !utils.TestBit(attributes, 7)

		/*
			// Bank the sprite data in is (CGB only)
			var bank uint16 = 0
			if l.Mb.Cgb && internal.IsBitSet(attributes, 3) {
			    bank = 1
			}
		*/
		scanline := ppu.ly

		// Set the line to draw based on if the sprite is flipped on the y
		line := scanline - yPos
		if yFlip {
			line = ySize - line - 1
		}

		// Load the data containing the sprite data for this line
		dataAddress := 0x8000 + (uint16(tileLocation) * 16) + (uint16(line * 2))

		data1 := ppu.Vram.Read(dataAddress)
		data2 := ppu.Vram.Read(dataAddress + 1)

		for tilePixel := uint8(0); tilePixel < 8; tilePixel++ {
			pixel := int16(xPos) + int16(7-tilePixel)
			if pixel < 0 || pixel >= 160 {
				continue
			}

			// Check if the pixel has priority.
			//  - In DMG this is determined by the sprite with the smallest X coordinate,
			//    then the first sprite in the OAM.
			//  - In CGB this is determined by the first sprite appearing in the OAM.
			// We add a fixed 100 to the xPos so we can use the 0 value as the absence of a sprite.
			if minx[pixel] != 0 && ( /*gb.IsCGB() || */ minx[pixel] <= xPos+spritePriorityOffset) {
				continue
			}

			colorBit := tilePixel

			// read the sprite in backwards for the x axis
			if xFlip {
				colorBit = uint8(int8(colorBit-7) * -1)
			}

			colorNum := (utils.BitValue(data2, colorBit) << 1) | utils.BitValue(data1, colorBit)
			// Colour 0 is transparent for sprites
			if colorNum == 0 {
				continue
			}

			// Determine the colour palette to use
			palette := palette0
			if utils.TestBit(attributes, 4) {
				palette = palette1
			}

			if priority || ppu.tileScanline[pixel] == 0 {
				drawSquare(ppu.Fb, &coloredRects[ppu.getColor(colorNum, palette)], 4, int(pixel)*Scale, int(ppu.ly)*Scale)
			}

			// Store the xpos of the sprite for this pixel for priority resolution
			minx[pixel] = xPos + spritePriorityOffset
		}

	}
}

// Get the RGB colour value for a colour num at an address using the current palette.
func (ppu *PPU) getColor(colourNum uint8, palette uint8) uint8 {
	hi := colourNum<<1 | 1
	lo := colourNum << 1
	index := (utils.BitValue(palette, hi) << 1) | utils.BitValue(palette, lo)
	return index
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
	case 0xFF47:
		return ppu.bgp
	case 0xFF48:
		return ppu.obp0
	case 0xFF49:
		return ppu.obp1
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
		println("writing to ppu LY 0")
		ppu.ly = 0
		//ppu.ly = value
	case 0xFF45:
		ppu.lyc = value
	case 0xFF47:
		ppu.bgp = value
	case 0xFF48:
		ppu.obp0 = value
	case 0xFF49:
		ppu.obp1 = value

	default:
		panic("not handled ppu write")
	}
}
