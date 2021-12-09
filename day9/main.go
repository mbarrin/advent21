package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const SIZE = 100

func main() {
	grid := [][]int64{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []int64{}
		split := strings.Split(scanner.Text(), "")
		for _, x := range split {
			num, _ := strconv.ParseInt(x, 10, 0)
			line = append(line, num)
		}
		grid = append(grid, line)
	}

	var total int64
	basins := []int{}
	for y, line := range grid {
		for x := range line {
			if isLowest(grid, y, x) {
				total += grid[y][x] + 1
				visited := [SIZE][SIZE]bool{}
				basins = append(basins, count(findBaisin(grid, y, x, visited)))
			}
		}
	}

	sort.Ints(basins)
	basinsTotal := 1

	for i := len(basins) - 3; i < len(basins); i++ {
		basinsTotal *= basins[i]
	}

	fmt.Printf("part 1: %d\n", total)
	fmt.Printf("part 2: %v\n", basinsTotal)
}

func count(visited [SIZE][SIZE]bool) int {
	total := 0
	for y, line := range visited {
		for x := range line {
			if visited[y][x] {
				total++
			}
		}
	}
	return total
}

func isLowest(grid [][]int64, y, x int) bool {
	center := grid[y][x]
	var up, down, left, right int64
	if y == 0 {
		up = 99
		down = grid[y+1][x]
	} else if y == len(grid)-1 {
		up = grid[y-1][x]
		down = 99
	} else {
		up = grid[y-1][x]
		down = grid[y+1][x]
	}

	if x == 0 {
		left = 99
		right = grid[y][x+1]
	} else if x == len(grid[0])-1 {
		left = grid[y][x-1]
		right = 99
	} else {
		left = grid[y][x-1]
		right = grid[y][x+1]
	}

	if up > center && down > center && left > center && right > center {
		return true
	}
	return false
}

func findBaisin(grid [][]int64, y, x int, visited [SIZE][SIZE]bool) [SIZE][SIZE]bool {
	if grid[y][x] == 9 {
		return visited
	}

	if visited[y][x] {
		return visited
	}

	visited[y][x] = true

	if y != 0 {
		visited = findBaisin(grid, y-1, x, visited)
	}
	if x != 0 {
		visited = findBaisin(grid, y, x-1, visited)
	}
	if y != len(grid)-1 {
		visited = findBaisin(grid, y+1, x, visited)
	}
	if x != len(grid[0])-1 {
		visited = findBaisin(grid, y, x+1, visited)
	}

	return visited
}
