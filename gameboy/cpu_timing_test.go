package gameboy

import (
	"strings"
	"testing"
)

func TestTiming(t *testing.T) {
	testTimingWithRom(
		"../../third_party/gb-test-roms/instr_timing/",
		"instr_timing.gb",
		t,
	)
}

func TestTimingMooneye(t *testing.T) {
	testTimingWithRom(
		"../../third_party/mooneye/acceptance/timer/",
		"tim00.gb",
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
	i := 1
	for {
		gb.Step()
		i++

		res := gb.SerialLink.GetLog()
		if res != "" {
			println(strings.Trim(res, "\n"))
		}
	}

	// ASSERT
	res := gb.SerialLink.GetLog()
	if strings.Trim(res[len(res)-7:], "\n") != "Passed" {
		t.Fatalf("%v did not return 'Passed'\n", romName)
	}
}
