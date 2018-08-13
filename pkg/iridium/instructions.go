package iridium

type Opcode uint8

const (
	OPCODE_HLT  Opcode = 0
	OPCODE_LOAD Opcode = 1
	OPCODE_IGL  Opcode = 255
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

	case 1:
		return OPCODE_LOAD

	default:
		return OPCODE_IGL
	}
}
