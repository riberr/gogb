package emulator

import (
	cpuStruct "gogb/emulator/cpu"
	"gogb/emulator/memory"
)

func Run() {
	romPath := "roms/cpu_instrs/individual/"
	romName := "06-ld r,r.gb"

	memory.CartLoad(romPath + romName)
	//fmt.Printf("%02x ", rom)

	cpu := cpuStruct.NewCPU()

	for {
		cpu.Step()
		//time.Sleep(time.Millisecond * 100)
	}
}
