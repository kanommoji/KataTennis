package tennis_test

import "testing"

func Test_SummaryTennisGame_Input_Score_Player_One_0_Player_Two_0_Should_Be_LOVE_LOVE(t *testing.T) {
	expected := "LOVE - LOVE"
	scorePlayerOne := 0
	scorePlayerTwo := 0

	actual := SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
