package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/slices"
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
	test := []int{1, 2, 3, 4}
	cloned := slices.Clone(test)

	fmt.Println(cloned)

	return 0
}

func part2(input string) int {
	return 0
}
