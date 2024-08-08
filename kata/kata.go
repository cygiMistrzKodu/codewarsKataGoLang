package kata

import (
	"strings"
	"unicode"
)

var NATO = map[string]string{
	"A": "Alfa", "B": "Bravo", "C": "Charlie", "D": "Delta", "E": "Echo",
	"F": "Foxtrot", "G": "Golf", "H": "Hotel", "I": "India", "J": "Juliett",
	"K": "Kilo", "L": "Lima", "M": "Mike", "N": "November", "O": "Oscar",
	"P": "Papa", "Q": "Quebec", "R": "Romeo", "S": "Sierra", "T": "Tango",
	"U": "Uniform", "V": "Victor", "W": "Whiskey", "X": "X-ray", "Y": "Yankee",
	"Z": "Zulu",
}

func ToNato(words string) string {

	natoLetters := []string{}
	for _, letter := range strings.Split(words, "") {

		sign := strings.ToUpper(strings.TrimSpace(letter))

		if len(sign) > 0 {
			if unicode.IsLetter(rune(sign[0])) {
				natoLetters = append(natoLetters, NATO[sign])
			} else {
				natoLetters = append(natoLetters, sign)
			}

		}

	}

	return strings.Join(natoLetters, " ")
}
