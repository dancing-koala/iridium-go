package iridium

import (
	"fmt"
	"testing"

	"github.com/bbuck/go-lexer"
)

func TestLexer_whitespace(t *testing.T) {
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

func TestLexer_opcodeToken(t *testing.T) {
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

	actual := getTokens(l, len(expected))

	assertTokens(t, expected, actual)
}

func TestLexer_registerToken(t *testing.T) {
	src := "$1  $234 \n\t$45\r \t $098 \n\n$9847\t"
	expected := []*lexer.Token{
		newToken(registerToken, "1"),
		newToken(registerToken, "234"),
		newToken(registerToken, "45"),
		newToken(registerToken, "098"),
		newToken(registerToken, "9847"),
	}

	l := lexer.New(src, registerState)
	l.StartSync()

	actual := getTokens(l, len(expected))

	assertTokens(t, expected, actual)
}

func TestLexer_intOperandToken(t *testing.T) {
	src := "#23 \t\n\t #987 \t  \r#007 \n\n\t #378 #123456 \n"
	expected := []*lexer.Token{
		newToken(intOperandToken, "23"),
		newToken(intOperandToken, "987"),
		newToken(intOperandToken, "007"),
		newToken(intOperandToken, "378"),
		newToken(intOperandToken, "123456"),
	}

	l := lexer.New(src, intOperandState)
	l.StartSync()

	actual := getTokens(l, len(expected))

	assertTokens(t, expected, actual)
}

func TestLexer_mixedTokens(t *testing.T) {
	src := "LOAD $1 #404\nADD $1 $2 $3\n\tJMPF #22\r\t\nJMPB #23 DIV $1 $2 $3"
	expected := []*lexer.Token{
		newToken(opcodeToken, "LOAD"),
		newToken(registerToken, "1"),
		newToken(intOperandToken, "404"),
		newToken(opcodeToken, "ADD"),
		newToken(registerToken, "1"),
		newToken(registerToken, "2"),
		newToken(registerToken, "3"),
		newToken(opcodeToken, "JMPF"),
		newToken(intOperandToken, "22"),
		newToken(opcodeToken, "JMPB"),
		newToken(intOperandToken, "23"),
		newToken(opcodeToken, "DIV"),
		newToken(registerToken, "1"),
		newToken(registerToken, "2"),
		newToken(registerToken, "3"),
	}

	l := lexer.New(src, opcodeState)
	l.StartSync()

	actual := getTokens(l, len(expected))

	assertTokens(t, expected, actual)
}

func newToken(t lexer.TokenType, value string) *lexer.Token {
	return &lexer.Token{
		Type:  t,
		Value: value,
	}
}

func getTokens(l *lexer.L, initialCap int) []*lexer.Token {
	result := make([]*lexer.Token, 0, initialCap)

	for {
		item, done := l.NextToken()

		if done {
			break
		}

		result = append(result, item)
	}

	return result
}

func assertTokens(t *testing.T, expected, actual []*lexer.Token) {
	if !assertIntMsg(t, "Comparing length", len(expected), len(actual)) {
		return
	}

	for i := range actual {
		assertTokenMsg(t, fmt.Sprintf("Comparing tokens at index <%d>", i), expected[i], actual[i])
	}
}

func TestStrToUint8(t *testing.T) {
	vals := []string{"128", "0", "255", "32", "33", "21", "1"}
	expected := []uint8{128, 0, 255, 32, 33, 21, 1}

	for i := 0; i < len(vals); i++ {
		actual, err := strToUint8(vals[i])
		assertNilMsg(t, fmt.Sprintf("Checking err at index <%d>", i), err)
		assertUint8Msg(t, fmt.Sprintf("Comparing values at index <%d>", i), expected[i], actual)
	}
}

func TestStrToInt32(t *testing.T) {
	vals := []string{"-12822", "404", "255012", "-31112", "323", "21", "1"}
	expected := []int32{-12822, 404, 255012, -31112, 323, 21, 1}

	for i := 0; i < len(vals); i++ {
		actual, err := strToInt32(vals[i])
		assertNilMsg(t, fmt.Sprintf("Checking err at index <%d>", i), err)
		assertInt32Msg(t, fmt.Sprintf("Comparing values at index <%d>", i), expected[i], actual)
	}
}
