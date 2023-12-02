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

	romPath := "third_party/mooneye/emulator-only/mbc2/"
	romName := "rom_1Mb.gb"

	//romPath := "third_party/mealybug-tearoom-tests/ppu/"
	//romName := "m3_wx_5_change.gb"

	gb := gameboy.New(false)
	if !gb.Bus.LoadCart(romPath, romName) {
		panic("error loading rom")
	}
	ebiten := ebiten2.New(gb)
	ebiten.Run()
}
