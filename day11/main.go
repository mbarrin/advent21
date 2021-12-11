package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func flash(grid [][]int64, flashed [][]bool, y, x, count int) ([][]int64, [][]bool, int) {
	if flashed[y][x] {
		return grid, flashed, count
	}

	grid[y][x]++

	if grid[y][x] <= 9 {
		return grid, flashed, count
	}

	flashed[y][x] = true
	count += 1

	//fmt.Printf("y: %d, x: %d\n", y, x)
	//for y := range grid {
	//	for x := range grid[y] {
	//		fmt.Printf("%02d ", grid[y][x])
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()

	if y > 0 && x > 0 {
		grid, flashed, count = flash(grid, flashed, y-1, x-1, count)
	}

	if y > 0 && x >= 0 {
		grid, flashed, count = flash(grid, flashed, y-1, x, count)
	}

	if y < len(grid)-1 && x < len(grid[y])-1 {
		grid, flashed, count = flash(grid, flashed, y+1, x+1, count)
	}

	if y < len(grid)-1 && x <= len(grid[y])-1 {
		grid, flashed, count = flash(grid, flashed, y+1, x, count)
	}

	if y > 0 && x < len(grid[y])-1 {
		grid, flashed, count = flash(grid, flashed, y-1, x+1, count)
	}

	if y >= 0 && x < len(grid[y])-1 {
		grid, flashed, count = flash(grid, flashed, y, x+1, count)
	}

	if y < len(grid)-1 && x > 0 {
		grid, flashed, count = flash(grid, flashed, y+1, x-1, count)
	}

	if y <= len(grid)-1 && x > 0 {
		grid, flashed, count = flash(grid, flashed, y, x-1, count)
	}

	grid[y][x] = 0

	return grid, flashed, count
}

func main() {
	grid := [][]int64{}
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []int64{}
		split := strings.Split(scanner.Text(), "")

		for _, char := range split {
			num, _ := strconv.ParseInt(char, 10, 0)
			line = append(line, num)
		}
		grid = append(grid, line)
	}

	grandTotal := 0
	for i := 0; i < 100; i++ {
		loopTotal := 0
		flashed := make([][]bool, 10)
		for i := range flashed {
			flashed[i] = make([]bool, 10)
		}

		for y := range grid {
			for x := range grid[y] {
				grid, flashed, loopTotal = flash(grid, flashed, y, x, loopTotal)
			}
		}

		grandTotal += loopTotal

	}
	fmt.Println(grandTotal)
}
