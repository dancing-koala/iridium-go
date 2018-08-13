package iridium

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

func (vm *VM) run() {
	for {
		if vm.pc >= len(vm.program) {
			break
		}

		switch vm.decodeOpcode() {
		case OPCODE_HLT:
			fmt.Println("HLT opcode")
			return

		default:
			fmt.Println("Unrecognized opcode found, terminating")
			return

		}
	}
}

func (vm *VM) decodeOpcode() Opcode {
	opcode := opcodeFor(vm.program[vm.pc])

	vm.pc++

	return opcode
}
