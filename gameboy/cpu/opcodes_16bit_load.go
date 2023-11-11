package cpu

import (
	"fmt"
	"gogb/utils"
)

var OpCodes16bitLoadGenerated = map[uint8]OpCode{
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~
	// 16BIT LOAD
	// ~~~~~~~~~~~~~~~~~~~~~~~~~~

	// Load 16-bit register / register pair
	0x01: NewOpCode(0x01, "LD BC,u16", 3, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.pc); cpu.pc++; cpu.regs.setBC(utils.ToUint16(lsb, msb)) },
	}),
	0x11: NewOpCode(0x11, "LD DE,u16", 3, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.pc); cpu.pc++; cpu.regs.setDE(utils.ToUint16(lsb, msb)) },
	}),
	0x21: NewOpCode(0x21, "LD HL,u16", 3, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.pc); cpu.pc++; cpu.regs.setHL(utils.ToUint16(lsb, msb)) },
	}),
	0x31: NewOpCode(0x31, "LD SP,u16", 3, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.pc); cpu.pc++; cpu.sp = utils.ToUint16(lsb, msb) },
	}),

	// Various stack
	0x08: NewOpCode(0x08, "LD (u16),SP", 3, 20, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) { cpu.bus.BusWrite(utils.ToUint16(lsb, msb), utils.Lsb(cpu.sp)) },
		func(cpu *CPU) { cpu.bus.BusWrite(utils.ToUint16(lsb, msb)+1, utils.Msb(cpu.sp)) },
	}),
	0xf9: NewOpCode(0xf9, "LD SP,HL", 1, 8, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp = cpu.regs.getHL() },
	}),
	0xf8: NewOpCode(0xf8, "LD HL,SP+i8 /*todo*/", 2, 12, []func(cpu *CPU){
		func(cpu *CPU) { e = cpu.bus.BusRead(cpu.pc); cpu.pc++ },
		func(cpu *CPU) {
			fmt.Printf("sp: %02x e: %02x\n", cpu.sp, e)
			println(e)
			cpu.regs.setHL(addSigned8(cpu, cpu.sp, e))
		},
	}),
	// Push to stack
	0xc5: NewOpCode(0xc5, "PUSH BC", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp-- },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.sp, utils.Msb(cpu.regs.getBC())); cpu.sp-- },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.sp, utils.Lsb(cpu.regs.getBC())) },
	}),
	0xd5: NewOpCode(0xd5, "PUSH DE", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp-- },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.sp, utils.Msb(cpu.regs.getDE())); cpu.sp-- },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.sp, utils.Lsb(cpu.regs.getDE())) },
	}),
	0xe5: NewOpCode(0xe5, "PUSH HL", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp-- },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.sp, utils.Msb(cpu.regs.getHL())); cpu.sp-- },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.sp, utils.Lsb(cpu.regs.getHL())) },
	}),
	0xf5: NewOpCode(0xf5, "PUSH AF", 1, 16, []func(cpu *CPU){
		func(cpu *CPU) { cpu.sp-- },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.sp, utils.Msb(cpu.regs.getAF())); cpu.sp-- },
		func(cpu *CPU) { cpu.bus.BusWrite(cpu.sp, utils.Lsb(cpu.regs.getAF())) },
	}),

	// Pop from stack
	0xc1: NewOpCode(0xc1, "POP BC", 1, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.regs.setBC(utils.ToUint16(lsb, msb)) },
	}),
	0xd1: NewOpCode(0xd1, "POP DE", 1, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.regs.setDE(utils.ToUint16(lsb, msb)) }}),
	0xe1: NewOpCode(0xe1, "POP HL", 1, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.regs.setHL(utils.ToUint16(lsb, msb)) }}),
	0xf1: NewOpCode(0xf1, "POP AF", 1, 12, []func(cpu *CPU){
		func(cpu *CPU) { lsb = cpu.bus.BusRead(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { msb = cpu.bus.BusRead(cpu.sp); cpu.sp++ },
		func(cpu *CPU) { cpu.regs.setAF(utils.ToUint16(lsb, msb) & 0xFFF0) }}), // lower 4 bits of F reg is always 0b0000
}
