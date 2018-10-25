package iridium

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/bbuck/go-lexer"
)

const (
	opcodeToken lexer.TokenType = iota
	registerToken
	intOperandToken
	numberToken
)

func lex(src string) {
	l := lexer.New(src, opcodeState)
	l.Start()
}

func opcodeState(l *lexer.L) lexer.StateFunc {
	r := l.Next()

	for isUpperAlpha(r) {
		r = l.Next()
	}

	l.Rewind()
	l.Emit(opcodeToken)

	return whitespaceState
}

func whitespaceState(l *lexer.L) lexer.StateFunc {
	r := l.Next()

	if r == lexer.EOFRune {
		return nil
	}

	if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
		l.Error(fmt.Sprintf("Unexpected token %q", r))
		return nil
	}

	l.Take(" \t\n\r")
	l.Ignore()

	r = l.Peek()

	switch {
	case isUpperAlpha(r):
		return opcodeState

	case isRegisterSymbol(r):
		return registerState

	case isIntOperandSymbol(r):
		return intOperandState
	}

	return nil
}

func registerState(l *lexer.L) lexer.StateFunc {
	r := l.Next()

	if !isRegisterSymbol(r) {
		l.Error(fmt.Sprintf("Not a register symbol: %q", r))
		return nil
	}

	l.Ignore() //Drop the symbol
	l.Take("0123456789")

	l.Emit(registerToken)

	return whitespaceState
}

func intOperandState(l *lexer.L) lexer.StateFunc {
	r := l.Next()

	if !isIntOperandSymbol(r) {
		l.Error(fmt.Sprintf("Not an int operand symbol: %q", r))
	}

	l.Ignore() // Drop the # symbol
	l.Take("0123456789")

	l.Emit(intOperandToken)

	return whitespaceState
}

func isUpperAlpha(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func isRegisterSymbol(r rune) bool {
	return r == '$'
}

func isIntOperandSymbol(r rune) bool {
	return r == '#'
}

func convertValue(t *lexer.Token) (interface{}, error) {

	switch t.Type {
	case opcodeToken:
		return strToOpcode(t.Value)
	case registerToken:
		return strToUint8(t.Value)

	case intOperandToken:
		return strToInt32(t.Value)
	}

	return nil, fmt.Errorf("Unknown token type: %+v", t)
}

func strToOpcode(val string) (Opcode, error) {
	switch val {
	case "HLT":
		return opcodeHLT, nil
	case "LOAD":
		return opcodeLOAD, nil
	case "ADD":
		return opcodeADD, nil
	case "SUB":
		return opcodeSUB, nil
	case "MUL":
		return opcodeMUL, nil
	case "DIV":
		return opcodeDIV, nil
	case "JMP":
		return opcodeJMP, nil
	case "JMPF":
		return opcodeJMPF, nil
	case "JMPB":
		return opcodeJMPB, nil
	case "EQ":
		return opcodeEQ, nil
	case "NEQ":
		return opcodeNEQ, nil
	case "GT":
		return opcodeGT, nil
	case "LT":
		return opcodeLT, nil
	case "GTQ":
		return opcodeGTQ, nil
	case "LTQ":
		return opcodeLTQ, nil
	case "JEQ":
		return opcodeJEQ, nil
	case "JNEQ":
		return opcodeJNEQ, nil
	case "IGL":
		return opcodeIGL, nil
	}

	return 0, errors.New("Unknown opcode: " + val)
}

func strToUint8(val string) (uint8, error) {
	result, err := strconv.ParseUint(val, 10, 8)
	return uint8(result), err
}

func strToInt32(val string) (int32, error) {
	result, err := strconv.ParseInt(val, 10, 32)
	return int32(result), err
}
