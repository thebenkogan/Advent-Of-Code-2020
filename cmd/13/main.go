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

func part1(input string) int {
	lines := strings.Split(input, "\n")

	timestamp, _ := strconv.Atoi(lines[0])
	minWaitTime := timestamp
	minId := 0
	for _, busId := range strings.Split(lines[1], ",") {
		if busId != "x" {
			id, _ := strconv.Atoi(busId)
			waitTime := id - (timestamp % id)
			if waitTime < minWaitTime {
				minId = id
				minWaitTime = waitTime
			}
		}
	}

	return minId * minWaitTime
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	busIds := make([]int, 0)
	for _, busId := range strings.Split(lines[1], ",") {
		if busId != "x" {
			id, _ := strconv.Atoi(busId)
			busIds = append(busIds, id)
		} else {
			busIds = append(busIds, 0)
		}
	}

	t := 0
	step := busIds[0]

	for i, busId := range busIds[1:] {
		if busId != 0 {
			for ((t + i + 1) % busId) != 0 {
				t += step
			}
			step *= busId
		}
	}

	return t
}
