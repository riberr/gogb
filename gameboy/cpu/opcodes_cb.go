package cpu

import (
	"gogb/utils"
)

func RLC(cpu *CPU, value uint8) uint8 {
	result := (value << 1) | (value >> 7)
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (value&(0b_0000_0001<<7)) == (0b_0000_0001<<7))
	return result
}

func RRC(cpu *CPU, value uint8) uint8 {
	result := (value >> 1) | (value << 7)
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (value&0b_0000_0001) == 1)
	return result
}

func RL(cpu *CPU, value uint8) uint8 {
	var c uint8
	if cpu.regs.getFlag(FLAG_CARRY_C_BIT) {
		c = 1
	} else {
		c = 0
	}
	result := (value << 1) | c
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (value&(0b_0000_0001<<7)) == (0b_0000_0001<<7))
	return result
}

func RR(cpu *CPU, value uint8) uint8 {
	var c uint8
	if cpu.regs.getFlag(FLAG_CARRY_C_BIT) {
		c = 1 << 7
	} else {
		c = 0
	}
	result := (value >> 1) | c
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (value&1) == 1)
	return result
}

func SLA(cpu *CPU, value uint8) uint8 {
	result := value << 1
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (value&(1<<7)) == (1<<7))
	return result
}

func SRA(cpu *CPU, value uint8) uint8 {
	result := (value >> 1) | (value & 0b_1000_0000)
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (value&1) == 1)
	return result
}

func SWAP(cpu *CPU, value uint8) uint8 {
	result := (value&0b_1111_0000)>>4 | (value&0b_0000_1111)<<4
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, false)
	return result
}

func SRL(cpu *CPU, value uint8) uint8 {
	result := value >> 1
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, result == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
	cpu.regs.setFlag(FLAG_CARRY_C_BIT, (value&0b_0000_0001) != 0)
	return result
}

func BIT(cpu *CPU, value uint8, bitPos int) {
	cpu.regs.setFlag(FLAG_ZERO_Z_BIT, ((value>>bitPos)&0b_0000_0001) == 0)
	cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
	cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, true)
	// r.FlagC -> unmodified
}

func RES(value uint8, bitPos int) uint8 {
	//return value & ^(0b_0000_0001 << bitPos)
	return utils.ClearBit(value, bitPos)
}

func SET(value uint8, bitPos int) uint8 {
	//return value | (0b_0000_0001 << bitPos)
	return utils.SetBit(value, bitPos)
}

var OpCodesCB = map[uint8]OpCode{
	// rotate left
	0x00: NewOpCode(0x00, "RLC B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RLC(cpu, cpu.regs.b) }}),
	0x01: NewOpCode(0x01, "RLC C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RLC(cpu, cpu.regs.c) }}),
	0x02: NewOpCode(0x02, "RLC D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RLC(cpu, cpu.regs.d) }}),
	0x03: NewOpCode(0x03, "RLC E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RLC(cpu, cpu.regs.e) }}),
	0x04: NewOpCode(0x04, "RLC H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RLC(cpu, cpu.regs.h) }}),
	0x05: NewOpCode(0x05, "RLC L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RLC(cpu, cpu.regs.l) }}),
	0x07: NewOpCode(0x07, "RLC A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RLC(cpu, cpu.regs.a) }}),
	0x06: NewOpCode(0x06, "RLC (HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RLC(cpu, e)) },
	}),

	// rotate right
	0x08: NewOpCode(0x08, "RRC B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RRC(cpu, cpu.regs.b) }}),
	0x09: NewOpCode(0x09, "RRC C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RRC(cpu, cpu.regs.c) }}),
	0x0a: NewOpCode(0x0a, "RRC D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RRC(cpu, cpu.regs.d) }}),
	0x0b: NewOpCode(0x0b, "RRC E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RRC(cpu, cpu.regs.e) }}),
	0x0c: NewOpCode(0x0c, "RRC H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RRC(cpu, cpu.regs.h) }}),
	0x0d: NewOpCode(0x0d, "RRC L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RRC(cpu, cpu.regs.l) }}),
	0x0f: NewOpCode(0x0f, "RRC A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RRC(cpu, cpu.regs.a) }}),
	0x0e: NewOpCode(0x0e, "RRC (HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RRC(cpu, e)) },
	}),

	// 	rotate left through carry
	0x10: NewOpCode(0x10, "RL B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RL(cpu, cpu.regs.b) }}),
	0x11: NewOpCode(0x11, "RL C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RL(cpu, cpu.regs.c) }}),
	0x12: NewOpCode(0x12, "RL D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RL(cpu, cpu.regs.d) }}),
	0x13: NewOpCode(0x13, "RL E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RL(cpu, cpu.regs.e) }}),
	0x14: NewOpCode(0x14, "RL H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RL(cpu, cpu.regs.h) }}),
	0x15: NewOpCode(0x15, "RL L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RL(cpu, cpu.regs.l) }}),
	0x17: NewOpCode(0x17, "RL A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RL(cpu, cpu.regs.a) }}),
	0x16: NewOpCode(0x16, "RL (HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RL(cpu, e)) },
	}),

	// rotate right through carry
	0x18: NewOpCode(0x18, "RR B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RR(cpu, cpu.regs.b) }}),
	0x19: NewOpCode(0x19, "RR C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RR(cpu, cpu.regs.c) }}),
	0x1a: NewOpCode(0x1a, "RR D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RR(cpu, cpu.regs.d) }}),
	0x1b: NewOpCode(0x1b, "RR E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RR(cpu, cpu.regs.e) }}),
	0x1c: NewOpCode(0x1c, "RR H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RR(cpu, cpu.regs.h) }}),
	0x1d: NewOpCode(0x1d, "RR L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RR(cpu, cpu.regs.l) }}),
	0x1f: NewOpCode(0x1f, "RR A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RR(cpu, cpu.regs.a) }}),
	0x1e: NewOpCode(0x1e, "RR (HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RR(cpu, e)) },
	}),

	// shift left arithmetic (b0=0)
	0x20: NewOpCode(0x20, "SLA B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SLA(cpu, cpu.regs.b) }}),
	0x21: NewOpCode(0x21, "SLA C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SLA(cpu, cpu.regs.c) }}),
	0x22: NewOpCode(0x22, "SLA D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SLA(cpu, cpu.regs.d) }}),
	0x23: NewOpCode(0x23, "SLA E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SLA(cpu, cpu.regs.e) }}),
	0x24: NewOpCode(0x24, "SLA H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SLA(cpu, cpu.regs.h) }}),
	0x25: NewOpCode(0x25, "SLA L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SLA(cpu, cpu.regs.l) }}),
	0x27: NewOpCode(0x27, "SLA A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SLA(cpu, cpu.regs.a) }}),
	0x26: NewOpCode(0x26, "SLA (HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SLA(cpu, e)) },
	}),

	// shift right arithmetic (b7=b7)
	0x28: NewOpCode(0x28, "SRA B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SRA(cpu, cpu.regs.b) }}),
	0x29: NewOpCode(0x29, "SRA C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SRA(cpu, cpu.regs.c) }}),
	0x2a: NewOpCode(0x2a, "SRA D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SRA(cpu, cpu.regs.d) }}),
	0x2b: NewOpCode(0x2b, "SRA E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SRA(cpu, cpu.regs.e) }}),
	0x2c: NewOpCode(0x2c, "SRA H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SRA(cpu, cpu.regs.h) }}),
	0x2d: NewOpCode(0x2d, "SRA L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SRA(cpu, cpu.regs.l) }}),
	0x2f: NewOpCode(0x2f, "SRA A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SRA(cpu, cpu.regs.a) }}),
	0x2e: NewOpCode(0x2e, "SRA (HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SRA(cpu, e)) },
	}),

	// exchange low/hi-nibble
	0x30: NewOpCode(0x30, "SWAP B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SWAP(cpu, cpu.regs.b) }}),
	0x31: NewOpCode(0x31, "SWAP C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SWAP(cpu, cpu.regs.c) }}),
	0x32: NewOpCode(0x32, "SWAP D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SWAP(cpu, cpu.regs.d) }}),
	0x33: NewOpCode(0x33, "SWAP E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SWAP(cpu, cpu.regs.e) }}),
	0x34: NewOpCode(0x34, "SWAP H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SWAP(cpu, cpu.regs.h) }}),
	0x35: NewOpCode(0x35, "SWAP L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SWAP(cpu, cpu.regs.l) }}),
	0x37: NewOpCode(0x37, "SWAP A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SWAP(cpu, cpu.regs.a) }}),
	0x36: NewOpCode(0x36, "SWAP (HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SWAP(cpu, e)) },
	}),

	// shift right logical (b7=0)
	0x38: NewOpCode(0x38, "SRL B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SRL(cpu, cpu.regs.b) }}),
	0x39: NewOpCode(0x39, "SRL C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SRL(cpu, cpu.regs.c) }}),
	0x3a: NewOpCode(0x3a, "SRL D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SRL(cpu, cpu.regs.d) }}),
	0x3b: NewOpCode(0x3b, "SRL E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SRL(cpu, cpu.regs.e) }}),
	0x3c: NewOpCode(0x3c, "SRL H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SRL(cpu, cpu.regs.h) }}),
	0x3d: NewOpCode(0x3d, "SRL L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SRL(cpu, cpu.regs.l) }}),
	0x3f: NewOpCode(0x3f, "SRL A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SRL(cpu, cpu.regs.a) }}),
	0x3e: NewOpCode(0x3e, "SRL (HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SRL(cpu, e)) },
	}),

	// test bit n
	0x40: NewOpCode(0x40, "BIT 0,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.b, 0) }}),
	0x41: NewOpCode(0x41, "BIT 0,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.c, 0) }}),
	0x42: NewOpCode(0x42, "BIT 0,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.d, 0) }}),
	0x43: NewOpCode(0x43, "BIT 0,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.e, 0) }}),
	0x44: NewOpCode(0x44, "BIT 0,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.h, 0) }}),
	0x45: NewOpCode(0x45, "BIT 0,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.l, 0) }}),
	0x47: NewOpCode(0x47, "BIT 0,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.a, 0) }}),
	0x46: NewOpCode(0x46, "BIT 0,(HL)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { BIT(cpu, cpu.bus.BusRead(cpu.regs.getHL()), 0) },
	}),

	0x48: NewOpCode(0x48, "BIT 1,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.b, 1) }}),
	0x49: NewOpCode(0x49, "BIT 1,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.c, 1) }}),
	0x4a: NewOpCode(0x4a, "BIT 1,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.d, 1) }}),
	0x4b: NewOpCode(0x4b, "BIT 1,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.e, 1) }}),
	0x4c: NewOpCode(0x4c, "BIT 1,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.h, 1) }}),
	0x4d: NewOpCode(0x4d, "BIT 1,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.l, 1) }}),
	0x4f: NewOpCode(0x4f, "BIT 1,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.a, 1) }}),
	0x4e: NewOpCode(0x4e, "BIT 1,(HL)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { BIT(cpu, cpu.bus.BusRead(cpu.regs.getHL()), 1) },
	}),

	0x50: NewOpCode(0x50, "BIT 2,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.b, 2) }}),
	0x51: NewOpCode(0x51, "BIT 2,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.c, 2) }}),
	0x52: NewOpCode(0x52, "BIT 2,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.d, 2) }}),
	0x53: NewOpCode(0x53, "BIT 2,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.e, 2) }}),
	0x54: NewOpCode(0x54, "BIT 2,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.h, 2) }}),
	0x55: NewOpCode(0x55, "BIT 2,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.l, 2) }}),
	0x57: NewOpCode(0x57, "BIT 2,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.a, 2) }}),
	0x56: NewOpCode(0x56, "BIT 2,(HL)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { BIT(cpu, cpu.bus.BusRead(cpu.regs.getHL()), 2) },
	}),

	0x58: NewOpCode(0x58, "BIT 3,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.b, 3) }}),
	0x59: NewOpCode(0x59, "BIT 3,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.c, 3) }}),
	0x5a: NewOpCode(0x5a, "BIT 3,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.d, 3) }}),
	0x5b: NewOpCode(0x5b, "BIT 3,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.e, 3) }}),
	0x5c: NewOpCode(0x5c, "BIT 3,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.h, 3) }}),
	0x5d: NewOpCode(0x5d, "BIT 3,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.l, 3) }}),
	0x5f: NewOpCode(0x5f, "BIT 3,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.a, 3) }}),
	0x5e: NewOpCode(0x5e, "BIT 3,(HL)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { BIT(cpu, cpu.bus.BusRead(cpu.regs.getHL()), 3) },
	}),

	0x60: NewOpCode(0x60, "BIT 4,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.b, 4) }}),
	0x61: NewOpCode(0x61, "BIT 4,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.c, 4) }}),
	0x62: NewOpCode(0x62, "BIT 4,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.d, 4) }}),
	0x63: NewOpCode(0x63, "BIT 4,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.e, 4) }}),
	0x64: NewOpCode(0x64, "BIT 4,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.h, 4) }}),
	0x65: NewOpCode(0x65, "BIT 4,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.l, 4) }}),
	0x67: NewOpCode(0x67, "BIT 4,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.a, 4) }}),
	0x66: NewOpCode(0x66, "BIT 4,(HL)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { BIT(cpu, cpu.bus.BusRead(cpu.regs.getHL()), 4) },
	}),

	0x68: NewOpCode(0x68, "BIT 5,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.b, 5) }}),
	0x69: NewOpCode(0x69, "BIT 5,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.c, 5) }}),
	0x6a: NewOpCode(0x6a, "BIT 5,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.d, 5) }}),
	0x6b: NewOpCode(0x6b, "BIT 5,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.e, 5) }}),
	0x6c: NewOpCode(0x6c, "BIT 5,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.h, 5) }}),
	0x6d: NewOpCode(0x6d, "BIT 5,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.l, 5) }}),
	0x6f: NewOpCode(0x6f, "BIT 5,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.a, 5) }}),
	0x6e: NewOpCode(0x6e, "BIT 5,(HL)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { BIT(cpu, cpu.bus.BusRead(cpu.regs.getHL()), 5) },
	}),

	0x70: NewOpCode(0x70, "BIT 6,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.b, 6) }}),
	0x71: NewOpCode(0x71, "BIT 6,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.c, 6) }}),
	0x72: NewOpCode(0x72, "BIT 6,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.d, 6) }}),
	0x73: NewOpCode(0x73, "BIT 6,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.e, 6) }}),
	0x74: NewOpCode(0x74, "BIT 6,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.h, 6) }}),
	0x75: NewOpCode(0x75, "BIT 6,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.l, 6) }}),
	0x77: NewOpCode(0x77, "BIT 6,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.a, 6) }}),
	0x76: NewOpCode(0x76, "BIT 6,(HL)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { BIT(cpu, cpu.bus.BusRead(cpu.regs.getHL()), 6) },
	}),

	0x78: NewOpCode(0x78, "BIT 7,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.b, 7) }}),
	0x79: NewOpCode(0x79, "BIT 7,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.c, 7) }}),
	0x7a: NewOpCode(0x7a, "BIT 7,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.d, 7) }}),
	0x7b: NewOpCode(0x7b, "BIT 7,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.e, 7) }}),
	0x7c: NewOpCode(0x7c, "BIT 7,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.h, 7) }}),
	0x7d: NewOpCode(0x7d, "BIT 7,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.l, 7) }}),
	0x7f: NewOpCode(0x7f, "BIT 7,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { BIT(cpu, cpu.regs.a, 7) }}),
	0x7e: NewOpCode(0x7e, "BIT 7,(HL)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { BIT(cpu, cpu.bus.BusRead(cpu.regs.getHL()), 7) },
	}),

	// reset bit n
	0x80: NewOpCode(0x80, "RES 0,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RES(cpu.regs.b, 0) }}),
	0x81: NewOpCode(0x81, "RES 0,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RES(cpu.regs.c, 0) }}),
	0x82: NewOpCode(0x82, "RES 0,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RES(cpu.regs.d, 0) }}),
	0x83: NewOpCode(0x83, "RES 0,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RES(cpu.regs.e, 0) }}),
	0x84: NewOpCode(0x84, "RES 0,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RES(cpu.regs.h, 0) }}),
	0x85: NewOpCode(0x85, "RES 0,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RES(cpu.regs.l, 0) }}),
	0x87: NewOpCode(0x87, "RES 0,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RES(cpu.regs.a, 0) }}),
	0x86: NewOpCode(0x86, "RES 0,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RES(e, 0)) },
	}),

	0x88: NewOpCode(0x88, "RES 1,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RES(cpu.regs.b, 1) }}),
	0x89: NewOpCode(0x89, "RES 1,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RES(cpu.regs.c, 1) }}),
	0x8a: NewOpCode(0x8a, "RES 1,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RES(cpu.regs.d, 1) }}),
	0x8b: NewOpCode(0x8b, "RES 1,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RES(cpu.regs.e, 1) }}),
	0x8c: NewOpCode(0x8c, "RES 1,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RES(cpu.regs.h, 1) }}),
	0x8d: NewOpCode(0x8d, "RES 1,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RES(cpu.regs.l, 1) }}),
	0x8f: NewOpCode(0x8f, "RES 1,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RES(cpu.regs.a, 1) }}),
	0x8e: NewOpCode(0x8e, "RES 1,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RES(e, 1)) },
	}),

	0x90: NewOpCode(0x90, "RES 2,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RES(cpu.regs.b, 2) }}),
	0x91: NewOpCode(0x91, "RES 2,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RES(cpu.regs.c, 2) }}),
	0x92: NewOpCode(0x92, "RES 2,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RES(cpu.regs.d, 2) }}),
	0x93: NewOpCode(0x93, "RES 2,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RES(cpu.regs.e, 2) }}),
	0x94: NewOpCode(0x94, "RES 2,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RES(cpu.regs.h, 2) }}),
	0x95: NewOpCode(0x95, "RES 2,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RES(cpu.regs.l, 2) }}),
	0x97: NewOpCode(0x97, "RES 2,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RES(cpu.regs.a, 2) }}),
	0x96: NewOpCode(0x96, "RES 2,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RES(e, 2)) },
	}),

	0x98: NewOpCode(0x98, "RES 3,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RES(cpu.regs.b, 3) }}),
	0x99: NewOpCode(0x99, "RES 3,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RES(cpu.regs.c, 3) }}),
	0x9a: NewOpCode(0x9a, "RES 3,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RES(cpu.regs.d, 3) }}),
	0x9b: NewOpCode(0x9b, "RES 3,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RES(cpu.regs.e, 3) }}),
	0x9c: NewOpCode(0x9c, "RES 3,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RES(cpu.regs.h, 3) }}),
	0x9d: NewOpCode(0x9d, "RES 3,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RES(cpu.regs.l, 3) }}),
	0x9f: NewOpCode(0x9f, "RES 3,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RES(cpu.regs.a, 3) }}),
	0x9e: NewOpCode(0x9e, "RES 3,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RES(e, 3)) },
	}),

	0xa0: NewOpCode(0xa0, "RES 4,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RES(cpu.regs.b, 4) }}),
	0xa1: NewOpCode(0xa1, "RES 4,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RES(cpu.regs.c, 4) }}),
	0xa2: NewOpCode(0xa2, "RES 4,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RES(cpu.regs.d, 4) }}),
	0xa3: NewOpCode(0xa3, "RES 4,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RES(cpu.regs.e, 4) }}),
	0xa4: NewOpCode(0xa4, "RES 4,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RES(cpu.regs.h, 4) }}),
	0xa5: NewOpCode(0xa5, "RES 4,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RES(cpu.regs.l, 4) }}),
	0xa7: NewOpCode(0xa7, "RES 4,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RES(cpu.regs.a, 4) }}),
	0xa6: NewOpCode(0xa6, "RES 4,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RES(e, 4)) },
	}),

	0xa8: NewOpCode(0xa8, "RES 5,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RES(cpu.regs.b, 5) }}),
	0xa9: NewOpCode(0xa9, "RES 5,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RES(cpu.regs.c, 5) }}),
	0xaa: NewOpCode(0xaa, "RES 5,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RES(cpu.regs.d, 5) }}),
	0xab: NewOpCode(0xab, "RES 5,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RES(cpu.regs.e, 5) }}),
	0xac: NewOpCode(0xac, "RES 5,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RES(cpu.regs.h, 5) }}),
	0xad: NewOpCode(0xad, "RES 5,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RES(cpu.regs.l, 5) }}),
	0xaf: NewOpCode(0xaf, "RES 5,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RES(cpu.regs.a, 5) }}),
	0xae: NewOpCode(0xae, "RES 5,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RES(e, 5)) },
	}),

	0xb0: NewOpCode(0xb0, "RES 6,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RES(cpu.regs.b, 6) }}),
	0xb1: NewOpCode(0xb1, "RES 6,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RES(cpu.regs.c, 6) }}),
	0xb2: NewOpCode(0xb2, "RES 6,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RES(cpu.regs.d, 6) }}),
	0xb3: NewOpCode(0xb3, "RES 6,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RES(cpu.regs.e, 6) }}),
	0xb4: NewOpCode(0xb4, "RES 6,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RES(cpu.regs.h, 6) }}),
	0xb5: NewOpCode(0xb5, "RES 6,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RES(cpu.regs.l, 6) }}),
	0xb7: NewOpCode(0xb7, "RES 6,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RES(cpu.regs.a, 6) }}),
	0xb6: NewOpCode(0xb6, "RES 6,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RES(e, 6)) },
	}),

	0xb8: NewOpCode(0xb8, "RES 7,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = RES(cpu.regs.b, 7) }}),
	0xb9: NewOpCode(0xb9, "RES 7,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = RES(cpu.regs.c, 7) }}),
	0xba: NewOpCode(0xba, "RES 7,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = RES(cpu.regs.d, 7) }}),
	0xbb: NewOpCode(0xbb, "RES 7,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = RES(cpu.regs.e, 7) }}),
	0xbc: NewOpCode(0xbc, "RES 7,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = RES(cpu.regs.h, 7) }}),
	0xbd: NewOpCode(0xbd, "RES 7,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = RES(cpu.regs.l, 7) }}),
	0xbf: NewOpCode(0xbf, "RES 7,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = RES(cpu.regs.a, 7) }}),
	0xbe: NewOpCode(0xbe, "RES 7,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), RES(e, 7)) },
	}),

	// set bit n
	0xc0: NewOpCode(0xc0, "SET 0,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SET(cpu.regs.b, 0) }}),
	0xc1: NewOpCode(0xc1, "SET 0,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SET(cpu.regs.c, 0) }}),
	0xc2: NewOpCode(0xc2, "SET 0,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SET(cpu.regs.d, 0) }}),
	0xc3: NewOpCode(0xc3, "SET 0,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SET(cpu.regs.e, 0) }}),
	0xc4: NewOpCode(0xc4, "SET 0,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SET(cpu.regs.h, 0) }}),
	0xc5: NewOpCode(0xc5, "SET 0,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SET(cpu.regs.l, 0) }}),
	0xc7: NewOpCode(0xc7, "SET 0,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SET(cpu.regs.a, 0) }}),
	0xc6: NewOpCode(0xc6, "SET 0,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SET(e, 0)) },
	}),

	0xc8: NewOpCode(0xc8, "SET 1,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SET(cpu.regs.b, 1) }}),
	0xc9: NewOpCode(0xc9, "SET 1,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SET(cpu.regs.c, 1) }}),
	0xca: NewOpCode(0xca, "SET 1,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SET(cpu.regs.d, 1) }}),
	0xcb: NewOpCode(0xcb, "SET 1,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SET(cpu.regs.e, 1) }}),
	0xcc: NewOpCode(0xcc, "SET 1,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SET(cpu.regs.h, 1) }}),
	0xcd: NewOpCode(0xcd, "SET 1,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SET(cpu.regs.l, 1) }}),
	0xcf: NewOpCode(0xcf, "SET 1,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SET(cpu.regs.a, 1) }}),
	0xce: NewOpCode(0xce, "SET 1,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SET(e, 1)) },
	}),

	0xd0: NewOpCode(0xd0, "SET 2,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SET(cpu.regs.b, 2) }}),
	0xd1: NewOpCode(0xd1, "SET 2,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SET(cpu.regs.c, 2) }}),
	0xd2: NewOpCode(0xd2, "SET 2,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SET(cpu.regs.d, 2) }}),
	0xd3: NewOpCode(0xd3, "SET 2,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SET(cpu.regs.e, 2) }}),
	0xd4: NewOpCode(0xd4, "SET 2,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SET(cpu.regs.h, 2) }}),
	0xd5: NewOpCode(0xd5, "SET 2,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SET(cpu.regs.l, 2) }}),
	0xd7: NewOpCode(0xd7, "SET 2,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SET(cpu.regs.a, 2) }}),
	0xd6: NewOpCode(0xd6, "SET 2,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SET(e, 2)) },
	}),

	0xd8: NewOpCode(0xd8, "SET 3,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SET(cpu.regs.b, 3) }}),
	0xd9: NewOpCode(0xd9, "SET 3,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SET(cpu.regs.c, 3) }}),
	0xda: NewOpCode(0xda, "SET 3,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SET(cpu.regs.d, 3) }}),
	0xdb: NewOpCode(0xdb, "SET 3,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SET(cpu.regs.e, 3) }}),
	0xdc: NewOpCode(0xdc, "SET 3,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SET(cpu.regs.h, 3) }}),
	0xdd: NewOpCode(0xdd, "SET 3,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SET(cpu.regs.l, 3) }}),
	0xdf: NewOpCode(0xdf, "SET 3,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SET(cpu.regs.a, 3) }}),
	0xde: NewOpCode(0xde, "SET 3,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SET(e, 3)) },
	}),

	0xe0: NewOpCode(0xe0, "SET 4,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SET(cpu.regs.b, 4) }}),
	0xe1: NewOpCode(0xe1, "SET 4,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SET(cpu.regs.c, 4) }}),
	0xe2: NewOpCode(0xe2, "SET 4,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SET(cpu.regs.d, 4) }}),
	0xe3: NewOpCode(0xe3, "SET 4,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SET(cpu.regs.e, 4) }}),
	0xe4: NewOpCode(0xe4, "SET 4,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SET(cpu.regs.h, 4) }}),
	0xe5: NewOpCode(0xe5, "SET 4,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SET(cpu.regs.l, 4) }}),
	0xe7: NewOpCode(0xe7, "SET 4,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SET(cpu.regs.a, 4) }}),
	0xe6: NewOpCode(0xe6, "SET 4,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SET(e, 4)) },
	}),

	0xe8: NewOpCode(0xe8, "SET 5,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SET(cpu.regs.b, 5) }}),
	0xe9: NewOpCode(0xe9, "SET 5,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SET(cpu.regs.c, 5) }}),
	0xea: NewOpCode(0xea, "SET 5,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SET(cpu.regs.d, 5) }}),
	0xeb: NewOpCode(0xeb, "SET 5,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SET(cpu.regs.e, 5) }}),
	0xec: NewOpCode(0xec, "SET 5,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SET(cpu.regs.h, 5) }}),
	0xed: NewOpCode(0xed, "SET 5,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SET(cpu.regs.l, 5) }}),
	0xef: NewOpCode(0xef, "SET 5,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SET(cpu.regs.a, 5) }}),
	0xee: NewOpCode(0xee, "SET 5,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SET(e, 5)) },
	}),

	0xf0: NewOpCode(0xf0, "SET 6,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SET(cpu.regs.b, 6) }}),
	0xf1: NewOpCode(0xf1, "SET 6,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SET(cpu.regs.c, 6) }}),
	0xf2: NewOpCode(0xf2, "SET 6,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SET(cpu.regs.d, 6) }}),
	0xf3: NewOpCode(0xf3, "SET 6,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SET(cpu.regs.e, 6) }}),
	0xf4: NewOpCode(0xf4, "SET 6,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SET(cpu.regs.h, 6) }}),
	0xf5: NewOpCode(0xf5, "SET 6,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SET(cpu.regs.l, 6) }}),
	0xf7: NewOpCode(0xf7, "SET 6,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SET(cpu.regs.a, 6) }}),
	0xf6: NewOpCode(0xf6, "SET 6,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SET(e, 6)) },
	}),

	0xf8: NewOpCode(0xf8, "SET 7,B", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = SET(cpu.regs.b, 7) }}),
	0xf9: NewOpCode(0xf9, "SET 7,C", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = SET(cpu.regs.c, 7) }}),
	0xfa: NewOpCode(0xfa, "SET 7,D", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = SET(cpu.regs.d, 7) }}),
	0xfb: NewOpCode(0xfb, "SET 7,E", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = SET(cpu.regs.e, 7) }}),
	0xfc: NewOpCode(0xfc, "SET 7,H", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = SET(cpu.regs.h, 7) }}),
	0xfd: NewOpCode(0xfd, "SET 7,L", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = SET(cpu.regs.l, 7) }}),
	0xff: NewOpCode(0xff, "SET 7,A", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = SET(cpu.regs.a, 7) }}),
	0xfe: NewOpCode(0xfe, "SET 7,(HL)", 2, 16, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), SET(e, 7)) },
	}),
}
