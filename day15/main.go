// Advent of Code 2024, Day 15
package main

import (
	"fmt"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 15: Warehouse Woes
// Part 1 answer: 1527563
// Part 2 answer: 1521635
func main() {
	fmt.Println("Advent of Code 2024, Day 15")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	grid, moves := readGridAndMoves(entries)
	robot := findRobot(grid)
	// Now process the moves
	for _, b := range moves {
		dir := convertMove(b)
		robot = moveRobot(grid, robot, dir)
	}
	return scoreGrid(grid)
}

func readGridAndMoves(entries []string) (common.Grid, []byte) {
	for i := 0; ; i++ {
		if len(entries[i]) == 0 {
			return common.ArraysGridFromLines(entries[:i]), []byte(strings.Join(entries[i+1:], ""))
		}
	}
}

func findRobot(grid common.Grid) common.Point {
	for p := range grid.AllPoints() {
		if grid.Get(p) == '@' {
			return p
		}
	}
	panic("no robot")
}

func convertMove(b byte) common.Point {
	switch b {
	case '^':
		return common.N
	case '<':
		return common.W
	case '>':
		return common.E
	case 'v':
		return common.S
	}
	panic("invalid move")
}

func moveRobot(grid common.Grid, robot common.Point, dir common.Point) common.Point {
	// Keep going in the direction until we get a wall or blank space
	p := robot
	for {
		p = p.Add(dir)
		v := grid.Get(p)
		if v == '#' {
			// nothing happens
			return robot
		}
		if v == '.' {
			// Move to next part
			break
		}
	}
	// Now move everything between p and robot
	for p != robot {
		grid.Set(p, grid.Get(p.Sub(dir)))
		p = p.Sub(dir)
	}
	grid.Set(robot, '.')
	return robot.Add(dir)
}

func scoreGrid(grid common.Grid) int {
	var total int
	for p := range grid.AllPoints() {
		if grid.Get(p) == 'O' {
			total += p.X() + 100*p.Y()
		}
	}
	return total
}

func part2(entries []string) int {
	// First, expand the entries
	for i := 0; len(entries[i]) > 0; i++ {
		entries[i] = expandGrid(entries[i])
	}
	grid, moves := readGridAndMoves(entries)
	robot := findRobot(grid)
	// Now process the moves
	for _, b := range moves {
		dir := convertMove(b)
		robot = moveRobot2(grid, robot, dir)
	}
	return scoreGrid2(grid)
}

var replacer = strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")

func expandGrid(s string) string {
	return replacer.Replace(s)
}

func moveRobot2(grid common.Grid, robot common.Point, dir common.Point) common.Point {
	// East and west is the same as before
	if dir == common.E || dir == common.W {
		return moveRobot(grid, robot, dir)
	}
	// North and south are more complicated
	var movingPoints []common.Point
	leadingEdge := mapset.NewThreadUnsafeSet[common.Point](robot)
	for !leadingEdge.IsEmpty() {
		newLeadingEdge := mapset.NewThreadUnsafeSet[common.Point]()
		for p := range leadingEdge.Iter() {
			movingPoints = append(movingPoints, p)
			p = p.Add(dir)
			v := grid.Get(p)
			if v == '#' {
				// Nothing can move
				return robot
			}
			if v == '[' {
				newLeadingEdge.Add(p)
				newLeadingEdge.Add(p.Add(common.E))
			} else if v == ']' {
				newLeadingEdge.Add(p)
				newLeadingEdge.Add(p.Add(common.W))
			}
		}
		leadingEdge = newLeadingEdge
	}
	// Cool, so we can move all the points in movingPoints (in reverse order)
	slices.Reverse(movingPoints)
	for _, p := range movingPoints {
		grid.Set(p.Add(dir), grid.Get(p))
		grid.Set(p, '.')
	}
	return robot.Add(dir)
}

func scoreGrid2(grid common.Grid) int {
	var total int
	for p := range grid.AllPoints() {
		if grid.Get(p) == '[' {
			total += p.X() + 100*p.Y()
		}
	}
	return total
}
