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

	defer file.Close()

	scanner := bufio.NewScanner(file)

	commands := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		commands = append(commands, line)
	}

	x, y, aim := 0, 0, 0
	for _, line := range commands {
		command := strings.Split(line, " ")

		direction := command[0]
		count, _ := strconv.Atoi(command[1])

		switch direction {
		case "up":
			y = y - count
		case "down":
			y = y + count
		case "forward":
			x = x + count
		}
	}

	fmt.Printf("part 1: %d\n", x*y)

	x, y, aim = 0, 0, 0
	for _, line := range commands {
		command := strings.Split(line, " ")

		direction := command[0]
		count, _ := strconv.Atoi(command[1])

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

	fmt.Printf("part 2: %d\n", x*y)
}
