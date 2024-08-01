package kata

func EachCons(arr []int, n int) [][]int {

	subsets := [][]int{}

	for index := range arr {

		if index+n > len(arr) {
			break
		}

		groupSet := arr[index : index+n]
		subsets = append(subsets, groupSet)

	}

	return subsets
}
