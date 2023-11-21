package utils

import "testing"

func TestMemoryHas(t *testing.T) {
	const from = 0xC000
	const to = 0xCFFF

	mem := NewSpace(from, to)

	if mem.Has(from) == false {
		t.Fatal("should be true")
	}

	if mem.Has(to) == false {
		t.Fatal("should be true")
	}

	if mem.Has(0xBFFF) == true {
		t.Fatal("should be false")
	}

	if mem.Has(0xD000) == true {
		t.Fatal("should be false")
	}
}

func TestMemoryAddressRange(t *testing.T) {
	const from = 0xC000
	const to = 0xCFFF

	mem := NewSpace(from, to)

	if mem.Read(from) != 0 {
		t.Fatal("should be 0")
	}
	mem.Write(from, 0xFF)
	if mem.Read(from) != 0xFF {
		t.Fatal("should be FF")
	}

	if mem.Read(to) != 0 {
		t.Fatal("should be 0")
	}
	mem.Write(to, 0xFF)
	if mem.Read(to) != 0xFF {
		t.Fatal("should be FF")
	}
}
