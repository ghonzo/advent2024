// Advent of Code 2024, Day 18
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
	"github.com/oleiade/lane/v2"
)

// Day 18: RAM Run
// Part 1 answer: 268
// Part 2 answer: 64,11
func main() {
	fmt.Println("Advent of Code 2024, Day 18")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries, 71, 1024))
	fmt.Printf("Part 2: %s\n", part2(entries, 71, 1024))
}

func part1(entries []string, gridSize, numRows int) int {
	grid := common.NewArraysGrid(gridSize, gridSize)
	exit := common.NewPoint(gridSize-1, gridSize-1)
	for _, line := range entries[:numRows] {
		values := common.ConvertToInts(line)
		grid.Set(common.NewPoint(values[0], values[1]), '#')
	}
	return findMinSteps(grid, exit)
}

// If -1, no path
func findMinSteps(grid common.Grid, exit common.Point) int {
	pq := lane.NewMinPriorityQueue[common.Point, int]()
	pq.Push(common.NewPoint(0, 0), 0)
	minScore := make(map[common.Point]int)
	for !pq.Empty() {
		pos, steps, _ := pq.Pop()
		if pos == exit {
			return steps
		}
		if curMin, ok := minScore[pos]; ok && steps >= curMin {
			// dead brach
			continue
		}
		minScore[pos] = steps
		// Now figure out the next moves
		for np := range pos.SurroundingCardinals() {
			if v, ok := grid.CheckedGet(np); ok && v != '#' {
				pq.Push(np, steps+1)
			}
		}
	}
	return -1
}

func part2(entries []string, gridSize, numRows int) string {
	grid := common.NewArraysGrid(gridSize, gridSize)
	exit := common.NewPoint(gridSize-1, gridSize-1)
	for _, line := range entries[:numRows] {
		values := common.ConvertToInts(line)
		grid.Set(common.NewPoint(values[0], values[1]), '#')
	}
	// Now add new points one at a time
	for _, line := range entries[numRows:] {
		values := common.ConvertToInts(line)
		grid.Set(common.NewPoint(values[0], values[1]), '#')
		if findMinSteps(grid, exit) < 0 {
			return fmt.Sprintf("%d,%d", values[0], values[1])
		}
	}
	panic("no solution")
}
