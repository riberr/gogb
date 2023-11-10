package gameboy

import (
	busPackage "gogb/gameboy/bus"
	cpuPackage "gogb/gameboy/cpu"
	"gogb/gameboy/seriallink"
	timer2 "gogb/gameboy/timer"
	"time"
)

func Run(debug bool) {

	// dependency injection
	timer := timer2.New()
	sl := seriallink.New()
	bus := busPackage.New(timer, sl)
	cpu := cpuPackage.New(bus, debug)

	romPath := "third_party/gb-test-roms/instr_timing/"
	romName := "instr_timing.gb"

	//bus2.CartLoad(romPath, romName)
	bus.LoadCart(romPath, romName)
	//fmt.Printf("%02x ", rom)

	for {
		timer.Tick()
		cpu.Step()
		time.Sleep(time.Millisecond * 100)
	}
}
