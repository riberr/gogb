package cpu

import (
	"fmt"
	"gogb/gameboy"
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
	gb := gameboy.New(false)

	// preconditions
	gb.Cpu.regs.setAF(0x0000)
	gb.Cpu.regs.setBC(0x010F)
	var value uint8 = 1

	// test
	sbc(gb.Cpu, value)

	// assert
	if gb.Cpu.regs.a != 0xFF {
		t.Errorf("A: got %02x, want %02x", gb.Cpu.regs.a, 0xFF)
	}

	if gb.Cpu.regs.f != 0x70 {
		t.Errorf("F: got %08b, want %08b", gb.Cpu.regs.f, 0x70)
	}
}

func TestRelativeJumpOP(t *testing.T) {
	gb := gameboy.New(false)

	// preconditions
	opcode := OpCodes[0x18]
	gb.Cpu.pc = 0xc2cb
	gb.Bus.Write(gb.Cpu.pc, 0xf4)

	// test
	for _, step := range opcode.steps {
		step(gb.Cpu)
	}

	truth := uint16(0xC2C0)
	if gb.Cpu.pc != truth {
		t.Fatalf("Got %02x, expected %02x", gb.Cpu.pc, truth)
	}
}

// LD HL,SP+i8
func Test0xF8(t *testing.T) {
	gb := gameboy.New(false)

	// preconditions
	opcode := OpCodes[0xF8]
	gb.Cpu.pc = 0xC2C5
	gb.Cpu.sp = 0xDFFD
	gb.Bus.Write(gb.Cpu.pc, 0xfe)

	// test
	for _, step := range opcode.steps {
		step(gb.Cpu)
	}

	fmt.Printf("result: %02x \n", gb.Cpu.regs.getHL())

	truth := uint16(0xDFFB)
	if gb.Cpu.regs.getHL() != truth {
		t.Fatalf("Got %02x, expected %02x", gb.Cpu.pc, truth)
	}
}
