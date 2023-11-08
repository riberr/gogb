package cpu

import "gogb/utils"

type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f uint8
	h uint8
	l uint8
}

type Flag int

const (
	FLAG_ZERO_Z_BIT        Flag = 7
	FLAG_SUBTRACTION_N_BIT Flag = 6
	FLAG_HALF_CARRY_H_BIT  Flag = 5
	FLAG_CARRY_C_BIT       Flag = 4
)

func NewRegisters() Registers {
	return Registers{
		a: 0x01,
		b: 0,
		c: 0x13,
		d: 0,
		e: 0xD8,
		f: 0xB0,
		h: 0x01,
		l: 0x4D,
	}
}

func (r *Registers) getFlag(flag Flag) bool {
	return utils.HasBit(r.f, int(flag))
}

func (r *Registers) setFlag(flag Flag, value bool) {
	if value {
		r.f = utils.SetBit(r.f, int(flag))
	} else {
		r.f = utils.ClearBit(r.f, int(flag))
	}
}

func (r *Registers) getAF() uint16 {
	return (uint16(r.a) << 8) | uint16(r.f)
}

func (r *Registers) setAF(value uint16) {
	r.a = uint8((value & 0xFF00) >> 8)
	r.f = uint8(value & 0xFF)
}

func (r *Registers) getBC() uint16 {
	return (uint16(r.b) << 8) | uint16(r.c)
}

func (r *Registers) setBC(value uint16) {
	r.b = uint8((value & 0xFF00) >> 8)
	r.c = uint8(value & 0xFF)
}

func (r *Registers) incBC() { r.setBC(r.getBC() + 1) }
func (r *Registers) decBC() { r.setBC(r.getBC() - 1) }

/*
func (r *Registers) setBC(lsb uint8, msb uint8) {
	r.b = msb
	r.c = lsb
}
*/

func (r *Registers) getDE() uint16 {
	return (uint16(r.d) << 8) | uint16(r.e)
}

func (r *Registers) setDE(value uint16) {
	r.d = uint8((value & 0xFF00) >> 8)
	r.e = uint8(value & 0xFF)
}

func (r *Registers) incDE() { r.setDE(r.getDE() + 1) }
func (r *Registers) decDE() { r.setDE(r.getDE() - 1) }

func (r *Registers) getHL() uint16 {
	return (uint16(r.h) << 8) | uint16(r.l)
}

func (r *Registers) setHL(value uint16) {
	r.h = uint8((value & 0xFF00) >> 8)
	r.l = uint8(value & 0xFF)
}

func (r *Registers) incHL() { r.setHL(r.getHL() + 1) }
func (r *Registers) decHL() { r.setHL(r.getHL() - 1) }
