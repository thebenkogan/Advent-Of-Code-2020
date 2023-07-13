package main

import (
	"fmt"
	"os"
	"sort"
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

func getSortedAdapters(input string) []int {
	nums := []int{0}

	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	sort.Ints(nums)

	builtIn := nums[len(nums)-1] + 3
	nums = append(nums, builtIn)

	return nums
}

func part1(input string) int {
	nums := getSortedAdapters(input)

	oneJolt := 0
	threeJolt := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == 1 {
			oneJolt++
		} else if nums[i+1]-nums[i] == 3 {
			threeJolt++
		}
	}

	return oneJolt * threeJolt
}

func part2(input string) int {
	nums := getSortedAdapters(input)

	dp := make([]int, len(nums))
	dp[0] = 1

	// opt(j) = number of ways to get to adapter j
	// opt(j) = opt(j-1) + opt(j-2) + opt(j-3)

	for i := 1; i < len(dp); i++ {
		total := 0

		if nums[i]-nums[i-1] == 1 || nums[i]-nums[i-1] == 2 || nums[i]-nums[i-1] == 3 {
			total += dp[i-1]
		}
		if i-2 >= 0 && nums[i]-nums[i-2] == 2 {
			total += dp[i-2]
		}
		if i-3 >= 0 && nums[i]-nums[i-3] == 3 {
			total += dp[i-3]
		}
		dp[i] = total
	}

	return dp[len(dp)-1]
}
