package gameboy

import (
	"testing"
)

func TestGbMicro(t *testing.T) {
	testGbMicro(
		"../third_party/gbmicrotest/bin/",
		"halt_bug.gb",
		t,
	)
}

func testGbMicro(
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

		res := gb.Cpu.Log

		//
		if res != lastLog {
			println(res)
		}
		lastLog = res

		if gb.Bus.Read(0xFF82) != 0 {
			println(gb.Bus.Read(0xFF80))
			println(gb.Bus.Read(0xFF81))
			println(gb.Bus.Read(0xFF82))
			break
		}
	}
	if gb.Bus.Read(0xFF82) != 1 {
		t.Fatalf("did not pass")
	}
}
