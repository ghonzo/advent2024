// Advent of Code 2024, Day 13
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
	"github.com/oleiade/lane/v2"
)

// Day 13:
// Part 1 answer:
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 13")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	//fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	for i := 0; i < len(entries); i += 4 {
		a := parsePoint(entries[i])
		b := parsePoint(entries[i+1])
		prize := parsePoint(entries[i+2])
		total += minTokens(a, b, prize)
	}
	return total
}

func parsePoint(s string) common.Point {
	values := common.ConvertToInts(s)
	return common.NewPoint(values[0], values[1])
}

type state struct {
	steps int
	pos   common.Point
	cost  int
}

const costA = 3
const costB = 1

func minTokens(a, b, prize common.Point) int {
	fmt.Println(a, b, prize)
	pq := lane.NewMinPriorityQueue[state, int]()
	pq.Push(state{}, 0)
	var minCost int
	for !pq.Empty() {
		s, _, _ := pq.Pop()
		if minCost > 0 && s.cost >= minCost {
			continue
		}
		if s.pos == prize {
			minCost = s.cost
			fmt.Println("New min ", minCost)
			continue
		}
		if s.steps < 100 {
			// Try A
			stateA := state{s.steps + 1, s.pos.Add(a), s.cost + costA}
			if stateA.pos.X() <= prize.X() && stateA.pos.Y() <= prize.Y() {
				pq.Push(stateA, prize.Sub(stateA.pos).ManhattanDistance())
			}
			// And B
			stateB := state{s.steps + 1, s.pos.Add(b), s.cost + costB}
			if stateB.pos.X() <= prize.X() && stateB.pos.Y() <= prize.Y() {
				pq.Push(stateB, prize.Sub(stateB.pos).ManhattanDistance())
			}
		}
	}
	// No solutions
	return 0
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
