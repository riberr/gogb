package cpu

import (
	"bufio"
	"fmt"
	"gogb/emulator/memory"
	"os"
	"strings"
	"testing"
)

/*
 uses blargg's cpu_instrs test roms and gameboy-logs
 https://gbdev.gg8.se/files/roms/blargg-gb-tests/
 https://github.com/wheremyfoodat/Gameboy-logs/tree/master
*/

func TestCpuOutput09(t *testing.T) {
	logFile, err := os.Open("../../Gameboy-logs-master/Blargg9LYStubbed/Blargg9.txt")
	if err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	defer logFile.Close()

	log := bufio.NewReader(logFile)

	if !memory.CartLoad("../../roms/cpu_instrs/individual/", "09-op r,r.gb") {
		t.Fatalf("error loading rom")
	}
	cpu := NewCPU(false)

	nrOfLines := 243272
	i := 1
	for {
		output := fmt.Sprintf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: 00:%04X (%02X %02X %02X %02X)\n",
			cpu.regs.a, cpu.regs.f, cpu.regs.b, cpu.regs.c, cpu.regs.d, cpu.regs.e, cpu.regs.h, cpu.regs.l, cpu.sp, cpu.pc,
			memory.BusRead(cpu.pc), memory.BusRead(cpu.pc+1), memory.BusRead(cpu.pc+2), memory.BusRead(cpu.pc+3),
		)
		cpu.Step()

		logLine, _, err := log.ReadLine()
		if err != nil {

			if err.Error() == "EOF" {
				break
			}

			fmt.Println("Error reading line:", err)
			return
		}

		if strings.Trim(string(logLine), "\n") != strings.Trim(output, "\n") {
			t.Fatalf("%v/%v: not equal!\ngot: \n%vwant: \n%v", i, nrOfLines, output, string(logLine))
		}
		i++
	}
}
