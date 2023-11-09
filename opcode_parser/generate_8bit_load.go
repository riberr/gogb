package main

import (
	"fmt"
	"os"
	"strings"
)

func generate8bitLoad(opCodes *OpCodes /*, output *os.File*/) {
	var loadRegFromReg []OpCode        // Load register (register)
	var loadRegFromBus []OpCode        // Load register (immediate)
	var loadRegFromIndirectHL []OpCode // Load register (indirect HL)
	var loadFromRegIndirectHL []OpCode // Load from register (indirect HL)
	var various []OpCode               // one offs...

	output, err := os.Create(OutputPath + "/generated_8bit_load.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	_, _ = output.WriteString("package cpu\n\n")
	_, _ = output.WriteString("import \"gogb/utils\"\n")
	_, _ = output.WriteString("import \"gogb/gameboy/bus\"\n\n")

	_, _ = output.WriteString("var GeneratedOpCodes8bitLoadGenerated = map[uint8]OpCode{\n")

	for i, opCode := range opCodes.Unprefixed {
		opCode.i = i
		x := i & 0x0F
		y := i & 0xF0 >> 4
		//fmt.Printf("0x%02x, 0x%02x\n", x, y)

		switch {

		case i == 0x36, i == 0x0A, i == 0x1A, i == 0x02, i == 0x12, i == 0xFA, i == 0xEA, i == 0xF2, i == 0xE2, i == 0xF0, i == 0xE0, i == 0x3A, i == 0x32, i == 0x2A, i == 0x22:
			various = append(various, opCode)
		case isBetween(0x70, i, 0x75), i == 0x77:
			loadFromRegIndirectHL = append(loadFromRegIndirectHL, opCode)
		case x == 0x06 && isBetween(0x04, y, 0x06), x == 0xE && isBetween(0x4, y, 0x7):
			loadRegFromIndirectHL = append(loadRegFromIndirectHL, opCode)
		case isBetween(0x40, i, 0x6F), isBetween(0x78, i, 0x7D), i == 0x7F:
			if i == 0x76 {
				// skip HALT
				continue
			}
			loadRegFromReg = append(loadRegFromReg, opCode)
		case x == 0x06 && isBetween(0x00, y, 0x03), x == 0x0E && isBetween(0x00, y, 0x04):
			loadRegFromBus = append(loadRegFromBus, opCode)
		}
	}

	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("    // 8BIT LOAD\n")
	_, _ = output.WriteString("    // ~~~~~~~~~~~~~~~~~~~~~~~~~~\n")
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    //  Load register (register)\n")
	for _, opCode := range loadRegFromReg {
		to := strings.ToLower(string(opCode.Name[3]))
		from := strings.ToLower(string(opCode.Name[5]))
		instr := fmt.Sprintf("cpu.regs.%v = cpu.regs.%v", to, from)
		writeCodeWithInstruction(opCode, output, instr)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Load register (immediate)\n")
	for _, opCode := range loadRegFromBus {
		// cpu.regs.b = bus.BusRead(cpu.pc); cpu.pc++
		to := strings.ToLower(string(opCode.Name[3]))
		instr := fmt.Sprintf("cpu.regs.%v = bus.BusRead(cpu.pc); cpu.pc++", to)
		writeCodeWithInstruction(opCode, output, instr)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Load register (indirect HL)\n")
	for _, opCode := range loadRegFromIndirectHL {
		// cpu.regs.b = bus.BusRead(cpu.regs.getHL())
		to := strings.ToLower(string(opCode.Name[3]))
		instr := fmt.Sprintf("cpu.regs.%v = bus.BusRead(cpu.regs.getHL())", to)
		writeCodeWithInstruction(opCode, output, instr)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Load from register (indirect HL)\n")
	for _, opCode := range loadFromRegIndirectHL {
		// bus.busWrite(cpu.regs.getHL(), cpu.regs.b)
		from := strings.ToLower(string(opCode.Name[8]))
		instr := fmt.Sprintf("bus.BusWrite(cpu.regs.getHL(), cpu.regs.%v)", from)
		writeCodeWithInstruction(opCode, output, instr)
	}
	_, _ = output.WriteString("\n")

	_, _ = output.WriteString("    // Various 8bit loads\n")
	for _, opCode := range various {
		//writeCode(opCode, output)
		switch opCode.Name {
		case "LD (BC),A":
			writeCodeWithInstruction(opCode, output, "bus.BusWrite(cpu.regs.getBC(), cpu.regs.a)")
		case "LD A,(BC)":
			writeCodeWithInstruction(opCode, output, "cpu.regs.a = bus.BusRead(cpu.regs.getBC())")
		case "LD (DE),A":
			writeCodeWithInstruction(opCode, output, "bus.BusWrite(cpu.regs.getDE(), cpu.regs.a)")
		case "LD A,(DE)":
			writeCodeWithInstruction(opCode, output, "cpu.regs.a = bus.BusRead(cpu.regs.getDE())")
		case "LD (HL+),A":
			writeCodeWithInstruction(opCode, output, "bus.BusWrite(cpu.regs.getHL(), cpu.regs.a); cpu.regs.incHL()")
		case "LD A,(HL+)":
			writeCodeWithInstruction(opCode, output, "cpu.regs.a = bus.BusRead(cpu.regs.getHL()); cpu.regs.incHL()")
		case "LD (HL-),A":
			writeCodeWithInstruction(opCode, output, "bus.BusWrite(cpu.regs.getHL(), cpu.regs.a); cpu.regs.decHL()")
		case "LD (HL),u8":
			writeCodeWithMultipleInstructions(opCode, output, "func(cpu *CPU) { lsb = bus.BusRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) {bus.BusWrite(cpu.regs.getHL(), lsb)}")
		case "LD A,(HL-)":
			writeCodeWithInstruction(opCode, output, "cpu.regs.a = bus.BusRead(cpu.regs.getHL()); cpu.pc--")
		case "LD (FF00+u8),A":
			writeCodeWithMultipleInstructions(opCode, output, "func(cpu *CPU) { lsb = bus.BusRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { bus.BusWrite(utils.ToUint16(lsb, 0xFF), cpu.regs.a) }")
		case "LD (FF00+C),A":
			writeCodeWithInstruction(opCode, output, "bus.BusWrite(utils.ToUint16(cpu.regs.c, 0xFF), cpu.regs.a)")
		case "LD (u16),A":
			writeCodeWithMultipleInstructions(opCode, output, "func(cpu *CPU) { lsb = bus.BusRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { msb = bus.BusRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { bus.BusWrite(utils.ToUint16(lsb, msb), cpu.regs.a) }")
		case "LD A,(FF00+u8)":
			writeCodeWithMultipleInstructions(opCode, output, "func(cpu *CPU) { lsb = bus.BusRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) {cpu.regs.a = bus.BusRead(utils.ToUint16(lsb, 0xFF))}")
		case "LD A,(FF00+C)":
			writeCodeWithInstruction(opCode, output, "cpu.regs.a = bus.BusRead(utils.ToUint16(cpu.regs.c, 0xFF))")
		case "LD A,(u16)":
			writeCodeWithMultipleInstructions(opCode, output, "func(cpu *CPU) {lsb = bus.BusRead(cpu.pc); cpu.pc++}, func(cpu *CPU) {msb = bus.BusRead(cpu.pc); cpu.pc++}, func(cpu *CPU) {cpu.regs.a = bus.BusRead(utils.ToUint16(lsb, msb))}")
		}
	}
	_, _ = output.WriteString("\n")
	_, _ = output.WriteString("}\n")

	// Verify

	var hits = 0
	for _, code := range opCodes.Unprefixed {
		if code.Group == "x8/lsm" {
			hits++
		}
	}

	if hits != (len(loadRegFromReg) + len(loadRegFromBus) + len(loadRegFromIndirectHL) + len(loadFromRegIndirectHL) + len(various)) {
		println("generate8bitLoad: Not all opcodes covered!")
	} else {
		println("generate8bitLoad: OK!")
	}
}
