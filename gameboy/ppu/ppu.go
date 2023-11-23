package ppu

import (
	"gogb/gameboy/utils"
	"image"
)

type PPU struct {
	Vram utils.Space
	Oam  utils.Space

	FbVram *image.RGBA // Framebuffer of vram tiles
}

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
