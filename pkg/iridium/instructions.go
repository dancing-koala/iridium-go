package iridium

type Opcode uint8

const (
	OPCODE_HLT  Opcode = 0
	OPCODE_LOAD Opcode = 1
	OPCODE_ADD  Opcode = 2
	OPCODE_SUB  Opcode = 3
	OPCODE_MUL  Opcode = 4
	OPCODE_DIV  Opcode = 5
	OPCODE_JMP  Opcode = 6
	OPCODE_JMPF Opcode = 7
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

	case 2:
		return OPCODE_ADD

	case 3:
		return OPCODE_SUB

	case 4:
		return OPCODE_MUL

	case 5:
		return OPCODE_DIV

	case 6:
		return OPCODE_JMP

	case 7:
		return OPCODE_JMPF

	default:
		return OPCODE_IGL
	}
}
