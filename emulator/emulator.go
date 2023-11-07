package emulator

import (
	cpu2 "gogb/emulator/cpu"
	"gogb/emulator/memory"
	"time"
)

func Run() {
	memory.CartLoad()
	//fmt.Printf("%02x ", rom)

	cpu := cpu2.NewCPU()

	for {
		cpu.Step()
		time.Sleep(time.Millisecond * 100)
	}
}
