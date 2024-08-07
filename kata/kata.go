package kata

func Dominator(a []int) int {

	groupByNumbers := make(map[int]int)

	for _, number := range a {
		groupByNumbers[number]++
	}

	gratestOccurrenceCount := 0
	gratesOccurrenceNumber := -1
	for number, times := range groupByNumbers {

		if gratestOccurrenceCount < times {
			gratestOccurrenceCount = times
			gratesOccurrenceNumber = number
		}

	}

	if gratestOccurrenceCount > len(a)/2 {
		return gratesOccurrenceNumber
	} else {
		return -1
	}

}
