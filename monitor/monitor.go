package monitor

import (
	"fmt"
	"github.com/karmakaze/irata65/disassembler"
	"github.com/karmakaze/irata65/processor"
)

func Dump(cpu *processor.CPU) {
	instruction := disassembler.DecodeInstruction(cpu.Ram, cpu.PC)

	fmt.Printf("A:%02x X:%02x Y:%02x P:%02x S:%02x PC:%04x |%02x%02x%02x| %s\n",
		cpu.A, cpu.X, cpu.Y, cpu.P, cpu.S, cpu.PC,
		cpu.Ram[cpu.PC], cpu.Ram[cpu.PC+1], cpu.Ram[cpu.PC+2], instruction)
}
