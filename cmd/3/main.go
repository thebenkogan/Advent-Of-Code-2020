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
	trees := 0
	x := 0
	for _, line := range strings.Split(input, "\n") {
		if line[x] == '#' {
			trees++
		}
		x = (x + 3) % len(line)
	}
	return trees
}

func countSlope(lines []string, right int, down int) int {
	trees := 0
	x, y := 0, 0
	for y < len(lines) {
		line := lines[y]
		if line[x] == '#' {
			trees++
		}
		x = (x + right) % len(line)
		y += down
	}
	return trees
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	return countSlope(lines, 1, 1) *
		countSlope(lines, 3, 1) *
		countSlope(lines, 5, 1) *
		countSlope(lines, 7, 1) *
		countSlope(lines, 1, 2)

}
