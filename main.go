package main

import (
	"bufio"
	"fmt"
	"os"

	"nstruck.dev/gomips/assembler"
)

func main() {

	fmt.Printf("# ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text() + "\n"

		tokens := assembler.Tokenize([]string{text})
		ast := assembler.Parse(tokens)
		fmt.Println(ast)

		fmt.Printf("# ")
	}

}
