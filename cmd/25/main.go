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

func transform(subjectNumber int, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= subjectNumber
		val %= 20201227
	}
	return val
}

func getLoopSize(subjectNumber int, publicKey int) int {
	val := 1
	loopSize := 1
	for {
		val *= subjectNumber
		val %= 20201227
		if val == publicKey {
			return loopSize
		}
		loopSize++
	}
}

func part1(input string) int {
	split := strings.Split(input, "\n")
	cardPublicKey, _ := strconv.Atoi(split[0])
	doorPublicKey, _ := strconv.Atoi(split[1])
	cardLoopSize := getLoopSize(7, cardPublicKey)
	return transform(doorPublicKey, cardLoopSize)
}

func part2(input string) int {
	return 0
}
