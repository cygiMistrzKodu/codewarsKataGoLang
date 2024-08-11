package kata

import (
	"strconv"
	"strings"
)

func reset(s string) string {
	m := 0
	a := []rune(s)
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '0' {
			a[i] = '1'
			m += 1
		}
	}
	return strings.Repeat("0", m) + string(a)[m:]
}

func NextHigher(n int) int {
	s := strconv.FormatInt(int64(n), 2)
	i := strings.LastIndex(s, "01")
	if i != -1 {
		s = s[0:i] + "10" + reset(s[i+2:])
	} else {
		s = "10" + reset(s[1:])
	}
	r, _ := strconv.ParseInt(s, 2, 32)
	return int(r)
}
