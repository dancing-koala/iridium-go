package iridium

import (
	"testing"

	"github.com/bbuck/go-lexer"
)

func assertIntMsg(t *testing.T, msg string, expected, actual int) bool {
	result := expected == actual

	if !result {
		t.Errorf("%s\nexpected <%d>, got <%d>", msg, expected, actual)
	}

	return result
}

func assertTokenMsg(t *testing.T, msg string, expected, actual *lexer.Token) bool {
	result := expected.Type == actual.Type && expected.Value == actual.Value

	if !result {
		t.Errorf("%s\nexpected <%+v>, got <%+v>", msg, expected, actual)
	}

	return result
}
