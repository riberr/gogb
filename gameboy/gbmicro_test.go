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
	gb := New(true)

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
