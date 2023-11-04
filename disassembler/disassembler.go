package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("roms/dmg_boot.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	br := bufio.NewReader(file)

	// infinite loop
	for {

		byte, err := br.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		// process the one byte b
		//fmt.Printf("%02x ", byte)
		switch byte {
		case 0x21:
			lo, _ := br.ReadByte()
			hi, _ := br.ReadByte()
			fmt.Printf("LD HL,$%02x%02x\n", hi, lo)
		case 0x31:
			lo, _ := br.ReadByte()
			hi, _ := br.ReadByte()
			fmt.Printf("LD SP,$%02x%02x\n", hi, lo)
		case 0x32:
			fmt.Printf("LD (HL-),A\n")

		case 0xAF:
			fmt.Printf("XOR A,A\n")
		default:
			fmt.Printf("Unhandled: %02x\n", byte)
			return
		}
		if err != nil {
			// end of file
			break
		}
	}
}
