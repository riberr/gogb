package main

import (
	"fmt"
	"os"
)

func generateMiscellaneous(opCodes *OpCodes) {
	var halt OpCode
	var stop OpCode
	var di OpCode
	var ei OpCode
	var nop OpCode

	output, err := os.Create(OutputPath + "/generated_misc.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	for i, opCode := range opCodes.Unprefixed {
		opCode.i = i

		switch {
		case i == 0x76:
			halt = opCode
		case i == 0x10:
			stop = opCode
		case i == 0xF3:
			di = opCode
		case i == 0xFB:
			ei = opCode
		case i == 0x00:
			nop = opCode
		}

	}

	_, _ = output.WriteString("package cpu\n\n")
	//_, _ = output.WriteString("import \"gogb/utils\"\n")
	//_, _ = output.WriteString("import \"gogb/emulator/memory\"\n\n")

	_, _ = output.WriteString("var OpCodesMisc = map[uint8]OpCode{\n")

	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("    // Miscellaneous \n")
	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Halt system clock\n")
	writeCode(halt, output)
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Stop system and main clocks\n")
	writeCode(stop, output)
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Disable interrupts\n")
	writeCode(di, output)
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Enable interrupts\n")
	writeCode(ei, output)
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // No operation\n")
	writeCode(nop, output)
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("\n")
	_, _ = output.WriteString("}\n")

	// Verify
	var hits = 0
	for _, code := range opCodes.Unprefixed {
		if code.Name == "PREFIX CB" {
			// ignore as we count this one as 'Rotates, shifts, and bit operations' / 'x8/rsb'
		} else if code.Group == "control/misc" {
			hits++
		}
	}

	if hits != 5 {
		println("generateMiscellaneous: Not all opcodes covered!")
	} else {
		println("generateMiscellaneous: OK!")
	}

}
