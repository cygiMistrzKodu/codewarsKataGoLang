package kata

func Quadratic(x1, x2 int) [3]int {

	c := x1 * x2
	b := ( 1 * -x2) - (-x1 * -1)

	return [3]int{1, b, c}
}
