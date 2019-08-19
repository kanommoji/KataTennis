package tennis

func SummaryTennisGame(scorePlayerOne int, scorePlayerTwo int) string {
	scores := [4]string{"LOVE","15","30","40"}
	if scorePlayerOne == 4 || scorePlayerTwo == 4{
		return checkWinner(scorePlayerOne,scorePlayerTwo)
	}
	return scores[scorePlayerOne]+" - "+scores[scorePlayerTwo]
}

func checkWinner(scorePlayerOne int, scorePlayerTwo int) string {
	if scorePlayerOne == 4 {
		return "Player 1 Win"
	}
	return "Player 2 Win"
	
}
