package memory

import (
	"fmt"
	"gogb/utils"
	"os"
)

var ROM_TYPES = [...]string{
	"ROM ONLY",
	"MBC1",
	"MBC1+RAM",
	"MBC1+RAM+BATTERY",
	"0x04 ???",
	"MBC2",
	"MBC2+BATTERY",
	"0x07 ???",
	"ROM+RAM 1",
	"ROM+RAM+BATTERY 1",
	"0x0A ???",
	"MMM01",
	"MMM01+RAM",
	"MMM01+RAM+BATTERY",
	"0x0E ???",
	"MBC3+TIMER+BATTERY",
	"MBC3+TIMER+RAM+BATTERY 2",
	"MBC3",
	"MBC3+RAM 2",
	"MBC3+RAM+BATTERY 2",
	"0x14 ???",
	"0x15 ???",
	"0x16 ???",
	"0x17 ???",
	"0x18 ???",
	"MBC5",
	"MBC5+RAM",
	"MBC5+RAM+BATTERY",
	"MBC5+RUMBLE",
	"MBC5+RUMBLE+RAM",
	"MBC5+RUMBLE+RAM+BATTERY",
	"0x1F ???",
	"MBC6",
	"0x21 ???",
	"MBC7+SENSOR+RUMBLE+RAM+BATTERY",
}

var LIC_CODE = map[uint16]string{
	0x00: "None",
	0x01: "Nintendo R&D1",
	0x08: "Capcom",
	0x13: "Electronic Arts",
	0x18: "Hudson Soft",
	0x19: "b-ai",
	0x20: "kss",
	0x22: "pow",
	0x24: "PCM Complete",
	0x25: "san-x",
	0x28: "Kemco Japan",
	0x29: "seta",
	0x30: "Viacom",
	0x31: "Nintendo",
	0x32: "Bandai",
	0x33: "Ocean/Acclaim",
	0x34: "Konami",
	0x35: "Hector",
	0x37: "Taito",
	0x38: "Hudson",
	0x39: "Banpresto",
	0x41: "Ubi Soft",
	0x42: "Atlus",
	0x44: "Malibu",
	0x46: "angel",
	0x47: "Bullet-Proof",
	0x49: "irem",
	0x50: "Absolute",
	0x51: "Acclaim",
	0x52: "Activision",
	0x53: "American sammy",
	0x54: "Konami",
	0x55: "Hi tech entertainment",
	0x56: "LJN",
	0x57: "Matchbox",
	0x58: "Mattel",
	0x59: "Milton Bradley",
	0x60: "Titus",
	0x61: "Virgin",
	0x64: "LucasArts",
	0x67: "Ocean",
	0x69: "Electronic Arts",
	0x70: "Infogrames",
	0x71: "Interplay",
	0x72: "Broderbund",
	0x73: "sculptured",
	0x75: "sci",
	0x78: "THQ",
	0x79: "Accolade",
	0x80: "misawa",
	0x83: "lozc",
	0x86: "Tokuma Shoten Intermedia",
	0x87: "Tsukuda Original",
	0x91: "Chunsoft",
	0x92: "Video system",
	0x93: "Ocean/Acclaim",
	0x95: "Varie",
	0x96: "Yonezawa/s’pal",
	0x97: "Kaneko",
	0x99: "Pack in soft",
	0xA4: "Konami (Yu-Gi-Oh!)",
}

var RAM_SIZE = map[uint8]string{
	0x00: "0",
	0x01: "-",
	0x02: "8 KiB",
	0x03: "32 KiB",
	0x04: "128 KiB",
	0x05: "64 KiB",
}

type Cart struct {
	fileName string
	size     int
	data     []uint8
	header   header
}

type header struct {
	entry          [4]uint8
	logo           [0x30]uint8
	title          [16]uint8
	cgbFlag        uint8
	newLicCode     [2]uint8
	sgbFlag        uint8
	cartType       uint8
	romSize        uint8
	ramSize        uint8
	destCode       uint8
	licCode        uint8
	version        uint8
	checksum       uint8
	globalChecksum [2]uint8
}

func newHeader(data []uint8) header {
	return header{
		entry:          ([4]uint8)(data[0x100 : 0x103+1]),
		logo:           ([0x30]uint8)(data[0x104 : 0x133+1]),
		title:          ([16]uint8)(data[0x134 : 0x143+1]),
		cgbFlag:        data[0x143+1],
		newLicCode:     ([2]uint8)(data[0x144 : 0x145+1]),
		sgbFlag:        data[0x146],
		cartType:       data[0x147],
		romSize:        data[0x148],
		ramSize:        data[0x149],
		destCode:       data[0x14A],
		licCode:        data[0x14B],
		version:        data[0x14C],
		checksum:       data[0x14D],
		globalChecksum: ([2]uint8)(data[0x14E : 0x14F+1]),
	}
}

var cart = Cart{}

func CartLoad(romPath string, romName string) bool {
	// Read the file into a byte array
	rom, err := os.ReadFile(romPath + romName)
	if err != nil {
		fmt.Printf("Failed to open: %s\n", romPath+romName)
		return false
	}

	cart.fileName = romName
	fmt.Printf("Opened: %s\n", cart.fileName)
	cart.size = len(rom)
	cart.data = rom

	cart.header = newHeader(cart.data)

	fmt.Printf("Cartridge Loaded:\n")
	fmt.Printf("\t Title    : %s\n", cart.header.title)
	fmt.Printf("\t Type     : %2.2X (%s)\n", cart.header.cartType, ROM_TYPES[cart.header.cartType])
	fmt.Printf("\t ROM Size : %d KB\n", 32<<cart.header.romSize)
	fmt.Printf("\t RAM Size : %2.2X (%s)\n", cart.header.ramSize, RAM_SIZE[cart.header.ramSize])
	fmt.Printf("\t LIC Code : %2.2X (%s)\n", cart.header.newLicCode, LIC_CODE[utils.ToUint16(cart.header.newLicCode[0], cart.header.newLicCode[1])])
	fmt.Printf("\t LIC Code : %2.2X (%s)\n", cart.header.licCode, LIC_CODE[uint16(cart.header.licCode)])
	fmt.Printf("\t ROM Vers : %2.2X\n", cart.header.version)

	var x uint8 = 0
	for i := 0x134; i <= 0x14C; i++ {
		x = x - cart.data[i] - 1
	}

	if x&0xFF != 0 {
		fmt.Printf("\t Checksum : %2.2X (%s)\n", cart.header.checksum, "PASSED")
	} else {
		fmt.Printf("\t Checksum : %2.2X (%s)\n", cart.header.checksum, "FAILED")
	}

	return true
}

func cartRead(address uint16) uint8 {
	return cart.data[address]
}

func cartWrite(address uint16, value uint8) {
	//todo
}
