// Advent of Code 2024, Day 1
package main

import (
	"fmt"
	"sort"

	"github.com/ghonzo/advent2024/common"
)

// Day 1: Historian Hysteria
// Part 1 answer: 2192892
// Part 2 answer: 22962826
func main() {
	fmt.Println("Advent of Code 2024, Day 1")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	left := make([]int, len(entries))
	right := make([]int, len(entries))
	for i, line := range entries {
		values := common.ConvertToInts(line)
		left[i] = values[0]
		right[i] = values[1]
	}
	sort.Ints(left)
	sort.Ints(right)
	for i, l := range left {
		total += common.Abs(l - right[i])
	}
	return total
}

func part2(entries []string) int {
	var total int
	left := make([]int, len(entries))
	rightMap := make(map[int]int)
	for i, line := range entries {
		values := common.ConvertToInts(line)
		left[i] = values[0]
		rightMap[values[1]]++
	}
	for _, l := range left {
		total += l * rightMap[l]
	}
	return total
}
