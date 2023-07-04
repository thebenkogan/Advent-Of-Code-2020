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

type Instruction struct {
	id  string
	arg int
}

func parseInstructions(input string) []Instruction {
	insts := make([]Instruction, 0)
	for _, inst := range strings.Split(input, "\n") {
		split := strings.Split(inst, " ")
		id := split[0]
		arg, _ := strconv.Atoi(split[1])
		insts = append(insts, Instruction{id, arg})
	}
	return insts
}

func simulate(insts []Instruction) (int, bool) {
	index := 0
	acc := 0
	visited := make(map[int]bool)
	for index < len(insts) {
		if visited[index] {
			return acc, false
		}
		visited[index] = true
		switch insts[index].id {
		case "nop":
			index++
		case "jmp":
			index += insts[index].arg
		case "acc":
			acc += insts[index].arg
			index++
		}
	}
	return acc, true
}

func part1(input string) int {
	insts := parseInstructions(input)
	acc, _ := simulate(insts)
	return acc
}

func part2(input string) int {
	insts := parseInstructions(input)
	for i := 0; i < len(insts); i++ {
		prevId := insts[i].id
		if insts[i].id == "acc" {
			continue
		} else if insts[i].id == "nop" {
			insts[i].id = "jmp"
		} else {
			insts[i].id = "nop"
		}
		if acc, ok := simulate(insts); ok {
			return acc
		}
		insts[i].id = prevId
	}
	return 0
}
