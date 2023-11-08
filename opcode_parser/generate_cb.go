package main

import (
	"fmt"
	"os"
	"strings"
)

func generateCB(opCodes *OpCodes) {
	output, err := os.Create(OutputPath + "/generated_cb.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	_, _ = output.WriteString("var OpCodesCB = map[uint8]OpCode{\n")

	for i, code := range opCodes.CBPrefixed {
		code.i = i
		switch {
		case strings.HasPrefix(code.Name, "RLC"):
			writeCB(code, output, 4, 3)
		case strings.HasPrefix(code.Name, "RRC"):
			writeCB(code, output, 4, 3)
		case strings.HasPrefix(code.Name, "RL "):
			writeCB(code, output, 3, 2)
		case strings.HasPrefix(code.Name, "RR "):
			writeCB(code, output, 3, 2)
		case strings.HasPrefix(code.Name, "SLA"):
			writeCB(code, output, 4, 3)
		case strings.HasPrefix(code.Name, "SRA"):
			writeCB(code, output, 4, 3)
		case strings.HasPrefix(code.Name, "SWAP"):
			writeCB(code, output, 5, 4)
		case strings.HasPrefix(code.Name, "SRL"):
			writeCB(code, output, 4, 3)
		case strings.HasPrefix(code.Name, "BIT 0"),
			strings.HasPrefix(code.Name, "BIT 1"),
			strings.HasPrefix(code.Name, "BIT 2"),
			strings.HasPrefix(code.Name, "BIT 3"),
			strings.HasPrefix(code.Name, "BIT 4"),
			strings.HasPrefix(code.Name, "BIT 5"),
			strings.HasPrefix(code.Name, "BIT 6"),
			strings.HasPrefix(code.Name, "BIT 7"),
			strings.HasPrefix(code.Name, "RES 0"),
			strings.HasPrefix(code.Name, "RES 1"),
			strings.HasPrefix(code.Name, "RES 2"),
			strings.HasPrefix(code.Name, "RES 3"),
			strings.HasPrefix(code.Name, "RES 4"),
			strings.HasPrefix(code.Name, "RES 5"),
			strings.HasPrefix(code.Name, "RES 6"),
			strings.HasPrefix(code.Name, "RES 7"),
			strings.HasPrefix(code.Name, "SET 0"),
			strings.HasPrefix(code.Name, "SET 1"),
			strings.HasPrefix(code.Name, "SET 2"),
			strings.HasPrefix(code.Name, "SET 3"),
			strings.HasPrefix(code.Name, "SET 4"),
			strings.HasPrefix(code.Name, "SET 5"),
			strings.HasPrefix(code.Name, "SET 6"),
			strings.HasPrefix(code.Name, "SET 7"):
			writeCBBitResSet(code, output, 6, 3, 4)
		}
	}
	_, _ = output.WriteString("\n")
	_, _ = output.WriteString("}\n")
}

func writeCB(code OpCode, output *os.File, regPos int, nameLength int) {
	reg := strings.ToLower(string(code.Name[regPos]))
	name := code.Name[0:nameLength]
	instr := fmt.Sprintf("cpu.regs.%v = %v(cpu.regs.%v)", reg, name, reg)
	writeCodeWithInstruction(code, output, instr)
}

func writeCBBitResSet(code OpCode, output *os.File, regPos int, nameLength int, bitPos int) {
	reg := strings.ToLower(string(code.Name[regPos]))
	bit := string(code.Name[bitPos])
	name := code.Name[0:nameLength]
	instr := fmt.Sprintf("cpu.regs.%v = %v(cpu.regs.%v, %v)", reg, name, reg, bit)
	writeCodeWithInstruction(code, output, instr)
}
