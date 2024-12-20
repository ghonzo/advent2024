// Advent of Code 2024, Day 20
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
)

// Day 20: Race Condition
// Part 1 answer: 1381
// Part 2 answer: 982124
func main() {
	fmt.Println("Advent of Code 2024, Day 20")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries, 100))
	fmt.Printf("Part 2: %d\n", part2(entries, 20, 100))
}

func part1(entries []string, minTimeSaved int) int {
	grid := common.ArraysGridFromLines(entries)
	path := findPath(grid)
	// The point on a path pointing to the picosecond
	stepMap := make(map[common.Point]int)
	for tick, p := range path {
		stepMap[p] = tick
	}
	var validCheats int
	// Now step along every point in the path and see if there are cheats
	for tick, cheatStart := range path[:len(path)-minTimeSaved-1] {
		// For each wall
		for p := range cheatStart.SurroundingCardinals() {
			if v, _ := grid.CheckedGet(p); v == '#' {
				// Now see if there are any shortcuts
				for cheatEnd := range p.SurroundingCardinals() {
					if cheatEndStep, ok := stepMap[cheatEnd]; ok {
						timeSaved := cheatEndStep - tick - 2
						if timeSaved >= minTimeSaved {
							validCheats++
						}
					}
				}
			}
		}
	}
	return validCheats
}

func findPath(grid common.Grid) []common.Point {
	var p common.Point
	for p = range grid.AllPoints() {
		if grid.Get(p) == 'S' {
			break
		}
	}
	var path []common.Point
	var previous common.Point
	for {
		path = append(path, p)
		if grid.Get(p) == 'E' {
			return path
		}
		for sc := range p.SurroundingCardinals() {
			if v, _ := grid.CheckedGet(sc); (sc != previous && v == '.') || v == 'E' {
				previous = p
				p = sc
				break
			}
		}
	}
}

// This also works for part 1 if cheatLimit is set to 2
func part2(entries []string, cheatLimit int, minTimeSaved int) int {
	grid := common.ArraysGridFromLines(entries)
	path := findPath(grid)
	// The point on a path pointing to the picosecond
	stepMap := make(map[common.Point]int)
	for i, p := range path {
		stepMap[p] = i
	}
	var validCheats int
	// Now step along every point in the path and find adjacent potential cheat starts
	for tick, cheatStart := range path[:len(path)-cheatLimit-1] {
		// Check all downstream path points and see if we can reach there in the time limit
		for _, cheatEnd := range path[tick+cheatLimit:] {
			dist := cheatEnd.Sub(cheatStart).ManhattanDistance()
			if dist <= cheatLimit {
				timeSaved := stepMap[cheatEnd] - tick - dist
				if timeSaved >= minTimeSaved {
					validCheats++
				}
			}
		}
	}
	return validCheats
}
