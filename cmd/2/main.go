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

type Validation struct {
	low      int
	high     int
	letter   string
	password string
}

func validationFromString(line string) Validation {
	split := strings.Split(line, " ")
	rangeStr := strings.Split(split[0], "-")
	low, _ := strconv.Atoi(string(rangeStr[0]))
	high, _ := strconv.Atoi(string(rangeStr[1]))
	return Validation{
		low:      low,
		high:     high,
		letter:   string(split[1][0]),
		password: split[2],
	}
}

func (v *Validation) validate() bool {
	count := strings.Count(v.password, v.letter)
	return count >= v.low && count <= v.high
}

func (v *Validation) tobogganValidate() bool {
	first := string(v.password[v.low-1]) == v.letter
	second := string(v.password[v.high-1]) == v.letter
	return first != second
}

func part1(input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		validation := validationFromString(line)
		if validation.validate() {
			total++
		}
	}
	return total
}

func part2(input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		validation := validationFromString(line)
		if validation.tobogganValidate() {
			total++
		}
	}
	return total
}
