package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func smallCave(cave string) bool {
	return unicode.IsLower(rune(cave[0]))
}

func contains(slice []string, s string) bool {
	for i := range slice {
		if slice[i] == s {
			return true
		}
	}
	return false
}

func findPaths(caveMap *map[string][]string, currentPath []string, allPaths [][]string, cave string) [][]string {
	if contains(currentPath, cave) && smallCave(cave) {
		return allPaths
	}

	currentPath = append(currentPath, cave)

	if cave == "end" {
		allPaths = append(allPaths, currentPath)
		return allPaths
	}

	for _, cave := range (*caveMap)[cave] {
		allPaths = findPaths(caveMap, currentPath, allPaths, cave)
	}

	return allPaths
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	caveMap := map[string][]string{}
	for scanner.Scan() {
		caves := strings.Split(scanner.Text(), "-")
		caveMap[caves[0]] = append(caveMap[caves[0]], caves[1])
		caveMap[caves[1]] = append(caveMap[caves[1]], caves[0])
	}

	foobar := findPaths(&caveMap, []string{}, [][]string{}, "start")

	fmt.Println(len(foobar))
}
