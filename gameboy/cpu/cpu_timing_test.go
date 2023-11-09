package cpu

import (
	bus2 "gogb/gameboy/bus"
	"gogb/gameboy/seriallink"
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

func testTimingWithRom(
	romPath string,
	romName string,
	t *testing.T,
) {
	// SETUP
	sl := seriallink.New()
	bus := bus2.New(sl)
	cpu := New(bus, false)

	if !bus.LoadCart(romPath, romName) {
		t.Fatalf("error loading rom")
	}

	// RUN TEST
	i := 1
	for {
		cpu.Step()
		i++

		res := sl.GetLog()
		println(strings.Trim(res, "\n"))
	}

	// ASSERT
	res := sl.GetLog()
	if strings.Trim(res[len(res)-7:], "\n") != "Passed" {
		t.Fatalf("%v did not return 'Passed'\n", romName)
	}
}
