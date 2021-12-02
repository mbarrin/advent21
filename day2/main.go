package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	x, y, aim := 0, 0, 0

	for scanner.Scan() {
		var direction string
		var count int

		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &count)

		switch direction {
		case "up":
			aim = aim - count
		case "down":
			aim = aim + count
		case "forward":
			x = x + count
			y = y + (aim * count)
		}
	}

	fmt.Printf("part 1: %d\n", x*aim)
	fmt.Printf("part 2: %d\n", x*y)
}
