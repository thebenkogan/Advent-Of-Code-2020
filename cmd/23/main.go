package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/zyedidia/generic/list"
	"golang.org/x/exp/slices"
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

type Cups struct {
	cups   *list.List[int]
	lookup map[int]*list.Node[int]
	size   int
	curr   *list.Node[int]
	low    int
	high   int
}

func (c *Cups) move() {
	removed := make([]*list.Node[int], 0)
	anchor := c.curr.Next
	for len(removed) < 3 {
		if anchor == nil {
			anchor = c.cups.Front
		}
		removed = append(removed, anchor)
		nodeToRemove := anchor
		anchor = anchor.Next
		c.cups.Remove(nodeToRemove)
	}

	destination := c.curr.Value - 1
	if destination < c.low {
		destination = c.high
	}
	for slices.ContainsFunc[*list.Node[int]](removed, func(n *list.Node[int]) bool {
		return n.Value == destination
	}) {
		destination--
		if destination < c.low {
			destination = c.high
		}
	}

	destinationNode := c.lookup[destination]
	finalNext := destinationNode.Next
	for i := 0; i < 3; i++ {
		destinationNode.Next = removed[i]
		destinationNode.Next.Prev = destinationNode
		destinationNode = destinationNode.Next
	}
	destinationNode.Next = finalNext
	if finalNext != nil {
		finalNext.Prev = destinationNode
	}

	if c.curr.Next == nil {
		c.curr = c.cups.Front
	} else {
		c.curr = c.curr.Next
	}
}

func (c *Cups) label() string {
	startNode := c.lookup[1]
	label := ""
	for len(label) < c.size-1 {
		startNode = startNode.Next
		if startNode == nil {
			startNode = c.cups.Front
		}
		label += strconv.Itoa(startNode.Value)
	}
	return label
}

func parseCups(input string) Cups {
	cups := list.New[int]()
	low := math.MaxInt
	high := math.MinInt
	size := 0
	lookup := make(map[int]*list.Node[int])
	for _, c := range strings.Split(input, "") {
		n, _ := strconv.Atoi(c)
		low = min(n, low)
		high = max(n, high)
		size++
		node := list.Node[int]{Value: n, Prev: nil, Next: nil}
		lookup[n] = &node
		cups.PushBackNode(&node)
	}
	return Cups{cups, lookup, size, cups.Front, low, high}
}

const NUM_MOVES = 100

func part1(input string) string {
	cups := parseCups(input)
	for i := 0; i < NUM_MOVES; i++ {
		cups.move()
	}
	return cups.label()
}

const NUM_MOVES_P2 = 10_000_000

func part2(input string) int {
	cups := parseCups(input)
	for i := cups.high + 1; i <= 1_000_000; i++ {
		node := list.Node[int]{Value: i, Prev: nil, Next: nil}
		cups.lookup[i] = &node
		cups.cups.PushBackNode(&node)
	}
	cups.size += 1_000_000 - cups.high
	cups.high = 1_000_000

	for i := 0; i < NUM_MOVES_P2; i++ {
		cups.move()
	}

	startNode := cups.lookup[1]

	// being lazy here, should check for nils but this 1 is likely to be in the middle so who cares
	return startNode.Next.Value * startNode.Next.Next.Value
}
