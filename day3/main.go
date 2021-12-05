package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func makeDecimal(s []int64) float64 {
	var end float64 = 0
	for i := range s {
		end += float64(s[i]) * math.Pow(2.0, float64(len(s)-i-1))
	}
	return end
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
}
