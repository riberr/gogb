package timer

import (
	"testing"
)

const cpuFreq = 4194304

func TestDIVSecond(t *testing.T) {
	// SETUP
	timer := New()
	var cycles = 0

	// 1 second
	for i := 0; i < cpuFreq; i++ {
		timer.Tick()
		if timer.Read(0xFF04) == 255 {
			cycles++
		}
	}

	// timer.div should be 0 as we have passed an even number of cycles
	if timer.Read(0xFF04) != 0 {
		t.Fatal("timer.div should be 0")
	}

	if cycles != 16384 {
		t.Fatalf("DIV should run at 16384 Hz")
	}

}

func TestDIVHalfSecond(t *testing.T) {
	// SETUP
	timer := New()
	var cycles = 0

	// 0.5 second
	for i := 0; i < cpuFreq/2; i++ {
		timer.Tick()
		if timer.Read(0xFF04) == 255 {
			cycles++
		}
	}

	// timer.div should be 0 as we have passed an even number of cycles
	if timer.Read(0xFF04) != 0 {
		t.Fatal("timer.div should be 0")
	}

	if cycles != 16384/2 {
		t.Fatalf("DIV should run at 16384 Hz")
	}
}

func TestTIMAClockDisabled(t *testing.T) {
	// SETUP
	timer := New()
	timer.Write(0xFF07, 0x00) // timer disabled

	initialTima := timer.Read(0xFF05)
	timer.Tick()
	steppedTima := timer.Read(0xFF05)

	println(initialTima)
	println(steppedTima)

	if initialTima != steppedTima {
		t.Fatalf("tima should not increment as timer is disabled")
	}
}

func TestTIMAClockDEnabled(t *testing.T) {
	// SETUP
	timer := New()
	timer.Write(0xFF07, 0b100) // TAC: timer enabled, 4096Hz, /1024

	cycles := 0
	initialTima := timer.Read(0xFF05)
	for i := 0; i < cpuFreq; i++ {
		timer.Tick()

	}
	if timer.Read(0xFF05) == 0 {
		cycles++
	}
	steppedTima := timer.Read(0xFF05)

	println(initialTima)
	println(steppedTima)
	println(cycles)

	if initialTima == steppedTima {
		t.Fatalf("tima should increment as timer is enabled")
	}
}
