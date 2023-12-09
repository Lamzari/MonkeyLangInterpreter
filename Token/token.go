package token

// @Performance: we might want to use int, but as a start a string is more versatile for debugging (I can just print it)
type TokenType string

// @Improvement: in a prod language I would attach filenames, line and column numbers to the token type in order to have nice error messages
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// literals and identifiers
	IDENT = "IDENT" // x, y, functionName...
	INT   = "INT"   // 123456

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LSQUIRLY = "{"
	RSQUIRLY = "}"

	// keywords
	FUNCTION = "FUNTCTION"
	LET      = "LET"
)