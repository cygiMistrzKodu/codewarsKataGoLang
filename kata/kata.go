package kata

type NumberIndex struct {
	number int
	index  int
}

func (numberIndex NumberIndex) GetIndex() int {
	return numberIndex.index
}

func (numberIndex NumberIndex) GetValue() int {
	return numberIndex.number
}

func QueueTime(customers []int, n int) int {

	if len(customers) <= 0 {
		return 0
	}

	queues := []int{}

	for position, customerTime := range customers {

		if position < n {
			queues = append(queues, customerTime)
		} else {

			lowestSumElement := LowestSum(queues)
			lowestSumTime := lowestSumElement.GetValue()
			indexOfMinSum := lowestSumElement.GetIndex()

			queues[indexOfMinSum] = lowestSumTime + customerTime

		}

	}

	return Max(queues)
}

func LowestSum(queues []int) NumberIndex {
	lowestSum := queues[0]
	position := 0
	for index, sum := range queues {
		if lowestSum > sum {
			lowestSum = sum
			position = index
		}
	}

	return NumberIndex{index: position, number: lowestSum}

}

func Max(queues []int) int {
	max := queues[0]

	for _, time := range queues {
		if time > max {
			max = time
		}
	}

	return max
}
