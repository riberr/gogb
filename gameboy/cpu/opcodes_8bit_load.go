package cpu

import (
	"gogb/utils"
)

var OpCodes8bitLoadGenerated = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 8BIT LOAD
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	//  Load register (register)
	0x40: NewOpCode(0x40, "LD B,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
	0x41: NewOpCode(0x41, "LD B,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.regs.c }}),
	0x42: NewOpCode(0x42, "LD B,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.regs.d }}),
	0x43: NewOpCode(0x43, "LD B,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.regs.e }}),
	0x44: NewOpCode(0x44, "LD B,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.regs.h }}),
	0x45: NewOpCode(0x45, "LD B,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.regs.l }}),
	0x47: NewOpCode(0x47, "LD B,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.regs.a }}),

	0x48: NewOpCode(0x48, "LD C,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = cpu.regs.b }}),
	0x49: NewOpCode(0x49, "LD C,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
	0x4a: NewOpCode(0x4a, "LD C,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = cpu.regs.d }}),
	0x4b: NewOpCode(0x4b, "LD C,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = cpu.regs.e }}),
	0x4c: NewOpCode(0x4c, "LD C,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = cpu.regs.h }}),
	0x4d: NewOpCode(0x4d, "LD C,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = cpu.regs.l }}),
	0x4f: NewOpCode(0x4f, "LD C,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = cpu.regs.a }}),

	0x50: NewOpCode(0x50, "LD D,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = cpu.regs.b }}),
	0x51: NewOpCode(0x51, "LD D,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = cpu.regs.c }}),
	0x52: NewOpCode(0x52, "LD D,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
	0x53: NewOpCode(0x53, "LD D,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = cpu.regs.e }}),
	0x54: NewOpCode(0x54, "LD D,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = cpu.regs.h }}),
	0x55: NewOpCode(0x55, "LD D,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = cpu.regs.l }}),
	0x57: NewOpCode(0x57, "LD D,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = cpu.regs.a }}),

	0x58: NewOpCode(0x58, "LD E,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = cpu.regs.b }}),
	0x59: NewOpCode(0x59, "LD E,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = cpu.regs.c }}),
	0x5a: NewOpCode(0x5a, "LD E,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = cpu.regs.d }}),
	0x5b: NewOpCode(0x5b, "LD E,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
	0x5c: NewOpCode(0x5c, "LD E,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = cpu.regs.h }}),
	0x5d: NewOpCode(0x5d, "LD E,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = cpu.regs.l }}),
	0x5f: NewOpCode(0x5f, "LD E,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = cpu.regs.a }}),

	0x60: NewOpCode(0x60, "LD H,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = cpu.regs.b }}),
	0x61: NewOpCode(0x61, "LD H,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = cpu.regs.c }}),
	0x62: NewOpCode(0x62, "LD H,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = cpu.regs.d }}),
	0x63: NewOpCode(0x63, "LD H,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = cpu.regs.e }}),
	0x64: NewOpCode(0x64, "LD H,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
	0x65: NewOpCode(0x65, "LD H,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = cpu.regs.l }}),
	0x67: NewOpCode(0x67, "LD H,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = cpu.regs.a }}),

	0x68: NewOpCode(0x68, "LD L,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = cpu.regs.b }}),
	0x69: NewOpCode(0x69, "LD L,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = cpu.regs.c }}),
	0x6a: NewOpCode(0x6a, "LD L,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = cpu.regs.d }}),
	0x6b: NewOpCode(0x6b, "LD L,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = cpu.regs.e }}),
	0x6c: NewOpCode(0x6c, "LD L,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = cpu.regs.h }}),
	0x6d: NewOpCode(0x6d, "LD L,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),
	0x6f: NewOpCode(0x6f, "LD L,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = cpu.regs.a }}),

	0x78: NewOpCode(0x78, "LD A,B", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.regs.b }}),
	0x79: NewOpCode(0x79, "LD A,C", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.regs.c }}),
	0x7a: NewOpCode(0x7a, "LD A,D", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.regs.d }}),
	0x7b: NewOpCode(0x7b, "LD A,E", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.regs.e }}),
	0x7c: NewOpCode(0x7c, "LD A,H", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.regs.h }}),
	0x7d: NewOpCode(0x7d, "LD A,L", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.regs.l }}),
	0x7f: NewOpCode(0x7f, "LD A,A", 1, 4, []func(cpu *CPU){func(cpu *CPU) {}}),

	// Load register (immediate)
	0x06: NewOpCode(0x06, "LD B,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.bus.BusRead(cpu.pc); cpu.pc++ }}),
	0x0e: NewOpCode(0x0e, "LD C,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = cpu.bus.BusRead(cpu.pc); cpu.pc++ }}),
	0x16: NewOpCode(0x16, "LD D,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = cpu.bus.BusRead(cpu.pc); cpu.pc++ }}),
	0x1e: NewOpCode(0x1e, "LD E,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = cpu.bus.BusRead(cpu.pc); cpu.pc++ }}),
	0x26: NewOpCode(0x26, "LD H,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = cpu.bus.BusRead(cpu.pc); cpu.pc++ }}),
	0x2e: NewOpCode(0x2e, "LD L,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = cpu.bus.BusRead(cpu.pc); cpu.pc++ }}),
	0x3e: NewOpCode(0x3e, "LD A,u8", 2, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(cpu.pc); cpu.pc++ }}),

	// Load register (indirect HL)
	0x46: NewOpCode(0x46, "LD B,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.b = cpu.bus.BusRead(cpu.regs.getHL()) }}),
	0x4e: NewOpCode(0x4e, "LD C,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.c = cpu.bus.BusRead(cpu.regs.getHL()) }}),
	0x56: NewOpCode(0x56, "LD D,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.d = cpu.bus.BusRead(cpu.regs.getHL()) }}),
	0x5e: NewOpCode(0x5e, "LD E,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.e = cpu.bus.BusRead(cpu.regs.getHL()) }}),
	0x66: NewOpCode(0x66, "LD H,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.h = cpu.bus.BusRead(cpu.regs.getHL()) }}),
	0x6e: NewOpCode(0x6e, "LD L,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.l = cpu.bus.BusRead(cpu.regs.getHL()) }}),
	0x7e: NewOpCode(0x7e, "LD A,(HL)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(cpu.regs.getHL()) }}),

	// Load from register (indirect HL)
	0x70: NewOpCode(0x70, "LD (HL),B", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.b) }}),
	0x71: NewOpCode(0x71, "LD (HL),C", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.c) }}),
	0x72: NewOpCode(0x72, "LD (HL),D", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.d) }}),
	0x73: NewOpCode(0x73, "LD (HL),E", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.e) }}),
	0x74: NewOpCode(0x74, "LD (HL),H", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.h) }}),
	0x75: NewOpCode(0x75, "LD (HL),L", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.l) }}),
	0x77: NewOpCode(0x77, "LD (HL),A", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.a) }}),

	// Various 8bit loads
	0x02: NewOpCode(0x02, "LD (BC),A", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getBC(), cpu.regs.a) }}),
	0x0a: NewOpCode(0x0a, "LD A,(BC)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(cpu.regs.getBC()) }}),
	0x12: NewOpCode(0x12, "LD (DE),A", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getDE(), cpu.regs.a) }}),
	0x1a: NewOpCode(0x1a, "LD A,(DE)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(cpu.regs.getDE()) }}),
	0x22: NewOpCode(0x22, "LD (HL+),A", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.a); cpu.regs.incHL() }}),
	0x2a: NewOpCode(0x2a, "LD A,(HL+)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(cpu.regs.getHL()); cpu.regs.incHL() }}),
	0x32: NewOpCode(0x32, "LD (HL-),A", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), cpu.regs.a); cpu.regs.decHL() }}),
	0x36: NewOpCode(0x36, "LD (HL),u8", 2, 12, []func(cpu *CPU){func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { cpu.bus.BusWrite(cpu.regs.getHL(), lsb) }}),
	0x3a: NewOpCode(0x3a, "LD A,(HL-)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(cpu.regs.getHL()); cpu.pc-- }}),
	0xe0: NewOpCode(0xe0, "LD (FF00+u8),A", 2, 12, []func(cpu *CPU){ // // Put memory address $FF00+n into A
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.bus.BusWrite(utils.ToUint16(lsb, 0xFF), cpu.regs.a) }}),
	0xe2: NewOpCode(0xe2, "LD (FF00+C),A", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.bus.BusWrite(utils.ToUint16(cpu.regs.c, 0xFF), cpu.regs.a) }}),
	0xea: NewOpCode(0xea, "LD (u16),A", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.bus.BusWrite(utils.ToUint16(lsb, msb), cpu.regs.a) }}),
	0xf0: NewOpCode(0xf0, "LD A,(FF00+u8)", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(utils.ToUint16(lsb, 0xFF)) }}),
	0xf2: NewOpCode(0xf2, "LD A,(FF00+C)", 1, 8, []func(cpu *CPU){func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(utils.ToUint16(cpu.regs.c, 0xFF)) }}),
	0xfa: NewOpCode(0xfa, "LD A,(u16)", 3, 16, []func(cpu *CPU){func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.pc); cpu.pc++ }, func(cpu *CPU) { cpu.regs.a = cpu.bus.BusRead(utils.ToUint16(lsb, msb)) }}),
}
