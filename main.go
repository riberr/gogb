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

	//romPath := "third_party/roms/"
	//romName := "Kirby.gb"

	//romPath := "third_party/scribbltests/"
	//romName := "winpos.gb"

	//romPath := "third_party/gb-test-roms/interrupt_time/"
	//romName := "interrupt_time.gb"

	//romPath := "third_party/"
	//romName := "dmg-acid2.gb"

	//romPath := "third_party/roms/games/"
	//romName := "Pokemon Red.gb"

	romPath := "third_party/mbc3/"
	romName := "rtc3test.gb"

	//romPath := "third_party/roms/games/"
	//romName := "Pokemon Red.gb"

	gb := gameboy.New(false)
	if !gb.Bus.LoadCart(romPath, romName) {
		panic("error loading rom")
	}
	ebiten := ebiten2.New(gb)
	ebiten.Run()
}
