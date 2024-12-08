package main

import (
	"testing"
)

func TestTokenizeLine(t *testing.T) {

	lines := []string{
		"beq $s2, 3, terminate\n",
	}

	tokens := Tokenize(lines)
	if len(tokens) != 7 {
		t.Fatalf(`len(tokens) = %d, want 7`, len(tokens))
	}
	isCorrectToken(t, tokens[0], "beq")
	isCorrectToken(t, tokens[1], "$s2")
	isCorrectToken(t, tokens[2], ",")
	isCorrectToken(t, tokens[3], "3")
	isCorrectToken(t, tokens[4], ",")
	isCorrectToken(t, tokens[5], "terminate")
	isCorrectToken(t, tokens[6], "\n")

}

func TestTokenizeLines(t *testing.T) {

	lines := []string{
		"beq $s2, 3, terminate\n",
		"la $a0, enterF # set argument a0 to enterF\n",
	}

	tokens := Tokenize(lines)
	if len(tokens) != 12 {
		t.Fatalf(`len(tokens) = %d, want 12`, len(tokens))
	}

	isCorrectToken(t, tokens[0], "beq")
	isCorrectToken(t, tokens[1], "$s2")
	isCorrectToken(t, tokens[2], ",")
	isCorrectToken(t, tokens[3], "3")
	isCorrectToken(t, tokens[4], ",")
	isCorrectToken(t, tokens[5], "terminate")
	isCorrectToken(t, tokens[6], "\n")
	isCorrectToken(t, tokens[7], "la")
	isCorrectToken(t, tokens[8], "$a0")
	isCorrectToken(t, tokens[9], ",")
	isCorrectToken(t, tokens[10], "enterF")
	isCorrectToken(t, tokens[11], "\n")

}

func TestTokenizeString(t *testing.T) {

	lines := []string{
		"sw $0, 0($6)\n",
		"label: .asciiz \"This is a test string\"\n",
	}

	tokens := Tokenize(lines)
	if len(tokens) != 13 {
		t.Fatalf(`len(tokens) = %d, want 14`, len(tokens))
	}

	isCorrectToken(t, tokens[0], "sw")
	isCorrectToken(t, tokens[1], "$0")
	isCorrectToken(t, tokens[2], ",")
	isCorrectToken(t, tokens[3], "0")
	isCorrectToken(t, tokens[4], "(")
	isCorrectToken(t, tokens[5], "$6")
	isCorrectToken(t, tokens[6], ")")
	isCorrectToken(t, tokens[7], "\n")
	isCorrectToken(t, tokens[8], "label")
	isCorrectToken(t, tokens[9], ":")
	isCorrectToken(t, tokens[10], ".asciiz")
	isCorrectToken(t, tokens[11], "This is a test string")
	isCorrectToken(t, tokens[12], "\n")

}

func isCorrectToken(t *testing.T, token Token, text string) {
	if token.Token != text {
		t.Fatalf(`token.Text = %s, want "%s"`, token.Token, text)
	}
}
