package main

import (
	"fmt"
	"os"
	"strconv"
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

func nthSpoken(input string, n int) int {
	lastSaidIndex := make(map[int]int)
	lastSaid, difference := 0, -1
	for i, s := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(s)
		lastSaidIndex[num] = i
		lastSaid = num
	}

	numSaid := len(lastSaidIndex)
	for numSaid < n {
		if difference == -1 {
			lastSaid = 0
		} else {
			lastSaid = difference
		}
		index, spokenBefore := lastSaidIndex[lastSaid]
		if spokenBefore {
			difference = numSaid - index
		} else {
			difference = -1
		}
		lastSaidIndex[lastSaid] = numSaid
		numSaid++
	}

	return lastSaid
}

func part1(input string) int {
	return nthSpoken(input, 2020)
}

func part2(input string) int {
	return nthSpoken(input, 30000000)
}
