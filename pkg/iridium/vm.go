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

func New() *VM {
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
	case OPCODE_HLT:
		fmt.Println("HLT opcode")
		return false

	case OPCODE_LOAD:
		reg := int(vm.next8Bits())
		number := vm.next16Bits()
		vm.registers[reg] = int32(number)

	case OPCODE_ADD:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		regTarget := int(vm.next8Bits())
		vm.registers[regTarget] = vm.registers[regA] + vm.registers[regB]

	case OPCODE_SUB:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		regTarget := int(vm.next8Bits())
		vm.registers[regTarget] = vm.registers[regA] - vm.registers[regB]

	case OPCODE_MUL:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		regTarget := int(vm.next8Bits())
		vm.registers[regTarget] = vm.registers[regA] * vm.registers[regB]

	case OPCODE_DIV:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		regTarget := int(vm.next8Bits())
		vm.registers[regTarget] = vm.registers[regA] / vm.registers[regB]
		vm.remainder = uint32(vm.registers[regA] % vm.registers[regB])

	case OPCODE_JMP:
		reg := int(vm.next8Bits())
		vm.pc = int(vm.registers[reg])

	case OPCODE_JMPF:
		vm.pc += int(vm.next8Bits())

	case OPCODE_JMPB:
		vm.pc -= int(vm.next8Bits())

	case OPCODE_EQ:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] == vm.registers[regB]
		vm.next8Bits()

	case OPCODE_NEQ:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] != vm.registers[regB]
		vm.next8Bits()

	case OPCODE_GT:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] > vm.registers[regB]
		vm.next8Bits()

	case OPCODE_LT:
		regA := int(vm.next8Bits())
		regB := int(vm.next8Bits())
		vm.equalFlag = vm.registers[regA] < vm.registers[regB]
		vm.next8Bits()

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
