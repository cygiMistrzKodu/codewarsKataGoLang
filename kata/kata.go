package kata

import (
	"strconv"
)

func MaxRot(n int64) int64 {
   max := n
	strn := strconv.FormatInt(n, 10)

	var left string
	var middle string
	var right string
	var currentNumber int64

	for index := range strn[:len(strn)-1] {
		left = strn[:index]
		middle = strn[index:index+1]
		right = strn[index+1:]

		strn = left+right+middle

		currentNumber,_ = strconv.ParseInt(strn,10,64)
		if max < currentNumber {max = currentNumber}
		}
	return max
}