package assembler

import (
	"fmt"
	"strings"
)

type ASTType int

const (
	Root ASTType = iota
	Label
	Instruction
	Operand
)

type AST struct {
	Type     ASTType
	Tokens   []Token
	Children []AST
}

func (a AST) Print(level int) {

	if level == 0 {
		fmt.Println("root:")
	}

	for _, token := range a.Tokens {
		fmt.Printf(strings.Repeat("  ", level))
		token.Print()
	}

	for _, child := range a.Children {
		child.Print(level + 1)
	}
}
