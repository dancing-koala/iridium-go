package iridium

import (
	"testing"
)

func TestNewVM(t *testing.T) {
	vm := NewVM()

	for i := range vm.registers {
		if vm.registers[i] != 0 {
			t.Errorf("Value at %d should be <0>, got <%d>", i, vm.registers[i])
		}
	}

}

func TestVmOpcode_HLT(t *testing.T) {
	vm := NewVM()

	testProgram := []uint8{0, 0, 0, 0}
	vm.program = testProgram

	vm.runOnce()

	if vm.pc != 1 {
		t.Errorf("expected <%d>, got <%d>", 1, vm.pc)
	}
}

func TestVmOpcode_IGL(t *testing.T) {
	vm := NewVM()

	testProgram := []uint8{200, 0, 0, 0}
	vm.program = testProgram

	vm.runOnce()

	if vm.pc != 1 {
		t.Errorf("expected <%d>, got <%d>", 1, vm.pc)
	}
}

func TestVmOpcode_LOAD(t *testing.T) {
	vm := NewVM()
	vm.program = []uint8{1, 0, 1, 244}
	vm.runOnce()

	if vm.registers[0] != 500 {
		t.Errorf("Expected <%d>, got <%d>", 500, vm.registers[0])
	}
}

func TestVmOpcode_ADD(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 8
	vm.registers[1] = 3
	vm.program = []uint8{2, 0, 1, 2}
	vm.runOnce()

	if vm.registers[2] != 11 {
		t.Errorf("Expected <%d>, got <%d>", 12, vm.registers[2])
	}
}

func TestVmOpcode_SUB(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 8
	vm.registers[1] = 3
	vm.program = []uint8{3, 0, 1, 2}
	vm.runOnce()

	if vm.registers[2] != 5 {
		t.Errorf("Expected <%d>, got <%d>", 12, vm.registers[2])
	}
}

func TestVmOpcode_MUL(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 8
	vm.registers[1] = 3
	vm.program = []uint8{4, 0, 1, 2}
	vm.runOnce()

	if vm.registers[2] != 24 {
		t.Errorf("Expected <%d>, got <%d>", 12, vm.registers[2])
	}
}

func TestVmOpcode_DIV(t *testing.T) {
	vm := NewVM()
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

func TestVmOpcode_JMP(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 2
	vm.program = []uint8{6, 0, 0, 0}
	vm.runOnce()

	if vm.pc != 2 {
		t.Errorf("Expected <%d>, got <%d>", 2, vm.pc)
	}
}

func TestVmOpcode_JMPF(t *testing.T) {
	vm := NewVM()
	vm.program = []uint8{7, 3, 0, 0, 0, 0, 9}
	vm.runOnce()

	if vm.pc != 5 {
		t.Errorf("Expected <%d>, got <%d>", 5, vm.pc)
	}
}

func TestVmOpcode_JMPB(t *testing.T) {
	vm := NewVM()
	vm.program = []uint8{8, 2, 0, 0, 0, 0, 9}
	vm.runOnce()

	if vm.pc != 0 {
		t.Errorf("Expected <%d>, got <%d>", 0, vm.pc)
	}
}

func TestVmOpcode_EQ_True(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 4
	vm.program = []uint8{9, 0, 1, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}
}

func TestVmOpcode_EQ_False(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 3
	vm.program = []uint8{9, 0, 1, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}
}

func TestVmOpcode_NEQ_True(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 3
	vm.program = []uint8{10, 0, 1, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}
}

func TestVmOpcode_NEQ_False(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 4
	vm.program = []uint8{10, 0, 1, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}
}

func TestVmOpcode_GT_True(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 3
	vm.program = []uint8{11, 0, 1, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}
}

func TestVmOpcode_GT_False(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 5
	vm.program = []uint8{11, 0, 1, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}
}

func TestVmOpcode_LT_True(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 3
	vm.registers[1] = 4
	vm.program = []uint8{12, 0, 1, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}
}

func TestVmOpcode_LT_False(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 3
	vm.program = []uint8{12, 0, 1, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}
}

func TestVmOpcode_GTQ_True(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 2
	vm.registers[1] = 2
	vm.program = []uint8{13, 0, 1, 0, 13, 0, 1, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}

	vm.registers[0] = 3

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}
}

func TestVmOpcode_GTQ_False(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 2
	vm.registers[1] = 3
	vm.program = []uint8{13, 0, 1, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}
}

func TestVmOpcode_LTQ_True(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 2
	vm.registers[1] = 2
	vm.program = []uint8{14, 0, 1, 0, 14, 0, 1, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}

	vm.registers[0] = 1

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}
}

func TestVmOpcode_LTQ_False(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 2
	vm.registers[1] = 1
	vm.program = []uint8{14, 0, 1, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}
}

func TestVmOpcode_JEQ_Jump(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 4
	vm.registers[2] = 8
	vm.program = []uint8{9, 0, 1, 0, 15, 2, 0, 0, 0, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}

	vm.runOnce()

	if vm.pc != 8 {
		t.Errorf("Expected <%d>, got <%d>", 8, vm.pc)
	}
}

func TestVmOpcode_JEQ_NoJump(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 5
	vm.registers[2] = 8
	vm.program = []uint8{9, 0, 1, 0, 15, 2, 0, 0, 0, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}

	vm.runOnce()

	if vm.pc != 6 {
		t.Errorf("Expected <%d>, got <%d>", 6, vm.pc)
	}
}

func TestVmOpcode_JNEQ_Jump(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 5
	vm.registers[2] = 8
	vm.program = []uint8{9, 0, 1, 0, 16, 2, 0, 0, 0, 0}
	vm.runOnce()

	if vm.equalFlag {
		t.Error("equalFlag should be false!")
	}

	vm.runOnce()

	if vm.pc != 8 {
		t.Errorf("Expected <%d>, got <%d>", 8, vm.pc)
	}
}

func TestVmOpcode_JNEQ_NoJump(t *testing.T) {
	vm := NewVM()
	vm.registers[0] = 4
	vm.registers[1] = 4
	vm.registers[2] = 8
	vm.program = []uint8{9, 0, 1, 0, 16, 2, 0, 0, 0, 0}
	vm.runOnce()

	if !vm.equalFlag {
		t.Error("equalFlag should be true!")
	}

	vm.runOnce()

	if vm.pc != 6 {
		t.Errorf("Expected <%d>, got <%d>", 6, vm.pc)
	}
}
