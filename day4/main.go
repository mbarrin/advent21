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

func removeIndex(slice []Board, index int) []Board {
	return append(slice[:index], slice[index+1:]...)
}

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

func (b *Board) won() bool {
	return b.colWin() || b.rowWin()
}

func (b *Board) rowWin() bool {
	for _, col := range b.marked {
		win := true
		for _, marked := range col {
			win = win && marked
		}
		if win {
			return true
		}
	}
	return false
}

func (b *Board) colWin() bool {
	for i := 0; i < len(b.marked); i++ {
		win := true
		for j := 0; j < len(b.marked); j++ {
			win = win && b.marked[j][i]
		}
		if win {
			return true
		}
	}
	return false
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
		board := Board{marked: [][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}}

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
			if board.won() {
				continue
			}
			board.contains(number)
			if board.won() {
				wins = append(wins, Win{board: board, number: number})
			}
		}
	}
	fmt.Printf("part 1: %d\n", wins[0].board.score()*wins[0].number)
	fmt.Printf("part 2: %d\n", wins[len(wins)-1].board.score()*wins[len(wins)-1].number)
}
