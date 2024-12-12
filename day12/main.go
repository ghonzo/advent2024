// Advent of Code 2024, Day 12
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 12: Garden Groups
// Part 1 answer: 1549354
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 12")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	//fmt.Printf("Part 2: %d\n", part2(entries))
}

type region struct {
	points    mapset.Set[common.Point]
	perimeter int
}

func part1(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	visited := mapset.NewThreadUnsafeSet[common.Point]()
	for p := range grid.AllPoints() {
		if visited.Contains(p) {
			continue
		}
		r := findRegion(grid, p)
		total += r.points.Cardinality() * r.perimeter
		visited = visited.Union(r.points)
	}
	return total
}

func findRegion(grid common.Grid, start common.Point) region {
	var r region
	plotType := grid.Get(start)
	r.points = mapset.NewThreadUnsafeSet[common.Point]()
	pointsToVisit := mapset.NewThreadUnsafeSet[common.Point](start)
	for !pointsToVisit.IsEmpty() {
		// Check each of the surrounding points
		currentPt, _ := pointsToVisit.Pop()
		r.points.Add(currentPt)
		for p := range currentPt.SurroundingCardinals() {
			if r.points.ContainsOne(p) {
				continue
			}
			if v, _ := grid.CheckedGet(p); v != plotType {
				r.perimeter++
			} else {
				pointsToVisit.Add(p)
			}
		}
	}
	return r
}

func part2(entries []string) int {
	var total int
	left := make([]int, len(entries))
	rightMap := make(map[int]int)
	for i, line := range entries {
		values := common.ConvertToInts(line)
		left[i] = values[0]
		rightMap[values[1]]++
	}
	for _, l := range left {
		total += l * rightMap[l]
	}
	return total
}
