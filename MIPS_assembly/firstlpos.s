.data
shouldben1:	.asciiz "Should be -1, and firstposshift and firstposmask returned: "
shouldbe0:	.asciiz "Should be 0 , and firstposshift and firstposmask returned: "
shouldbe16:	.asciiz "Should be 16, and firstposshift and firstposmask returned: "
shouldbe31:	.asciiz "Should be 31, and firstposshift and firstposmask returned: "

.text
main:
	la	$a0, shouldbe31
	jal	print_str
	lui	$a0, 0x8000	# should be 31
	jal	first1posshift
	move	$a0, $v0
	jal	print_int
	jal	print_space

	lui	$a0, 0x8000
	jal	first1posmask
	move	$a0, $v0
	jal	print_int
	jal	print_newline

	la	$a0, shouldbe16
	jal	print_str
	lui	$a0, 0x0001	# should be 16
	jal	first1posshift
	move	$a0, $v0
	jal	print_int
	jal	print_space

	lui	$a0, 0x0001
	jal	first1posmask
	move	$a0, $v0
	jal	print_int
	jal	print_newline

	la	$a0, shouldbe0
	jal	print_str
	li	$a0, 1		# should be 0
	jal	first1posshift
	move	$a0, $v0
	jal	print_int
	jal	print_space

	li	$a0, 1
	jal	first1posmask
	move	$a0, $v0
	jal	print_int
	jal	print_newline

	la	$a0, shouldben1
	jal	print_str
	move	$a0, $0		# should be -1
	jal	first1posshift
	move	$a0, $v0
	jal	print_int
	jal	print_space

	move	$a0, $0
	jal	first1posmask
	move	$a0, $v0
	jal	print_int
	jal	print_newline

	li	$v0, 10
	syscall

first1posshift:
	### YOUR CODE HERE ###
	addu 	$sp, $sp, -4
	sw	$ra, 0($sp)
	move	$s0, $zero
	li	$s1, -1
lop:	addi 	$s1, $s1, 1
	beq	$s1, 32, return1
	sllv	$s0, $a0, $s1
	bltu	$s0, 0x80000000, lop
	li	$v0, 31
	sub	$v0, $v0, $s1 
	
return:	lw	$ra, 0($sp)
	addu	$sp, $sp, 4
	jr 	$ra
return1:
	li	$v0, -1
	lw	$ra, 0($sp)
	addu	$sp, $sp, 4
	jr 	$ra
first1posmask:
	### YOUR CODE HERE ###
	addu 	$sp, $sp, -4
	sw	$ra, 0($sp)
	move	$s0, $zero
	li	$s1, 32
	
lp:	addu	$s1, $s1, -1
	andi	$s0, $a0, 0x80000000
	beq 	$s1, -1, return2
	bne	$s0, $zero, return3
	sll	$a0, $a0, 1
	j	lp
		
return2:		#-1
	li	$v0, -1	
	lw	$ra, 0($sp)
	addu	$sp, $sp, 4
	jr 	$ra
	
return3:
	addu	$v0, $s1, 0	
	lw	$ra, 0($sp)
	addu	$sp, $sp, 4
	jr 	$ra

print_int:
	move	$a0, $v0
	li	$v0, 1
	syscall
	jr	$ra

print_str:
	li	$v0, 4
	syscall
	jr	$ra

print_space:
	li	$a0, ' '
	li	$v0, 11
	syscall
	jr	$ra

print_newline:
	li	$a0, '\n'
	li	$v0, 11
	syscall
	jr	$ra
