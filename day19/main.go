// Advent of Code 2024, Day 1
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2024/common"
	"github.com/oleiade/lane/v2"
)

// Day 19: Linen Layout
// Part 1 answer: 278
// Part 2 answer: 569808947758890
func main() {
	fmt.Println("Advent of Code 2024, Day 19")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	towels := strings.Split(entries[0], ", ")
	var total int
	for _, line := range entries[2:] {
		if isPossible(line, towels) {
			total++
		}
	}
	return total
}

func isPossible(design string, towels []string) bool {
	// string left to match, with the score the length of that string
	pq := lane.NewMinPriorityQueue[string, int]()
	seen := make(map[string]bool)
	pq.Push(design, len(design))
	for !pq.Empty() {
		s, _, _ := pq.Pop()
		if len(s) == 0 {
			return true
		}
		for _, t := range towels {
			if remaining, ok := strings.CutPrefix(s, t); ok && !seen[remaining] {
				pq.Push(remaining, len(remaining))
				seen[remaining] = true
			}
		}
	}
	return false
}

func part2(entries []string) uint64 {
	towels := strings.Split(entries[0], ", ")
	var total uint64
	for _, line := range entries[2:] {
		total += numPossible(line, towels)
	}
	return total
}

func numPossible(design string, towels []string) uint64 {
	numPaths := make(map[string]uint64)
	seen := make(map[string]bool)
	var total uint64
	// string left to match with the score the length of the string
	pq := lane.NewMaxPriorityQueue[string, int]()
	pq.Push(design, len(design))
	numPaths[design]++
	for !pq.Empty() {
		s, _, _ := pq.Pop()
		curNumPaths := numPaths[s]
		for _, t := range towels {
			if remaining, ok := strings.CutPrefix(s, t); ok {
				if len(remaining) == 0 {
					total += curNumPaths
				} else {
					numPaths[remaining] += curNumPaths
					if !seen[remaining] {
						seen[remaining] = true
						pq.Push(remaining, len(remaining))
					}
				}
			}
		}
	}
	return total
}
