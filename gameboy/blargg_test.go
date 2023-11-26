package gameboy

import (
	"errors"
	"strings"
	"testing"
)

/*
 uses blargg's cpu_instrs test roms and gameboy-logs
 https://gbdev.gg8.se/files/roms/blargg-gb-tests/
 https://github.com/wheremyfoodat/Gameboy-logs/tree/master
*/

func TestOneBlargg(t *testing.T) {
	res, err := testBlarggRom(
		"../third_party/gb-test-roms/cpu_instrs/individual/",
		"02-interrupts.gb",
	)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !res {
		t.Errorf("Failed")
	}
}

func TestBlarggCpuInstr(t *testing.T) {
	path := "../third_party/gb-test-roms/cpu_instrs/individual/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testBlarggRom(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func TestBlarggInstrTiming(t *testing.T) {
	path := "../third_party/gb-test-roms/instr_timing/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testBlarggRom(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

// ; Tests interrupt handling time for slow and fast CPU.
func TestBlarggInterruptTime(t *testing.T) {
	path := "../third_party/gb-test-roms/interrupt_time/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testBlarggRom(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func TestBlarggMemTiming(t *testing.T) {
	path := "../third_party/gb-test-roms/mem_timing/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testBlarggRom(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func TestBlarggMemTiming2(t *testing.T) {
	path := "../third_party/gb-test-roms/mem_timing-2/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testBlarggRom(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func testBlarggRom(
	romPath string,
	romName string,
) (bool, error) {
	gb := New(false)

	if !gb.Bus.LoadCart(romPath, romName) {
		return false, errors.New("could not load rom")
	}

	// iterate until detecting finish loop or until max iterations
	for i := 0; i < 20000000; i++ {
		gb.Step()
		if strings.Contains(gb.SerialLink.GetLog(), "Passed") {
			return true, nil
		}
	}
	println(gb.SerialLink.GetLog())
	return false, nil
}
