package bus

type Space struct {
	from uint16
	to   uint16
	size uint16
	data []uint8
}

func NewSpace(from uint16, to uint16) Space {
	size := to - from + 1
	mem := Space{
		from: from,
		to:   to,
		size: size,
		data: make([]uint8, size),
	}

	return mem
}

func (m *Space) write(address uint16, value uint8) {
	m.data[address-m.from] = value
}

func (m *Space) read(address uint16) uint8 {
	return m.data[address-m.from]
}

func (m *Space) has(address uint16) bool {
	if m.from <= address && address <= m.to {
		return true
	} else {
		return false
	}
}
