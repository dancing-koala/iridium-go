package iridium

import (
	"fmt"

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
