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
	charMap := map[string]string{
		"}": "{",
		"]": "[",
		">": "<",
		")": "(",
	}
	charScore := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	badChars := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		charList := []string{}
		chars := strings.Split(line, "")
		for _, char := range chars {
			if char == "{" || char == "(" || char == "<" || char == "[" {
				charList = append(charList, char)
			} else {
				if charList[len(charList)-1] != charMap[char] {
					badChars = append(badChars, char)
					break
				} else {
					charList = charList[:len(charList)-1]
				}
			}
		}
	}
	total := 0
	for _, char := range badChars {
		total += charScore[char]
	}
	fmt.Println(total)
}
