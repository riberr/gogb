package main

import (
	"testing"
)

func TestFlagSetAndRead(t *testing.T) {
	register := NewRegisters()

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
