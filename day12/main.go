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

// A fence is on a particular plot (pos) on a particular side (dir)
type fence struct {
	pos, dir common.Point
}

type region struct {
	points mapset.Set[common.Point]
	fences mapset.Set[fence]
}

func (r region) area() int {
	return r.points.Cardinality()
}

func (r region) perimeter() int {
	return r.fences.Cardinality()
}

func (r region) sides() int {
	var sides int
	fencesToCount := r.fences.Clone()
	for !fencesToCount.IsEmpty() {
		f, _ := fencesToCount.Pop()
		sides++
		if f.dir == common.N || f.dir == common.S {
			// travel east and west, removing contiguous fences
			removeFences(fencesToCount, f, common.E)
			removeFences(fencesToCount, f, common.W)
		} else {
			// travel north and south, removing contiguous fences
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
		// if the fence is on the same side of the plot, remove it
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
		visited.Append(r.points.ToSlice()...)
	}
	return total
}

func findRegion(grid common.Grid, start common.Point) region {
	var r region
	r.points = mapset.NewThreadUnsafeSet[common.Point]()
	r.fences = mapset.NewThreadUnsafeSet[fence]()
	pointsToVisit := mapset.NewThreadUnsafeSet[common.Point](start)
	plotType := grid.Get(start)
	for !pointsToVisit.IsEmpty() {
		// Pick a plot ... any plot
		currentPt, _ := pointsToVisit.Pop()
		r.points.Add(currentPt)
		// Check each of the surrounding points
		for p := range currentPt.SurroundingCardinals() {
			if r.points.ContainsOne(p) {
				// Already visited that one
				continue
			}
			if v, _ := grid.CheckedGet(p); v != plotType {
				// Different type (or out of bounds), so add a fence
				r.fences.Add(fence{currentPt, p.Sub(currentPt)})
			} else {
				// Same type, so add it to the region
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
		visited.Append(r.points.ToSlice()...)
	}
	return total
}
