package kata

func Between(a, b int) []int {

	var numbers []int
	for number := a; number <= b; number++ {
		numbers = append(numbers, number)
	}

	return numbers
}
