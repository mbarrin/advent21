package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func transpose(s [][]string) [][]string {
	height := len(s)
	width := len(s[0])
	newSlice := make([][]string, width)

	for i := 0; i < width; i++ {
		newSlice[i] = make([]string, height)
		for j := 0; j < height; j++ {
			newSlice[i][j] = s[j][i]
		}
	}
	return newSlice
}

func tally(s []string) map[string]int {
	count := make(map[string]int)

	for _, x := range s {
		count[x]++
	}

	return count
}

func mostCommon(s []string) (string, int) {
	var mostCommonKey string

	count := make(map[string]int)

	for _, x := range s {
		count[x]++
	}

	for k, v := range count {
		if v > count[mostCommonKey] {
			mostCommonKey = k
		}
	}

	return mostCommonKey, count[mostCommonKey]
}

func leastCommon(s []string) (string, int) {
	var leastCommonKey string

	count := make(map[string]int)

	for _, x := range s {
		count[x]++
	}

	for k, v := range count {
		if count[leastCommonKey] == 0 {
			leastCommonKey = k
		} else if v < count[leastCommonKey] {
			leastCommonKey = k
		}
	}

	return leastCommonKey, count[leastCommonKey]
}

func makeDecimal(s []string) int64 {
	num, _ := strconv.ParseInt(strings.Join(s, ""), 2, 0)
	return num
}

func removeLines(bits [][]string, index int, method string) [][]string {
	if len(bits) == 1 {
		return bits
	}
	eugh := transpose(bits)

	var s string

	mCommon, mCount := mostCommon(eugh[index])
	lCommon, lCount := leastCommon(eugh[index])

	if mCount > lCount {
		s = mCommon
	} else {
		s = lCommon
	}

	if mCount == len(eugh[index])/2 && lCount == len(eugh[index])/2 {
		if method == "most" {
			s = "1"
		} else {
			s = "0"
		}
	}

	if len(bits) == 20000 {
		fmt.Printf("\nindex: %v\n", index)
		fmt.Printf("s: %v\n", s)
		fmt.Printf("bits: %v\n", bits)
		fmt.Printf("eugh: %v\n", eugh)
		fmt.Printf("method: %v\n", method)
		fmt.Printf("mCount: %v\n", mCount)
		fmt.Printf("lCount: %v\n", lCount)
		fmt.Printf("mCommon: %v\n", mCommon)
		fmt.Printf("lCommon: %v\n\n", lCommon)
	}

	stripped := [][]string{}

	for _, number := range bits {
		if number[index] == s {
			stripped = append(stripped, number)
		}
	}

	fmt.Printf("stripped: %v\n", stripped)

	return stripped
}

func main() {
	file, _ := os.Open("test.txt")

	scanner := bufio.NewScanner(file)

	input := [][]string{}

	gamma := []string{}
	epsilon := []string{}

	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}

	oxygen := input
	co2 := input

	transposed := transpose(input)

	for i, x := range transposed {
		count := tally(x)
		fmt.Println(count)
		most, _ := mostCommon(x)
		least, _ := leastCommon(x)

		//fmt.Printf("oxygen %v\n", oxygen)
		//fmt.Printf("co2 %v\n", co2)
		//fmt.Printf("i %v\n", i)
		//fmt.Printf("most %v\n", most)

		oxygen = removeLines(oxygen, i, "most")
		co2 = removeLines(co2, i, "least")

		gamma = append(gamma, most)
		epsilon = append(epsilon, least)
	}

	fmt.Println(makeDecimal(oxygen[0]))
	fmt.Println(makeDecimal(co2[0]))

	fmt.Printf("part 1: %d\n", makeDecimal(gamma)*makeDecimal(epsilon))
	fmt.Printf("part 2: %d\n", makeDecimal(oxygen[0])*makeDecimal(co2[0]))
}
