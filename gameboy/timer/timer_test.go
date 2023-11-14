package timer

import (
	"gogb/gameboy/interrupts"
	"testing"
)

const cpuFreq = 4194304

func TestDIVSecond(t *testing.T) {
	// SETUP
	timer := New(interrupts.New())
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
	timer := New(interrupts.New())
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
	timer := New(interrupts.New())
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
	timer := New(interrupts.New())
	timer.Write(0xFF07, 0b100) // TAC: timer enabled, 4096Hz, /1024

	cycles := 0
	initialTima := timer.Read(0xFF05)
	// not a full period
	for i := 0; i < 6667; i++ {
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

func TestDivIncrease(t *testing.T) {
	timer := New(interrupts.New())

	for i := 0; i < 255; i++ {
		timer.Tick()
	}

	println(timer.Read(0xFF04))
	if timer.Read(0xFF04) != 0 {
		t.Fatalf("Should be 0")
	}

	timer.Tick()

	if timer.Read(0xFF04) != 1 {
		t.Fatalf("Should be 1")
	}
}

func TestTimaIncrease(t *testing.T) {
	if 1024 != addTima(0b_0000_00100) {
		t.Fatalf("should be 1024")
	}
	if 16 != addTima(0b_0000_00101) {
		t.Fatalf("should be 16")
	}

	if 64 != addTima(0b_0000_00110) {
		t.Fatalf("should be 64")
	}
	println(addTima(0b_0000_00111))
	if 256 != addTima(0b_0000_00111) {
		t.Fatalf("should be 256")
	}
}

func addTima(TAC uint8) uint16 {
	timer := New(interrupts.New())

	timer.Write(0xFF07, TAC)

	for {
		timer.Tick()

		if timer.Read(0xFF05) == 1 {
			break
		}
	}

	return timer.sysclk
}

/*
func TestClockDisabled(t *testing.T) {
InterruptManager intManager = new InterruptManager();
Timer timer = new Timer(intManager);

for (int i = 0; i < 2056; i++) {
timer.Tick();
}
// TIMA should not count
Assert.Equal(0, timer.TIMA);
// DIV doesn't care if the clock is enabled
Assert.Equal(8, timer.DIV);
}
/*


func TestTimaOverflow(t *testing.T) {
InterruptManager intManager = new InterruptManager();
Timer timer = new Timer(intManager);

timer.TAC = 0b_0000_00101;  // frequency = 16

int max = 16 * 256;
for (int i = 0; i < max; i++) {
timer.Tick();
}

// the interruption should not happen immediately
Assert.True(((intManager.IF >> 2) & 1) == 0);

timer.Tick();
timer.Tick();
timer.Tick();

Assert.Equal(4, timer.ticksSinceOverflow);
Assert.True(((intManager.IF >> 2) & 1) == 1);

timer.Tick();
Assert.Equal(5, timer.ticksSinceOverflow);
timer.Tick();
Assert.False(timer.overflow);

}

func AddTima(u8 TAC) int {
InterruptManager intManager = new InterruptManager();
Timer timer = new Timer(intManager);

timer.TAC = TAC;

while (true) {
timer.Tick();

if (timer.TIMA == 1) {
break;
}
}

return timer.timerCounter;
}

*/
