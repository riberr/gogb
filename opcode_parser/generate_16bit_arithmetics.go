package main

import (
	"fmt"
	"os"
)

func generate16bitArithmetics(opCodes *OpCodes) {
	var add []OpCode // Add (register?), Add SP?
	var inc []OpCode // Increment (register), Increment SP
	var dec []OpCode // Decrement (register), Decrement SP

	output, err := os.Create(OutputPath + "/generated_16bit_arithmetics.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	for i, opCode := range opCodes.Unprefixed {
		opCode.i = i
		x := i & 0x0F
		y := i & 0xF0 >> 4
		//fmt.Printf("0x%02x, 0x%02x\n", x, y)

		switch {
		case x == 0x9 && isBetween(0x0, y, 0x3), i == 0xE8:
			add = append(add, opCode)
		case x == 0x3 && isBetween(0x0, y, 0x3):
			inc = append(inc, opCode)
		case x == 0xB && isBetween(0x0, y, 0x3):
			dec = append(dec, opCode)
		}
	}

	_, _ = output.WriteString("package cpu\n\n")
	//_, _ = output.WriteString("import \"gogb/utils\"\n")
	//_, _ = output.WriteString("import \"gogb/emulator/memory\"\n\n")

	_, _ = output.WriteString("var GeneratedOpCodes16bitArithmeticsGenerated = map[uint8]OpCode{\n")

	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("    // 16BIT ARITHMETICS & LOGICAL\n")
	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Add (register?), Add SP?\n")
	for _, opCode := range add {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Increment (register), Increment SP\n")
	for _, opCode := range inc {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Decrement (register), Decrement SP\n")
	for _, opCode := range dec {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("\n")
	_, _ = output.WriteString("}\n")

	// Verify
	var hits = 0
	for _, code := range opCodes.Unprefixed {
		if code.Group == "x16/alu" {
			hits++
		}

		// We count 0xF8 (LD HL,SP+i8) as x16/lsm instead of x16/alu
		if code.Name == "LD HL,SP+i8" {
			hits--
		}
	}

	if hits != (len(add) + len(inc) + len(dec)) {
		println("generate16bitArithmetics: Not all opcodes covered!")
	} else {
		println("generate16bitArithmetics: OK!")
	}

}
