package timer

// https://github.com/davidwhitney/CoreBoy/blob/master/CoreBoy/timer/Timer.cs
// https://robertovaccari.com/blog/2020_09_26_gameboy/
// https://github.com/rvaccarim/FrozenBoy/blob/master/FrozenBoyCore/Processor/Timer.cs
// 1 machine cycle = 4 clock cycles

type timer struct {
	div  uint8 //0xFF04 divider register
	tima uint8 //0xFF05 timer register
	tma  uint8 //0xFF06 timer register overflow

	// 0xFF07 Timer Control
	// Bits 1-0 - Input Clock Select
	//            00: 4096   Hz
	//            01: 262144 Hz
	//            10: 65536  Hz
	//            11: 16384  Hz
	// Bit  2   - Timer Enable
	//
	// Note: The "Timer Enable" bit only affects TIMA,
	// DIV is ALWAYS counting.
	tac uint8
}

const (
	cpuFreq = 4194304 // Hz
	divFreq = 16384   // Hz
)

func (t timer) Tick() {
	// todo inc clock cycle
}
