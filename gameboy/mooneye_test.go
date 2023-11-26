package gameboy

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestOneMooneye(t *testing.T) {
	res, err := testMooneye(
		"../third_party/mooneye/acceptance/timer/",
		"div_write.gb",
	)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if !res {
		t.Errorf("Failed")
	}
}

func TestMooneyeTimer(t *testing.T) {
	path := "../third_party/mooneye/acceptance/timer/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testMooneye(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func TestMooneyeInterrupts(t *testing.T) {
	path := "../third_party/mooneye/acceptance/interrupts/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testMooneye(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func TestMooneyeMBC1(t *testing.T) {
	path := "../third_party/mooneye/emulator-only/mbc1/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testMooneye(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func TestMooneyeTimingAndVarious(t *testing.T) {
	path := "../third_party/mooneye/acceptance/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testMooneye(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func TestMooneyePPU(t *testing.T) {
	path := "../third_party/mooneye/acceptance/ppu/"
	roms := getRoms(path, "")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			result, err := testMooneye(path, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if !result {
				t.Errorf("Failed")
			}
		})
	}
}

func testMooneye(
	romPath string,
	romName string,
) (bool, error) {
	gb := New(false)

	if !gb.Bus.LoadCart(romPath, romName) {
		return false, errors.New("could not load rom")
	}

	// iterate until detecting finish loop or until max iterations
	for i := 0; i < 10000000; i++ {
		gb.Step()
		if inFinishLoop(gb) {
			println("finish loop")
			break
		}
	}

	if !didPass(gb) {
		fmt.Printf("registers not fibonacci: %v", gb.Cpu.GetInternalString())
		return false, nil
	}

	return true, nil
}

func inFinishLoop(gb *GameBoy) bool {
	return gb.Bus.Read(gb.Cpu.GetPC()) == 0x00 &&
		gb.Bus.Read(gb.Cpu.GetPC()+1) == 0x18 &&
		gb.Bus.Read(gb.Cpu.GetPC()+2) == 0xFD
}

func didPass(gb *GameBoy) bool {
	regs := gb.Cpu.GetInternalString()
	if strings.HasPrefix(regs, "A:00 F:C0 B:03 C:05 D:08 E:0D H:15 L:22") {
		return true
	}
	return false
}
