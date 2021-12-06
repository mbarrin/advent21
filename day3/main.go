package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func makeDecimal(bits []int64) float64 {
	var end float64 = 0
	for i := range bits {
		end += float64(bits[i]) * math.Pow(2.0, float64(len(bits)-i-1))
	}
	return end
}

func keepLines(grid [][]int64, index, toKeep int) [][]int64 {
	if len(grid) == 1 {
		return grid
	}
	newGrid := [][]int64{}
	for y := range grid {
		if grid[y][index] == int64(toKeep) {
			newGrid = append(newGrid, grid[y])
		}
	}
	return newGrid
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	input := [][]int64{}

	gamma := []int64{}
	epsilon := []int64{}

	for scanner.Scan() {
		bits := []int64{}
		line := strings.Split(scanner.Text(), "")
		for _, x := range line {
			bit, _ := strconv.ParseInt(x, 2, 0)
			bits = append(bits, bit)
		}
		input = append(input, bits)
	}

	oxygen := input
	co2 := input

	for x := 0; x < len(input[0]); x++ {
		count := make(map[int64]int)
		for y := 0; y < len(input); y++ {
			count[input[y][x]] += 1
		}

		if count[1] > count[0] {
			gamma = append(gamma, 1)
			epsilon = append(epsilon, 0)
		} else {
			gamma = append(gamma, 0)
			epsilon = append(epsilon, 1)
		}
	}

	fmt.Printf("part 1: %.0f\n", makeDecimal(gamma)*makeDecimal(epsilon))

	for i := 0; i < len(input[0]); i++ {
		o2count := make(map[int64]int)
		co2count := make(map[int64]int)

		for y := 0; y < len(oxygen); y++ {
			o2count[oxygen[y][i]] += 1
		}

		for y := 0; y < len(co2); y++ {
			co2count[co2[y][i]] += 1
		}

		if o2count[1] >= o2count[0] {
			oxygen = keepLines(oxygen, i, 1)
		} else if o2count[1] < o2count[0] {
			oxygen = keepLines(oxygen, i, 0)
		}

		if co2count[1] < co2count[0] {
			co2 = keepLines(co2, i, 1)
		} else if co2count[1] >= co2count[0] {
			co2 = keepLines(co2, i, 0)
		}
	}

	fmt.Printf("part 2: %.0f\n", makeDecimal(oxygen[0])*makeDecimal(co2[0]))
}
