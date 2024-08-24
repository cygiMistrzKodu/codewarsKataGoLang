package kata

import "strconv"

func WhatCentury(year string) string {

	endingsPostfix := map[int]string{
		1: "st",
		2: "nd",
		3: "rd",
		4: "th",
		5: "th",
		6: "th",
		7: "th",
		8: "th",
		9: "th",
		0: "th",
	}

	yearNumber, _ := strconv.ParseInt(year, 10, 64)

	if yearNumber%1000 == 0 || yearNumber%100 == 0 {
		centaury := (yearNumber / 100)

		if centaury >= 11 && centaury <= 13 {
			return strconv.Itoa(int(centaury)) + "th"
		}
		return strconv.Itoa(int(centaury)) + endingsPostfix[int(centaury%10)]

	}

	centaury := (yearNumber / 100) + 1

	if centaury >= 11 && centaury <= 13 {
		return strconv.Itoa(int(centaury)) + "th"
	}

	return strconv.Itoa(int(centaury)) + endingsPostfix[int(centaury%10)]
}
