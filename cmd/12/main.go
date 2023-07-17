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

type Position struct {
	x   int
	y   int
	dir int // 0 = east, 1 = south, 2 = west, 3 = north
}

func (p *Position) move(dir int, amount int) {
	switch dir {
	case 0:
		p.x += amount
	case 1:
		p.y -= amount
	case 2:
		p.x -= amount
	case 3:
		p.y += amount
	}
}

func (p *Position) rotate(degrees int) {
	p.dir = ((p.dir+degrees/90)%4 + 4) % 4
}

func part1(input string) float64 {
	startPosition := Position{0, 0, 0}

	for _, line := range strings.Split(input, "\n") {
		letter := line[0:1]
		amount, _ := strconv.Atoi(line[1:])

		switch letter {
		case "N":
			startPosition.move(3, amount)
		case "S":
			startPosition.move(1, amount)
		case "E":
			startPosition.move(0, amount)
		case "W":
			startPosition.move(2, amount)
		case "L":
			startPosition.rotate(-amount)
		case "R":
			startPosition.rotate(amount)
		case "F":
			startPosition.move(startPosition.dir, amount)
		}
	}

	return math.Abs(float64(startPosition.x)) + math.Abs(float64(startPosition.y))
}

type Position2 struct {
	x        int
	y        int
	waypontX int
	waypontY int
}

func (p *Position2) moveWaypoint(dir int, amount int) {
	switch dir {
	case 0:
		p.waypontX += amount
	case 1:
		p.waypontY -= amount
	case 2:
		p.waypontX -= amount
	case 3:
		p.waypontY += amount
	}
}

func (p *Position2) rotate(degrees int) {
	times := (degrees/90 + 4) % 4
	for i := 0; i < times; i++ {
		p.waypontX, p.waypontY = p.waypontY, -p.waypontX
	}
}

func part2(input string) float64 {
	startPosition := Position2{0, 0, 10, 1}

	for _, line := range strings.Split(input, "\n") {
		letter := line[0:1]
		amount, _ := strconv.Atoi(line[1:])

		switch letter {
		case "N":
			startPosition.moveWaypoint(3, amount)
		case "S":
			startPosition.moveWaypoint(1, amount)
		case "E":
			startPosition.moveWaypoint(0, amount)
		case "W":
			startPosition.moveWaypoint(2, amount)
		case "L":
			startPosition.rotate(-amount)
		case "R":
			startPosition.rotate(amount)
		case "F":
			startPosition.x += startPosition.waypontX * amount
			startPosition.y += startPosition.waypontY * amount
		}
	}

	return math.Abs(float64(startPosition.x)) + math.Abs(float64(startPosition.y))
}
