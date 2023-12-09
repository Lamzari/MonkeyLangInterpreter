package lexer

import (
	"fmt"

	"github.com/Lamzari/MonkeyLangInterpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in the input (current char)
	readPosition int  // current reading position in the input (so position + 1)
	ch           byte // current char we are examining
}

func New(input string) *Lexer {
	var l *Lexer = &Lexer{input: input}
	l.readChar()
	return l
}

// @Improvement: as of now we only support ASCII, to support unicode we should change ch to a rune type and change how we read input (each rune might be of variable length in unicode (utf-8))
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 is the value for the ASCII 'NUL', in our case means either we haven't read anything or EOF
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	fmt.Printf("char:%q", l.ch)

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LSQUIRLY, l.ch)
	case '}':
		tok = newToken(token.RSQUIRLY, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
