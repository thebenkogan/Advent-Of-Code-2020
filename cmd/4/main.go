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

func part1(input string) int {
	valid := 0
	regex, _ := regexp.Compile("(byr|iyr|eyr|hgt|hcl|ecl|pid)")
	for _, passport := range strings.Split(input, "\n\n") {
		if len(regex.FindAllString(passport, -1)) == 7 {
			valid++
		}
	}
	return valid
}

func validateField(field string, value string) bool {
	switch field {
	case "byr":
		val, err := strconv.Atoi(value)
		return err == nil && val >= 1920 && val <= 2002
	case "iyr":
		val, err := strconv.Atoi(value)
		return err == nil && val >= 2010 && val <= 2020
	case "eyr":
		val, err := strconv.Atoi(value)
		return err == nil && val >= 2020 && val <= 2030
	case "hgt":
		if strings.HasSuffix(value, "cm") {
			val, err := strconv.Atoi(value[:len(value)-2])
			return err == nil && val >= 150 && val <= 193
		} else if strings.HasSuffix(value, "in") {
			val, err := strconv.Atoi(value[:len(value)-2])
			return err == nil && val >= 59 && val <= 76
		}
		return false
	case "hcl":
		regex, _ := regexp.Compile(`^#[0-9a-f]{6}$`)
		return regex.MatchString(value)
	case "ecl":
		regex, _ := regexp.Compile(`amb|blu|brn|gry|grn|hzl|oth`)
		return regex.MatchString(value)
	case "pid":
		regex, _ := regexp.Compile(`^[0-9]{9}$`)
		return regex.MatchString(value)
	default:
		return false
	}
}

func part2(input string) int {
	valid := 0
	regex, _ := regexp.Compile(`(byr:\S+|iyr:\S+|eyr:\S+|hgt:\S+|hcl:\S+|ecl:\S+|pid:\S+)`)
	for _, passport := range strings.Split(input, "\n\n") {
		matches := regex.FindAllString(passport, -1)
		if len(matches) != 7 {
			continue
		}
		isValid := true
		for _, match := range matches {
			split := strings.Split(match, ":")
			if !validateField(split[0], split[1]) {
				isValid = false
				break
			}
		}
		if isValid {
			valid++
		}
	}
	return valid
}
