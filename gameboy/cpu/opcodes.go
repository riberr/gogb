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

func add16(cpu *CPU, value16 uint16, value8 uint8) uint16 {
	result := value16 + uint16(value8)
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, uint8(result) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, ((value16&0xF)+(uint16(value8)&0xF)) > 0xF)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result>>8) != 0)
	return result
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
	result := uint16(int(value16) + int(int8(value8)))
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, false)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	if int8(value8) >= 0 {
		cpu.regs.setFlag(FLAG_CARRY_C_BIT, uint16(int(value16&0xFF)+int(int8(value8))) > 0xFF)
		cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, uint16(int(value16&0xF)+int(int8(value8&0xf))) > 0xF)
	} else {
		cpu.regs.setFlag(FLAG_CARRY_C_BIT, (result&0xFF) <= (value16&0xFF))
		cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, (result&0xF) <= (value16&0xF))
	}
	return result
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

func daa(cpu *CPU) {
	if cpu.regs.getFlag(FLAG_SUBTRACTION_N_BIT) {
		if cpu.regs.getFlag(FLAG_CARRY_C_BIT) {
			cpu.regs.a -= 0x60
		}
		if cpu.regs.getFlag(FLAG_HALF_CARRY_H_BIT) {
			cpu.regs.a -= 0x6
		}
	} else {
		if cpu.regs.getFlag(FLAG_CARRY_C_BIT) || cpu.regs.a > 0x99 {
			cpu.regs.a += 0x60
			cpu.regs.setFlag(FLAG_CARRY_C_BIT, true)
		}
		if cpu.regs.getFlag(FLAG_HALF_CARRY_H_BIT) || (cpu.regs.a&0xF) > 0x9 {
			cpu.regs.a += 0x6
		}
	}
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, cpu.regs.a == 0)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
}

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
