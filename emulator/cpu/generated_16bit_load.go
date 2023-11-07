package cpu

var OpCodes16bitLoadGenerated = map[uint8]OpCode{
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~
    // 16BIT LOAD
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~

    // Load 16-bit register / register pair
    0x01: NewOpCode(0x01, "LD BC,u16", 3, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0x11: NewOpCode(0x11, "LD DE,u16", 3, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0x21: NewOpCode(0x21, "LD HL,u16", 3, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0x31: NewOpCode(0x31, "LD SP,u16", 3, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),

    // Various stack
    0x08: NewOpCode(0x08, "LD (u16),SP", 3, 20, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0xf8: NewOpCode(0xf8, "LD HL,SP+i8", 2, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0xf9: NewOpCode(0xf9, "LD SP,HL", 1, 8, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),

    // Push to stack
    0xc5: NewOpCode(0xc5, "PUSH BC", 1, 16, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0xd5: NewOpCode(0xd5, "PUSH DE", 1, 16, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0xe5: NewOpCode(0xe5, "PUSH HL", 1, 16, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0xf5: NewOpCode(0xf5, "PUSH AF", 1, 16, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),

    // Pop from stack
    0xc1: NewOpCode(0xc1, "POP BC", 1, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0xd1: NewOpCode(0xd1, "POP DE", 1, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0xe1: NewOpCode(0xe1, "POP HL", 1, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),
    0xf1: NewOpCode(0xf1, "POP AF", 1, 12, []func(cpu *CPU){func(cpu *CPU) {/*todo*/}}),


}
