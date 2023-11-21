package gameboy

import (
	"strings"
	"testing"
)

/*
func TestTiming(t *testing.T) {
	t.Fatalf("TODO")
	testTimingWithRom(
		"../third_party/gb-test-roms/instr_timing/",
		"instr_timing.gb",
		t,
	)
}
*/

func TestMooneye(t *testing.T) {
	testTimingWithRom(
		"../third_party/mooneye/emulator-only/mbc1/",
		"bits_bank1.gb",
		t,
	)
}

func testTimingWithRom(
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
	//lastDebug := ""
	for {
		gb.Step()

		res := gb.SerialLink.GetLog()
		//resDebug := gb.Cpu.
		//
		if res != lastLog {
			println(strings.Trim(res, "\n"))
			println(gb.Timer.Read(0xFF05))
		}
		//if resDebug != lastDebug {
		//	println(lastDebug)
		//}
		//println(strings.Trim(gb.Cpu.Log, "\n"))
		//lastLog = res
		//lastDebug = resDebug
	}

	// ASSERT
	res := gb.SerialLink.GetLog()
	if strings.Trim(res[len(res)-7:], "\n") != "Passed" {
		t.Fatalf("%v did not return 'Passed'\n", romName)
	}
}
