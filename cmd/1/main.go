package main

import (
	"fmt"
	"os"
)

const DAY = 1

func readInput(name string) string {
	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s/cmd/%d/%s.txt", pwd, DAY, name)
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func main() {
	var inputFileName string
	if len(os.Args) > 1 && os.Args[1] == "test" {
		inputFileName = os.Args[1]
	} else {
		inputFileName = "in"
	}
	input := readInput(inputFileName)
	fmt.Println(input)

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v", part2(input))
}
