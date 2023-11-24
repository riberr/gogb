package ppu

import (
	"image"
	"image/color"
	"image/draw"
)

const (
	scale    = 4
	vramAddr = 0x8000
)

var sp = image.Point{}

func drawSquare(img *image.RGBA, c /*color.Color*/ *image.Uniform, size int, x, y int) {
	draw.Draw(img, image.Rect(x, y, x+size, y+size), c, sp, draw.Over)
}

// slower??
func drawSquare2(img *image.RGBA, c color.Color, size int, x, y int) {
	for dy := 0; dy < size; dy++ {
		for dx := 0; dx < size; dx++ {
			img.Set(x+dx, y+dy, c)
		}
	}
}
