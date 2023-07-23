package main

import (
	"fmt"
	"os"
	"regexp"
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

func parseTicket(ticketStr string) []int {
	ticket := make([]int, 0)
	for _, s := range strings.Split(ticketStr, ",") {
		val, _ := strconv.Atoi(s)
		ticket = append(ticket, val)
	}
	return ticket
}

func parseTickets(ticketsStr string) [][]int {
	ticketsSplit := strings.Split(ticketsStr, "\n")[1:]
	tickets := make([][]int, 0)
	for _, ticketStr := range ticketsSplit {
		tickets = append(tickets, parseTicket(ticketStr))
	}
	return tickets
}

var rangeRegex = regexp.MustCompile(`(\d+)-(\d+)`)

type Range struct {
	low  int
	high int
}

func parseRanges(input string) map[string][]Range {
	ranges := make(map[string][]Range)
	for _, fieldStr := range strings.Split(input, "\n") {
		field := strings.Split(fieldStr, ":")[0]
		fieldRanges := make([]Range, 0)
		for _, rangeStr := range rangeRegex.FindAllStringSubmatch(fieldStr, -1) {
			low, _ := strconv.Atoi(rangeStr[1])
			high, _ := strconv.Atoi(rangeStr[2])
			fieldRanges = append(fieldRanges, Range{low, high})
		}
		ranges[field] = fieldRanges
	}
	return ranges
}

func (r *Range) contains(n int) bool {
	return r.low <= n && r.high >= n
}

func validateNum(n int, ranges []Range) bool {
	for _, r := range ranges {
		if r.contains(n) {
			return true
		}
	}
	return false
}

func part1(input string) int {
	sections := strings.Split(input, "\n\n")
	fieldRanges := parseRanges(sections[0])
	ranges := make([]Range, 0)
	for _, rs := range fieldRanges {
		ranges = append(ranges, rs...)
	}

	nearbyTickets := parseTickets(sections[2])
	total := 0
	for _, ticket := range nearbyTickets {
		for _, val := range ticket {
			valid := validateNum(val, ranges)
			if !valid {
				total += val
			}
		}
	}

	return total
}

func part2(input string) int {
	sections := strings.Split(input, "\n\n")
	fieldRanges := parseRanges(sections[0])
	ranges := make([]Range, 0)
	for _, rs := range fieldRanges {
		ranges = append(ranges, rs...)
	}

	// get the valid tickets
	nearbyTickets := parseTickets(sections[2])
	validTickets := make([][]int, 0)
	for _, ticket := range nearbyTickets {
		validTicket := true
		for _, val := range ticket {
			valid := validateNum(val, ranges)
			if !valid {
				validTicket = false
				break
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	// field name -> set of ticket column indices that could be that field
	matches := make(map[string]map[int]bool)
	for field := range fieldRanges {
		matches[field] = make(map[int]bool)
	}

	// populate matches map by finding which fields correspond to which columns
	var single string
	for field, ranges := range fieldRanges {
		for i := 0; i < len(validTickets[0]); i++ {
			isMatch := true
			for _, ticket := range validTickets {
				if !validateNum(ticket[i], ranges) {
					isMatch = false
					break
				}
			}
			if isMatch {
				matches[field][i] = true
			}
		}
		if len(matches[field]) == 1 {
			single = field
		}
	}

	// figure out the order of fields by popping the match with 1 index
	// and deleting that index from every other index set, repeat
	order := make([]string, len(validTickets[0]))
	for len(matches) > 0 {
		var toDelete int
		for k := range matches[single] {
			toDelete = k
		}
		order[toDelete] = single
		delete(matches, single)

		for field, idxs := range matches {
			delete(idxs, toDelete)
			if len(idxs) == 1 {
				single = field
			}
		}
	}

	// for all fields with "departure" prefix, multipy by my ticket's value
	myTicket := parseTicket(strings.Split(sections[1], "\n")[1])
	product := 1
	for i, field := range order {
		if strings.HasPrefix(field, "departure") {
			product *= myTicket[i]
		}
	}

	return product
}
