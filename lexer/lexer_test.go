package lexer

import (
	"monkeylang/token"
	"os"
	"testing"
)

type LexerTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestVarAssign(t *testing.T) {
	input := readSourceFile(t, "source_01.mke")
	tests := []LexerTest{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}

	tryLexerTestsOnInput(t, input, tests)
}

func TestFunctionCreate(t *testing.T) {
	input := readSourceFile(t, "source_02.mke")
	tests := []LexerTest{
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
	}

	tryLexerTestsOnInput(t, input, tests)
}

func TestNextToken(t *testing.T) {

	input := readSourceFile(t, "source_03.mke")
	tests := []LexerTest{

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "3"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.EOF, ""},
	}
	tryLexerTestsOnInput(t, input, tests)
}

func tryLexerTestsOnInput(t testing.TB, input string, tests []LexerTest) {
	l := New(input)
	for i, lexerTest := range tests {
		tok := l.NextToken()
		if tok.Type != lexerTest.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q, Position=%v", i, lexerTest.expectedType, tok.Type, l.readPosition)
		}
		if tok.Literal != lexerTest.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q, Position=%v", i, lexerTest.expectedLiteral, tok.Literal, l.readPosition)
		}
	}
}

func readSourceFile(t testing.TB, fileName string) string {
	pwd, _ := os.Getwd()
	fileUri := pwd + "/../data/" + fileName
	input, err := os.ReadFile(fileUri)

	if err != nil {
		t.Errorf("Could not read file content: %v", fileUri)
	}

	return string(input)
}
