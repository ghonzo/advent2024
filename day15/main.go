// Advent of Code 2024, Day 15
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ghonzo/advent2024/common"
)

// Day 15: Warehouse Woes
// Part 1 answer: 1527563
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 15")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var grid common.Grid
	var moves []byte
	for i := 0; ; i++ {
		if len(entries[i]) == 0 {
			grid = common.ArraysGridFromLines(entries[:i])
			moves = []byte(strings.Join(entries[i+1:], ""))
			break
		}
	}
	var robot common.Point
	for p := range grid.AllPoints() {
		if grid.Get(p) == '@' {
			robot = p
			break
		}
	}
	// Now process the moves
	for _, b := range moves {
		dir := convertMove(b)
		robot = moveRobot(grid, robot, dir)
	}
	return scoreGrid(grid)
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
	var grid common.Grid
	var moves []byte
	for i := 0; ; i++ {
		// Expand entries
		entries[i] = expandGrid(entries[i])
		if len(entries[i]) == 0 {
			grid = common.ArraysGridFromLines(entries[:i])
			moves = []byte(strings.Join(entries[i+1:], ""))
			break
		}
	}
	var robot common.Point
	for p := range grid.AllPoints() {
		if grid.Get(p) == '@' {
			robot = p
			break
		}
	}
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
	allMovedPoints := []common.Point{}
	movingPoints := []common.Point{robot}
	for {
		var newMovingPoints []common.Point
		nonspace := false
		for _, mp := range movingPoints {
			nmp := mp.Add(dir)
			v := grid.Get(nmp)
			if v == '#' {
				// Nothing moves
				return robot
			}
			newMovingPoints = append(newMovingPoints, nmp)
			if v == '[' {
				newMovingPoints = append(newMovingPoints, nmp.Add(common.E))
				nonspace = true
			} else if v == ']' {
				newMovingPoints = append(newMovingPoints, nmp.Add(common.W))
				nonspace = true
			}
		}
		movingPoints = newMovingPoints
		allMovedPoints = append(allMovedPoints, movingPoints...)
		if !nonspace {
			break
		}
	}
	// Now we shift everything that moved
	slices.Reverse(allMovedPoints)
	for _, p := range allMovedPoints {
		grid.Set(p, grid.Get(p.Sub(dir)))
	}
	grid.Set(robot, '.')
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