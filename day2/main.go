// Advent of Code 2024, Day 2
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ghonzo/advent2024/common"
)

// Day 2: Red-Nosed Reports
// Part 1 answer: 306
// Part 2 answer: 366
func main() {
	fmt.Println("Advent of Code 2024, Day 2")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var safe int
	for _, line := range entries {
		fields := strings.Fields(line)
		levels := make([]int, len(fields))
		for i, f := range fields {
			levels[i] = common.Atoi(f)
		}
		if isSafe(levels) {
			safe++
		}
	}
	return safe
}

func isSafe(levels []int) bool {
	inc := levels[0] < levels[1]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if (inc && (diff < 1 || diff > 3)) ||
			(!inc && (diff > -1 || diff < -3)) {
			return false
		}
	}
	return true
}

func part2(entries []string) int {
	var safe int
	for _, line := range entries {
		fields := strings.Fields(line)
		levels := make([]int, len(fields))
		for i, f := range fields {
			levels[i] = common.Atoi(f)
		}
		if isSafe(levels) {
			safe++
		} else {
		inner:
			for i := 0; i < len(levels); i++ {
				if isSafe(slices.Concat(levels[:i], levels[i+1:])) {
					safe++
					break inner
				}
			}
		}
	}
	return safe
}
