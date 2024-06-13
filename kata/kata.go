package kata

func Xor(a, b bool) bool {

	if a && !b || !a && b {
		return true
	}

	return false
}
