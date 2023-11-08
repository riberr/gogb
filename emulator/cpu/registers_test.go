package cpu

import (
	"testing"
)

func TestFlagSetAndRead(t *testing.T) {
	register := NewRegisters()
	register.f = 0x00
	if register.getFlag(FLAG_ZERO_Z_BIT) {
		t.Fatal("Should be false")
	}
	register.setFlag(FLAG_ZERO_Z_BIT, true)
	if !register.getFlag(FLAG_ZERO_Z_BIT) {
		t.Fatal("Should be true")
	}

	if register.getFlag(FLAG_CARRY_C_BIT) {
		t.Fatal("Should be false")
	}
	register.setFlag(FLAG_CARRY_C_BIT, true)
	if !register.getFlag(FLAG_CARRY_C_BIT) {
		t.Fatal("Should be true")
	}

	if register.getFlag(FLAG_SUBTRACTION_N_BIT) {
		t.Fatal("Should be false")
	}
	register.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	if !register.getFlag(FLAG_SUBTRACTION_N_BIT) {
		t.Fatal("Should be true")
	}

	if register.getFlag(FLAG_HALF_CARRY_H_BIT) {
		t.Fatal("Should be false")
	}
	register.setFlag(FLAG_HALF_CARRY_H_BIT, true)
	if !register.getFlag(FLAG_HALF_CARRY_H_BIT) {
		t.Fatal("Should be true")
	}
}

func TestRegisterBC(t *testing.T) {
	register := NewRegisters()

	if register.getBC() != 0 {
		t.Fatal("Should be 0")
	}

	value := uint16(0xFE12)
	register.setBC(value)

	if register.getBC() != value {
		t.Fatalf("Should be %v", value)
	}
}
