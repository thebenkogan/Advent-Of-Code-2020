package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

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

func reverse[T any](data []T) []T {
	out := make([]T, len(data))
	for i, v := range data {
		out[len(data)-i-1] = v
	}
	return out
}

type Tile struct {
	id     int
	top    []string   // the top edge from left to right
	left   []string   // the left edge from bottom to top
	right  []string   // the right edge from bottom to top
	bottom []string   // the bottom edge from left to right
	inner  [][]string // the tile without the edges
}

func (t *Tile) matchesSomeEdge(other Tile) bool {
	return slices.Equal(t.top, other.bottom) ||
		slices.Equal(t.bottom, other.top) ||
		slices.Equal(t.right, other.left) ||
		slices.Equal(t.left, other.right)
}

func (t *Tile) matchesOnRight(other Tile) bool {
	return slices.Equal(t.right, other.left)
}

func (t *Tile) matchesOnBottom(other Tile) bool {
	return slices.Equal(t.bottom, other.top)
}

func (t *Tile) toFlipped() Tile {
	return Tile{
		id:     t.id,
		top:    t.bottom,
		bottom: t.top,
		left:   reverse(t.left),
		right:  reverse(t.right),
		inner:  reverse(t.inner),
	}
}

func (t *Tile) toRotatedRight() Tile {
	inner := make([][]string, 0)
	for i := 0; i < len(t.inner); i++ {
		newRow := make([]string, 0)
		for _, row := range reverse(t.inner) {
			newRow = append(newRow, row[i])
		}
		inner = append(inner, newRow)
	}
	return Tile{
		id:     t.id,
		top:    t.left,
		bottom: t.right,
		left:   reverse(t.bottom),
		right:  reverse(t.top),
		inner:  inner,
	}
}

func (t *Tile) toRotatedLeft() Tile {
	inner := make([][]string, 0)
	for i := 0; i < len(t.inner); i++ {
		newRow := make([]string, 0)
		for _, row := range t.inner {
			newRow = append(newRow, row[len(row)-i-1])
		}
		inner = append(inner, newRow)
	}
	return Tile{
		id:     t.id,
		top:    reverse(t.right),
		bottom: reverse(t.left),
		left:   t.top,
		right:  t.bottom,
		inner:  inner,
	}
}

func (t *Tile) toRotated180() Tile {
	inner := make([][]string, 0)
	for _, row := range reverse(t.inner) {
		inner = append(inner, reverse(row))
	}
	return Tile{
		id:     t.id,
		top:    reverse(t.bottom),
		bottom: reverse(t.top),
		left:   reverse(t.right),
		right:  reverse(t.left),
		inner:  inner,
	}
}

func orientations(t *Tile) []Tile {
	rotations := []Tile{
		*t,
		t.toRotatedLeft(),
		t.toRotatedRight(),
		t.toRotated180(),
	}

	result := make([]Tile, 0)
	for _, rotation := range rotations {
		result = append(result, rotation, rotation.toFlipped())
	}

	return result
}

func parseTile(tileInfoStr string) Tile {
	idStr, tileStr, _ := strings.Cut(tileInfoStr, "\n")
	secondPart := strings.Split(idStr, " ")[1]
	id, _ := strconv.Atoi(secondPart[:len(secondPart)-1])
	lines := strings.Split(tileStr, "\n")

	top := strings.Split(lines[0], "")
	bottom := strings.Split(lines[len(lines)-1], "")
	right := make([]string, len(lines))
	left := make([]string, len(lines))
	inner := make([][]string, 0)
	for i, line := range lines {
		split := strings.Split(line, "")
		left[len(lines)-i-1] = string(split[0])
		right[len(lines)-i-1] = string(split[len(line)-1])
		if i != 0 && i != len(lines)-1 {
			inner = append(inner, split[1:len(split)-1])
		}
	}

	tile := Tile{
		id:     id,
		top:    top,
		bottom: bottom,
		left:   left,
		right:  right,
		inner:  inner,
	}

	return tile
}

func part1(input string) int {
	tiles := make(map[int][]Tile)
	for _, tileInfo := range strings.Split(input, "\n\n") {
		tile := parseTile(tileInfo)
		tiles[tile.id] = orientations(&tile)
	}

	cornerIds := make([]int, 0)
	for tileId := range tiles {
		matches := 0
		tile := tiles[tileId][0]
		for otherTileId := range tiles {
			if tileId == otherTileId {
				continue
			}
			for _, orientation := range tiles[otherTileId] {
				if tile.matchesSomeEdge(orientation) {
					matches++
					break
				}
			}
		}
		if matches == 2 {
			cornerIds = append(cornerIds, tile.id)
		}
	}

	product := 1
	for _, id := range cornerIds {
		product *= id
	}
	return product
}

var tiles map[int][]Tile
var grid [][]Tile
var ROWS int

func search(row int, col int, placedIds map[int]bool) bool {
	if row == ROWS {
		return true
	}
	for tileId, orientations := range tiles {
		if placedIds[tileId] {
			continue
		}
		for _, tile := range orientations {
			valid := true
			if col > 0 {
				valid = valid && grid[row][col-1].matchesOnRight(tile)
			}
			if row > 0 {
				valid = valid && grid[row-1][col].matchesOnBottom(tile)
			}
			if valid {
				placedIds[tileId] = true
				grid[row][col] = tile
				nextCol := col + 1
				nextRow := row
				if nextCol == ROWS {
					nextCol = 0
					nextRow++
				}
				if search(nextRow, nextCol, placedIds) {
					return true
				} else {
					placedIds[tileId] = false
				}
			}
		}
	}
	return false // nothing fits here, go back!
}

type Coord struct {
	row int
	col int
}

// assuming this is the top right point of the monster, returns coordinates of all squares in monster
func seaMonsterCoords(row int, col int) []Coord {
	return []Coord{
		{row, col},
		{row + 1, col - 18},
		{row + 1, col - 13},
		{row + 1, col - 12},
		{row + 1, col - 7},
		{row + 1, col - 6},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
		{row + 2, col - 17},
		{row + 2, col - 14},
		{row + 2, col - 11},
		{row + 2, col - 8},
		{row + 2, col - 5},
		{row + 2, col - 2},
	}
}

func part2(input string) int {
	tiles = make(map[int][]Tile)
	for _, tileInfo := range strings.Split(input, "\n\n") {
		tile := parseTile(tileInfo)
		tiles[tile.id] = orientations(&tile)
	}
	ROWS = int(math.Sqrt(float64(len((tiles)))))

	grid = make([][]Tile, 0)
	for i := 0; i < ROWS; i++ {
		grid = append(grid, make([]Tile, ROWS))
	}

	search(0, 0, make(map[int]bool))

	combined := Tile{
		inner: make([][]string, 0),
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < len(grid[i][0].inner); j++ {
			row := make([]string, 0)
			for _, tile := range grid[i] {
				row = append(row, tile.inner[j]...)
			}
			combined.inner = append(combined.inner, row)
		}
	}

	numNonMonster := 0
	for _, tile := range orientations(&combined) {
		image := tile.inner
		found := 0
		for row := 0; row < len(image); row++ {
			for col := 0; col < len(image[row]); col++ {
				if col >= 18 && row < len(image)-2 {
					valid := true
					for _, coord := range seaMonsterCoords(row, col) {
						if image[coord.row][coord.col] == "." {
							valid = false
							break
						}
					}
					if valid {
						found++
						for _, coord := range seaMonsterCoords(row, col) {
							image[coord.row][coord.col] = "O"
						}
					}
				}
			}
		}

		if found > 0 {
			for row := 0; row < len(image); row++ {
				for col := 0; col < len(image[row]); col++ {
					if image[row][col] == "#" {
						numNonMonster++
					}
				}
			}
		}
	}

	return numNonMonster
}
