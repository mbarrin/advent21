package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func tally(s []string) map[string]int {
	output := map[string]int{}

	for _, x := range s {
		output[x]++
	}

	return output
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}
	transformMap := map[string]string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	start := lines[0]

	for _, x := range lines[2:] {
		chars := strings.Split(x, "")
		transformMap[strings.Join(chars[0:2], "")] = chars[6]
	}

	split := strings.Split(start, "")

	for i := 0; i < 10; i++ {
		newSplit := []string{split[0]}

		for j := 0; j < len(split)-1; j++ {
			key := strings.Join(split[j:j+2], "")
			newSplit = append(newSplit, transformMap[key], split[j+1])
		}

		split = newSplit
	}

	vals := []int{}
	info := tally(split)

	for _, v := range info {
		vals = append(vals, v)
	}

	sort.Ints(vals)

	fmt.Println("part 1:", vals[len(vals)-1]-vals[0])

}
