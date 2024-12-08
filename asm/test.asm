.data
enterF: .asciiz "Enter a value for f: "
enterG: .asciiz "Enter a value for g: "
answer: .asciiz "Answer for f = g - (f + 5): "

.text
main:
	beq $s2, 3, terminate
	la $a0, enterF # set argument a0 to enterF 
	jal readInput # run subroutine that displays enterF text and gets input
	move $s1, $s0
	
	la $a0, enterG # set argument a0 to enterG 
	jal readInput # run subroutine that displays enterG text and gets input
	move $t0, $s0
	
	jal printAnswer # print the answer of f and g
	
	addi $s2, $s2, 1
	j main
	
terminate:
	li $v0, 10 # terminate program
    syscall   
	
readInput:
	li $v0, 4 # print $a0 to the screen
	syscall
	
	li $v0, 5 # read an integer from the screen
	syscall
	move $s0, $v0
	
	jr $ra
	
printAnswer:
	li $v0, 4 # print the answer text to the screen
	la $a0, answer
	syscall
	
	add $t2, $s1, 5 # compute (f + 5)
	sub $s1, $t0, $t2 # compute g - (f + 5)
	
	li $v0, 1 # print result of $t3
	move $a0, $s1
	syscall
	
	li $v0, 11 # print new line character
	la $a0, '\n'
	syscall
	
	jr $ra	
	