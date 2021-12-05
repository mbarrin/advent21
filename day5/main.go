package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var x1, x2, y1, y2 int

	grid := [1000][1000]int{}

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		var high, low int

		if x1 == x2 {
			if y1 > y2 {
				high, low = y1, y2
			} else {
				high, low = y2, y1
			}

			for i := low; i <= high; i++ {
				grid[i][x1] += 1
			}
		} else if y1 == y2 {
			if x1 > x2 {
				high, low = x1, x2
			} else {
				high, low = x2, x1
			}
			for i := low; i <= high; i++ {
				grid[y1][i] += 1
			}
		}
	}
	dangerous := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				dangerous += 1
			}
		}
	}

	fmt.Printf("part 1: %d\n", dangerous)
}
