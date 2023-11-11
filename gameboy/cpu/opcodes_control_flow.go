package cpu

import (
	"gogb/utils"
)

var OpCodesControlFlow = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// Control flow
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Relative jump
	0x18: NewOpCode(0x18, "JR i8", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.pc = uint16(int(cpu.pc) + int(e)) }, // todo: correct casting? e = signed_8(read(PC++)); PC = PC + e
	}),

	// Jump
	0xc3: NewOpCode(0xc3, "JP u16", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),

	// Jump to HL
	0xe9: NewOpCode(0xe9, "JP HL", 1, 4, []func(cpu *CPU){func(cpu *CPU) { cpu.pc = cpu.regs.getHL() }}),

	// Jump (conditional)
	0xc2: NewOpCode(0xc2, "JP NZ,u16", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++; stop = cpu.regs.getFlag(FLAG_ZERO_Z_BIT) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xca: NewOpCode(0xca, "JP Z,u16", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++; stop = !cpu.regs.getFlag(FLAG_ZERO_Z_BIT) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xd2: NewOpCode(0xd2, "JP NC,u16", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++; stop = cpu.regs.getFlag(FLAG_CARRY_C_BIT) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xda: NewOpCode(0xda, "JP C,u16", 3, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++; stop = !cpu.regs.getFlag(FLAG_CARRY_C_BIT) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),

	// Relative jump (conditional)
	0x20: NewOpCode(0x20, "JR NZ,i8", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.Read(cpu.pc); cpu.pc++; stop = cpu.regs.getFlag(FLAG_ZERO_Z_BIT) },
		func(cpu *CPU) { cpu.pc = uint16(int(cpu.pc) + int(int8(e))) },
	}),
	0x28: NewOpCode(0x28, "JR Z,i8", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.Read(cpu.pc); cpu.pc++; stop = !cpu.regs.getFlag(FLAG_ZERO_Z_BIT) },
		func(cpu *CPU) { cpu.pc = uint16(int(cpu.pc) + int(int8(e))) },
	}),
	0x30: NewOpCode(0x30, "JR NC,i8", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.Read(cpu.pc); cpu.pc++; stop = cpu.regs.getFlag(FLAG_CARRY_C_BIT) },
		func(cpu *CPU) { cpu.pc = uint16(int(cpu.pc) + int(int8(e))) },
	}),
	0x38: NewOpCode(0x38, "JR C,i8", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.Read(cpu.pc); cpu.pc++; stop = !cpu.regs.getFlag(FLAG_CARRY_C_BIT) },
		func(cpu *CPU) { cpu.pc = uint16(int(cpu.pc) + int(int8(e))) },
	}),

	// Call function,  Call function (conditional)
	0xcd: NewOpCode(0xcd, "CALL u16", 3, 24, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xc4: NewOpCode(0xc4, "CALL NZ,u16", 3, 24, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++; stop = cpu.regs.getFlag(FLAG_ZERO_Z_BIT) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xcc: NewOpCode(0xcc, "CALL Z,u16", 3, 24, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++; stop = !cpu.regs.getFlag(FLAG_ZERO_Z_BIT) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xd4: NewOpCode(0xd4, "CALL NC,u16", 3, 24, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++; stop = cpu.regs.getFlag(FLAG_CARRY_C_BIT) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xdc: NewOpCode(0xdc, "CALL C,u16", 3, 24, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.pc); cpu.pc++; stop = !cpu.regs.getFlag(FLAG_CARRY_C_BIT) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),

	// Return from function, Return from function (conditional), Return from interrupt handler
	0xc9: NewOpCode(0xc9, "RET", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xc0: NewOpCode(0xc0, "RET NZ", 1, 20, []func(cpu *CPU){
		func(cpu *CPU) { stop = cpu.regs.getFlag(FLAG_ZERO_Z_BIT) },
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xd0: NewOpCode(0xd0, "RET NC", 1, 20, []func(cpu *CPU){
		func(cpu *CPU) { stop = cpu.regs.getFlag(FLAG_CARRY_C_BIT) },
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xc8: NewOpCode(0xc8, "RET Z", 1, 20, []func(cpu *CPU){
		func(cpu *CPU) { stop = !cpu.regs.getFlag(FLAG_ZERO_Z_BIT) },
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xd8: NewOpCode(0xd8, "RET C", 1, 20, []func(cpu *CPU){
		func(cpu *CPU) { stop = !cpu.regs.getFlag(FLAG_CARRY_C_BIT) },
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) },
	}),
	0xd9: NewOpCode(0xd9, "RETI /*todo*/", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.Read(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.pc = utils.ToUint16(lsb, msb) /*todo intManager.IME = true;*/ },
	}),

	// Restart / Call function (implied)
	0xc7: NewOpCode(0xc7, "RST 00h", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = 0x00 },
	}),
	0xcf: NewOpCode(0xcf, "RST 08h", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = 0x08 },
	}),
	0xd7: NewOpCode(0xd7, "RST 10h /*todo*/", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = 0x10 },
	}),
	0xdf: NewOpCode(0xdf, "RST 18h", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = 0x18 },
	}),
	0xe7: NewOpCode(0xe7, "RST 20h", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = 0x20 },
	}),
	0xef: NewOpCode(0xef, "RST 28h", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = 0x28 },
	}),
	0xf7: NewOpCode(0xf7, "RST 30h", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = 0x30 },
	}),
	0xff: NewOpCode(0xff, "RST 38h", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Msb(cpu.pc)) },
		func(cpu *CPU) { cpu.sp--; cpu.bus.Write(cpu.sp, utils.Lsb(cpu.pc)) },
		func(cpu *CPU) { cpu.pc = 0x38 },
	}),
}
