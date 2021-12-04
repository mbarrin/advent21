package main

import "testing"

func TestRowWin(t *testing.T) {
	input := [][]bool{
		{false, false, false},
		{false, true, true},
		{true, true, true},
	}

	badInput := [][]bool{
		{false, false, false},
		{false, true, true},
		{false, true, true},
	}

	expectedOutput := true

	board := Board{marked: input}
	badBoard := Board{marked: badInput}

	if expectedOutput != board.Win(true) {
		t.Fail()
		t.Log("FAIL")
	}

	if badBoard.Win(true) != false {
		t.Fail()
		t.Log("FAIL")
	}
}

func TestColWin(t *testing.T) {
	input := [][]bool{
		{true, false, true},
		{true, false, false},
		{true, false, true},
	}

	badInput := [][]bool{
		{false, false, true},
		{true, false, false},
		{true, false, true},
	}

	board := Board{marked: input}
	badBoard := Board{marked: badInput}

	expectedOutput := true

	if expectedOutput != board.Win(false) {
		t.Fail()
		t.Log("FAIL")
	}

	if badBoard.Win(false) != false {
		t.Fail()
		t.Log("FAIL")
	}
}
