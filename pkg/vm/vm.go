package vm

import (
	"fmt"
)

type VM struct {
	registers [32]int32
	pc        int
	program   []uint8
}

func New() *VM {
	return &VM{
		pc:      0,
		program: make([]uint8, 0, 2),
	}
}

func (self *VM) run() {
	for {
		if self.pc >= len(self.program) {
			break
		}

		switch self.decodeOpcode() {
		case OPCODE_HLT:
			fmt.Println("HLT opcode")
			return

		default:
			fmt.Println("Unrecognized opcode found, terminating")
			return

		}
	}
}

func (self *VM) decodeOpcode() Opcode {
	opcode := opcodeFor(self.program[self.pc])

	self.pc++

	return opcode
}
