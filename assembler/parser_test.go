package assembler

import (
	"testing"
)

func TestParser1(t *testing.T) {

	lines := []string{
		"main:\n",
		"	beq $s2, 3, terminate\n",
	}

	tokens := Tokenize(lines)
	ast := Parse(tokens)
	ast.Print(0)
}

func TestParser2(t *testing.T) {

	lines := []string{
		"main:\n",
		"	beq $s2, 3, terminate\n",
		"	la $a0, enterF # set argument a0 to enterF\n",
		"	jal readInput # run subroutine that displays enterF text and gets input\n",
		"	move $s1, $s0\n",
		"	la $a0, enterG # set argument a0 to enterG\n",
		"	jal readInput # run subroutine that displays enterG text and gets input\n",
		"	move $t0, $s0\n",
		"	jal printAnswer # print the answer of f and g\n",
		"	addi $s2, $s2, 1\n",
		"	j main\n",
	}

	tokens := Tokenize(lines)
	ast := Parse(tokens)
	ast.Print(0)
}
