package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	grid   [][]int64
	marked [][]bool
}

type Win struct {
	board  Board
	number int64
}

var instructions []int64

func (b *Board) score() int64 {
	var score int64

	for j, col := range b.marked {
		for i := range col {
			if !b.marked[j][i] {
				score += b.grid[j][i]
			}
		}
	}
	return score
}

func (b *Board) contains(num int64) bool {
	for j, col := range b.grid {
		for i := range col {
			if b.grid[j][i] == num {
				b.marked[j][i] = true
				return true
			}
		}
	}
	return false
}

func (b *Board) boardWin() bool {
	return b.Win(true) || b.Win(false)
}

func (b *Board) Win(row bool) bool {
	for i := 0; i < len(b.marked); i++ {
		win := true
		for j := 0; j < len(b.marked); j++ {
			if row {
				win = win && b.marked[i][j]
			} else {
				win = win && b.marked[j][i]
			}
		}
		if win {
			return true
		}
	}
	return false
}

func newBoard() Board {
	return Board{marked: [][]bool{
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
	}}
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	input := []string{}
	boards := []Board{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for _, num := range strings.Split(input[0], ",") {
		instruction, _ := strconv.ParseInt(num, 10, 0)
		instructions = append(instructions, instruction)
	}

	for i := 2; i < len(input); i += 6 {
		board := newBoard()
		for j := i; j < i+5; j++ {
			line := []int64{}
			for _, num := range strings.Fields(input[j]) {
				number, _ := strconv.ParseInt(num, 10, 0)
				line = append(line, number)
			}
			board.grid = append(board.grid, line)
		}
		boards = append(boards, board)
	}

	wins := []Win{}

	for _, number := range instructions {
		for _, board := range boards {
			if board.boardWin() {
				continue
			}
			board.contains(number)
			if board.boardWin() {
				wins = append(wins, Win{board: board, number: number})
			}
		}
	}
	fmt.Printf("part 1: %d\n", wins[0].board.score()*wins[0].number)
	fmt.Printf("part 2: %d\n", wins[len(wins)-1].board.score()*wins[len(wins)-1].number)
}
