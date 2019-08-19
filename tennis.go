package tennis

func SummaryTennisGame(scorePlayerOne int, scorePlayerTwo int) string {
	score := [5]string{"LOVE","15"}
	playerOne := score[scorePlayerOne]
	playerTwo := score[scorePlayerTwo]
	return playerOne+" - "+playerTwo
}
