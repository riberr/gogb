package mbc

type MBC interface {
	Read(address uint16) uint8
	WriteRom(address uint16, value uint8)
	WriteRam(address uint16, value uint8)
}
