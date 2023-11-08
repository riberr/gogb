package memory

type Memory struct {
	from uint16
	to   uint16
	size uint16
	data []uint8
}

func NewMemory(from uint16, to uint16) Memory {
	size := to - from + 1
	mem := Memory{
		from: from,
		to:   to,
		size: size,
		data: make([]uint8, size),
	}

	return mem
}

func (m *Memory) write(address uint16, value uint8) {
	m.data[address-m.from] = value
}

func (m *Memory) read(address uint16) uint8 {
	return m.data[address-m.from]
}

func (m *Memory) has(address uint16) bool {
	if m.from <= address && address <= m.to {
		return true
	} else {
		return false
	}
}
