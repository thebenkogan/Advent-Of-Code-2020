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

type Direction int

const (
	East Direction = iota
	West
	Southeast
	Southwest
	Northeast
	Northwest
)

var directionsMap = map[string]Direction{
	"e":  East,
	"se": Southeast,
	"sw": Southwest,
	"w":  West,
	"nw": Northwest,
	"ne": Northeast,
}

type Coord[T any] struct {
	x T
	y T
}

var directions = []Coord[float32]{
	{-1, 0},
	{1, 0},
	{0.5, -1},
	{-0.5, -1},
	{0.5, 1},
	{-0.5, 1},
}

func lex(input string) []Direction {
	dirs := make([]Direction, 0)
	i := 0
	for i < len(input) {
		end := i + 1
		for end < len(input) {
			if _, ok := directionsMap[input[i:end]]; ok {
				break
			} else {
				end++
			}
		}
		dir := directionsMap[input[i:end]]
		dirs = append(dirs, dir)
		i = end
	}
	return dirs
}

func getCoord(dirs []Direction) Coord[float32] {
	var x float32
	var y float32
	for _, dir := range dirs {
		switch dir {
		case East:
			x += 1
		case West:
			x -= 1
		case Southeast:
			x += 0.5
			y -= 1
		case Southwest:
			x -= 0.5
			y -= 1
		case Northeast:
			x += 0.5
			y += 1
		case Northwest:
			x -= 0.5
			y += 1
		}
	}
	return Coord[float32]{x, y}
}

func hash(coord Coord[float32]) string {
	return fmt.Sprintf("%f#%f", coord.x, coord.y)
}

func fromHash(hash string) Coord[float32] {
	ns := strings.Split(hash, "#")
	x, _ := strconv.ParseFloat(ns[0], 32)
	y, _ := strconv.ParseFloat(ns[1], 32)
	return Coord[float32]{float32(x), float32(y)}
}

// map of coordinate hash to true if the tile is black, false if white
func getTiles(input string) map[string]bool {
	tiles := make(map[string]bool)

	for _, line := range strings.Split(input, "\n") {
		dirs := lex(line)
		coord := getCoord(dirs)
		hash := hash(coord)
		tiles[hash] = !tiles[hash]
	}

	return tiles
}

func countBlack(tiles map[string]bool) int {
	numBlack := 0
	for _, isBlack := range tiles {
		if isBlack {
			numBlack++
		}
	}
	return numBlack
}

func part1(input string) int {
	tiles := getTiles(input)
	return countBlack(tiles)
}

func flip(tiles map[string]bool) {
	blacksToFlip := make([]string, 0)
	whiteCounts := make(map[string]int)
	for tileHash, isBlack := range tiles {
		if !isBlack {
			continue
		}
		coord := fromHash(tileHash)
		blackNeighbors := 0
		for _, dir := range directions {
			neighborHash := hash(Coord[float32]{coord.x + dir.x, coord.y + dir.y})
			if tiles[neighborHash] {
				blackNeighbors++
			} else {
				whiteCounts[neighborHash]++
			}
		}
		if blackNeighbors == 0 || blackNeighbors > 2 {
			blacksToFlip = append(blacksToFlip, tileHash)
		}
	}
	for _, tileHash := range blacksToFlip {
		tiles[tileHash] = !tiles[tileHash]
	}
	for tileHash, count := range whiteCounts {
		if count == 2 {
			tiles[tileHash] = !tiles[tileHash]
		}
	}
}

func part2(input string) int {
	tiles := getTiles(input)
	for i := 0; i < 100; i++ {
		flip(tiles)
	}
	return countBlack(tiles)
}
