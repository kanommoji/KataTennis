package tennis

func SummaryTennisGame(scorePlayerOne int, scorePlayerTwo int) string {
	score := [5]string{"LOVE","15","30","40","Win"}
	playerOne := score[scorePlayerOne]
	playerTwo := score[scorePlayerTwo]
	if playerOne == score[4] {
		return "Player 1 Win"
	}
	if playerTwo == score[4] {
		return "Player 2 Win"
	}
	return playerOne+" - "+playerTwo
}
