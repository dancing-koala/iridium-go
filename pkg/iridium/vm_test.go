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

	vm.run()

	if vm.pc != 1 {
		t.Errorf("expected <%d>, got <%d>", 1, vm.pc)
	}
}
func TestVmRun_IGL(t *testing.T) {
	vm := New()

	testProgram := []uint8{200, 0, 0, 0}
	vm.program = testProgram

	vm.run()

	if vm.pc != 1 {
		t.Errorf("expected <%d>, got <%d>", 1, vm.pc)
	}
}
