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

type Cube struct {
	x int
	y int
	z int
}

func (c *Cube) hash() string {
	return fmt.Sprintf("%d#%d#%d", c.x, c.y, c.z)
}

func fromHash(hash string) Cube {
	split := strings.Split(hash, "#")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	z, _ := strconv.Atoi(split[2])
	return Cube{x, y, z}
}

func (c *Cube) surroundingCubes() []Cube {
	cubes := make([]Cube, 0)
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				if x == c.x && y == c.y && z == c.z {
					continue
				}
				cubes = append(cubes, Cube{x, y, z})
			}
		}
	}
	return cubes
}

func applyCycle(activeCubes map[string]Cube) map[string]Cube {
	newCubes := make(map[string]Cube)

	inactiveCubeCounts := make(map[string]int)
	for hash, cube := range activeCubes {
		numActive := 0
		for _, sc := range cube.surroundingCubes() {
			if _, ok := activeCubes[sc.hash()]; ok {
				numActive++
			} else {
				inactiveCubeCounts[sc.hash()] += 1
			}
		}
		if numActive == 2 || numActive == 3 {
			newCubes[hash] = cube
		}
	}

	for hash, count := range inactiveCubeCounts {
		if count == 3 {
			newCubes[hash] = fromHash(hash)
		}
	}

	return newCubes
}

func part1(input string) int {
	cubes := make(map[string]Cube)

	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if string(c) == "#" {
				cube := Cube{x, y, 0}
				cubes[cube.hash()] = cube
			}
		}
	}

	for i := 0; i < 6; i++ {
		cubes = applyCycle(cubes)
	}

	return len(cubes)
}

type Cube4 struct {
	x int
	y int
	z int
	w int
}

func (c *Cube4) hash() string {
	return fmt.Sprintf("%d#%d#%d#%d", c.x, c.y, c.z, c.w)
}

func fromHash4(hash string) Cube4 {
	split := strings.Split(hash, "#")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	z, _ := strconv.Atoi(split[2])
	w, _ := strconv.Atoi(split[3])
	return Cube4{x, y, z, w}
}

func (c *Cube4) surroundingCubes() []Cube4 {
	cubes := make([]Cube4, 0)
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				for w := c.w - 1; w <= c.w+1; w++ {
					if x == c.x && y == c.y && z == c.z && w == c.w {
						continue
					}
					cubes = append(cubes, Cube4{x, y, z, w})
				}
			}
		}
	}
	return cubes
}

func applyCycle4(activeCubes map[string]Cube4) map[string]Cube4 {
	newCubes := make(map[string]Cube4)

	inactiveCubeCounts := make(map[string]int)
	for hash, cube := range activeCubes {
		numActive := 0
		for _, sc := range cube.surroundingCubes() {
			if _, ok := activeCubes[sc.hash()]; ok {
				numActive++
			} else {
				inactiveCubeCounts[sc.hash()] += 1
			}
		}
		if numActive == 2 || numActive == 3 {
			newCubes[hash] = cube
		}
	}

	for hash, count := range inactiveCubeCounts {
		if count == 3 {
			newCubes[hash] = fromHash4(hash)
		}
	}

	return newCubes
}

func part2(input string) int {
	cubes := make(map[string]Cube4)

	for y, line := range strings.Split(input, "\n") {
		for x, c := range line {
			if string(c) == "#" {
				cube := Cube4{x, y, 0, 0}
				cubes[cube.hash()] = cube
			}
		}
	}

	for i := 0; i < 6; i++ {
		cubes = applyCycle4(cubes)
	}

	return len(cubes)
}
