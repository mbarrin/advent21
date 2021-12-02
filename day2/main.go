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

	commands := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		commands = append(commands, line)
	}

	x, y, aim := 0, 0, 0
	var direction string
	var count int

	for _, line := range commands {
		fmt.Sscanf(line, "%s %d", &direction, &count)

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
