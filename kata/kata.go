package kata

func Amidakuji(ar []string) []int {

	nums := make([]int, len(ar[0])+1)
	for i := range nums {
		nums[i] = i
	}
	for _, row := range ar {
		for i, s := range row {
			if s == '1' {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	return nums
}
