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
	fmt.Printf("Part 2: %d\n", part2(entries, 100))
}

type cheat [2]common.Point

func part1(entries []string, cheatLimit int) int {
	grid, start, end := readGrid(entries)
	path := findPath(grid, start, end)
	// The point on a path pointing to the picosecond
	stepMap := make(map[common.Point]int)
	for i, p := range path {
		stepMap[p] = i
	}
	validCheats := make(map[cheat]bool)
	// Now step along every point in the path and see if there are cheats
	for i, p := range path[:len(path)-cheatLimit-1] {
		// For each wall
		for cheatStart := range p.SurroundingCardinals() {
			if v, _ := grid.CheckedGet(cheatStart); v == '#' {
				// Now see if there are any shortcuts
				for cheatEnd := range cheatStart.SurroundingCardinals() {
					if cheatEndStep, ok := stepMap[cheatEnd]; ok {
						timeSaved := cheatEndStep - i - 2
						if timeSaved >= cheatLimit {
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

func part2(entries []string, cheatLimit int) int {
	grid, start, end := readGrid(entries)
	path := findPath(grid, start, end)
	// The point on a path pointing to the picosecond
	stepMap := make(map[common.Point]int)
	for i, p := range path {
		stepMap[p] = i
	}
	validCheats := make(map[cheat]bool)
	// Now step along every point in the path and see if there are cheats
	for i, p := range path[:len(path)-cheatLimit-1] {
		fmt.Println(i, p, len(validCheats))
		// For each wall
		for cheatStart := range p.SurroundingCardinals() {
			if v, _ := grid.CheckedGet(cheatStart); v == '#' {
				cheatPath := mapset.NewThreadUnsafeSet[common.Point](cheatStart)
				nextPoints := mapset.NewThreadUnsafeSet[common.Point](cheatStart)
				for cheatTime := 1; cheatTime < 20 && !nextPoints.IsEmpty(); cheatTime++ {
					nextNextPoints := mapset.NewThreadUnsafeSet[common.Point]()
					for p, ok := nextPoints.Pop(); ok; {
						cheatPath.Add(p)
						for sc := range p.SurroundingCardinals() {
							v, _ := grid.CheckedGet(sc)
							if v == '#' && !cheatPath.Contains(sc) {
								nextNextPoints.Add(sc)
							} else if cheatEndStep, found := stepMap[sc]; found {
								timeSaved := cheatEndStep - i - cheatTime - 1
								if timeSaved >= cheatLimit {
									validCheats[[2]common.Point{cheatStart, sc}] = true
								}
							}
						}
					}
					nextPoints = nextNextPoints
				}
			}
		}
	}
	return len(validCheats)
}
