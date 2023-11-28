package main

import (
	ebiten2 "gogb/ebiten"
	"gogb/gameboy"
)

func main() {
	//romPath := "third_party/games/"
	//romName := "Yoshi.gb"

	//romPath := "third_party/games/"
	//romName := "Tetris.gb"

	romPath := "third_party/gb-test-roms/instr_timing/"
	romName := "instr_timing.gb"

	//romPath := "third_party/"
	//romName := "dmg-acid2.gb"

	//romPath := "third_party/roms/games/"
	//romName := "Pokemon Red.gb"

	//romPath := "third_party/mooneye/manual-only/"
	//romName := "sprite_priority.gb"

	gb := gameboy.New(false)
	if !gb.Bus.LoadCart(romPath, romName) {
		panic("error loading rom")
	}
	ebiten := ebiten2.New(gb)
	ebiten.Run()
}
