package iridium

import (
	"fmt"
)

type VM struct {
	registers [32]int32
	pc        int
	program   []uint8
	remainder uint32
	equalFlag bool
}

func NewVM() *VM {
	return &VM{
		program: make([]uint8, 0, 2),
	}
}

func (vm *VM) run() {
	for vm.executeInstruction() {
		//Intentionally empty
	}
}

func (vm *VM) runOnce() {
	vm.executeInstruction()
}

func (vm *VM) executeInstruction() bool {
	if vm.pc >= len(vm.program) {
		return false
	}

	switch vm.decodeOpcode() {
	case opcodeHLT:
		fmt.Println("HLT opcode")
		return false

	case opcodeLOAD:
		reg := int(vm.next8Bits())
		number := vm.next16Bits()
		vm.registers[reg] = int32(number)

	case opcodeADD:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		regTarget := int(vm.next8Bits())
		vm.registers[regTarget] = vm.registers[regA] + vm.registers[regB]

	case opcodeSUB:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		regTarget := int(vm.next8Bits())
		vm.registers[regTarget] = vm.registers[regA] - vm.registers[regB]

	case opcodeMUL:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		regTarget := int(vm.next8Bits())
		vm.registers[regTarget] = vm.registers[regA] * vm.registers[regB]

	case opcodeDIV:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		regTarget := int(vm.next8Bits())
		vm.registers[regTarget] = vm.registers[regA] / vm.registers[regB]
		vm.remainder = uint32(vm.registers[regA] % vm.registers[regB])

	case opcodeJMP:
		reg := int(vm.next8Bits())
		vm.pc = int(vm.registers[reg])

	case opcodeJMPF:
		vm.pc += int(vm.next8Bits())

	case opcodeJMPB:
		vm.pc -= int(vm.next8Bits())

	case opcodeEQ:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] == vm.registers[regB]
		vm.next8Bits()

	case opcodeNEQ:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] != vm.registers[regB]
		vm.next8Bits()

	case opcodeGT:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] > vm.registers[regB]
		vm.next8Bits()

	case opcodeLT:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] < vm.registers[regB]
		vm.next8Bits()

	case opcodeGTQ:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] >= vm.registers[regB]
		vm.next8Bits()

	case opcodeLTQ:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] <= vm.registers[regB]
		vm.next8Bits()

	case opcodeJEQ:
		target := int(vm.next8Bits())

		if vm.equalFlag {
			vm.pc = int(vm.registers[target])
		}

	case opcodeJNEQ:
		target := int(vm.next8Bits())

		if !vm.equalFlag {
			vm.pc = int(vm.registers[target])
		}

	default:
		fmt.Println("Unrecognized opcode found, terminating")
		return false

	}

	return true
}

func (vm *VM) decodeOpcode() Opcode {
	opcode := opcodeFor(vm.program[vm.pc])

	vm.pc++

	return opcode
}

func (vm *VM) next8Bits() uint8 {
	val := vm.program[vm.pc]

	vm.pc++

	return val
}

func (vm *VM) next16Bits() uint16 {
	val := (uint16(vm.program[vm.pc]) << 8) | uint16(vm.program[vm.pc+1])

	vm.pc += 2

	return val
}

func (vm *VM) addToProgram(values ...uint8) {
	for _, val := range values {
		vm.program = append(vm.program, val)
	}
}
