// Advent of Code 2024, Day 20
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 20: Race Condition
// Part 1 answer: 1381
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 20")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries, 100))
	fmt.Printf("Part 2: %d\n", part2(entries, 20, 100))
}

type cheat [2]common.Point

func part1(entries []string, minTimeSaved int) int {
	grid, start, end := readGrid(entries)
	path := findPath(grid, start, end)
	// The point on a path pointing to the picosecond
	stepMap := make(map[common.Point]int)
	for i, p := range path {
		stepMap[p] = i
	}
	validCheats := make(map[cheat]bool)
	// Now step along every point in the path and see if there are cheats
	for i, p := range path[:len(path)-minTimeSaved-1] {
		// For each wall
		for cheatStart := range p.SurroundingCardinals() {
			if v, _ := grid.CheckedGet(cheatStart); v == '#' {
				// Now see if there are any shortcuts
				for cheatEnd := range cheatStart.SurroundingCardinals() {
					if cheatEndStep, ok := stepMap[cheatEnd]; ok {
						timeSaved := cheatEndStep - i - 2
						if timeSaved >= minTimeSaved {
							validCheats[[2]common.Point{cheatStart, cheatEnd}] = true
						}
					}
				}
			}
		}
	}
	return len(validCheats)
}

func readGrid(entries []string) (grid common.Grid, start, end common.Point) {
	grid = common.ArraysGridFromLines(entries)
	for p := range grid.AllPoints() {
		v := grid.Get(p)
		if v == 'S' {
			start = p
		} else if v == 'E' {
			end = p
		}
	}
	return
}

func findPath(grid common.Grid, start, end common.Point) []common.Point {
	path := []common.Point{start}
	for p := start; p != end; p = nextPoint(grid, p) {
		path = append(path, p)
		grid.Set(p, 'O')
	}
	path = append(path, end)
	return path
}

func nextPoint(grid common.Grid, p common.Point) common.Point {
	for sc := range p.SurroundingCardinals() {
		if v, _ := grid.CheckedGet(sc); v == '.' || v == 'E' {
			return sc
		}
	}
	panic("No next point")
}

func part2(entries []string, cheatLimit int, minTimeSaved int) int {
	grid, start, end := readGrid(entries)
	path := findPath(grid, start, end)
	// The point on a path pointing to the picosecond
	stepMap := make(map[common.Point]int)
	for i, p := range path {
		stepMap[p] = i
	}
	var validCheats int
	// This is for deugging only . Time saved pointing to number of cheats
	timeSavedMap := make(map[int]int)
	cheatStartDone := mapset.NewThreadUnsafeSet[common.Point]()
	// Now step along every point in the path and find adjacent potential cheat starts
	for tick, p := range path[:len(path)-cheatLimit-1] {
		for cheatStart := range p.SurroundingCardinals() {
			if v, _ := grid.CheckedGet(cheatStart); v == '#' && !cheatStartDone.Contains(cheatStart) {
				cheatStartDone.Add(cheatStart)
				// Check all downstream path points and see if we can reach there in 20 ticks or less
				for _, cheatEnd := range path[tick+cheatLimit:] {
					dist := cheatEnd.Sub(cheatStart).ManhattanDistance()
					if dist <= cheatLimit-1 {
						timeSaved := stepMap[cheatEnd] - tick - dist - 1
						if timeSaved >= minTimeSaved {
							validCheats++
							// For debug
							timeSavedMap[timeSaved]++
						}
					}
				}
			}
		}
	}
	// Debug
	for k, v := range timeSavedMap {
		fmt.Printf("There are %d cheats that save %d picoseconds\n", v, k)
	}
	return validCheats
}
