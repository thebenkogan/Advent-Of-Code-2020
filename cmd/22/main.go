package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	dll "github.com/monitor1379/yagods/lists/doublylinkedlist"
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

func playTillWinner(p1 *dll.List[int], p2 *dll.List[int]) *dll.List[int] {
	for !p1.Empty() && !p2.Empty() {
		p1Val, _ := p1.Get(0)
		p1.Remove(0)
		p2Val, _ := p2.Get(0)
		p2.Remove(0)
		if p1Val > p2Val {
			p1.Append(p1Val)
			p1.Append(p2Val)
		} else {
			p2.Append(p2Val)
			p2.Append(p1Val)
		}
	}
	if !p1.Empty() {
		return p1
	} else {
		return p2
	}
}

func parseCards(input string) (*dll.List[int], *dll.List[int]) {
	players := strings.Split(input, "\n\n")
	p1 := dll.New[int]()
	for _, s := range strings.Split(players[0], "\n")[1:] {
		n, _ := strconv.Atoi(s)
		p1.Append(n)
	}
	p2 := dll.New[int]()
	for _, s := range strings.Split(players[1], "\n")[1:] {
		n, _ := strconv.Atoi(s)
		p2.Append(n)
	}
	return p1, p2
}

func part1(input string) int {
	p1, p2 := parseCards(input)

	winner := playTillWinner(p1, p2)
	total := 0

	winner.Each(func(i int, v int) {
		total += v * (winner.Size() - i)
	})

	return total
}

type Game struct {
	p1         *dll.List[int]
	p2         *dll.List[int]
	gameHashes map[string]bool
}

func (g *Game) copy(p1Amount int, p2Amount int) Game {
	p1Copy := dll.New[int]()
	for i := 0; i < p1Amount && i < g.p1.Size(); i++ {
		val, _ := g.p1.Get(i)
		p1Copy.Append(val)
	}
	p2Copy := dll.New[int]()
	for i := 0; i < p2Amount && i < g.p2.Size(); i++ {
		val, _ := g.p2.Get(i)
		p2Copy.Append(val)
	}
	return Game{
		p1Copy,
		p2Copy,
		make(map[string]bool),
	}
}

func (g *Game) hash() string {
	p1Str := ""
	g.p1.Each(func(index, value int) {
		p1Str += strconv.Itoa(value) + ","
	})
	p2Str := ""
	g.p2.Each(func(index, value int) {
		p2Str += strconv.Itoa(value) + ","
	})
	return p1Str + "." + p2Str
}

func (g *Game) drawCards() (int, int) {
	p1Val, _ := g.p1.Get(0)
	g.p1.Remove(0)
	p2Val, _ := g.p2.Get(0)
	g.p2.Remove(0)
	return p1Val, p2Val
}

func (g *Game) p1WinRound(p1Card int, p2Card int) {
	g.p1.Append(p1Card)
	g.p1.Append(p2Card)
}

func (g *Game) p2WinRound(p1Card int, p2Card int) {
	g.p2.Append(p2Card)
	g.p2.Append(p1Card)
}

var resultsStore = make(map[string]bool)

// true if player 1 won the game
func (g *Game) play() bool {
	gameHash := g.hash()
	if res, ok := resultsStore[gameHash]; ok {
		return res
	}

	for !g.p1.Empty() && !g.p2.Empty() {
		hash := g.hash()
		if g.gameHashes[hash] {
			resultsStore[gameHash] = true
			return true
		}
		g.gameHashes[hash] = true

		p1Card, p2Card := g.drawCards()

		var p1Won bool
		if p1Card <= g.p1.Size() && p2Card <= g.p2.Size() {
			subGame := g.copy(p1Card, p2Card)
			p1Won = subGame.play()
		} else {
			p1Won = p1Card > p2Card
		}

		if p1Won {
			g.p1WinRound(p1Card, p2Card)
		} else {
			g.p2WinRound(p1Card, p2Card)
		}
	}

	resultsStore[gameHash] = !g.p1.Empty()
	return !g.p1.Empty()
}

func part2(input string) int {
	p1, p2 := parseCards(input)
	game := Game{p1, p2, make(map[string]bool)}

	p1Won := game.play()
	winningCards := game.p1
	if !p1Won {
		winningCards = game.p2
	}

	total := 0
	winningCards.Each(func(i int, v int) {
		total += v * (winningCards.Size() - i)
	})
	return total
}
