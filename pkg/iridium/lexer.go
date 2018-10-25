package iridium

import (
	"fmt"

	"github.com/bbuck/go-lexer"
)

const (
	opcodeToken lexer.TokenType = iota
	registerToken
	intOperandToken
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

	if isUpperAlpha(r) {
		return opcodeState
	}

	return nil
}

func isUpperAlpha(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
