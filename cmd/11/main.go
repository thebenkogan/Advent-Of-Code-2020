package main

import (
	"fmt"
	"os"
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

type Direction struct {
	x int
	y int
}

var directions = []Direction{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func countOccupied(seats [][]string) int {
	count := 0
	for i := range seats {
		for j := range seats[i] {
			if seats[i][j] == "#" {
				count++
			}
		}
	}
	return count
}

func copySeats(seats [][]string) [][]string {
	newSeats := make([][]string, len(seats))
	for i := range seats {
		newSeats[i] = make([]string, len(seats[i]))
		copy(newSeats[i], seats[i])
	}
	return newSeats
}

func inBounds(seats [][]string, i int, j int) bool {
	return i >= 0 && i < len(seats) && j >= 0 && j < len(seats[i])
}

func step(seats [][]string) ([][]string, int) {
	newSeats := copySeats(seats)
	seatsChanged := 0

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == "L" {
				unoccupied := true
				for _, dir := range directions {
					if inBounds(seats, i+dir.x, j+dir.y) {
						if seats[i+dir.x][j+dir.y] == "#" {
							unoccupied = false
							break
						}
					}
				}
				if unoccupied {
					newSeats[i][j] = "#"
					seatsChanged++
				}
			}

			if seats[i][j] == "#" {
				numOccupied := 0
				for _, dir := range directions {
					if inBounds(seats, i+dir.x, j+dir.y) {
						if seats[i+dir.x][j+dir.y] == "#" {
							numOccupied++
						}
					}
				}
				if numOccupied >= 4 {
					newSeats[i][j] = "L"
					seatsChanged++
				}
			}
		}
	}

	return newSeats, seatsChanged
}

func part1(input string) int {
	seats := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		seats = append(seats, strings.Split(line, ""))
	}

	seatsChanged := 1
	for seatsChanged > 0 {
		seats, seatsChanged = step(seats)
	}

	return countOccupied(seats)
}

func stepPart2(seats [][]string) ([][]string, int) {
	newSeats := copySeats(seats)
	seatsChanged := 0

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == "L" {
				unoccupied := true
				for _, dir := range directions {
					x := i + dir.x
					y := j + dir.y
					for inBounds(seats, x, y) && seats[x][y] == "." {
						x += dir.x
						y += dir.y
					}
					if inBounds(seats, x, y) && seats[x][y] == "#" {
						unoccupied = false
						break
					}
				}
				if unoccupied {
					newSeats[i][j] = "#"
					seatsChanged++
				}
			}

			if seats[i][j] == "#" {
				numOccupied := 0
				for _, dir := range directions {
					x := i + dir.x
					y := j + dir.y
					for inBounds(seats, x, y) && seats[x][y] == "." {
						x += dir.x
						y += dir.y
					}
					if inBounds(seats, x, y) && seats[x][y] == "#" {
						numOccupied++
					}
				}
				if numOccupied >= 5 {
					newSeats[i][j] = "L"
					seatsChanged++
				}
			}
		}
	}

	return newSeats, seatsChanged
}

func part2(input string) int {
	seats := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		seats = append(seats, strings.Split(line, ""))
	}

	seatsChanged := 1
	for seatsChanged > 0 {
		seats, seatsChanged = stepPart2(seats)
	}

	return countOccupied(seats)
}
