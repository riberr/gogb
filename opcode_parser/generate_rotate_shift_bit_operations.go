package main

import "os"

func generateRotateShiftBitOperations(opCodes *OpCodes, output *os.File) {
	var rlca OpCode
	var rla OpCode
	var rrca OpCode
	var rra OpCode
	var cb OpCode

	for i, opCode := range opCodes.Unprefixed {
		opCode.i = i

		switch {
		case i == 0x07:
			rlca = opCode
		case i == 0x17:
			rla = opCode
		case i == 0x0F:
			rrca = opCode
		case i == 0x1F:
			rra = opCode
		case i == 0xCB:
			cb = opCode
		}
	}

	// TODO: add CB-prefixed instructions

	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("    // Rotates, Shifts, Bit operations \n")
	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // RLCA\n")
	writeCode(rlca, output)
	writeCode(rla, output)
	writeCode(rrca, output)
	writeCode(rra, output)

	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Increment (register), Increment SP\n")
	writeCode(cb, output)
	_, _ = output.WriteString("\n")

	// Verify
	var hits = 0
	for _, code := range opCodes.Unprefixed {
		if code.Group == "x8/rsb" {
			hits++
		}

		// We count 0xCB (CB op) as x16/rsb instead of 'control/misc'
		if code.Name == "PREFIX CB" {
			hits++
		}
	}

	if hits != 5 {
		println("generateRotateShiftBitOperations: Not all opcodes covered!")
	} else {
		println("generateRotateShiftBitOperations: OK!")
	}

}
