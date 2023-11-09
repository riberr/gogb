package cpu

var OpCodes8bitArithmeticsGenerated = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 8BIT ARITHMETICS & LOGICAL
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Add (register), Add (indirect HL), Add (immediate)
	0x80: NewOpCode(0x80, "ADD A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.regs.b) }}),
	0x81: NewOpCode(0x81, "ADD A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.regs.c) }}),
	0x82: NewOpCode(0x82, "ADD A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.regs.d) }}),
	0x83: NewOpCode(0x83, "ADD A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.regs.e) }}),
	0x84: NewOpCode(0x84, "ADD A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.regs.h) }}),
	0x85: NewOpCode(0x85, "ADD A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.regs.l) }}),
	0x86: NewOpCode(0x86, "ADD A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.bus.BusRead(cpu.regs.getHL())) }}),
	0x87: NewOpCode(0x87, "ADD A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.regs.a) }}),
	0xc6: NewOpCode(0xc6, "ADD A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { add(cpu, cpu.bus.BusRead(cpu.pc)); cpu.pc++ }}),

	// Add with carry (register), Add with carry (indirect HL), Add with carry (immediate)
	0x88: NewOpCode(0x88, "ADC A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.regs.b) }}),
	0x89: NewOpCode(0x89, "ADC A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.regs.c) }}),
	0x8a: NewOpCode(0x8a, "ADC A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.regs.d) }}),
	0x8b: NewOpCode(0x8b, "ADC A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.regs.e) }}),
	0x8c: NewOpCode(0x8c, "ADC A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.regs.h) }}),
	0x8d: NewOpCode(0x8d, "ADC A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.regs.l) }}),
	0x8e: NewOpCode(0x8e, "ADC A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.bus.BusRead(cpu.regs.getHL())) }}),
	0x8f: NewOpCode(0x8f, "ADC A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.regs.a) }}),
	0xce: NewOpCode(0xce, "ADC A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { adc(cpu, cpu.bus.BusRead(cpu.pc)); cpu.pc++ }}),

	// Subtract (register), Subtract (indirect HL), Subtract (immediate)
	0x90: NewOpCode(0x90, "SUB A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.regs.b) }}),
	0x91: NewOpCode(0x91, "SUB A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.regs.c) }}),
	0x92: NewOpCode(0x92, "SUB A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.regs.d) }}),
	0x93: NewOpCode(0x93, "SUB A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.regs.e) }}),
	0x94: NewOpCode(0x94, "SUB A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.regs.h) }}),
	0x95: NewOpCode(0x95, "SUB A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.regs.l) }}),
	0x96: NewOpCode(0x96, "SUB A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.bus.BusRead(cpu.regs.getHL())) }}),
	0x97: NewOpCode(0x97, "SUB A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.regs.a) }}),
	0xd6: NewOpCode(0xd6, "SUB A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { sub(cpu, cpu.bus.BusRead(cpu.pc)); cpu.pc++ }}),

	// Subtract with carry (register),  Subtract with carry (indirect HL), Subtract with carry (immediate)
	0x98: NewOpCode(0x98, "SBC A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.regs.b) }}),
	0x99: NewOpCode(0x99, "SBC A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.regs.c) }}),
	0x9a: NewOpCode(0x9a, "SBC A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.regs.d) }}),
	0x9b: NewOpCode(0x9b, "SBC A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.regs.e) }}),
	0x9c: NewOpCode(0x9c, "SBC A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.regs.h) }}),
	0x9d: NewOpCode(0x9d, "SBC A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.regs.l) }}),
	0x9e: NewOpCode(0x9e, "SBC A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.bus.BusRead(cpu.regs.getHL())) }}),
	0x9f: NewOpCode(0x9f, "SBC A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.regs.a) }}),
	0xde: NewOpCode(0xde, "SBC A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { sbc(cpu, cpu.bus.BusRead(cpu.pc)); cpu.pc++ }}),

	// Compare (register), Compare (indirect HL), Compare (immediate)
	0xb8: NewOpCode(0xb8, "CP A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.regs.b) }}),
	0xb9: NewOpCode(0xb9, "CP A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.regs.c) }}),
	0xba: NewOpCode(0xba, "CP A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.regs.d) }}),
	0xbb: NewOpCode(0xbb, "CP A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.regs.e) }}),
	0xbc: NewOpCode(0xbc, "CP A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.regs.h) }}),
	0xbd: NewOpCode(0xbd, "CP A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.regs.l) }}),
	0xbe: NewOpCode(0xbe, "CP A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.bus.BusRead(cpu.regs.getHL())) }}),
	0xbf: NewOpCode(0xbf, "CP A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.regs.a) }}),
	0xfe: NewOpCode(0xfe, "CP A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cp(cpu, cpu.bus.BusRead(cpu.pc)); cpu.pc++ }}),

	// Increment (register),Increment (indirect HL)
	0x04: NewOpCode(0x04, "INC B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = inc(cpu, cpu.regs.b) }}),
	0x0c: NewOpCode(0x0c, "INC C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = inc(cpu, cpu.regs.c) }}),
	0x14: NewOpCode(0x14, "INC D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = inc(cpu, cpu.regs.d) }}),
	0x1c: NewOpCode(0x1c, "INC E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = inc(cpu, cpu.regs.e) }}),
	0x24: NewOpCode(0x24, "INC H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = inc(cpu, cpu.regs.h) }}),
	0x2c: NewOpCode(0x2c, "INC L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = inc(cpu, cpu.regs.l) }}),
	0x3c: NewOpCode(0x3c, "INC A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = inc(cpu, cpu.regs.a) }}),
	0x34: NewOpCode(0x34, "INC (HL)", 1, 12, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), inc(cpu, e)) },
	}),

	// Decrement (register), Decrement (indirect HL)
	0x05: NewOpCode(0x05, "DEC B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = dec(cpu, cpu.regs.b) }}),
	0x0d: NewOpCode(0x0d, "DEC C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = dec(cpu, cpu.regs.c) }}),
	0x15: NewOpCode(0x15, "DEC D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = dec(cpu, cpu.regs.d) }}),
	0x1d: NewOpCode(0x1d, "DEC E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = dec(cpu, cpu.regs.e) }}),
	0x25: NewOpCode(0x25, "DEC H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = dec(cpu, cpu.regs.h) }}),
	0x2d: NewOpCode(0x2d, "DEC L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = dec(cpu, cpu.regs.l) }}),
	0x3d: NewOpCode(0x3d, "DEC A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = dec(cpu, cpu.regs.a) }}),
	0x35: NewOpCode(0x35, "DEC (HL)", 1, 12, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.regs.getHL()) },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), dec(cpu, e)) },
	}),

	// Bitwise AND (register), Bitwise AND (indirect HL), Bitwise AND (immediate)
	0xa0: NewOpCode(0xa0, "AND A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.regs.b) }}),
	0xa1: NewOpCode(0xa1, "AND A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.regs.c) }}),
	0xa2: NewOpCode(0xa2, "AND A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.regs.d) }}),
	0xa3: NewOpCode(0xa3, "AND A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.regs.e) }}),
	0xa4: NewOpCode(0xa4, "AND A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.regs.h) }}),
	0xa5: NewOpCode(0xa5, "AND A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.regs.l) }}),
	0xa6: NewOpCode(0xa6, "AND A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.bus.BusRead(cpu.regs.getHL())) }}),
	0xa7: NewOpCode(0xa7, "AND A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.regs.a) }}),
	0xe6: NewOpCode(0xe6, "AND A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { and(cpu, cpu.bus.BusRead(cpu.pc)); cpu.pc++ }}),

	// Bitwise OR (register),  Bitwise OR (indirect HL), Bitwise OR (immediate)
	0xb0: NewOpCode(0xb0, "OR A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.regs.b) }}),
	0xb1: NewOpCode(0xb1, "OR A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.regs.c) }}),
	0xb2: NewOpCode(0xb2, "OR A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.regs.d) }}),
	0xb3: NewOpCode(0xb3, "OR A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.regs.e) }}),
	0xb4: NewOpCode(0xb4, "OR A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.regs.h) }}),
	0xb5: NewOpCode(0xb5, "OR A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.regs.l) }}),
	0xb6: NewOpCode(0xb6, "OR A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.bus.BusRead(cpu.regs.getHL())) }}),
	0xb7: NewOpCode(0xb7, "OR A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.regs.a) }}),
	0xf6: NewOpCode(0xf6, "OR A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { or(cpu, cpu.bus.BusRead(cpu.pc)); cpu.pc++ }}),

	// Bitwise XOR (register),  Bitwise XOR (indirect HL), Bitwise XOR (immediate)
	0xa8: NewOpCode(0xa8, "XOR A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.regs.b) }}),
	0xa9: NewOpCode(0xa9, "XOR A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.regs.c) }}),
	0xaa: NewOpCode(0xaa, "XOR A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.regs.d) }}),
	0xab: NewOpCode(0xab, "XOR A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.regs.e) }}),
	0xac: NewOpCode(0xac, "XOR A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.regs.h) }}),
	0xad: NewOpCode(0xad, "XOR A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.regs.l) }}),
	0xae: NewOpCode(0xae, "XOR A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.bus.BusRead(cpu.regs.getHL())) }}),
	0xaf: NewOpCode(0xaf, "XOR A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.regs.a) }}),
	0xee: NewOpCode(0xee, "XOR A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { xor(cpu, cpu.bus.BusRead(cpu.pc)); cpu.pc++ }}),

	// CCF: Complement carry flag, SCF: Set carry flag, DAA: Decimal adjust accumulator, CPL: Complement accumulator
	0x27: NewOpCode(0x27, "DAA todo", 1, 4, []func(cpu *CPU){func(cpu *CPU) { /*todo*/ }}),
	0x2f: NewOpCode(0x2f, "CPL", 1, 4, []func(cpu *CPU){
		// flip all bits
		func(cpu *CPU) {
			cpu.regs.a = ^cpu.regs.a
			cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, true)
			cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, true)
		},
	}),
	0x37: NewOpCode(0x37, "SCF", 1, 4, []func(cpu *CPU){
		// Set carry flag
		func(cpu *CPU) {
			cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
			cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
			cpu.regs.setFlag(FLAG_CARRY_C_BIT, true)
		},
	}),
	0x3f: NewOpCode(0x3f, "CCF", 1, 4, []func(cpu *CPU){
		// Complement carry flag
		func(cpu *CPU) {
			cpu.regs.setFlag(FLAG_SUBTRACTION_N_BIT, false)
			cpu.regs.setFlag(FLAG_HALF_CARRY_H_BIT, false)
			cpu.regs.setFlag(FLAG_CARRY_C_BIT, !cpu.regs.getFlag(FLAG_CARRY_C_BIT))
		},
	}),
}
