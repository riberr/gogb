package gameboy

import (
	"strings"
	"testing"
)

func TestBlarggTiming(t *testing.T) {
	testBlargg(
		"../third_party/gb-test-roms/instr_timing/",
		"instr_timing.gb",
		t,
	)
}

func testBlargg(
	romPath string,
	romName string,
	t *testing.T,
) {
	// SETUP
	gb := New(false)

	if !gb.Bus.LoadCart(romPath, romName) {
		t.Fatalf("error loading rom")
	}

	// RUN TEST
	lastLog := ""
	for {
		gb.Step()

		res := gb.SerialLink.GetLog()

		//
		if res != lastLog {

			println(strings.Trim(res, "\n"))
			//println(gb.Timer.Read(0xFF05))
		}
		//println(strings.Trim(gb.Cpu.Log, "\n"))
		lastLog = res

		if gb.Bus.Read(gb.Cpu.GetPC()) == 0x00 && gb.Bus.Read(gb.Cpu.GetPC()+1) == 0x18 && gb.Bus.Read(gb.Cpu.GetPC()+2) == 0xFD {
			panic("finish loop!")
		}
	}

	// ASSERT
	res := gb.SerialLink.GetLog()
	if strings.Trim(res[len(res)-7:], "\n") != "Passed" {
		t.Fatalf("%v did not return 'Passed'\n", romName)
	}
}
