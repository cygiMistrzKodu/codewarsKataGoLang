package kata

func Amidakuji(ar []string) []int {

	players := make([]int, len(ar[0])+1)
	for playerNumber := range players {
		players[playerNumber] = playerNumber
		playerNumber++
	}

	playersFinalPostions := make([]int, len(ar[0])+1)

	var leftSide rune
	var rightSide rune

	for playerPostion, playerNumber := range players {

		playerFinalPostion := playerPostion

		for _, ladderFloor := range ar {

			if playerFinalPostion > 0 {

				leftSide = rune(ladderFloor[playerFinalPostion-1])
			} else {
				leftSide = '0'
			}

			if playerFinalPostion < len(ladderFloor) {

				rightSide = rune(ladderFloor[playerFinalPostion])
			} else {
				rightSide = '0'
			}

			if leftSide != 0 {

				if leftSide == '1' {
					playerFinalPostion--
				}

			}

			if rightSide != 0 {

				if rightSide == '1' {
					playerFinalPostion++

				}
			}

		}

		playersFinalPostions[playerFinalPostion] = playerNumber
	}

	return playersFinalPostions
}
