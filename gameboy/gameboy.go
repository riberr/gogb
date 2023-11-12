package gameboy

import (
	busPackage "gogb/gameboy/bus"
	cpuPackage "gogb/gameboy/cpu"
	interrupts2 "gogb/gameboy/interrupts"
	"gogb/gameboy/seriallink"
	timerPackage "gogb/gameboy/timer"
)

type GameBoy struct {
	Timer      *timerPackage.Timer
	SerialLink *seriallink.SerialLink
	Bus        *busPackage.Bus
	Cpu        *cpuPackage.CPU
}

func New(debug bool) *GameBoy {
	interrupts := interrupts2.New()
	timer := timerPackage.New(interrupts)
	sl := seriallink.New()
	bus := busPackage.New(interrupts, timer, sl)
	cpu := cpuPackage.New(bus, interrupts, debug)

	return &GameBoy{
		Timer:      timer,
		SerialLink: sl,
		Bus:        bus,
		Cpu:        cpu,
	}
}

func (gb *GameBoy) Step() {
	gb.Timer.Tick()
	gb.Cpu.Step()
}

func Run() {

	romPath := "third_party/gb-test-roms/instr_timing/"
	romName := "instr_timing.gb"
	gb := New(false)
	gb.Bus.LoadCart(romPath, romName)

	for {
		gb.Step()
	}
}
