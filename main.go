package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Lamzari/MonkeyLangInterpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey %s! This is the language REPL\n", user.Username)
	fmt.Printf("start typing commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
