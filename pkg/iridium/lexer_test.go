package iridium

import (
	"fmt"
	"testing"

	"github.com/bbuck/go-lexer"
)

func TestLex_whitespace(t *testing.T) {
	src := " \n \t \r  \n "
	expectedLen := 0

	l := lexer.New(src, whitespaceState)
	l.StartSync()

	actual := make([]*lexer.Token, 0, expectedLen)

	for {
		item, done := l.NextToken()

		if done {
			break
		}

		actual = append(actual, item)
	}

	if !assertIntMsg(t, "Comparing length", expectedLen, len(actual)) {
		return
	}
}

func TestLex_opcodeToken(t *testing.T) {
	src := "ABC  COUCOU \n\tHIBOU   \r  HI\nBOU "
	expected := []*lexer.Token{
		newToken(opcodeToken, "ABC"),
		newToken(opcodeToken, "COUCOU"),
		newToken(opcodeToken, "HIBOU"),
		newToken(opcodeToken, "HI"),
		newToken(opcodeToken, "BOU"),
	}

	l := lexer.New(src, opcodeState)
	l.StartSync()

	actual := make([]*lexer.Token, 0, len(expected))

	for {
		item, done := l.NextToken()

		if done {
			break
		}

		actual = append(actual, item)
	}

	if !assertIntMsg(t, "Comparing length", len(expected), len(actual)) {
		return
	}

	for i := range actual {
		assertTokenMsg(t, fmt.Sprintf("Comparing tokens at index <%d>", i), expected[i], actual[i])
	}
}

func newToken(t lexer.TokenType, value string) *lexer.Token {
	return &lexer.Token{
		Type:  t,
		Value: value,
	}
}
