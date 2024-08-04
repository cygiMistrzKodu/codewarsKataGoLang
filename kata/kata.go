package kata

import (
	"time"
)

func UnluckyDays(year int) int {

	fridays13InYear := 0
	for month := time.January; month <= time.December; month++ {
		thirteenthDay := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
		
		if thirteenthDay.Weekday() == time.Friday {
			fridays13InYear++
		}

	}

	return fridays13InYear
}
