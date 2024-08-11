package kata



func NextHigher(n int) int {
	firstOneMask := n & (-n)
	result := firstOneMask + n
	ones := result ^ n
	ones = (ones >> 2) / firstOneMask   
	return (result | ones )
}
