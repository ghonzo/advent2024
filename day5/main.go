// Advent of Code 2024, Day 5
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 5: Print Queue
// Part 1 answer: 4135
// Part 2 answer: 5285
func main() {
	fmt.Println("Advent of Code 2024, Day 5")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	orderingRules, updates := readEntries(entries)
	for _, update := range updates {
		if isOrdered(orderingRules, update) {
			total += update[len(update)/2]
		}
	}
	return total
}

func readEntries(entries []string) (map[int]mapset.Set[int], [][]int) {
	// Key is the page number, values must be before that page
	orderingRules := make(map[int]mapset.Set[int])
	updates := [][]int{}
	var i int
	var line string
	for i, line = range entries {
		if len(line) == 0 {
			break
		}
		rule := common.ConvertToInts(line)
		s, ok := orderingRules[rule[1]]
		if !ok {
			s = mapset.NewSet[int]()
			orderingRules[rule[1]] = s
		}
		s.Add(rule[0])
	}
	// Now for the second part
	for _, line = range entries[i+1:] {
		updates = append(updates, common.ConvertToInts(line))
	}
	return orderingRules, updates
}

func isOrdered(orderingRules map[int]mapset.Set[int], pages []int) bool {
	for i, pageNum := range pages[:len(pages)-1] {
		if preconditions, ok := orderingRules[pageNum]; ok && preconditions.ContainsAny(pages[i+1:]...) {
			return false
		}
	}
	return true
}

func part2(entries []string) int {
	var total int
	orderingRules, updates := readEntries(entries)
	for _, update := range updates {
		if pages, outOfOrder := orderPages(orderingRules, update); outOfOrder {
			total += pages[len(pages)/2]
		}
	}
	return total
}

// Returns the ordered pages, and true if we had to make any changes
func orderPages(orderingRules map[int]mapset.Set[int], pages []int) ([]int, bool) {
	for i, pageNum := range pages[:len(pages)-1] {
		for j, pageNum2 := range pages[i+1:] {
			if preconditions, ok := orderingRules[pageNum]; ok && preconditions.Contains(pageNum2) {
				pages[i], pages[i+j+1] = pages[i+j+1], pages[i]
				pages, _ = orderPages(orderingRules, pages)
				return pages, true
			}
		}
	}
	return pages, false
}
