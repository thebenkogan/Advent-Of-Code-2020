package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/monitor1379/yagods/stacks/arraystack"
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

func countOuterBags(adj map[string][]string, start string) int {
	total := 0
	stack := arraystack.New[string]()
	visited := make(map[string]bool)
	stack.Push(start)
	for stack.Size() > 0 {
		next, _ := stack.Pop()
		for _, neighbor := range adj[next] {
			if !visited[neighbor] {
				stack.Push(neighbor)
				visited[neighbor] = true
				total++
			}
		}
	}
	return total
}

func part1(input string) int {
	adj := make(map[string][]string)
	containRegex, _ := regexp.Compile(`\d+ (\w+ \w+) bags?`)
	for _, rule := range strings.Split(input, "\n") {
		id := strings.Split(rule, " bags")[0]
		contains := containRegex.FindAllStringSubmatch(rule, -1)
		for _, contain := range contains {
			adj[contain[1]] = append(adj[contain[1]], id)
		}
	}
	return countOuterBags(adj, "shiny gold")
}

type Bag struct {
	id    string
	count int
}

func countInnerBags(adj map[string][]Bag, start string) int {
	total := 0
	stack := arraystack.New[Bag]()
	stack.Push(Bag{id: start, count: 1})
	for stack.Size() > 0 {
		next, _ := stack.Pop()
		if neighbors, ok := adj[next.id]; ok {
			total += next.count
			for _, neighbor := range neighbors {
				stack.Push(Bag{id: neighbor.id, count: neighbor.count * next.count})
			}
		} else {
			total += next.count
		}
	}
	return total - 1 // subtract out start bag
}

func part2(input string) int {
	adj := make(map[string][]Bag)
	containRegex, _ := regexp.Compile(`(\d+) (\w+ \w+) bags?`)
	for _, rule := range strings.Split(input, "\n") {
		id := strings.Split(rule, " bags")[0]
		contains := containRegex.FindAllStringSubmatch(rule, -1)
		for _, contain := range contains {
			count, _ := strconv.Atoi(contain[1])
			adj[id] = append(adj[id], Bag{id: contain[2], count: count})
		}
	}

	return countInnerBags(adj, "shiny gold")
}
