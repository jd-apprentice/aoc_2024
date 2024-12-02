package day_02

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

/*
--- Day 2: Red-Nosed Reports ---

Fortunately, the first location The Historians want to search isn't a long walk from the Chief Historian's office.
While the Red-Nosed Reindeer nuclear fusion/fission plant appears to contain no sign of the Chief Historian, the engineers there run up to you as soon as they see you. Apparently, they still talk about the time Rudolph was saved through molecular synthesis from a single electron.
They're quick to add that - since you're already here - they'd really appreciate your help analyzing some unusual data from the Red-Nosed reactor. You turn to check if The Historians are waiting for you, but they seem to have already divided into groups that are currently searching every corner of the facility. You offer to help with the unusual data.
The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9

This example data contains six reports each containing five levels.
The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:
    The levels are either all increasing or all decreasing.
    Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

    7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
    1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
    9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
    1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
    8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
    1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.

So, in this example, 2 reports are safe.
Analyze the unusual data from the engineers. How many reports are safe?

*/

func isSafe(report []int) bool {
	var isIncreasing bool = true
	var isDecreasing bool = true

	for index := 0; index < len(report)-1; index++ {
		var diff int = report[index+1] - report[index]

		if diff > 3 || diff < -3 || diff == 0 {
			return false
		}

		if diff < 0 {
			isIncreasing = false
		}

		if diff > 0 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func canBeSafeWithOneRemoval(report []int) bool {
	for index := 0; index < len(report); index++ {
		var modifiedReport []int = make([]int, 0, len(report)-1)

		var beforeIndex []int = report[:index]
		modifiedReport = append(modifiedReport, beforeIndex...)

		var afterIndex []int = report[index+1:]
		modifiedReport = append(modifiedReport, afterIndex...)

		if isSafe(modifiedReport) {
			return true
		}
	}

	return false
}

func Part1(input []string) string {
	var start time.Time = time.Now()

	var safeCount int = 0

	for _, line := range input {
		var levels []string = strings.Fields(line)
		var report []int = make([]int, len(levels))

		for index, level := range levels {
			num, _ := strconv.Atoi(level)
			report[index] = num
		}

		if isSafe(report) {
			safeCount++
		}
	}

	var elapsed time.Duration = time.Since(start)
	fmt.Printf("Execution time: %v ms\n", elapsed.Microseconds())

	return strconv.Itoa(safeCount)
}

func Part2(input []string) string {
	var start time.Time = time.Now()

	var safeCount int = 0

	for _, line := range input {
		var levels []string = strings.Fields(line)
		var report []int = make([]int, len(levels))

		for index, level := range levels {
			num, _ := strconv.Atoi(level)
			report[index] = num
		}

		if isSafe(report) || canBeSafeWithOneRemoval(report) {
			safeCount++
		}
	}

	var elapsed time.Duration = time.Since(start)
	fmt.Printf("Execution time: %v ms\n", elapsed.Microseconds())

	return strconv.Itoa(safeCount)
}
