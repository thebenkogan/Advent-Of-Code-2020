package main

import (
	"strconv"
	"strings"
)

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
