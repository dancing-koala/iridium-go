package vm

import (
	"testing"
)

func TestNewVM(t *testing.T) {
	testVM := New()

	for i := range testVM.registers {
		if testVM.registers[i] != 0 {
			t.Errorf("Value at %d should be <0>, got <%d>", i, testVM.registers[i])
		}
	}

}
