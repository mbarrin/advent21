package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func min(s []int64) int64 {
	var min int64 = s[0]
	for _, value := range s {
		if min > value {
			min = value
		}
	}
	return min
}

func fuelCost(f float64) float64 {
	var i, previous, cost float64

	for i = 0; i < f; i++ {
		previous++
		cost += previous
	}

	return cost
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	crabs := make(map[float64]float64)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		for _, x := range split {
			num, _ := strconv.ParseFloat(x, 64)
			crabs[num]++
		}
	}

	cost := []int64{}
	newCost := []int64{}

	var loc float64
	for loc = 0; loc < 2000; loc++ {
		var fuel, newFuel float64
		for k, v := range crabs {
			steps := (k - loc)
			fuel += math.Abs(steps * v)
			newFuel += fuelCost(math.Abs(steps)) * v
		}
		cost = append(cost, int64(fuel))
		newCost = append(newCost, int64(newFuel))
	}

	fmt.Printf("part 1: %v\n", min(cost))
	fmt.Printf("part 2: %v\n", min(newCost))
}
