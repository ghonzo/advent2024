// Advent of Code 2024, Day 6
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 6:
// Part 1 answer: 5080
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 6")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	visited := mapset.NewSet[common.Point]()
	grid := common.ArraysGridFromLines(entries)
	dir := common.N
	startPt := findStart(grid)
	visited.Add(startPt)
	for p := startPt; ; p = p.Add(dir) {
		v, ok := grid.CheckedGet(p)
		if !ok {
			break
		}
		if v == '#' {
			// undo the movement
			p = p.Sub(dir)
			// turn right
			switch dir {
			case common.N:
				dir = common.E
			case common.E:
				dir = common.S
			case common.S:
				dir = common.W
			case common.W:
				dir = common.N
			}
		} else {
			visited.Add(p)
		}
	}
	return visited.Cardinality()
}

func findStart(g common.Grid) common.Point {
	for p := range g.AllPoints() {
		if g.Get(p) == '^' {
			return p
		}
	}
	panic("no start")
}

func part2(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	startPt := findStart(grid)
	for p := range grid.AllPoints() {
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
	visited := mapset.NewSet[posAndDir]()
	dir := common.N
	for p := startPt; ; p = p.Add(dir) {
		v, ok := grid.CheckedGet(p)
		if !ok {
			// escaped
			return false
		}
		if v == '#' {
			// undo the movement
			p = p.Sub(dir)
			// turn right
			switch dir {
			case common.N:
				dir = common.E
			case common.E:
				dir = common.S
			case common.S:
				dir = common.W
			case common.W:
				dir = common.N
			}
		} else {
			pad := posAndDir{pos: p, dir: dir}
			if visited.Contains(pad) {
				// loop
				return true
			}
			visited.Add(pad)
		}
	}
}
