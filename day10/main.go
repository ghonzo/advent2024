// Advent of Code 2024, Day 10
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 10:
// Part 1 answer:
// Part 2 answer:
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
			total += findNumTrailheads(grid, p)
		}
	}
	return total
}

func findNumTrailheads(grid common.Grid, trailhead common.Point) int {
	pointSet := mapset.NewThreadUnsafeSet[common.Point]()
	pointSet.Add(trailhead)
	for height := '1'; height <= '9'; height++ {
		newPointSet := mapset.NewThreadUnsafeSet[common.Point]()
		for p := range pointSet.Iter() {
			newPointSet.Append(findSurroundingPoints(grid, p, byte(height))...)
		}
		pointSet = newPointSet
	}
	return pointSet.Cardinality()
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

func part2(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	for p := range grid.AllPoints() {
		if grid.Get(p) == '0' {
			total += findNumTrailheads2(grid, p)
		}
	}
	return total
}

func findNumTrailheads2(grid common.Grid, trailhead common.Point) int {
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
	var trailheads int
	for _, v := range pointMap {
		trailheads += v
	}
	return trailheads
}
