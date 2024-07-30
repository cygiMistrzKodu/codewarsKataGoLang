package kata

func combat(health, damage float64) float64 {

	newHealt := health - damage

	if newHealt < 0 {
		return 0
	} else {
		return newHealt
	}

}
