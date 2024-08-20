package kata

import "slices"

func QueueTime(customers []int, n int) int {

	if len(customers) <= 0 {
		return 0
	}

	queues := []int{}

	for position, customerTime := range customers {

		if position < n {
			queues = append(queues, customerTime)
		} else {

			lowestSumTime := slices.Min(queues)
			lowestSumIndex := slices.Index(queues, lowestSumTime)
			queues[lowestSumIndex] = lowestSumTime + customerTime

		}

	}

	return slices.Max(queues)
}
