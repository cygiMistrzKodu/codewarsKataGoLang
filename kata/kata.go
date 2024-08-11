package kata



func NextHigher(n int) int {
	r := n & -n
  p := n + r
  q := (n ^ p) / (4 * r)
  return p | q
}
