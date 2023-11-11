package cpu

import (
	"fmt"
	bus2 "gogb/gameboy/bus"
	interrupts2 "gogb/gameboy/interrupts"
	"gogb/gameboy/seriallink"
	timer2 "gogb/gameboy/timer"
	"testing"
)

func TestSBC(t *testing.T) {
	/*
		input:
		A: 00 F: 00 B: 01 C: 0F D: 7F E: 80 H: 10 L: 1F SP: DFF1 PC: 00:DEF8 (98 00 00 C3) [Z: false, N: false, H: false, C: false]
		got:
		A: FF F: 50 B: 01 C: 0F D: 7F E: 80 H: 10 L: 1F SP: DFF1 PC: 00:DEF9 (00 00 C3 B5)
		want:
		A: FF F: 70 B: 01 C: 0F D: 7F E: 80 H: 10 L: 1F SP: DFF1 PC: 00:DEF9 (00 00 C3 B5)
	*/

	// setup
	interrupts := interrupts2.New()
	timer := timer2.New(interrupts)
	sl := seriallink.New()
	bus := bus2.New(interrupts, timer, sl)
	cpu := New(bus, interrupts, false)

	// preconditions
	cpu.regs.setAF(0x0000)
	cpu.regs.setBC(0x010F)
	var value uint8 = 1

	// test
	sbc(cpu, value)

	// assert
	if cpu.regs.a != 0xFF {
		t.Errorf("A: got %02x, want %02x", cpu.regs.a, 0xFF)
	}

	if cpu.regs.f != 0x70 {
		t.Errorf("F: got %08b, want %08b", cpu.regs.f, 0x70)
	}
}

func TestRelativeJumpOP(t *testing.T) {
	// di
	interrupts := interrupts2.New()
	timer := timer2.New(interrupts)
	sl := seriallink.New()
	bus := bus2.New(interrupts, timer, sl)
	cpu := New(bus, interrupts, false)

	// preconditions
	opcode := OpCodes[0x18]
	cpu.pc = 0xc2cb
	bus.Write(cpu.pc, 0xf4)

	// test
	for _, step := range opcode.steps {
		step(cpu)
	}

	fmt.Printf("result: %02x \n", cpu.pc)

	truth := uint16(0xC2C0)
	if cpu.pc != truth {
		t.Fatalf("Got %02x, expected %02x", cpu.pc, truth)
	}
}
