package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		split := strings.SplitAfter(scanner.Text(), "|")
		output := strings.Fields(split[1])

		for _, num := range output {
			switch len(num) {
			// 2=1, 3=7, 4=4, 7=9
			case 2, 3, 4, 7:
				total++
			}
		}
	}

	fmt.Printf("part 1: %d\n", total)
}
