package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func min(s []float64) float64 {
	var min float64 = s[0]
	for _, value := range s {
		if min > value {
			min = value
		}
	}
	return min
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	crabs := make(map[int64]int64)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		for _, x := range split {
			num, _ := strconv.ParseInt(x, 10, 0)
			crabs[num]++
		}
	}

	cost := []float64{}

	for loc := 0; loc < 1885; loc++ {
		var fuel float64
		var foo int64 = int64(loc)
		for k, v := range crabs {
			tmp := float64((k - foo) * v)
			fuel += math.Abs(tmp)
		}
		cost = append(cost, fuel)
	}

	fmt.Println(min(cost))
}
