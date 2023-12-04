package utils

type Space struct {
	Start uint16
	End   uint16
	size  uint16
	data  []uint8
}

func NewSpace(from uint16, to uint16) Space {
	size := to - from + 1
	mem := Space{
		Start: from,
		End:   to,
		size:  size,
		data:  make([]uint8, size),
	}

	return mem
}

func (m *Space) WriteDirect(address uint16, value uint8) {
	m.data[address] = value
}

func (m *Space) Write(address uint16, value uint8) {
	m.data[address-m.Start] = value
}

func (m *Space) Read(address uint16) uint8 {
	return m.data[address-m.Start]
}

func (m *Space) ReadIndex(index uint16) uint8 {
	return m.data[index]
}

func (m *Space) Has(address uint16) bool {
	if m.Start <= address && address <= m.End {
		return true
	} else {
		return false
	}
}
