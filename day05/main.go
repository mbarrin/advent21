package main

import (
	"bufio"
	"fmt"
	"os"
)

func makeRange(start, end int) []int {
	if start < end {
		a := make([]int, end-start+1)

		for i := range a {
			a[i] = start + i
		}
		return a
	} else {
		a := make([]int, start-end+1)

		for i := range a {
			a[i] = start - i
		}
		return a
	}
}

func main() {
	file, _ := os.Open("input.txt")
	const size = 1000

	scanner := bufio.NewScanner(file)

	var x1, x2, y1, y2 int

	gridOne := [size][size]int{}
	gridTwo := [size][size]int{}

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		var highY, lowY, highX, lowX int

		if y1 > y2 {
			highY, lowY = y1, y2
		} else {
			highY, lowY = y2, y1
		}

		if x1 > x2 {
			highX, lowX = x1, x2
		} else {
			highX, lowX = x2, x1
		}

		if x1 == x2 || y1 == y2 {
			for i := lowY; i <= highY; i++ {
				for j := lowX; j <= highX; j++ {
					gridOne[i][j] += 1
					gridTwo[i][j] += 1
				}
			}
		} else {
			xRange := makeRange(x1, x2)
			yRange := makeRange(y1, y2)

			for index := range yRange {
				gridTwo[yRange[index]][xRange[index]] += 1
			}
		}
	}

	dangerousOne, dangerousTwo := 0, 0

	for i := range gridOne {
		for j := range gridOne[i] {
			if gridOne[i][j] > 1 {
				dangerousOne += 1
			}
		}
	}

	for i := range gridTwo {
		for j := range gridTwo[i] {
			if gridTwo[i][j] > 1 {
				dangerousTwo += 1
			}
		}
	}

	fmt.Printf("part 1: %d\n", dangerousOne)
	fmt.Printf("part 2: %d\n", dangerousTwo)
}
