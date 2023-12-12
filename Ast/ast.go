/*  The basic structure of the ast is the following:
*   The difference between a statement and an expression is that a statement doesn't return any value, while an expression does
*   statement: let x = 5;
*   expression: someMethod();
    ┌───────────────────────────────┐
    │                               │
    │          *ast.Program         │
    │                               │
    ├───────────────────────────────┤
    │                               │
    │          []Statements         │
    │                               │
    └───────────────┬───────────────┘
                    │
    ┌───────────────┴───────────────┐
    │        *ast.LetStatement      │
    ├───────────────────────────────┤
┌───┤              Name             │
│   ├───────────────────────────────┤
│   │              Value            ├────┐
│   └───────────────────────────────┘    │
│                                        │
│   ┌────────────────────┐               │
│   │                    │               │
│   │   *ast.Identifier  │               │
└─► │                    │               │
    └────────────────────┘               │
                                         │
              ┌──────────────────────┐   │
              │                      │   │
              │   *ast.Expression    │ ◄─┘
              │                      │
              └──────────────────────┘
*/

package ast

import "github.com/Lamzari/MonkeyLangInterpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//@@Consider: consider grouping statements each in its own file and group them in a folder
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
