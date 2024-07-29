package kata

import "strconv"

func BinToDec(bin string) int {

	deciaml, err := strconv.ParseInt(bin, 2, 64)

	if err != nil {
		return -1
	}

	return int(deciaml)
}
