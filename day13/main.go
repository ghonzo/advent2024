// Advent of Code 2024, Day 13
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
)

// Day 13: Claw Contraption
// Part 1 answer: 31589
// Part 2 answer: 98080815200063
func main() {
	fmt.Println("Advent of Code 2024, Day 13")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	for i := 0; i < len(entries); i += 4 {
		a := parsePoint(entries[i])
		b := parsePoint(entries[i+1])
		prize := parsePoint(entries[i+2])
		// fyi, you can swap this out with part2 solution function "solve" and it also works
		total += search(a, b, prize)
	}
	return total
}

func parsePoint(s string) common.Point {
	values := common.ConvertToInts(s)
	return common.NewPoint(values[0], values[1])
}

const costA = 3
const costB = 1

// Do a brute-force search of the problems space. Return the minimum
// cost or 0 if no solution.
func search(a, b, prize common.Point) int {
	var minCost int
	for aPress := 0; aPress < 100; aPress++ {
		for bPress := 0; bPress < 100; bPress++ {
			if a.Times(aPress).Add(b.Times(bPress)) == prize {
				cost := aPress*costA + bPress*costB
				if minCost == 0 || cost < minCost {
					minCost = cost
				}
			}
		}
	}
	// No solutions
	return minCost
}

func part2(entries []string) int {
	// Brute force isn't going to cut it like part 1.
	// Went down a rabbit hole, looking at Extended Euclidean
	// Algorithm before realizing that this is just a simple system
	// of two equations that we just have to make sure has an integer
	// solution. So Cramer's Rule helped with that.
	var total int
	adjustment := common.NewPoint(10000000000000, 10000000000000)
	for i := 0; i < len(entries); i += 4 {
		a := parsePoint(entries[i])
		b := parsePoint(entries[i+1])
		prize := parsePoint(entries[i+2]).Add(adjustment)
		total += solve(a, b, prize)
	}
	return total
}

// Returns 0 if no solution, otherwise returns the total cost
func solve(a, b, prize common.Point) int {
	// The hardest part of this is keeping the variable names straight!
	// This is Cramer's Rule (https://en.wikipedia.org/wiki/Cramer%27s_rule#Explicit_formulas_for_small_systems)
	denom := a.X()*b.Y() - b.X()*a.Y()
	aPress := (prize.X()*b.Y() - b.X()*prize.Y()) / denom
	bPress := (a.X()*prize.Y() - prize.X()*a.Y()) / denom
	// Now check to see if we ended up truncating to get to integers
	if aPress*a.X()+bPress*b.X() != prize.X() || aPress*a.Y()+bPress*b.Y() != prize.Y() {
		return 0
	}
	return aPress*costA + bPress*costB
}
