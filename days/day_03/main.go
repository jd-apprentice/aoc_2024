package day_03

import (
	"fmt"
	"regexp"
	"strconv"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	var regex *regexp.Regexp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	var totalSum int = 0

	for _, line := range input {
		for _, match := range regex.FindAllStringSubmatch(line, -1) {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])

			totalSum += x * y
		}
	}

	return strconv.Itoa(totalSum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
