// Advent of Code 2024, Day 10
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
)

// Day 10: Hoof It
// Part 1 answer: 820
// Part 2 answer: 1786
func main() {
	fmt.Println("Advent of Code 2024, Day 10")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	for p := range grid.AllPoints() {
		if grid.Get(p) == '0' {
			// We just care about the number of trailends
			total += len(findTrailends(grid, p))
		}
	}
	return total
}

func part2(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	for p := range grid.AllPoints() {
		if grid.Get(p) == '0' {
			// In this case, we just care about the number of paths
			for _, v := range findTrailends(grid, p) {
				total += v
			}
		}
	}
	return total
}

// Returns the points that are valid ends for this trail, and the number of ways you can get there
func findTrailends(grid common.Grid, trailhead common.Point) map[common.Point]int {
	pointMap := make(map[common.Point]int)
	pointMap[trailhead] = 1
	for height := '1'; height <= '9'; height++ {
		newPointMap := make(map[common.Point]int)
		for p, num := range pointMap {
			for _, sp := range findSurroundingPoints(grid, p, byte(height)) {
				newPointMap[sp] += num
			}
		}
		pointMap = newPointMap
	}
	return pointMap
}

func findSurroundingPoints(grid common.Grid, p common.Point, v byte) []common.Point {
	var points []common.Point
	for sp := range p.SurroundingCardinals() {
		if spv, _ := grid.CheckedGet(sp); spv == v {
			points = append(points, sp)
		}
	}
	return points
}
