package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Flags struct {
	Z string `json:"Z"`
	N string `json:"N"`
	H string `json:"H"`
	C string `json:"C"`
}

type TimingNoBranch struct {
	Type    string `json:"Type"`
	Comment string `json:"Comment"`
}

type OpCode struct {
	Name            string           `json:"Name"`
	Group           string           `json:"Group"`
	TCyclesBranch   int              `json:"TCyclesBranch"`
	TCyclesNoBranch int              `json:"TCyclesNoBranch"`
	Length          int              `json:"Length"`
	Flags           Flags            `json:"Flags"`
	TimingNoBranch  []TimingNoBranch `json:"TimingNoBranch"`
	i               int
}

type OpCodes struct {
	Unprefixed []OpCode `json:"Unprefixed"`
	CBPrefixed []OpCode `json:"CBPrefixed"`
}

/*
target:

var OpCodes = map[uint8]OpCode{
	0xC3: NewOpCode(0xC3, "JP u16", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
}
*/

/*
inspiration: https://github.com/rvaccarim/FrozenBoy/blob/master/FrozenBoyCore/Processor/Opcode/OpcodeHandler.cs
*/
func main() {
	// Open the JSON file for reading
	file, err := os.Open("opcode_parser/dmgops.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Open a file for writing, creating it if it doesn't exist or truncating it if it does.
	output, err := os.Create("opcode_parser/output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Create a variable to hold the decoded data
	var opCodes OpCodes

	// Decode the JSON data into the data variable
	if err := decoder.Decode(&opCodes); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Parse
	_, _ = output.WriteString("var OpCodesGenerated = map[uint8]OpCode{\n")

	generate8bitLoad(&opCodes, output)
	/*
		generate16bitLoad(&opCodes, output)
		generate8bitArithmetics(&opCodes, output)
		generate16bitArithmetics(&opCodes, output)
		generateRotateShiftBitOperations(&opCodes, output)
		generateControlFlow(&opCodes, output)
		generateMiscellaneous(&opCodes, output)
	*/
	/*
		for i, opCode := range opCodes.Unprefixed {
			_, _ = fmt.Fprintf(output, "    0x%02x: NewOpCode(0x%02x, \"%s\", %d, %d, []func(cpu *CPU){\n",
				i, i, opCode.Name, opCode.Length, opCode.TCyclesBranch)

			switch {
			case i == 0x00: // NOP

			case i == 0x10: // STOP

			case i == 0x20, i == 0x30: // JR NC,i8

			case i == 0x01, i == 0x11, i == 0x21, i == 0x31: // LOAD u16

			case i == 0x02, i == 0x12, i == 0x22, i == 0x32: // LOAD from A

			case i == 0x06, i == 0x16, i == 0x26, i == 0x36: // LD B,u8

			case i == 0x46, i == 0x56, i == 0x66: // LD B,(HL)

			case i == 0x76: //HALT

			case isBetween(0x40, i, 0x7F): // LOAD u8
				from := strings.ToLower(opCode.Name)[3]
				to := strings.ToLower(opCode.Name)[5]
				_, _ = fmt.Fprintf(output, "        func(cpu *CPU) { cpu.regs.%c = cpu.regs.%c },\n", from, to)

			case isBetween(0x80, i, 0xBF): // MATH

				switch {
				case isBetween(0x80, i, 0x85): // ADD
					from := strings.ToLower(opCode.Name)[6]
					_, _ = fmt.Fprintf(output, "        func(cpu *CPU) { add(cpu.regs.%c) },\n", from)
				default:
					_, _ = fmt.Fprintf(output, "        func(cpu *CPU) {  MATH  },\n")
				}

			case opCode.Name == "UNUSED":
				_, _ = fmt.Fprintf(output, "    }\n")
			}

			_, _ = output.WriteString("    }),\n")
		}
	*/

	_, _ = output.WriteString("}\n")
}

func isBetween(low int, value int, high int) bool {
	return value >= low && value <= high
}

func writeCode(opCode OpCode, output *os.File) {
	_, _ = fmt.Fprintf(output, "    0x%02x: NewOpCode(0x%02x, \"%s\", %d, %d, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),\n",
		opCode.i, opCode.i, opCode.Name, opCode.Length, opCode.TCyclesBranch)
}

func writeCodeWithInstruction(opCode OpCode, output *os.File, instruction string) {
	_, _ = fmt.Fprintf(output, "    0x%02x: NewOpCode(0x%02x, \"%s\", %d, %d, []func(cpu *CPU){func(cpu *CPU) {%v}}),\n",
		opCode.i, opCode.i, opCode.Name, opCode.Length, opCode.TCyclesBranch, instruction)
}

func writeCodeWithMultipleInstructions(opCode OpCode, output *os.File, instruction string) {
	_, _ = fmt.Fprintf(output, "    0x%02x: NewOpCode(0x%02x, \"%s\", %d, %d, []func(cpu *CPU){%v}),\n",
		opCode.i, opCode.i, opCode.Name, opCode.Length, opCode.TCyclesBranch, instruction)
}
