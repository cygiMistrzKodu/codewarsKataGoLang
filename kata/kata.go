package kata

import (
	"strings"
)

func Contamination(text, char string) string {

	letters := strings.Split(text, "")

	for i := 0; i < len(letters); i++ {
		letters[i] = char
	}

	return strings.Join(letters, "")

}
