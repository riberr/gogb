package main

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
		a: 0,
		b: 0,
		c: 0,
		d: 0,
		e: 0,
		f: 0,
		h: 0,
		l: 0,
	}
}

func (r *Registers) getFlag(flag Flag) bool {
	return utils.HasBit(r.f, int(flag))
}

func (r *Registers) setFlag(flag Flag, value bool) {
	//r.f |= 1 << flag
	r.f = utils.SetBit(r.f, int(flag))
}

func (r *Registers) getBC() uint16 {
	return (uint16(r.b) << 8) | uint16(r.c)
}

func (r *Registers) setBC(value uint16) {
	r.b = uint8((value & 0xFF00) >> 8)
	r.c = uint8(value & 0xFF)
}

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

func (r *Registers) getHL() uint16 {
	return (uint16(r.h) << 8) | uint16(r.l)
}

func (r *Registers) setHL(value uint16) {
	r.h = uint8((value & 0xFF00) >> 8)
	r.l = uint8(value & 0xFF)
}
