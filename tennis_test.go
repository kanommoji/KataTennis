package tennis_test

import (
	"tennis"
	"testing"
)

func Test_SummaryTennisGame_Input_Score_Player_One_0_Player_Two_0_Should_Be_LOVE_LOVE(t *testing.T) {
	expected := "LOVE - LOVE"
	scorePlayerOne := 0
	scorePlayerTwo := 0

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_1_Player_Two_0_Should_Be_15_LOVE(t *testing.T) {
	expected := "15 - LOVE"
	scorePlayerOne := 1
	scorePlayerTwo := 0

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_0_Should_Be_30_LOVE(t *testing.T) {
	expected := "30 - LOVE"
	scorePlayerOne := 2
	scorePlayerTwo := 0

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_1_Should_Be_30_15(t *testing.T) {
	expected := "30 - 15"
	scorePlayerOne := 2
	scorePlayerTwo := 1

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_2_Should_Be_30_30(t *testing.T) {
	expected := "30 - 30"
	scorePlayerOne := 2
	scorePlayerTwo := 2

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_3_Should_Be_30_40(t *testing.T) {
	expected := "30 - 40"
	scorePlayerOne := 2
	scorePlayerTwo := 3

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_4_Should_Be_Player_2_Win(t *testing.T) {
	expected := "Player 2 Win"
	scorePlayerOne := 2
	scorePlayerTwo := 4

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_3_Player_Two_0_Should_Be_40_LOVE(t *testing.T) {
	expected := "40 - LOVE"
	scorePlayerOne := 3
	scorePlayerTwo := 0

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_4_Player_Two_0_Should_Be_Player_1_Win(t *testing.T) {
	expected := "Player 1 Win"
	scorePlayerOne := 4
	scorePlayerTwo := 0

	actual := tennis.SummaryTennisGame(scorePlayerOne, scorePlayerTwo)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
