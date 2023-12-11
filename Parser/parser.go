package parser

import (
	"github.com/Lamzari/MonkeyLangInterpreter/ast"
	"github.com/Lamzari/MonkeyLangInterpreter/lexer"
	"github.com/Lamzari/MonkeyLangInterpreter/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.currentToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	//@@Doing: starting to write the code for the recursive descent parser
	//         pseudo code in ./simple_recursive_descent_parser_pseudocode.txt
	return nil
}
