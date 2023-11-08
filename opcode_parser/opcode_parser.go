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

const OutputPath = "generated/"

func main() {
	// Open the JSON file for reading
	file, err := os.Open("opcode_parser/dmgops.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Open a file for writing, creating it if it doesn't exist or truncating it if it does.
	/*
		output, err := os.Create("opcode_parser/output/output.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer output.Close()
	*/

	// Create a JSON decoder
	decoder := json.NewDecoder(file)

	// Create a variable to hold the decoded data
	var opCodes OpCodes

	// Decode the JSON data into the data variable
	if err := decoder.Decode(&opCodes); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	generate8bitLoad(&opCodes /*, output*/)
	generate16bitLoad(&opCodes)
	generate8bitArithmetics(&opCodes)
	generate16bitArithmetics(&opCodes)
	generateRotateShiftBitOperations(&opCodes)
	//generateControlFlow(&opCodes)		// using non-generated code. Don't override!
	generateMiscellaneous(&opCodes)

	generateCB(&opCodes)

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
