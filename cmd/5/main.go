package main

import (
	"fmt"
	"os"
	"sort"
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

func isUpperHalf(char string) bool {
	return char == "B" || char == "R"
}

func searchRange(input string, lower int, upper int) int {
	for _, char := range input {
		if isUpperHalf(string(char)) {
			lower = (lower+upper)/2 + 1
		} else {
			upper = (lower + upper) / 2
		}
	}
	return lower
}

func part1(input string) int {
	max := 0
	for _, boardingPass := range strings.Split(input, "\n") {
		row := searchRange(boardingPass[:7], 0, 127)
		col := searchRange(boardingPass[7:], 0, 7)
		seatId := row*8 + col
		if seatId > max {
			max = seatId
		}
	}
	return max
}

func part2(input string) int {
	seatCounts := make(map[int]int)
	for _, boardingPass := range strings.Split(input, "\n") {
		row := searchRange(boardingPass[:7], 0, 127)
		col := searchRange(boardingPass[7:], 0, 7)
		seatId := row*8 + col
		seatCounts[seatId+1]++
		seatCounts[seatId-1]++
	}

	oneCounts := make([]int, 0)
	for seatId, count := range seatCounts {
		if count == 1 {
			oneCounts = append(oneCounts, seatId)
		}
	}

	sort.Slice(oneCounts, func(i, j int) bool {
		return oneCounts[i] < oneCounts[j]
	})

	return oneCounts[2] + 1
}
