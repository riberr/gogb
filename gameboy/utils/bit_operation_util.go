package utils

// SetBit sets the bit at pos in the integer n.
func SetBit(n uint8, pos int) uint8 {
	n |= 1 << pos
	return n
}

// ClearBit clears the bit at pos in n.
func ClearBit(n uint8, pos int) uint8 {
	mask := uint8(^(1 << pos))
	n &= mask
	return n
}

// TestBit checks whether a bit is set
func TestBit(n uint8, pos int) bool {
	val := n & (1 << pos)
	return val > 0
}

// TestBit checks whether a bit is set
func TestBit16(n uint16, pos int) bool {
	val := n & (1 << pos)
	return val > 0
}

func BitValue(value uint8, bit uint8) uint8 {
	return (value >> bit) & 1
}

func ToUint16(lsb uint8, msb uint8) uint16 {
	return uint16(lsb) | (uint16(msb) << 8)
}

func Msb(value uint16) uint8 {
	return uint8((value & 0xFF00) >> 8)
}

func Lsb(value uint16) uint8 {
	return uint8(value & 0x00FF)
}

// ToInt transforms a bool into a 1/0 value.
func ToInt(val bool) uint8 {
	if val {
		return 1
	}
	return 0
}
