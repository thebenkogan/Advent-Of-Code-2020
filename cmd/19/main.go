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

type Resolver struct {
	rules    map[int][][]int
	resolved map[int][]string
}

func parseRules(rulesStr string) Resolver {
	rules := make(map[int][][]int)
	resolved := make(map[int][]string)
	for _, line := range strings.Split(rulesStr, "\n") {
		split := strings.Split(line, ":")
		ruleNumber, _ := strconv.Atoi(split[0])
		if string(split[1][1]) == "\"" {
			resolved[ruleNumber] = []string{string(split[1][2])}
		} else {
			matches := make([][]int, 0)
			matchesStr := strings.Split(split[1], "|")
			for _, match := range matchesStr {
				numsStr := strings.Split(strings.Trim(match, " "), " ")
				nums := make([]int, 0)
				for _, numStr := range numsStr {
					num, _ := strconv.Atoi(numStr)
					nums = append(nums, num)
				}
				matches = append(matches, nums)
			}
			rules[ruleNumber] = matches
		}
	}
	return Resolver{rules, resolved}
}

// returns all possible strings by taking one substring from each array
func getAllStrings(matches [][]string, prefix string) []string {
	if len(matches) == 0 {
		return []string{prefix}
	}

	combos := make([]string, 0)
	for _, match := range matches[0] {
		combos = append(combos, getAllStrings(matches[1:], prefix+match)...)
	}
	return combos
}

func (r *Resolver) resolve(ruleNumber int) []string {
	if resolved, ok := r.resolved[ruleNumber]; ok {
		return resolved
	}

	matches := make([]string, 0)
	rules := r.rules[ruleNumber]
	for _, rule := range rules {
		ruleMatches := make([][]string, 0)
		for _, num := range rule {
			ruleMatches = append(ruleMatches, r.resolve(num))
		}
		matches = append(matches, getAllStrings(ruleMatches, "")...)
	}

	r.resolved[ruleNumber] = matches
	return matches
}

func part1(input string) int {
	sections := strings.Split(input, "\n\n")
	resolver := parseRules(sections[0])

	matches := resolver.resolve(0)
	matchSet := make(map[string]bool)
	for _, match := range matches {
		matchSet[match] = true
	}

	total := 0
	for _, msg := range strings.Split(sections[1], "\n") {
		if _, ok := matchSet[msg]; ok {
			total += 1
		}
	}
	return total
}

func part2(input string) int {
	sections := strings.Split(input, "\n\n")
	resolver := parseRules(sections[0])
	resolver.rules[8] = [][]int{{42}, {42, 8}}
	resolver.rules[11] = [][]int{{42, 31}, {42, 11, 31}}

	// 8 is all messages with any number of 42's
	// 11 is all messages as 42 42 42 ... 31 31 31 ... (equal # of each)
	// 8 | 11 is then any message as 42 42 42 42 ... 31 31 31 ...
	// i.e. any number of 42's followed by one or more 31's
	// AND number of 31's < number of 42's <-- SNEAKY SNEAKY

	matches42 := resolver.resolve(42)
	matches31 := resolver.resolve(31)
	matches42Set := make(map[string]bool)
	matches31Set := make(map[string]bool)
	for _, match := range matches42 {
		matches42Set[match] = true
	}
	for _, match := range matches31 {
		matches31Set[match] = true
	}
	matchSize := len(matches42[0])

	total := 0
	for _, msg := range strings.Split(sections[1], "\n") {
		valid := true
		matching42 := true
		num42 := 0
		num31 := 0
		for i := 0; i < len(msg); i += matchSize {
			_, in42 := matches42Set[msg[i:i+matchSize]]
			_, in31 := matches31Set[msg[i:i+matchSize]]
			if in42 {
				num42++
			}
			if in31 {
				num31++
			}
			if matching42 && !in42 {
				matching42 = false
				if !in31 {
					valid = false
					break
				}
			} else if !matching42 && !in31 {
				valid = false
				break
			}
		}
		if valid && !matching42 && num31 < num42 {
			total++
		}
	}
	return total
}
