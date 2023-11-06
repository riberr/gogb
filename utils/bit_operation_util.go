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

// HasBit checks whether a bit is set
func HasBit(n uint8, pos int) bool {
	val := n & (1 << pos)
	return val > 0
}

func ToUint16(lsb uint8, msb uint8) uint16 {
	return uint16(lsb) | (uint16(msb) << 8)
}
