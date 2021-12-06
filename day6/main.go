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
	fish := []int64{}

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")

		for _, x := range input {
			num, _ := strconv.ParseInt(x, 10, 0)
			fish = append(fish, num)
		}
	}

	for i := 0; i < 80; i++ {
		newFish := []int64{}
		for j := 0; j < len(fish); j++ {
			if fish[j] == 0 {
				fish[j] = 6
				newFish = append(newFish, 8)
			} else {
				fish[j] -= 1
			}
		}
		fish = append(fish, newFish...)
	}

	fmt.Printf("part 1: %d\n", len(fish))
}
