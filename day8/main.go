// Advent of Code 2024, Day 8
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 8: Resonant Collinearity
// Part 1 answer: 265
// Part 2 answer: 962
func main() {
	fmt.Println("Advent of Code 2024, Day 8")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	grid := common.ArraysGridFromLines(entries)
	antinodeLocations := mapset.NewThreadUnsafeSet[common.Point]()
	for p1 := range grid.AllPoints() {
		v1 := grid.Get(p1)
		if v1 == '.' {
			continue
		}
		for p2 := range grid.AllPoints() {
			if p1 == p2 {
				continue
			}
			if grid.Get(p2) == v1 {
				// Same frequency ... add an antinode
				diff := p2.Sub(p1)
				antinode := p1.Sub(diff)
				if inBounds(grid, antinode) {
					antinodeLocations.Add(antinode)
				}
			}
		}
	}
	return antinodeLocations.Cardinality()
}

func inBounds(grid common.Grid, p common.Point) bool {
	_, ok := grid.CheckedGet(p)
	return ok
}

func part2(entries []string) int {
	grid := common.ArraysGridFromLines(entries)
	antinodeLocations := mapset.NewThreadUnsafeSet[common.Point]()
	for p1 := range grid.AllPoints() {
		v1 := grid.Get(p1)
		if v1 == '.' {
			continue
		}
		for p2 := range grid.AllPoints() {
			if p1 == p2 {
				continue
			}
			if grid.Get(p2) == v1 {
				// Same frequency ... add antinodes
				diff := p2.Sub(p1)
				for p := p1; inBounds(grid, p); p = p.Sub(diff) {
					antinodeLocations.Add(p)
				}
			}
		}
	}
	return antinodeLocations.Cardinality()
}
