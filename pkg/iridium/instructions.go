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
	OPCODE_JMPB Opcode = 8
	OPCODE_EQ   Opcode = 9
	OPCODE_NEQ  Opcode = 10
	OPCODE_GT   Opcode = 11
	OPCODE_LT   Opcode = 12
	OPCODE_GTQ  Opcode = 13
	OPCODE_LTQ  Opcode = 14
	OPCODE_JEQ  Opcode = 15
	OPCODE_JNEQ Opcode = 16
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
	// This not really useful right now but this gives us more flexibility.
	// This might disappear in the future.
	return Opcode(val)
}
