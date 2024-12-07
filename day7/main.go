// Advent of Code 2024, Day 7
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2024/common"
)

// Day 7: Bridge Repair
// Part 1 answer: 6083020304036
// Part 2 answer: 59002246504791
func main() {
	fmt.Println("Advent of Code 2024, Day 7")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	for _, line := range entries {
		parts := strings.Split(line, ": ")
		result := common.Atoi(parts[0])
		numbers := common.ConvertToInts(parts[1])
		if isValid(result, numbers[0], numbers[1:]) {
			total += result
		}
	}
	return total
}

func isValid(result int, intermediate int, remaining []int) bool {
	if len(remaining) == 0 {
		return result == intermediate
	}
	return isValid(result, intermediate+remaining[0], remaining[1:]) ||
		isValid(result, intermediate*remaining[0], remaining[1:])
}

func part2(entries []string) int {
	var total int
	for _, line := range entries {
		parts := strings.Split(line, ": ")
		result := common.Atoi(parts[0])
		numbers := common.ConvertToInts(parts[1])
		if isValid2(result, numbers[0], numbers[1:]) {
			total += result
		}
	}
	return total
}

func isValid2(result int, intermediate int, remaining []int) bool {
	if len(remaining) == 0 {
		return result == intermediate
	}
	return isValid2(result, intermediate+remaining[0], remaining[1:]) ||
		isValid2(result, intermediate*remaining[0], remaining[1:]) ||
		isValid2(result, concat(intermediate, remaining[0]), remaining[1:])
}

func concat(left, right int) int {
	return common.Atoi(fmt.Sprintf("%d%d", left, right))
}
