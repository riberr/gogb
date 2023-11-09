package gameboy

import (
	busPackage "gogb/gameboy/bus"
	cpuPackage "gogb/gameboy/cpu"
	"gogb/gameboy/seriallink"
	"time"
)

func Run(debug bool) {

	// dependency injection
	sl := seriallink.New()
	bus := busPackage.New(sl)
	cpu := cpuPackage.New(bus, debug)

	romPath := "roms/cpu_instrs/individual/"
	romName := "06-ld r,r.gb"

	//bus2.CartLoad(romPath, romName)
	bus.LoadCart(romPath, romName)
	//fmt.Printf("%02x ", rom)

	for {
		cpu.Step()
		time.Sleep(time.Millisecond * 100)
	}
}