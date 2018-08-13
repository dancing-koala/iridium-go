package iridium

import (
	"testing"
)

func TestNewInstruction(t *testing.T) {
	instr := NewInstruction(OPCODE_HLT)

	if instr.opcode != OPCODE_HLT {
		t.Errorf("expected opcode <%d>, got <%d>", OPCODE_HLT, instr.opcode)
	}
}
