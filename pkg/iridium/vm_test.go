package iridium

import (
	"testing"
)

func TestNew_initialization(t *testing.T) {
	vm := New()

	for i := range vm.registers {
		if vm.registers[i] != 0 {
			t.Errorf("Value at %d should be <0>, got <%d>", i, vm.registers[i])
		}
	}

}

func TestVmRun_HLT(t *testing.T) {
	vm := New()

	testProgram := []uint8{0, 0, 0, 0}
	vm.program = testProgram

	vm.runOnce()

	if vm.pc != 1 {
		t.Errorf("expected <%d>, got <%d>", 1, vm.pc)
	}
}

func TestVmRun_IGL(t *testing.T) {
	vm := New()

	testProgram := []uint8{200, 0, 0, 0}
	vm.program = testProgram

	vm.runOnce()

	if vm.pc != 1 {
		t.Errorf("expected <%d>, got <%d>", 1, vm.pc)
	}
}

func TestVmRun_LOAD(t *testing.T) {
	vm := New()
	vm.program = []uint8{1, 0, 1, 244}
	vm.runOnce()

	if vm.registers[0] != 500 {
		t.Errorf("Expected <%d>, got <%d>", 500, vm.registers[0])
	}
}

func TestVmRun_ADD(t *testing.T) {
	vm := New()
	vm.registers[0] = 4
	vm.registers[1] = 8
	vm.program = []uint8{2, 0, 1, 2}
	vm.runOnce()

	if vm.registers[2] != 12 {
		t.Errorf("Expected <%d>, got <%d>", 12, vm.registers[0])
	}
}
