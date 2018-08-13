package vm

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
