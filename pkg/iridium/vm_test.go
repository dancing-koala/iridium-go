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
	vm.registers[0] = 8
	vm.registers[1] = 3
	vm.program = []uint8{2, 0, 1, 2}
	vm.runOnce()

	if vm.registers[2] != 11 {
		t.Errorf("Expected <%d>, got <%d>", 12, vm.registers[2])
	}
}

func TestVmRun_SUB(t *testing.T) {
	vm := New()
	vm.registers[0] = 8
	vm.registers[1] = 3
	vm.program = []uint8{3, 0, 1, 2}
	vm.runOnce()

	if vm.registers[2] != 5 {
		t.Errorf("Expected <%d>, got <%d>", 12, vm.registers[2])
	}
}

func TestVmRun_MUL(t *testing.T) {
	vm := New()
	vm.registers[0] = 8
	vm.registers[1] = 3
	vm.program = []uint8{4, 0, 1, 2}
	vm.runOnce()

	if vm.registers[2] != 24 {
		t.Errorf("Expected <%d>, got <%d>", 12, vm.registers[2])
	}
}

func TestVmRun_DIV(t *testing.T) {
	vm := New()
	vm.registers[0] = 8
	vm.registers[1] = 3
	vm.program = []uint8{5, 0, 1, 2}
	vm.runOnce()

	if vm.registers[2] != 2 {
		t.Errorf("Expected <%d>, got <%d>", 2, vm.registers[2])
	}

	if vm.remainder != 2 {
		t.Errorf("Expected remainder <%d>, got <%d>", 2, vm.remainder)
	}
}

func TestVmRun_JMP(t *testing.T) {
	vm := New()
	vm.registers[0] = 2
	vm.program = []uint8{6, 0, 0, 0}
	vm.runOnce()

	if vm.pc != 2 {
		t.Errorf("Expected <%d>, got <%d>", 2, vm.pc)
	}
}

func TestVmRun_JMPF(t *testing.T) {
	vm := New()
	vm.program = []uint8{7, 3, 0, 0, 0, 0, 9}
	vm.runOnce()

	if vm.pc != 5 {
		t.Errorf("Expected <%d>, got <%d>", 5, vm.pc)
	}
}

func TestVmRun_JMPB(t *testing.T) {
	vm := New()
	vm.program = []uint8{8, 2, 0, 0, 0, 0, 9}
	vm.runOnce()

	if vm.pc != 0 {
		t.Errorf("Expected <%d>, got <%d>", 0, vm.pc)
	}
}

func TestVmRun_EQ_True(t *testing.T) {
	vm := New()
	vm.registers[0] = 4
	vm.registers[1] = 4
	vm.program = []uint8{9, 0, 1, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}
}

func TestVmRun_EQ_False(t *testing.T) {
	vm := New()
	vm.registers[0] = 4
	vm.registers[1] = 3
	vm.program = []uint8{9, 0, 1, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}
}
