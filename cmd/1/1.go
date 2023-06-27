package main

import (
	"strconv"
	"strings"
)

func part1(input string) int {
	pairs := make(map[int]bool)
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		other := 2020 - num
		if pairs[other] {
			return num * other
		}
		pairs[num] = true
	}
	return 0
}
