package main

import (
	"fmt"
	"math"
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

var memRegex = regexp.MustCompile(`mem\[(\d+)\]`)

type Mask struct {
	set   int64
	clear int64
}

func maskFromStr(maskStr string) Mask {
	setStr := ""
	clearStr := ""
	for _, c := range maskStr {
		if string(c) == "X" {
			setStr += "0"
			clearStr += "1"
		} else if string(c) == "1" {
			setStr += "1"
			clearStr += "1"
		} else {
			setStr += "0"
			clearStr += "0"
		}
	}

	set, _ := strconv.ParseInt(setStr, 2, 64)
	clear, _ := strconv.ParseInt(clearStr, 2, 64)

	return Mask{
		set:   set,
		clear: clear,
	}
}

func part1(input string) int64 {
	memory := make(map[int]int64)
	lines := strings.Split(input, "\n")
	var mask Mask
	for _, line := range lines {
		valueStr := strings.Split(line, " ")[2]
		if line[:3] == "mem" {
			address, _ := strconv.Atoi(memRegex.FindStringSubmatch(line)[1])
			value, _ := strconv.ParseInt(valueStr, 10, 64)
			value |= mask.set
			value &= mask.clear
			memory[address] = value
		} else {
			mask = maskFromStr(valueStr)
		}
	}

	var total int64
	for _, v := range memory {
		total += v
	}

	return total
}

type MaskV2 struct {
	set         int64
	floatSets   []int64
	floatClears []int64
}

func maskFromStrV2(maskStr string) MaskV2 {
	setStr := ""
	floatSets := make([]int64, 0)
	floatClears := make([]int64, 0)
	index := 35
	for _, c := range maskStr {
		if string(c) == "X" {
			setStr += string("0")
			floatSets = append(floatSets, 1<<index)
			floatClears = append(floatClears, math.MaxInt64&(^(1 << index)))
		} else {
			setStr += string(c)
		}
		index--
	}

	set, _ := strconv.ParseInt(setStr, 2, 64)

	return MaskV2{
		set:         set,
		floatSets:   floatSets,
		floatClears: floatClears,
	}
}

func (m *MaskV2) generateAddresses(address int64) []int64 {
	address |= m.set
	addresses := []int64{address}
	for i := 0; i < len(m.floatSets); i++ {
		newAddresses := make([]int64, 0)
		for _, address := range addresses {
			newAddresses = append(newAddresses, address|m.floatSets[i], address&m.floatClears[i])
		}
		addresses = newAddresses
	}
	return addresses
}

func part2(input string) int {
	memory := make(map[int64]int)
	lines := strings.Split(input, "\n")
	var mask MaskV2
	for _, line := range lines {
		valueStr := strings.Split(line, " ")[2]
		if line[:3] == "mem" {
			addressBase, _ := strconv.ParseInt(memRegex.FindStringSubmatch(line)[1], 10, 64)
			addresses := mask.generateAddresses(addressBase)
			value, _ := strconv.Atoi(valueStr)
			for _, address := range addresses {
				memory[address] = value
			}
		} else {
			mask = maskFromStrV2(valueStr)
		}
	}

	var total int
	for _, v := range memory {
		total += v
	}

	return total
}
