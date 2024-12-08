package main

import (
	"fmt"
)

type Tokens rune
type Token struct {
	Token  string
	Line   int
	Column int
}

func (t Token) Print() string {
	token := t.Token
	if token == "\n" {
		token = "\\n"
	}
	return fmt.Sprintf("%d:%d %s", t.Line, t.Column, token)
}

func NewToken(token string, line int, column int) Token {
	return Token{
		Token:  token,
		Line:   line,
		Column: column,
	}
}

func Tokenize(lines []string) []Token {

	tokens := []Token{}

	for y, line := range lines {
		for start, x := 0, 0; x < len(line); x++ {
			switch line[x] {
			case '#':
				tokens = append(tokens, NewToken(string('\n'), y, x))
				x = len(line)
				break
			case ',', ':', '\n':
				if start < x {
					tokens = append(tokens, NewToken(line[start:x], y, start))
				}
				tokens = append(tokens, NewToken(string(line[x]), y, x))
				start = x + 1
				break
			case ' ', '	':
				if start < x {
					tokens = append(tokens, NewToken(line[start:x], y, start))
				}
				start = x + 1
				break
			}
		}
	}

	return tokens
}
