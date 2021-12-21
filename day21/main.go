package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type player struct {
	score    int
	position int
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	players := []player{}

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "")

		num, _ := strconv.Atoi(input[len(input)-1])

		newPlayer := player{score: 0, position: num}

		players = append(players, newPlayer)
	}

	counter := 0
	for {
		moves := ((counter + 2) * 3) % 100

		players[counter%2].position += moves

		players[counter%2].position = players[counter%2].position % 10

		if players[counter%2].position == 0 {
			players[counter%2].score += 10
		} else {
			players[counter%2].score += players[counter%2].position
		}

		if players[counter%2].score >= 1000 {
			foo := counter + 3
			bar := players[(counter%2)+1].score

			fmt.Println(foo * bar)
			break
		}

		counter += 3
	}
}
