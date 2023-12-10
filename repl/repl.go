package repl

import (
	"bufio"
	"fmt"
	"github.com/Lamzari/MonkeyLangInterpreter/lexer"
	tok "github.com/Lamzari/MonkeyLangInterpreter/token"
	"io"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		for token := lexer.NextToken(); token.Type != tok.EOF; token = lexer.NextToken() {
			fmt.Printf("%+v\n", token)
		}
	}
}
