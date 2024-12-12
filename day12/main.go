// Advent of Code 2024, Day 12
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 12: Garden Groups
// Part 1 answer: 1549354
// Part 2 answer: 937032
func main() {
	fmt.Println("Advent of Code 2024, Day 12")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type fence struct {
	pos, dir common.Point
}

type region struct {
	points mapset.Set[common.Point]
	fence  mapset.Set[fence]
}

func (r region) area() int {
	return r.points.Cardinality()
}

func (r region) perimeter() int {
	return r.fence.Cardinality()
}

func (r region) sides() int {
	var sides int
	fencesToCount := r.fence.Clone()
	for !fencesToCount.IsEmpty() {
		f, _ := fencesToCount.Pop()
		sides++
		if f.dir == common.N || f.dir == common.S {
			// travel east and west
			removeFences(fencesToCount, f, common.E)
			removeFences(fencesToCount, f, common.W)
		} else {
			removeFences(fencesToCount, f, common.N)
			removeFences(fencesToCount, f, common.S)
		}
	}
	return sides
}

func removeFences(fences mapset.Set[fence], f fence, dir common.Point) {
	for {
		// f is a copy we can modify it
		f.pos = f.pos.Add(dir)
		if fences.Contains(f) {
			fences.Remove(f)
		} else {
			return
		}
	}
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
		total += r.area() * r.perimeter()
		visited = visited.Union(r.points)
	}
	return total
}

func findRegion(grid common.Grid, start common.Point) region {
	var r region
	plotType := grid.Get(start)
	r.points = mapset.NewThreadUnsafeSet[common.Point]()
	r.fence = mapset.NewThreadUnsafeSet[fence]()
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
				r.fence.Add(fence{currentPt, p.Sub(currentPt)})
			} else {
				pointsToVisit.Add(p)
			}
		}
	}
	return r
}

func part2(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	visited := mapset.NewThreadUnsafeSet[common.Point]()
	for p := range grid.AllPoints() {
		if visited.Contains(p) {
			continue
		}
		r := findRegion(grid, p)
		total += r.area() * r.sides()
		visited = visited.Union(r.points)
	}
	return total
}
