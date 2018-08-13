package iridium

type Opcode uint8

const (
	OPCODE_HLT Opcode = 0
	OPCODE_IGL Opcode = 1
)

type Instruction struct {
	opcode Opcode
}

func NewInstruction(opcode Opcode) *Instruction {
	return &Instruction{
		opcode: opcode,
	}
}

func opcodeFor(val uint8) Opcode {
	switch val {
	case 0:
		return OPCODE_HLT
	default:
		return OPCODE_IGL
	}
}
