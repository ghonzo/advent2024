// Advent of Code 2024, Day 14
package main

import (
	"fmt"
	"regexp"

	"github.com/ghonzo/advent2024/common"
)

// Day 14: Restroom Redoubt
// Part 1 answer: 229632480
// Part 2 answer: 7051

// The real size of the space for the robots
var spaceDim = common.NewPoint(101, 103)

func main() {
	fmt.Println("Advent of Code 2024, Day 14")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries, spaceDim))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type robot struct {
	p, v common.Point
}

func part1(entries []string, bounds common.Point) int {
	halfBounds := common.NewPoint(bounds.X()/2, bounds.Y()/2)
	quadrantCount := make(map[common.Point]int)
	for _, r := range readRobots(entries) {
		// find final pos
		pos := r.p.Add(r.v.Times(100))
		// mod it to the bounds of the space
		pos = pointMod(pos, bounds)
		// Find the quadrant
		pos = pos.Sub(halfBounds)
		quadrant := common.NewPoint(common.Sgn(pos.X()), common.Sgn(pos.Y()))
		quadrantCount[quadrant]++
	}
	return quadrantCount[common.NW] * quadrantCount[common.NE] * quadrantCount[common.SW] * quadrantCount[common.SE]
}

var robotRegexp = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func readRobots(lines []string) []*robot {
	robots := make([]*robot, len(lines))
	for i, line := range lines {
		group := robotRegexp.FindStringSubmatch(line)
		robots[i] = &robot{common.NewPoint(common.Atoi(group[1]), common.Atoi(group[2])),
			common.NewPoint(common.Atoi(group[3]), common.Atoi(group[4]))}
	}
	return robots
}

func part2(entries []string) int {
	robots := readRobots(entries)
	// Keep looping until they are not overlapping
	var step int
	for step = 1; ; step++ {
		// Number of robots at a given point
		locMap := make(map[common.Point]int)
		// Set to true if more than one robot at a point
		overlap := false
		// Update robot positions
		for _, r := range robots {
			r.p = pointMod(r.p.Add(r.v), spaceDim)
			locMap[r.p]++
			if locMap[r.p] > 1 {
				overlap = true
			}
		}
		if !overlap {
			break
		}
	}
	// Let's have a look!
	grid := common.NewSparseGrid()
	for _, r := range robots {
		grid.Set(r.p, '*')
	}
	fmt.Print(common.RenderGrid(grid, '.'))
	return step
}

func pointMod(p, bounds common.Point) common.Point {
	return common.NewPoint(common.Mod(p.X(), bounds.X()), common.Mod(p.Y(), bounds.Y()))
}
