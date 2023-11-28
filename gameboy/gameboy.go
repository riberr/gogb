package gameboy

import (
	busPackage "gogb/gameboy/bus"
	cpuPackage "gogb/gameboy/cpu"
	interruptsPackage "gogb/gameboy/interrupts"
	joypadPackage "gogb/gameboy/joypad"
	ppuPackage "gogb/gameboy/ppu"
	"gogb/gameboy/seriallink"
	timerPackage "gogb/gameboy/timer"
)

type GameBoy struct {
	Interrupts *interruptsPackage.Interrupts2
	Timer      *timerPackage.Timer
	SerialLink *seriallink.SerialLink
	Bus        *busPackage.Bus
	Cpu        *cpuPackage.CPU
	Ppu        *ppuPackage.PPU
	JoyPad     *joypadPackage.JoyPad
}

func New(debug bool) *GameBoy {
	interrupts := interruptsPackage.NewInterrupts2()
	timer := timerPackage.New(interrupts)
	sl := seriallink.New()
	ppu := ppuPackage.New(interrupts)
	joyPad := joypadPackage.New(interrupts)
	bus := busPackage.New(interrupts, timer, sl, ppu, joyPad)
	cpu := cpuPackage.New(bus, interrupts, debug)

	return &GameBoy{
		Interrupts: interrupts,
		Timer:      timer,
		SerialLink: sl,
		Bus:        bus,
		Cpu:        cpu,
		Ppu:        ppu,
		JoyPad:     joyPad,
	}
}

func (gb *GameBoy) Step() int {
	cycles := 0
	cyclesOp := 0

	cyclesOp = gb.Cpu.Step()

	cycles += cyclesOp

	gb.Timer.UpdateTimers(cyclesOp)

	gb.Ppu.Update(cyclesOp)

	cycles += gb.Cpu.DoInterrupts()

	//gb.Sound.Buffer(cyclesOp, gb.getSpeed())
	return cycles
}

func (gb *GameBoy) GenerateGraphics() {
	gb.Ppu.GenerateDebugGraphics()
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
