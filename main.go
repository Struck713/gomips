package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Printf("# ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text() + "\n"
		for _, token := range Tokenize([]string{text}) {
			fmt.Printf("[%s] ", token.Print())
		}
		fmt.Println()
		fmt.Printf("# ")
	}

}
