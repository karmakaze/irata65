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
	opcode := cpu.Ram[cpu.PC]
	cpu.PC++
	u8 := cpu.Ram[cpu.PC]
	u16 := uint16(cpu.Ram[(cpu.PC + 1)])
	u16 = (u16 << 8) | uint16(u8)

	switch opcode {
	case 0x60: // rts
		cpu.PC = u16
	}
}
