package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	charMap := map[string]string{
		"}": "{",
		"]": "[",
		">": "<",
		")": "(",
	}
	reverseCharMap := map[string]string{
		"{": "}",
		"[": "]",
		"<": ">",
		"(": ")",
	}
	charScore := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	secondCharScore := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	badChars := []string{}
	scores := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		badLine := false
		charList := []string{}
		chars := strings.Split(line, "")
		for _, char := range chars {
			if char == "{" || char == "(" || char == "<" || char == "[" {
				charList = append(charList, char)
			} else {
				if charList[len(charList)-1] != charMap[char] {
					badChars = append(badChars, char)
					badLine = true
					break
				} else {
					charList = charList[:len(charList)-1]
				}
			}
		}
		if !badLine {
			total := 0
			for i := len(charList) - 1; i >= 0; i-- {
				char := reverseCharMap[charList[i]]
				total = (total * 5) + secondCharScore[char]
			}
			scores = append(scores, total)
		}
	}

	total := 0
	for _, char := range badChars {
		total += charScore[char]
	}
	fmt.Printf("part 1: %d\n", total)

	sort.Ints(scores)
	fmt.Printf("part 2: %d\n", scores[len(scores)/2])
}
