package tennis

var scoreOne, scoreTwo []int

func SummaryTennisGame(scorePlayerOne int, scorePlayerTwo int) string {
	newCourt()
	winner := checkWinner(scorePlayerOne, scorePlayerTwo)
	if winner != "" {
		return winner
	}
	scores := [4]string{"LOVE", "15", "30", "40"}
	return scores[scorePlayerOne] + " - " + scores[scorePlayerTwo]
}

func newCourt() {
	scoreOne = []int{0}
	scoreTwo = []int{0}
}

func checkWinner(scorePlayerOne int, scorePlayerTwo int) string {
	if scorePlayerOne == 4 {
		return "Player 1 Win"
	}
	if scorePlayerTwo == 4 {
		return "Player 2 Win"
	}
	return ""
}

func AddScore(scorePlayerOne int, scorePlayerTwo int) (int, int) {
	scoreOne = append(scoreOne, scorePlayerOne)
	scoreTwo = append(scoreTwo, scorePlayerTwo)
	var one, two int
	for _, score := range scoreOne {
		one += score
	}
	for _, score := range scoreTwo {
		two += score
	}
	return one, two
}
