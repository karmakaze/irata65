package main

import (
	"fmt"
	"os"

	"github.com/karmakaze/irata65/monitor"
	"github.com/karmakaze/irata65/processor"
)

func main() {
	fmt.Printf("args: %+v\n", os.Args[1:])

	cpu := processor.NewCPU()
	cpu.Ram[0x0200] = 0x60
	cpu.Ram[0x0201] = 0x45
	cpu.Ram[0x0202] = 0x03
	monitor.Dump(cpu)
	cpu.Step()
	monitor.Dump(cpu)
}
