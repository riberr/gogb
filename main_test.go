package main

import (
	ebiten2 "gogb/ebiten"
	"gogb/gameboy"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	romPath := "third_party/games/"
	romName := "Tetris.gb"

	//romPath := "third_party/gb-test-roms/cpu_instrs/individual/"
	//romName := "01-special.gb"

	gb := gameboy.New(false)
	if !gb.Bus.LoadCart(romPath, romName) {
		panic("error loading rom")
	}
	ebiten := ebiten2.New(gb)
	ebiten.Run()
}
