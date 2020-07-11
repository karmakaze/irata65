package disassembler

func DecodeInstruction(bytes[] uint8, pc uint16) string {
	instruction := "?"

	opcode := bytes[pc]
	lobyte := bytes[pc+1]
	hibyte := bytes[pc+2]
	_ = (uint16(hibyte) << 8) | uint16(lobyte)

	switch opcode {
	case 0x60:
		instruction = "rts"
	}

	return instruction
}
