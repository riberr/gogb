package cpu

import (
	"gogb/utils"
)

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
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (cpu.regs.a&0xF)+(value&0xF) > 0x0F)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result>>8) != 0)
	cpu.regs.a = uint8(result)
}

func addHL(cpu *CPU, value uint16) {
	result := int(cpu.regs.getHL()) + int(value)
	// regs.FlagZ = Unmodified
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, ((cpu.regs.getHL()&0x0FFF)+(value&0x0FFF)) > 0x0FFF)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, result>>16 != 0)
	cpu.regs.setHL(uint16(result))
}

func addSigned8(cpu *CPU, value16 uint16, value8 uint8) uint16 {
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, false)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, ((value16&0xF)+(uint16(value8)&0xF)) > 0xF)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, ((int(value16)+int(value8))>>8) != 0)
	return uint16(int(value16) + int(value8))
}

func adc(cpu *CPU, value uint8) {
	carry := utils.ToInt(cpu.regs.getFlag(FLAG_CARRY_C_BIT))
	result := int(cpu.regs.a) + int(carry) + int(value) // cast to int to be able to detect overflow
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, uint8(result) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (cpu.regs.a&0xF)+(value&0xF)+carry > 0xF)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result>>8) != 0)
	cpu.regs.a = uint8(result)
}

func sub(cpu *CPU, value uint8) {
	dirtySum := int16(cpu.regs.a) - int16(value) // cast to int16 to be able to detect overflow
	total := uint8(dirtySum)
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, uint8(total) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, int16(cpu.regs.a&0xF)-int16(value&0xF) < 0)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, dirtySum < 0)
	cpu.regs.a = total
}

func sbc(cpu *CPU, value uint8) {
	carry := utils.ToInt(cpu.regs.getFlag(FLAG_CARRY_C_BIT))
	result := int(cpu.regs.a) - int(value) - int(carry) // cast to int to be able to detect overflow
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, uint8(result) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (int(cpu.regs.a)&0xF)-(int(value)&0xF)-int(utils.ToInt(cpu.regs.getFlag(FLAG_CARRY_C_BIT))) < 0)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, result < 0)
	cpu.regs.a = uint8(result)
}

func cp(cpu *CPU, value uint8) {
	result := int(cpu.regs.a) - int(value)
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (0x0F&value) > (0x0F&cpu.regs.a))
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result>>8) != 0)
}

func inc(cpu *CPU, value uint8) uint8 {
	//fmt.Printf("%08b \n", cpu.regs.f)
	result := value + 1
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (0x0F&value) == 0x0F)
	// flag c unmodified
	return result
}

func dec(cpu *CPU, value uint8) uint8 {
	result := value - 1
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (0x0F&value) == 0x00)
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
var ee uint16 = 0               // temp value when executing opcodes
var stop = false                // used for conditional jumps

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
