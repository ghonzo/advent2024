// Advent of Code 2024, Day 4
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
)

// Day 4: Ceres Search
// Part 1 answer: 2414
// Part 2 answer: 1871
func main() {
	fmt.Println("Advent of Code 2024, Day 4")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	for p := range grid.AllPoints() {
		if grid.Get(p) == 'X' {
			// Now see if there's an M surrounding and follow it
			for sp := range p.SurroundingPoints() {
				if v, _ := grid.CheckedGet(sp); v == 'M' {
					dir := sp.Sub(p)
					sp = sp.Add(dir)
					if v, _ = grid.CheckedGet(sp); v == 'A' {
						sp = sp.Add(dir)
						if v, _ = grid.CheckedGet(sp); v == 'S' {
							total++
						}
					}
				}
			}
		}
	}
	return total
}

func part2(entries []string) int {
	var total int
	grid := common.ArraysGridFromLines(entries)
	for x := 0; x < grid.Size().X()-2; x++ {
		for y := 0; y < grid.Size().Y()-2; y++ {
			if grid.Get(common.NewPoint(x+1, y+1)) == 'A' {
				if ((grid.Get(common.NewPoint(x, y)) == 'M' && grid.Get(common.NewPoint(x+2, y+2)) == 'S') ||
					(grid.Get(common.NewPoint(x, y)) == 'S' && grid.Get(common.NewPoint(x+2, y+2)) == 'M')) &&
					((grid.Get(common.NewPoint(x+2, y)) == 'M' && grid.Get(common.NewPoint(x, y+2)) == 'S') ||
						(grid.Get(common.NewPoint(x+2, y)) == 'S' && grid.Get(common.NewPoint(x, y+2)) == 'M')) {
					total++
				}
			}
		}
	}
	return total
}
