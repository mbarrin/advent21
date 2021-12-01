package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var singlePrevious int
	var singleIncrease int

	var batchPrevious int
	var batchIncrease int

	var lines []int

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}

	for i := 0; i < len(lines); i++ {
		batch := 0

		for _, j := range lines[i : i+3] {
			batch += j
		}

		if lines[i] > singlePrevious && singlePrevious != 0 {
			singleIncrease++
		}

		if batch > batchPrevious && batchPrevious != 0 {
			batchIncrease++
		}

		singlePrevious = lines[i]
		batchPrevious = batch
	}

	fmt.Printf("part 1: %d\n", singleIncrease)
	fmt.Printf("part 2: %d\n", batchIncrease)
}
