package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func totalFish(fish map[int64]int) int {
	count := 0
	for _, v := range fish {
		count += v
	}
	return count
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	fish := make(map[int64]int)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")

		for _, x := range input {
			num, _ := strconv.ParseInt(x, 10, 0)
			fish[num] += 1
		}
	}

	for i := 0; i < 256; i++ {
		if i == 80 {
			fmt.Printf("part 1: %d\n", totalFish(fish))
		}

		tmpFish := make(map[int64]int)
		for k, v := range fish {
			if k == 0 {
				tmpFish[6] += v
				tmpFish[8] += v
			} else {
				tmpFish[k-1] += v
			}
		}
		fish = tmpFish
	}
	fmt.Printf("part 2: %d\n", totalFish(fish))
}
