package gameboy

import (
	busPackage "gogb/gameboy/bus"
	cpuPackage "gogb/gameboy/cpu"
	interrupts2 "gogb/gameboy/interrupts"
	"gogb/gameboy/seriallink"
	timerPackage "gogb/gameboy/timer"
)

type GameBoy struct {
	Interrupts *interrupts2.Interrupts
	Timer      *timerPackage.Timer
	timer2     *timerPackage.Timer2
	SerialLink *seriallink.SerialLink
	Bus        *busPackage.Bus
	Cpu        *cpuPackage.CPU
}

func New(debug bool) *GameBoy {
	interrupts := interrupts2.New()
	timer := timerPackage.New(interrupts)
	timer2 := timerPackage.NewTimer2(interrupts)
	sl := seriallink.New()
	bus := busPackage.New(interrupts, timer, timer2, sl)
	cpu := cpuPackage.New(bus, interrupts, debug)

	return &GameBoy{
		Interrupts: interrupts,
		Timer:      timer,
		timer2:     timer2,
		SerialLink: sl,
		Bus:        bus,
		Cpu:        cpu,
	}
}

func (gb *GameBoy) Step() int {
	cycles := 0
	cyclesOp := 0

	cyclesOp = gb.Cpu.Step()

	cycles += cyclesOp

	gb.Timer.UpdateTimers(cyclesOp)

	//gb.updateGraphics(cyclesOp)
	cycles += gb.Cpu.DoInterrupts()

	//gb.Sound.Buffer(cyclesOp, gb.getSpeed())
	return cycles
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
