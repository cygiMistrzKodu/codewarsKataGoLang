package kata

import (
	"math"
)

func WallPaper(l, w, h float64) string {

	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven",
	"eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen",
	"sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	
	if (l * w * h == 0) {
		return numbers[0]
	}
	
	wallArea := 2*((w*h) + (l * h))
	areaOfOneRoll := 0.52 * 10

	more15Percent := 1.15
	rollRequire := int(math.Ceil((wallArea/areaOfOneRoll) * more15Percent))

	return numbers[rollRequire]
}
