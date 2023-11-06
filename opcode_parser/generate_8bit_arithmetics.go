package main

import "os"

func generate8bitArithmetics(opCodes *OpCodes, output *os.File) {
	var add []OpCode     // Add (register), Add (indirect HL), Add (immediate)
	var adc []OpCode     // Add with carry (register), Add with carry (indirect HL), Add with carry (immediate)
	var sub []OpCode     // Subtract (register), Subtract (indirect HL), Subtract (immediate)
	var sbc []OpCode     // Subtract with carry (register),  Subtract with carry (indirect HL), Subtract with carry (immediate)
	var cp []OpCode      // Compare (register), Compare (indirect HL), Compare (immediate)
	var inc []OpCode     // Increment (register),Increment (indirect HL),
	var dec []OpCode     // Decrement (register), Decrement (indirect HL)
	var and []OpCode     // Bitwise AND (register), Bitwise AND (indirect HL), Bitwise AND (immediate)
	var or []OpCode      // Bitwise OR (register),  Bitwise OR (indirect HL), Bitwise OR (immediate)
	var xor []OpCode     // Bitwise XOR (register),  Bitwise XOR (indirect HL), Bitwise XOR (immediate)
	var various []OpCode // CCF: Complement carry flag, SCF: Set carry flag, DAA: Decimal adjust accumulator, CPL: Complement accumulator

	for i, opCode := range opCodes.Unprefixed {
		opCode.i = i
		x := i & 0x0F
		y := i & 0xF0 >> 4
		//fmt.Printf("0x%02x, 0x%02x\n", x, y)

		switch {
		case isBetween(0x80, i, 0x87), i == 0xC6:
			add = append(add, opCode)
		case isBetween(0x88, i, 0x8F), i == 0xCE:
			adc = append(adc, opCode)
		case isBetween(0x90, i, 0x97), i == 0xD6:
			sub = append(sub, opCode)
		case isBetween(0x98, i, 0x9F), i == 0xDE:
			sbc = append(sbc, opCode)
		case isBetween(0xB8, i, 0xBF), i == 0xFE:
			cp = append(cp, opCode)
		case x == 0x4 && isBetween(0x0, y, 0x3), x == 0xC && isBetween(0x0, y, 0x3):
			inc = append(inc, opCode)
		case x == 0x5 && isBetween(0x0, y, 0x3), x == 0xD && isBetween(0x0, y, 0x3):
			dec = append(dec, opCode)
		case isBetween(0xA0, i, 0xA7), i == 0xE6:
			and = append(and, opCode)
		case isBetween(0xB0, i, 0xB7), i == 0xF6:
			or = append(or, opCode)
		case isBetween(0xA8, i, 0xAF), i == 0xEE:
			xor = append(xor, opCode)
		case i == 0x3F, i == 0x37, i == 0x27, i == 0x2F:
			various = append(various, opCode)
		}
	}

	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("    // 8BIT ARITHMETICS & LOGICAL\n")
	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Add (register), Add (indirect HL), Add (immediate)\n")
	for _, opCode := range add {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Add with carry (register), Add with carry (indirect HL), Add with carry (immediate)\n")
	for _, opCode := range adc {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Subtract (register), Subtract (indirect HL), Subtract (immediate)\n")
	for _, opCode := range sub {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Subtract with carry (register),  Subtract with carry (indirect HL), Subtract with carry (immediate)\n")
	for _, opCode := range sbc {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Compare (register), Compare (indirect HL), Compare (immediate)\n")
	for _, opCode := range cp {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Increment (register),Increment (indirect HL)\n")
	for _, opCode := range inc {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Decrement (register), Decrement (indirect HL)\n")
	for _, opCode := range dec {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Bitwise AND (register), Bitwise AND (indirect HL), Bitwise AND (immediate)\n")
	for _, opCode := range and {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Bitwise OR (register),  Bitwise OR (indirect HL), Bitwise OR (immediate)\n")
	for _, opCode := range or {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Bitwise XOR (register),  Bitwise XOR (indirect HL), Bitwise XOR (immediate)\n")
	for _, opCode := range xor {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // CCF: Complement carry flag, SCF: Set carry flag, DAA: Decimal adjust accumulator, CPL: Complement accumulator\n")
	for _, opCode := range various {
		writeCode(opCode, output)
	}
	_, _ = output.WriteString("\n")

	// Verify
	var hits = 0
	for _, code := range opCodes.Unprefixed {
		if code.Group == "x8/alu" {
			hits++
		}
	}

	if hits != (len(add) + len(adc) + len(sub) + len(sbc) + len(cp) + len(inc) + len(dec) + len(and) + len(or) + len(xor) + len(various)) {
		println("generate8bitArithmetics: Not all opcodes covered!")
	} else {
		println("generate8bitArithmetics: OK!")
	}

}
