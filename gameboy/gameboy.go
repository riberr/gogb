package gameboy

import (
	busPackage "gogb/gameboy/bus"
	cpuPackage "gogb/gameboy/cpu"
	interrupts2 "gogb/gameboy/interrupts"
	"gogb/gameboy/seriallink"
	timer2 "gogb/gameboy/timer"
	"time"
)

func Run(debug bool) {

	// dependency injection
	interrupts := interrupts2.New()
	timer := timer2.New(interrupts)
	sl := seriallink.New()
	bus := busPackage.New(interrupts, timer, sl)
	cpu := cpuPackage.New(bus, interrupts, debug)

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
