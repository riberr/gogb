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

const romsPath = "../third_party/gbmicrotest/bin/"

func TestOneRom(t *testing.T) {
	got, want, _ := testGbMicro(
		romsPath,
		"int_timer_halt.gb",
	)

	if got != want {
		t.Fatalf("Got %d, Want %d\n", got, want)
	}
	fmt.Printf("Correct result: %d\n", got)
}

// Passes all except 3: timer_tima_write_a, timer_tima_write_c, timer_tima_write_e. Very close though.
func TestTimer(t *testing.T) {
	roms := getRoms(romsPath, "timer_")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPath, test.rom)
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
	roms := getRoms(romsPath, "int_timer")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPath, test.rom)
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
	roms := getRoms(romsPath, "dma_")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPath, test.rom)
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

// passes 0. GoBoy passes halt_op_dupe and halt_op_dupe_delay
func TestHalt(t *testing.T) {
	roms := getRoms(romsPath, "halt_")

	for _, test := range roms {
		t.Run(test.rom, func(t *testing.T) {
			//t.Parallel()
			got, want, err := testGbMicro(romsPath, test.rom)
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
		if strings.HasPrefix(e.Name(), filterHasPrefix) {
			res = append(res, test{rom: e.Name()})
		}
	}

	return res
}
