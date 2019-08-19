package tennis_test

import (
	"tennis"
	"testing"
)

func Test_SummaryTennisGame_Input_Score_Player_One_0_Player_Two_0_Should_Be_LOVE_LOVE(t *testing.T) {
	expected := "LOVE - LOVE"

	actual := tennis.SummaryTennisGame(0, 0)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_1_Player_Two_0_Should_Be_15_LOVE(t *testing.T) {
	expected := "15 - LOVE"

	actual := tennis.SummaryTennisGame(tennis.AddScore(1, 0))

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_0_Should_Be_30_LOVE(t *testing.T) {
	expected := "30 - LOVE"
	tennis.AddScore(1, 0)

	actual := tennis.SummaryTennisGame(tennis.AddScore(1, 0))

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_1_Should_Be_30_15(t *testing.T) {
	expected := "30 - 15"
	tennis.AddScore(1, 0)
	tennis.AddScore(1, 0)

	actual := tennis.SummaryTennisGame(tennis.AddScore(0, 1))

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_2_Should_Be_30_30(t *testing.T) {
	expected := "30 - 30"
	tennis.AddScore(1, 0)
	tennis.AddScore(1, 0)
	tennis.AddScore(0, 1)

	actual := tennis.SummaryTennisGame(tennis.AddScore(0, 1))

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_3_Should_Be_30_40(t *testing.T) {
	expected := "30 - 40"
	tennis.AddScore(1, 0)
	tennis.AddScore(1, 0)
	tennis.AddScore(0, 1)
	tennis.AddScore(0, 1)

	actual := tennis.SummaryTennisGame(tennis.AddScore(0, 1))

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_2_Player_Two_4_Should_Be_Player_2_Win(t *testing.T) {
	expected := "Player 2 Win"
	tennis.AddScore(1, 0)
	tennis.AddScore(1, 0)
	tennis.AddScore(0, 1)
	tennis.AddScore(0, 1)
	tennis.AddScore(0, 1)

	actual := tennis.SummaryTennisGame(tennis.AddScore(0, 1))

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_3_Player_Two_0_Should_Be_40_LOVE(t *testing.T) {
	expected := "40 - LOVE"
	tennis.AddScore(1, 0)
	tennis.AddScore(1, 0)

	actual := tennis.SummaryTennisGame(tennis.AddScore(1, 0))

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_SummaryTennisGame_Input_Score_Player_One_4_Player_Two_0_Should_Be_Player_1_Win(t *testing.T) {
	expected := "Player 1 Win"
	tennis.AddScore(1, 0)
	tennis.AddScore(1, 0)
	tennis.AddScore(1, 0)

	actual := tennis.SummaryTennisGame(tennis.AddScore(1, 0))

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}
