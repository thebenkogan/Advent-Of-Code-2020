package main

import (
	"fmt"
	"math"
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

const PREAMBLE_LENGTH = 25

func parseNums(input string) []int {
	nums := make([]int, 0)
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return nums
}

func twoSum(nums []int, target int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[target-num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func part1(input string) int {
	nums := parseNums(input)

	for i := PREAMBLE_LENGTH; i < len(nums); i++ {
		if !twoSum(nums[i-PREAMBLE_LENGTH:i], nums[i]) {
			return nums[i]
		}
	}

	return 0
}

func minMax(nums []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func findWeakness(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == target {
				min, max := minMax(nums[i : j+1])
				return min + max
			}
			if sum > target {
				break
			}
		}
	}
	return 0
}

func part2(input string) int {
	invalidNum := part1(input)
	nums := parseNums(input)
	return findWeakness(nums, invalidNum)
}
