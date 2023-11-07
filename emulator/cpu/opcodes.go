package cpu

type OpCode struct {
	value   uint8            // for instance 0xC3
	label   string           // JP {0:x4}
	length  int              // in bytes
	tCycles int              // clock cycles
	mCycles int              // machine cycles
	steps   []func(cpu *CPU) // function array
}

func NewOpCode(value uint8, label string, length int, tCycles int, steps []func(cpu *CPU)) OpCode {
	return OpCode{
		value:   value,
		label:   label,
		length:  length,
		tCycles: tCycles,
		mCycles: tCycles / 4,
		steps:   steps,
	}
}

func add(cpu *CPU, value uint8) {
	result := int(cpu.regs.a) + int(value) // cast to int to be able to detect overflow
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, uint8(result) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (cpu.regs.a&0xF)+(value+0xF) > 0xF)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result>>8) != 0)
	cpu.regs.a = uint8(result)
}

func adc(cpu *CPU, value uint8) {
	carry := uint8(0)
	if cpu.regs.getFlag(FLAG_CARRY_C_BIT) {
		carry = 1
	}
	result := int(cpu.regs.a) + int(carry) + int(value) // cast to int to be able to detect overflow
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, uint8(result) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (cpu.regs.a&0xF)+(value+0xF)+carry > 0xF)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result>>8) != 0)
	cpu.regs.a = uint8(result)
}

func sub(cpu *CPU, value uint8) {
	result := int(cpu.regs.a) - int(value) // cast to int to be able to detect overflow
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, uint8(result) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (value+0xF) > (cpu.regs.a&0xF))
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result>>8) != 0)
	cpu.regs.a = uint8(result)
}

func sbc(cpu *CPU, value uint8) {
	carry := uint8(0)
	if cpu.regs.getFlag(FLAG_CARRY_C_BIT) {
		carry = 1
	}
	result := int(cpu.regs.a) - int(value) - int(carry) // cast to int to be able to detect overflow
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, uint8(result) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, cpu.regs.a^value^uint8(result&0xFF)&(1<<4) != 0)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, result < 0)
	cpu.regs.a = uint8(result)
}

func cp(cpu *CPU, value uint8) {
	result := int(cpu.regs.a) - int(value)
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (0xF&value) > (0xF&cpu.regs.a))
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result>>8) != 0)
}

func inc(cpu *CPU, value uint8) uint8 {
	result := value + 1
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (0xF&value) == 0xF)
	// flag c unmodified
	return result
}

func dec(cpu *CPU, value uint8) uint8 {
	result := value - 1
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (0xF&value) == 0x00)
	// flag c unmodified
	return result
}

func and(cpu *CPU, value uint8) {
	result := cpu.regs.a & value
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, true)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, false)
	cpu.regs.a = result
}

func or(cpu *CPU, value uint8) {
	result := cpu.regs.a | value
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, false)
	cpu.regs.a = result
}

func xor(cpu *CPU, value uint8) {
	result := cpu.regs.a ^ value
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, false)
	cpu.regs.a = result
}

var lsb, msb, e uint8 = 0, 0, 0 // temp values when executing opcodes
var stop = false                // used for conditional jumps

//var OpCodes = map[uint8]OpCode{}

var OpCodes = merge(
	OpCodes8bitLoadGenerated,
	OpCodes16bitLoadGenerated,
	OpCodes8bitArithmeticsGenerated,
	OpCodes16bitArithmeticsGenerated,
	OpCodesRotateShiftBitoperations,
	OpCodesControlFlow,
	OpCodesMisc,
)

func merge(ms ...map[uint8]OpCode) map[uint8]OpCode {
	res := map[uint8]OpCode{}

	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

/*
var OpCodes = map[uint8]OpCode{
	0x00: NewOpCode(0x00, "NOP", 1, 4, []func(cpu *CPU){}),
	0x03: NewOpCode(0x03, "INC BC", 1, 8, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),
	0x0B: NewOpCode(0x0B, "DEC BC", 1, 8, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),

	0x11: NewOpCode(0x11, "LD DE,u16", 3, 12, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.e = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.regs.d = busRead(cpu.pc); cpu.pc++ },
	}),

	0x21: NewOpCode(0x21, "LD HL,u16", 3, 12, []func(cpu *CPU){
		func(cpu *CPU) { cpu.regs.l = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.regs.h = busRead(cpu.pc); cpu.pc++ },
	}),

	0x41: NewOpCode(0x41, "LD B,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.regs.c }}),
	0x47: NewOpCode(0x47, "LD B,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.regs.a }}),

	0x46: NewOpCode(0x46, "LD B,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = busRead(cpu.regs.getHL()) }}),

	0x66: NewOpCode(0x66, "LD H,(HL)", 1, 8, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),

	0x70: NewOpCode(0x70, "LD (HL),B", 1, 8, []func(cpu *CPU){func(cpu *CPU) { busWrite(cpu.regs.getHL(), cpu.regs.b) }}),

	0x73: NewOpCode(0x73, "LD (HL),E", 1, 8, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),

	0x83: NewOpCode(0x83, "ADD A,E", 1, 4, []func(cpu *CPU){
		//func(cpu *CPU) { cpu.regs.pc++ },
	}),

	0xC3: NewOpCode(0xC3, "JP u16", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = busRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xCC: NewOpCode(0xCC, "CALL Z,u16", 3, 12 , []func(cpu *CPU){
		func(cpu *CPU) { cpu.pc = cpu.pc + 2 },
	}),
	0xCE: NewOpCode(0xCE, "ADC A,u8", 2, 8, []func(cpu *CPU){
		func(cpu *CPU) { cpu.pc++ },
	}),

	0xe0: NewOpCode(0xe0, "LD (FF00+u8),A", 2, 12, []func(cpu *CPU){func(cpu *CPU) { lsb = busRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { busWrite(utils.ToUint16(lsb, 0xFF), cpu.regs.a) }}),
	0xea: NewOpCode(0xea, "LD (u16),A", 3, 16, []func(cpu *CPU){func(cpu *CPU) { lsb = busRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { msb = busRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { busWrite(utils.ToUint16(lsb, msb), cpu.regs.a) }}),
	0xf0: NewOpCode(0xf0, "LD A,(FF00+u8)", 2, 12, []func(cpu *CPU){func(cpu *CPU) { lsb = busRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { cpu.regs.a = busRead(utils.ToUint16(lsb, 0xFF)) }}),
	0xfa: NewOpCode(0xfa, "LD A,(u16)", 3, 16, []func(cpu *CPU){func(cpu *CPU) { lsb = busRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { msb = busRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { cpu.regs.a = busRead(utils.ToUint16(lsb, msb)) }}),
}
*/
