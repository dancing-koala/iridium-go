package vm

import (
	"testing"
)

func TestNew_initialization(t *testing.T) {
	testVM := New()

	for i := range testVM.registers {
		if testVM.registers[i] != 0 {
			t.Errorf("Value at %d should be <0>, got <%d>", i, testVM.registers[i])
		}
	}

}

func TestVmRun_HLT(t *testing.T) {
	testVM := New()

	testProgram := []uint8{0, 0, 0, 0}
	testVM.program = testProgram

	testVM.run()

	if testVM.pc != 1 {
		t.Errorf("expected <%d>, got <%d>", 1, testVM.pc)
	}
}
func TestVmRun_IGL(t *testing.T) {
	testVM := New()

	testProgram := []uint8{200, 0, 0, 0}
	testVM.program = testProgram

	testVM.run()

	if testVM.pc != 1 {
		t.Errorf("expected <%d>, got <%d>", 1, testVM.pc)
	}
}
