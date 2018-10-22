package iridium

type Opcode uint8

const (
	opcodeHLT  Opcode = 0
	opcodeLOAD Opcode = 1
	opcodeADD  Opcode = 2
	opcodeSUB  Opcode = 3
	opcodeMUL  Opcode = 4
	opcodeDIV  Opcode = 5
	opcodeJMP  Opcode = 6
	opcodeJMPF Opcode = 7
	opcodeJMPB Opcode = 8
	opcodeEQ   Opcode = 9
	opcodeNEQ  Opcode = 10
	opcodeGT   Opcode = 11
	opcodeLT   Opcode = 12
	opcodeGTQ  Opcode = 13
	opcodeLTQ  Opcode = 14
	opcodeJEQ  Opcode = 15
	opcodeJNEQ Opcode = 16
	opcodeIGL  Opcode = 255
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
