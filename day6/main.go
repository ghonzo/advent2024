// Advent of Code 2024, Day 6
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 6: Guard Gallivant
// Part 1 answer: 5080
// Part 2 answer: 1919
func main() {
	fmt.Println("Advent of Code 2024, Day 6")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	grid := common.ArraysGridFromLines(entries)
	startPt := findStart(grid)
	return len(findAllVisitedPoints(grid, startPt))
}

func findStart(g common.Grid) common.Point {
	for p := range g.AllPoints() {
		if g.Get(p) == '^' {
			return p
		}
	}
	panic("no start")
}

func findAllVisitedPoints(g common.Grid, startPt common.Point) []common.Point {
	visited := mapset.NewThreadUnsafeSet[common.Point]()
	visited.Add(startPt)
	for p, dir := startPt, common.N; ; p = p.Add(dir) {
		v, ok := g.CheckedGet(p)
		if !ok {
			// escaped
			break
		}
		if v == '#' {
			// undo the movement
			p = p.Sub(dir)
			// and turn right
			dir = dir.Right()
		} else {
			visited.Add(p)
		}
	}
	return visited.ToSlice()
}

func part2(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	startPt := findStart(grid)
	// obstacle must be placed on a location we would have visited
	for _, p := range findAllVisitedPoints(grid, startPt) {
		if grid.Get(p) == '.' {
			grid.Set(p, '#')
			if stuckInLoop(grid, startPt) {
				total++
			}
			grid.Set(p, '.')
		}
	}
	return total
}

type posAndDir struct {
	pos common.Point
	dir common.Point
}

func stuckInLoop(grid common.Grid, startPt common.Point) bool {
	visited := mapset.NewThreadUnsafeSet[posAndDir]()
	for pad := (posAndDir{pos: startPt, dir: common.N}); ; pad.pos = pad.pos.Add(pad.dir) {
		v, ok := grid.CheckedGet(pad.pos)
		if !ok {
			// escaped
			return false
		}
		if v == '#' {
			// undo the movement
			pad.pos = pad.pos.Sub(pad.dir)
			// turn right
			pad.dir = pad.dir.Right()
		} else {
			if !visited.Add(pad) {
				// not added, so loop detected
				return true
			}
		}
	}
}
