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
	dotMap := map[string]int{}
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

	fmt.Sscanf(instructions[0], "fold along %s", &foldInfo)
	foo := strings.Split(foldInfo, "=")

	foldAxis := foo[0]
	foldLine, _ := strconv.ParseInt(foo[1], 10, 0)

	for i := range dots {
		var y, x int64
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
		dotMap[fmt.Sprintf("%d,%d", y, x)]++
	}

	total := 0
	for range dotMap {
		total++
	}
	fmt.Printf("part 1: %d\n", total)
}
