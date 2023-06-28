package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(day string, name string) string {
	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/cmd/%s/%s.txt", pwd, day, name)
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func main() {
	day := os.Args[1]
	var inputFileName string
	if len(os.Args) > 2 && os.Args[2] == "test" {
		inputFileName = "test"
	} else {
		inputFileName = "in"
	}
	input := readInput(day, inputFileName)

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v", part2(input))
}

func part1(input string) int {
	pairs := make(map[int]bool)
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		other := 2020 - num
		if pairs[other] {
			return num * other
		}
		pairs[num] = true
	}
	return 0
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines); i++ {
		first, _ := strconv.Atoi(lines[i])
		target := 2020 - first
		pairs := make(map[int]bool)
		for j := i + 1; j < len(lines); j++ {
			num, _ := strconv.Atoi(lines[j])
			other := target - num
			if pairs[other] {
				return first * num * other
			}
			pairs[num] = true
		}
	}
	return 0
}
