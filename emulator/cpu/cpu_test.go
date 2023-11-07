package cpu

import (
	"gogb/emulator/memory"
	"testing"
)

/*
 uses blargg's cpu_instrs test roms and gameboy-logs
 https://gbdev.gg8.se/files/roms/blargg-gb-tests/
 https://github.com/wheremyfoodat/Gameboy-logs/tree/master
*/

func TestCpuOutput06(t *testing.T) {
	memory.CartLoad("roms/cpu_instrs/individual/", "06-ld r,r.gb")

	cpu := NewCPU()

	for {
		cpu.Step()
		//time.Sleep(time.Millisecond * 100)
	}
}

func CompareOutputToTruth06(t *testing.T) {

}
