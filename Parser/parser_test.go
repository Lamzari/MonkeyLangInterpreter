package parser

import (
	"testing"

	"github.com/Lamzari/MonkeyLangInterpreter/ast"
	"github.com/Lamzari/MonkeyLangInterpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	var input string = `
  let x = 10;
  let y = 69;
  let gincobiloba = 420;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesn't contain 3 statements, it contains %d statements", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"gincobiloba"},
	}

	for idx, currTest := range tests {
		statement := program.Statements[idx]
		if !testLetStatement(t, statement, currTest.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let' got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}
