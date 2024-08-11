package kata



func NextHigher(x int) int {
	// right most set bit
	rightOne := x & -x

	// reset the pattern and set next higher bit
	// left part of x will be here
	nextHigherOneBit := x + rightOne

	// nextHigherOneBit is now part [D] of the above explanation.

	// isolate the pattern
	rightOnesPattern := x ^ nextHigherOneBit

	// right adjust pattern
	rightOnesPattern = (rightOnesPattern) / rightOne

	// correction factor
	rightOnesPattern >>= 2

	// rightOnesPattern is now part [A] of the above explanation.

	// integrate new pattern (Add [D] and [A])
	next := nextHigherOneBit | rightOnesPattern
	return next
}
