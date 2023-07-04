package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	day := os.Args[1]
	var inputFileName string
	if len(os.Args) > 2 && os.Args[2] == "test" {
		inputFileName = "test"
	} else {
		inputFileName = "in"
	}
	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/cmd/%s/%s.txt", pwd, day, inputFileName)
	b, _ := os.ReadFile(path)
	input := string(b)

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v", part2(input))
}

func part1(input string) int {
	total := 0
	for _, group := range strings.Split(input, "\n\n") {
		letters := make(map[rune]bool)
		for _, c := range strings.Join(strings.Split(group, "\n"), "") {
			letters[c] = true
		}
		total += len(letters)
	}
	return total
}

func intersection(a, b map[rune]bool) map[rune]bool {
	c := make(map[rune]bool)
	for k := range a {
		if b[k] {
			c[k] = true
		}
	}
	return c
}

func part2(input string) int {
	total := 0
	for _, group := range strings.Split(input, "\n\n") {
		var intersect map[rune]bool
		for _, person := range strings.Split(group, "\n") {
			letters := make(map[rune]bool)
			for _, c := range person {
				letters[c] = true
			}
			if intersect == nil {
				intersect = letters
			} else {
				intersect = intersection(intersect, letters)
			}
		}
		total += len(intersect)
	}
	return total
}
