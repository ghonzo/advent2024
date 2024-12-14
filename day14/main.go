// Advent of Code 2024, Day 14
package main

import (
	"fmt"
	"regexp"

	"github.com/ghonzo/advent2024/common"
)

// Day 14: Restroom Redoubt
// Part 1 answer: 229632480
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 14")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries, 101, 103))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

var robotRegexp = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

type robot struct {
	p, v common.Point
}

func part1(entries []string, width, height int) int {
	var nw, ne, se, sw int
	for _, line := range entries {
		group := robotRegexp.FindStringSubmatch(line)
		r := robot{common.NewPoint(common.Atoi(group[1]), common.Atoi(group[2])), common.NewPoint(common.Atoi(group[3]), common.Atoi(group[4]))}
		// find final pos
		pos := r.p.Add(r.v.Times(100))
		// find the quadrant
		x, y := common.Mod(pos.X(), width), common.Mod(pos.Y(), height)
		if x < (width)/2 && y < (height)/2 {
			nw++
		} else if x < (width)/2 && y > (height)/2 {
			sw++
		} else if x > (width)/2 && y < (height)/2 {
			ne++
		} else if x > (width)/2 && y > (height)/2 {
			se++
		}
	}
	return nw * sw * ne * se
}

func part2(entries []string) int {
	bounds := common.NewPoint(101, 103)
	halfWidth := bounds.X() / 2
	var robots []robot
	for _, line := range entries {
		group := robotRegexp.FindStringSubmatch(line)
		robots = append(robots, robot{common.NewPoint(common.Atoi(group[1]), common.Atoi(group[2])), common.NewPoint(common.Atoi(group[3]), common.Atoi(group[4]))})
	}
	// Keep looping until they are all contiguous and not overlapping
	var step int
outer:
	for {
		step++
		locMap := make(map[common.Point]int)
		// Update robot positions
		for _, r := range robots {
			r.p = pointMod(r.p.Add(r.v), bounds)
			if _, ok := locMap[r.p]; ok {
				continue outer
			} else {
				locMap[r.p]++
			}
		}
		// Now that they are all there, confirm that there's another one symmetrical across middle
		for p := range locMap {
			diff := p.X() - halfWidth
			if _, ok := locMap[common.NewPoint(halfWidth-diff, p.Y())]; !ok {
				continue outer
			}
		}
		// We might have it
		grid := common.NewArraysGrid(101, 103)
		for p := range locMap {
			grid.Set(p, '*')
		}
		fmt.Print(common.RenderGrid(grid, '.'))
		break
	}
	return step
}

func pointMod(p, bounds common.Point) common.Point {
	return common.NewPoint(common.Mod(p.X(), bounds.X()), common.Mod(p.Y(), bounds.Y()))
}
