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

	if expectedOutput != board.rowWin() {
		t.Fail()
		t.Log("FAIL")
	}

	if badBoard.rowWin() != false {
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

	if expectedOutput != board.colWin() {
		t.Fail()
		t.Log("FAIL")
	}

	if badBoard.colWin() != false {
		t.Fail()
		t.Log("FAIL")
	}
}
