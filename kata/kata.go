package kata

func Incrementer(n []int) []int {

	lastDigitsAfterInrements := []int{}
	for index, number := range n {
		lastDigitsAfterInrements = append(lastDigitsAfterInrements, (number+index+1)%10)
	}

	return lastDigitsAfterInrements
}
