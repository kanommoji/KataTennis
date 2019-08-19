package tennis

func SummaryTennisGame(scorePlayerOne int, scorePlayerTwo int) string {
	score := [5]string{"LOVE","15","30","40","Win"}
	if scorePlayerOne == 4 {
		return "Player 1 Win"
	}
	if scorePlayerTwo == 4 {
		return "Player 2 Win"
	}
	return score[scorePlayerOne]+" - "+score[scorePlayerTwo]
}
