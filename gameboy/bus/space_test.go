package bus

import "testing"

func TestMemoryHas(t *testing.T) {
	const from = 0xC000
	const to = 0xCFFF

	mem := NewMemory(from, to)

	if mem.has(from) == false {
		t.Fatal("should be true")
	}

	if mem.has(to) == false {
		t.Fatal("should be true")
	}

	if mem.has(0xBFFF) == true {
		t.Fatal("should be false")
	}

	if mem.has(0xD000) == true {
		t.Fatal("should be false")
	}
}

func TestMemoryAddressRange(t *testing.T) {
	const from = 0xC000
	const to = 0xCFFF

	mem := NewMemory(from, to)

	if mem.read(from) != 0 {
		t.Fatal("should be 0")
	}
	mem.write(from, 0xFF)
	if mem.read(from) != 0xFF {
		t.Fatal("should be FF")
	}

	if mem.read(to) != 0 {
		t.Fatal("should be 0")
	}
	mem.write(to, 0xFF)
	if mem.read(to) != 0xFF {
		t.Fatal("should be FF")
	}
}
