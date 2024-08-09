package kata

import (
	"regexp"
)

func Solution(str string) []string {

	regexPattern := regexp.MustCompile(`..|.`)
	result := regexPattern.FindAllString(str, -1)

	for index, text := range result {
		if len(text) < 2 {
			result[index] = text + "_"
		}

	}
	return result
}
