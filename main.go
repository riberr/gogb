package main

import (
	ebiten2 "gogb/ebiten"
	"gogb/gameboy"
)

func main() {
	//romPath := "third_party/roms/"
	//romName := "Tetris.gb"

	//romPath := "third_party/roms/"
	//romName := "Yoshi.gb"

	romPath := "third_party/roms/"
	romName := "Kirby.gb"

	//romPath := "third_party/scribbltests/"
	//romName := "winpos.gb"

	//romPath := "third_party/gb-test-roms/mem_timing/"
	//romName := "mem_timing.gb"

	//romPath := "third_party/"
	//romName := "dmg-acid2.gb"

	//romPath := "third_party/roms/games/"
	//romName := "Pokemon Red.gb"

	//romPath := "third_party/mooneye/emulator-only/mbc1/"
	//romName := "bits_bank1.gb"

	gb := gameboy.New(false)
	if !gb.Bus.LoadCart(romPath, romName) {
		panic("error loading rom")
	}
	ebiten := ebiten2.New(gb)
	ebiten.Run()
}
