package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getCount(lines []string, interations int) int {
	transformMap := map[string][]string{}

	start := lines[0]
	split := strings.Split(start, "")

	for _, x := range lines[2:] {
		chars := strings.Split(x, "")
		transformMap[strings.Join(chars[0:2], "")] = []string{fmt.Sprintf("%s%s", chars[0], chars[6]), fmt.Sprintf("%s%s", chars[6], chars[1])}
	}

	codes := map[string]int{}
	first := split[0]
	last := split[len(split)-1]

	for i := 0; i < len(split)-1; i++ {
		codes[strings.Join(split[i:i+2], "")]++
	}

	for i := 0; i < interations; i++ {
		newCodes := map[string]int{}
		for k, v := range codes {
			for _, x := range transformMap[k] {
				newCodes[x] += 1 * v
			}
		}
		codes = newCodes
	}

	output := map[string]int{}

	for k, v := range codes {
		split := strings.Split(k, "")

		for _, x := range split {
			output[x] += v
		}
	}

	output[first]++
	output[last]++

	counts := []int{}

	for _, v := range output {
		counts = append(counts, v/2)
	}

	sort.Ints(counts)

	return counts[len(counts)-1] - counts[0]
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("part 1:", getCount(lines, 10))
	fmt.Println("part 2:", getCount(lines, 40))
}
