package ppu

import (
	"image"
	"image/color"
)

const (
	vramWidth  = (16 * 8 * Scale) + (16 * Scale)
	vramHeight = (24 * 8 * Scale) + (24 * Scale)
)

var colors = []color.RGBA{
	{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
	{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xFF},
	{R: 0x55, G: 0x55, B: 0x55, A: 0xFF},
	{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
}
var coloredRects = []image.Uniform{
	{C: colors[0]},
	{C: colors[1]},
	{C: colors[2]},
	{C: colors[3]},
}

func (ppu *PPU) GenerateDebugVram() {
	xDraw, yDraw, tileNum := 0, 0, uint16(0)

	//384 tiles, 24 x 16
	for y := 0; y < 24; y++ {
		for x := 0; x < 16; x++ {
			ppu.drawTile(ppu.FbVram, tileNum, xDraw+(x*Scale), yDraw+(y*Scale))
			xDraw += 8 * Scale
			tileNum++
		}

		yDraw += 8 * Scale
		xDraw = 0
	}
}

// https://github.com/rockytriton/LLD_gbemu/blob/main/part11/lib/ui.c
func (ppu *PPU) drawTile(img *image.RGBA, tileNum uint16, x int, y int) {
	for tileY := uint16(0); tileY < 16; tileY += 2 {
		b1 := ppu.Vram.Read(VramAddr + (tileNum * 16) + tileY)
		b2 := ppu.Vram.Read(VramAddr + (tileNum * 16) + tileY + 1)

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

			//colorBit := uint8(int8((bit%8)-7) * -1)
			//c := (utils.BitValue(b1, colorBit) << 1) | utils.BitValue(b2, colorBit)

			xx := x + ((7 - bit) * Scale)
			yy := y + (int(tileY) / 2 * Scale)

			drawSquare(img /*colors[c]*/, &coloredRects[c], 4, xx, yy)
			//drawSquare2(img /*colors[c]*/, colors[c], 4, xx, yy)

		}
	}
}
