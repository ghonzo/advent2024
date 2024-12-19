// Advent of Code 2024, Day 1
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2024/common"
	"github.com/oleiade/lane/v2"
)

// Day 1: Linen Layout
// Part 1 answer: 278
// Part 2 answer: 75758894579712 is too low
func main() {
	fmt.Println("Advent of Code 2024, Day 19")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	towels := strings.Split(entries[0], ", ")
	var total int
	for i, line := range entries[2:] {
		fmt.Println("Line ", i)
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
		np := numPossible(line, towels)
		total += np
		fmt.Println(line, " ", np)
	}
	return total
}

type partialDesign struct {
	previous  *partialDesign
	remaining string
	extra     int
}

func (pd *partialDesign) score() int {
	return len(pd.remaining)
}

func (pd *partialDesign) combos() uint64 {
	combos := uint64(1)
	for p := pd.previous; p.previous != nil; p = p.previous {
		//combos *= (p.extra + 1)
		combos <<= p.extra
	}
	return combos
}

func numPossible(design string, towels []string) uint64 {
	// remaining -> partial design
	designMap := make(map[string]*partialDesign)
	var leaves []*partialDesign
	root := &partialDesign{remaining: design}
	pq := lane.NewMinPriorityQueue[*partialDesign, int]()
	pq.Push(root, root.score())
	for !pq.Empty() {
		pd, _, _ := pq.Pop()
		for _, t := range towels {
			if remaining, ok := strings.CutPrefix(pd.remaining, t); ok {
				if len(remaining) == 0 {
					leaves = append(leaves, &partialDesign{previous: pd})
				} else if v, ok := designMap[remaining]; ok {
					v.extra++
				} else {
					newPd := &partialDesign{pd, remaining, 0}
					designMap[remaining] = newPd
					pq.Push(newPd, newPd.score())
				}
			}
		}
	}
	var total uint64
	for _, pd := range leaves {
		total += pd.combos()
	}
	return total
}
