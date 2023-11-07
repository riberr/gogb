package main

import (
	"fmt"
	"os"
)

func generateControlFlow(opCodes *OpCodes) {
	var variousJump []OpCode             // Add (register?), Add SP?
	var jumpConditional []OpCode         // Jump (conditional)
	var relativeJumpConditional []OpCode // Relative jump (conditional)
	var variousCall []OpCode             // Call function
	var variousRet []OpCode              // Return from function, Return from function (conditional), Return from interrupt handler
	var restart []OpCode                 // Restart / Call function (implied)

	output, err := os.Create(OutputPath + "/generated_control_flow.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	for i, opCode := range opCodes.Unprefixed {
		opCode.i = i

		switch {
		case i == 0xC3, i == 0xE9, i == 0x18:
			variousJump = append(variousJump, opCode)
		case i == 0xC2, i == 0xD2, i == 0xCA, i == 0xDA:
			jumpConditional = append(jumpConditional, opCode)
		case i == 0x20, i == 0x30, i == 0x28, i == 0x38:
			relativeJumpConditional = append(relativeJumpConditional, opCode)
		case i == 0xCD, i == 0xC4, i == 0xD4, i == 0xCC, i == 0xDC:
			variousCall = append(variousCall, opCode)
		case i == 0xC9, i == 0xC0, i == 0xD0, i == 0xC8, i == 0xD8, i == 0xD9:
			variousRet = append(variousRet, opCode)
		case i == 0xC7, i == 0xD7, i == 0xE7, i == 0xF7, i == 0xCF, i == 0xDF, i == 0xEF, i == 0xFF:
			restart = append(restart, opCode)
		}

	}

	_, _ = output.WriteString("package cpu\n\n")
	//_, _ = output.WriteString("import \"gogb/utils\"\n")
	//_, _ = output.WriteString("import \"gogb/emulator/memory\"\n\n")

	_, _ = output.WriteString("var OpCodesControlFlow = map[uint8]OpCode{\n")

	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("    // Control flow \n")
	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Jump, Jump to HL, Relative jump\n")
	for _, opCode := range variousJump {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Jump (conditional) \n")
	for _, opCode := range jumpConditional {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Relative jump (conditional) \n")
	for _, opCode := range relativeJumpConditional {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Call function,  Call function (conditional) \n")
	for _, opCode := range variousCall {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Return from function, Return from function (conditional), Return from interrupt handler \n")
	for _, opCode := range variousRet {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Restart / Call function (implied) \n")
	for _, opCode := range restart {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("\n")
	_, _ = output.WriteString("}\n")

	// Verify
	var hits = 0
	for _, code := range opCodes.Unprefixed {
		if code.Group == "control/br" {
			hits++
		}

	}

	if hits != (len(variousJump) + len(jumpConditional) + len(relativeJumpConditional) + len(variousCall) + len(variousRet) + len(restart)) {
		println("generateControlFlow: Not all opcodes covered!")
	} else {
		println("generateControlFlow: OK!")
	}

}
