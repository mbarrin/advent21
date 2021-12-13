package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	instructions := []string{}
	dots := []string{}
	var foldInfo string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold along") {
			instructions = append(instructions, line)
		} else {
			dots = append(dots, line)
		}
	}
	var y, x int64

	for i := range instructions {
		newDots := []string{}
		dotMap := map[string]int{}

		fmt.Sscanf(instructions[i], "fold along %s", &foldInfo)
		foldInfo := strings.Split(foldInfo, "=")

		foldAxis := foldInfo[0]
		foldLine, _ := strconv.ParseInt(foldInfo[1], 10, 0)

		for i := range dots {
			fmt.Sscanf(dots[i], "%d,%d", &x, &y)
			if foldAxis == "y" {
				if y > foldLine {
					y = y - ((y - foldLine) * 2)
				}
			} else {
				if x > foldLine {
					x = x - ((x - foldLine) * 2)
				}
			}
			coords := fmt.Sprintf("%d,%d", x, y)
			dotMap[coords]++
			newDots = append(newDots, coords)
		}

		total := 0
		for range dotMap {
			total++
		}
		if i == 0 {
			fmt.Printf("part 1: %d\n", total)
		}
		dots = newDots
	}

	letters := [6][39]string{}

	for i := range letters {
		for j := range letters[i] {
			letters[i][j] = " "
		}
	}

	for i := range dots {
		fmt.Sscanf(dots[i], "%d,%d", &x, &y)
		letters[y][x] = "*"
	}

	fmt.Println("part 2: ")
	for i := range letters {
		fmt.Println(letters[i])
	}
}
