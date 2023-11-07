package cpu

var OpCodesRotateShiftBitoperations = map[uint8]OpCode{
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~
    // Rotates, Shifts, Bit operations 
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~

    // RLCA
    0x07: NewOpCode(0x07, "RLCA", 1, 4, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0x17: NewOpCode(0x17, "RLA", 1, 4, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0x0f: NewOpCode(0x0f, "RRCA", 1, 4, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0x1f: NewOpCode(0x1f, "RRA", 1, 4, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),

    // Increment (register), Increment SP
    0xcb: NewOpCode(0xcb, "PREFIX CB", 1, 4, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),


}
