package gameboy

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

type test struct {
	rom string
}

const romsPathGbMicro = "../third_party/gbmicrotest/bin/"

func TestOneRom(t *testing.T) {
	got, want, _ := testGbMicro(
		romsPathGbMicro,
		"int_hblank_halt_bug_a.gb",
	)

	if got != want {
		t.Fatalf("Got %d, Want %d\n", got, want)
	}
	fmt.Printf("Correct result: %d\n", got)
}

// Passes all except 3: timer_tima_write_a, timer_tima_write_c, timer_tima_write_e. Very close though.
func TestTimer(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "timer_")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

// passes 2/7. GoBoy passes 3 (diff is int_timer_incs.gb)
func TestInterruptsTimer(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "int_timer")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

func TestInterruptsHBlank(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "int_hblank_halt_s")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

// passes 1. GoBoy passes halt_op_dupe and halt_op_dupe_delay
func TestHalt(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "halt_")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

// passes 5/7. not dma_basic nor dma_timing_a
func TestDma(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "dma_")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

// passes 8/16
func TestVram(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "vram")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

func TestOam(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "oam_write")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

func TestMBC1Banking(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "mbc1_")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

func TestSprite(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "sprite")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

func TestLCDOn(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "lcdon")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

func TestLYWhileLCDOff(t *testing.T) {
	roms := getRoms(romsPathGbMicro, "ly_while")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPathGbMicro, test.rom)
			if err != nil {
				t.Fatalf("error: %v", err)
			}
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			fmt.Printf("Correct result: %d\n", got)
		})
	}
}

func testGbMicro(
	romPath string,
	romName string,
	// t *testing.T,
) (uint8, uint8, error) {
	// SETUP
	gb := New(false)

	if !gb.Bus.LoadCart(romPath, romName) {
		return 0, 0, errors.New("could not load rom")
	}

	// RUN TEST
	for {
		gb.Step()

		if gb.Bus.Read(0xFF82) != 0 {
			if gb.Bus.Read(0xFF80) != gb.Bus.Read(0xFF81) {
				fmt.Printf("%v: Did not pass: test result: %v, expected result: %v",
					romName, gb.Bus.Read(0xFF80), gb.Bus.Read(0xFF81))
			}
			return gb.Bus.Read(0xFF80), gb.Bus.Read(0xFF81), nil
		}
	}
}

func getRoms(path string, filterHasPrefix string) []test {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var res []test
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), filterHasPrefix) && strings.HasSuffix(e.Name(), ".gb") {
			res = append(res, test{rom: e.Name()})
		}
	}

	return res
}
