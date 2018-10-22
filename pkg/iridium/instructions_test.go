package iridium

import (
	"testing"
)

func TestNewInstruction(t *testing.T) {
	instr := NewInstruction(opcodeHLT)

	if instr.opcode != opcodeHLT {
		t.Errorf("expected opcode <%d>, got <%d>", opcodeHLT, instr.opcode)
	}
}
