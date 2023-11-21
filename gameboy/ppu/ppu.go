package ppu

import (
	"gogb/gameboy/utils"
)

type PPU struct {
	Vram utils.Space
	Oam  utils.Space
}

func New() *PPU {
	return &PPU{
		Vram: utils.NewSpace(0x8000, 0x9FFF), // Video RAM
		Oam:  utils.NewSpace(0xFE00, 0xFE9F), // Object attribute bus
	}
}

/*
func (ppu *PPU) VramHas(address uint16) bool {
	return ppu.vram.Has(address)
}

func (ppu *PPU) VramWrite(address uint16, value uint8) {
	ppu.vram.Write(address, value)
}

func (ppu *PPU) VramRead(address uint16) uint8 {
	return ppu.vram.Read(address)
}

func (ppu *PPU) OamHas(address uint16) bool {
	return ppu.oam.Has(address)
}

func (ppu *PPU) OamWrite(address uint16, value uint8) {
	ppu.oam.Write(address, value)
}

func (ppu *PPU) OamRead(address uint16) uint8 {
	return ppu.oam.Read(address)
}
*/
