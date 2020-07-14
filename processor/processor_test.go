package processor

import "testing"

func TestCLI(t *testing.T) {
	cpu := NewCPU()
	cpu.P = 0xFF
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x58 // cli

	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.P = 0xFB // N V - B D i Z C
	expectRegs.PC = 0x0201
	if expectRegs != cpu.Registers {
		t.Errorf("Registers CLI\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestSEI(t *testing.T) {
	cpu := NewCPU()
	cpu.P = 0x00
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x78 // sei

	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.P = 0x04 // n v - b d I z c
	expectRegs.PC = 0x0201
	if expectRegs != cpu.Registers {
		t.Errorf("Registers SEI\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestCLV(t *testing.T) {
	cpu := NewCPU()
	cpu.P = 0xFF
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0xB8 // clv

	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.P = 0xBF // N v - B D I Z C
	expectRegs.PC = 0x0201
	if expectRegs != cpu.Registers {
		t.Errorf("Registers CLV\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestSEC(t *testing.T) {
	cpu := NewCPU()
	cpu.P = 0x00
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x38 // sec

	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.P = 0x01 // n v - b d i z C
	expectRegs.PC = 0x0201
	if expectRegs != cpu.Registers {
		t.Errorf("Registers SEC\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestCLC(t *testing.T) {
	cpu := NewCPU()
	cpu.P = 0xFF
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x18 // clc

	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.P = 0xFE // N V - B D I Z c
	expectRegs.PC = 0x0201
	if expectRegs != cpu.Registers {
		t.Errorf("Registers CLC\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestSED(t *testing.T) {
	cpu := NewCPU()
	cpu.P = 0x00
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0xF8 // sed

	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.P = 0x08 // n v - b D i z c
	expectRegs.PC = 0x0201
	if expectRegs != cpu.Registers {
		t.Errorf("Registers SED\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestCLD(t *testing.T) {
	cpu := NewCPU()
	cpu.P = 0xFF
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0xD8 // cld

	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.P = 0xF7 // N V - M d I Z C
	expectRegs.PC = 0x0201
	if expectRegs != cpu.Registers {
		t.Errorf("Registers CLD\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestRTS(t *testing.T) {
	cpu := NewCPU()
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x60 // rts

	cpu.S = 0xFD
	cpu.Ram[0x01FD] = 0x45 // lo
	cpu.Ram[0x01FE] = 0x03 // hi

	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.PC = 0x0345
	expectRegs.S = 0xFF
	if expectRegs != cpu.Registers {
		t.Errorf("Registers RTS\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestADCBinaryImmediate(t *testing.T) {
	cpu := NewCPU()
	// test carry flag
	cpu.P = 0x00
	cpu.A = 0xA0
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x69
	cpu.Ram[0x0201] = 0x99
	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.PC = 0x0202
	expectRegs.A = 0x39 // 0xA0 + 0x99
	expectRegs.P = 0x01 // n v - b d i z C
	if expectRegs != cpu.Registers {
		t.Errorf("Registers ADC (binary C)\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}

	// test zero fiag
	cpu = NewCPU()
	cpu.P = 0x00
	cpu.A = 0x00
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x69
	cpu.Ram[0x0201] = 0x00
	expectRegs = cpu.Registers

	cpu.Step()

	expectRegs.PC = 0x0202
	expectRegs.A = 0x00 // 0x00 + 0x00
	expectRegs.P = 0x02 // n v - b d i Z c
	if expectRegs != cpu.Registers {
		t.Errorf("Registers ADC (binary Z)\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}

func TestADCDecimalImmediate(t *testing.T) {
	cpu := NewCPU()
	cpu.P = 0x00
	cpu.A = 0x00
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x69
	cpu.Ram[0x0201] = 0x00
	expectRegs := cpu.Registers

	cpu.Step()

	expectRegs.PC = 0x0202
	expectRegs.A = 0x00 // 0x00 + 0x00
	expectRegs.P = 0x02 // n v - b d i Z c
	if expectRegs != cpu.Registers {
		t.Errorf("Registers ADC (decimal C)\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}

	// test zero fiag
	cpu = NewCPU()
	cpu.P = 0x00
	cpu.A = 0x00
	cpu.PC = 0x0200
	cpu.Ram[0x0200] = 0x69
	cpu.Ram[0x0201] = 0x00
	expectRegs = cpu.Registers

	cpu.Step()

	expectRegs.PC = 0x0202
	expectRegs.A = 0x00 // 0x00 + 0x00
	expectRegs.P = 0x02 // n v - b d i Z c
	if expectRegs != cpu.Registers {
		t.Errorf("Registers ADC (decimal Z)\nexpected: %+v\n  actual: %+v", expectRegs, cpu.Registers)
	}
}
