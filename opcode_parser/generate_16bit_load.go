package main

import (
	"fmt"
	"os"
)

func generate16bitLoad(opCodes *OpCodes) {
	var load16bitReg []OpCode // Load 16-bit register / register pair
	var stackVarious []OpCode
	var pushToStack []OpCode  // Push to stack
	var popFromStack []OpCode // Pop from stack

	output, err := os.Create(OutputPath + "/generated_16bit_load.go")
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
		case x == 0x01 && isBetween(0x00, y, 0x03):
			load16bitReg = append(load16bitReg, opCode)
		case i == 0x08, i == 0xF9, i == 0xF8:
			stackVarious = append(stackVarious, opCode)

		case x == 0x05 && isBetween(0xC, y, 0xF):
			pushToStack = append(pushToStack, opCode)
		case x == 0x01 && isBetween(0xC, y, 0xF):
			popFromStack = append(popFromStack, opCode)

		}
	}

	_, _ = output.WriteString("package cpu\n\n")
	//_, _ = output.WriteString("import \"gogb/utils\"\n")
	//_, _ = output.WriteString("import \"gogb/emulator/memory\"\n\n")

	_, _ = output.WriteString("var OpCodes16bitLoadGenerated = map[uint8]OpCode{\n")

	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("    // 16BIT LOAD\n")
	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Load 16-bit register / register pair\n")
	for _, opCode := range load16bitReg {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Various stack\n")
	for _, opCode := range stackVarious {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Push to stack\n")
	for _, opCode := range pushToStack {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Pop from stack\n")
	for _, opCode := range popFromStack {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("\n")
	_, _ = output.WriteString("}\n")

	// Verify

	var hits = 0
	for _, code := range opCodes.Unprefixed {
		if code.Group == "x16/lsm" {
			hits++
		}

		// gekkio doc says this opcode is x16/lsm. gbops says it is x16/alu
		if code.Name == "LD HL,SP+i8" {
			hits++
		}
	}

	if hits != (len(load16bitReg) + len(stackVarious) + len(pushToStack) + len(popFromStack)) {
		println("generate16bitLoad: Not all opcodes covered!")
	} else {
		println("generate16bitLoad: OK!")
	}

}
