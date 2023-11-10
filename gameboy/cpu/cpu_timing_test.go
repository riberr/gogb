package cpu

import (
	bus2 "gogb/gameboy/bus"
	"gogb/gameboy/seriallink"
	timer2 "gogb/gameboy/timer"
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
	timer := timer2.New()
	sl := seriallink.New()
	bus := bus2.New(timer, sl)
	cpu := New(bus, true)

	if !bus.LoadCart(romPath, romName) {
		t.Fatalf("error loading rom")
	}

	// RUN TEST
	i := 1
	for {
		timer.Tick()
		cpu.Step()
		i++

		res := sl.GetLog()
		if res != "" {
			println(strings.Trim(res, "\n"))
		}
	}

	// ASSERT
	res := sl.GetLog()
	if strings.Trim(res[len(res)-7:], "\n") != "Passed" {
		t.Fatalf("%v did not return 'Passed'\n", romName)
	}
}
