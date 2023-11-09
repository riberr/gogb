package cpu

import (
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

	// preconditions
	cpu := NewCPU(true)
	cpu.regs.setAF(0x0000)
	cpu.regs.setBC(0x010F)
	var value uint8 = 1

	// test
	sbc(&cpu, value)

	//fmt.Printf("a: %02x \n", cpu.regs.a)
	//fmt.Printf("f: %02x \n", cpu.regs.f)

	// assert
	if cpu.regs.a != 0xFF {
		t.Fatalf("A: got %02x, want %02x", cpu.regs.a, 0xFF)
	}

	if cpu.regs.f != 0x70 {
		t.Fatalf("F: got %08b, want %08b", cpu.regs.f, 0x70)
	}
}
