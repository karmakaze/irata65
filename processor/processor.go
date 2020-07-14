package processor

type Registers struct {
	A  uint8
	X  uint8
	Y  uint8
	P  uint8
	S  uint8
	PC uint16
}

type Memory struct {
	Ram []uint8
}

type CPU struct {
	Registers
	Memory
}

func NewCPU() *CPU {
	return &CPU{
		Registers: Registers{
			A:  0,
			X:  0,
			Y:  0,
			P:  0x04, // N V - B D I Z C
			S:  0xFD,
			PC: 0x0200,
		},
		Memory: Memory{
			Ram: make([]uint8, 1024, 1024),
		},
	}
}

func (cpu *CPU) Step() {
	opcode := cpu.FetchInstructionByte()

	switch opcode {
	case 0x60: // rts
		cpu.PC = cpu.PopStackWord()
	case 0x69: // adc
		m8 := cpu.FetchInstructionByte()
		if uint16(cpu.A)+uint16(m8) >= 0x0100 {
			cpu.P |= 0x01
		}
		cpu.A += m8
		if cpu.A == 0 {
			cpu.P |= 0x02 // n v - b d i Z c
		}
	case 0x18: // clc
		cpu.P &= 0xFE // N V - B D I Z c
	case 0x38: // sec
		cpu.P |= 0x01 // n v - b d i z C

	case 0xB8: // clv
		cpu.P &= 0xBF // N v - B D I Z C

	case 0x58: // cli
		cpu.P &= 0xFB // N V - B D i Z C
	case 0x78: // sei
		cpu.P |= 0x04 // n v - b d I z c

	case 0xD8: // cld
		cpu.P &= 0xF7 // N V - B d I Z C
	case 0xF8: // sed
		cpu.P |= 0x08 // n v - b D i z c
	}
}

func (cpu *CPU) FetchInstructionByte() uint8 {
	d8 := cpu.Ram[cpu.PC]
	cpu.PC++
	return d8
}

func (cpu *CPU) FetchInstructionWord() uint16 {
	d8 := cpu.Ram[cpu.PC]
	d16 := uint16(cpu.Ram[cpu.PC+1])
	d16 = (d16 << 8) | uint16(d8)
	cpu.PC += 2
	return d16
}

func (cpu *CPU) PopStackWord() uint16 {
	s8 := cpu.Ram[0x0100+uint16(cpu.S)]
	s16 := uint16(cpu.Ram[0x0100+uint16(cpu.S+1)])
	s16 = (s16 << 8) | uint16(s8)
	cpu.S += 2
	return s16
}

func (cpu *CPU) PushStackWord() {

}
