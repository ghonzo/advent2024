// Advent of Code 2024, Day 16
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
	"github.com/oleiade/lane/v2"
)

// Day 16:
// Part 1 answer: 95444
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 16")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	//fmt.Printf("Part 2: %d\n", part2(entries))
}

type state struct {
	pos, dir common.Point
}

func part1(entries []string) int {
	maze := common.ArraysGridFromLines(entries)
	var start, end common.Point
	for p := range maze.AllPoints() {
		switch maze.Get(p) {
		case 'S':
			start = p
		case 'E':
			end = p
		}
	}
	minCost := make(map[state]int)
	pq := lane.NewMinPriorityQueue[state, int]()
	pq.Push(state{start, common.E}, 0)
	for !pq.Empty() {
		curState, score, _ := pq.Pop()
		// Finish state
		if curState.pos == end {
			return score
		}
		// Find all the possible new states
		// Move forward
		newState := state{curState.pos.Add(curState.dir), curState.dir}
		newScore := score + 1
		if v, ok := maze.CheckedGet(newState.pos); ok && v != '#' {
			if possibleNewState(minCost, newState, newScore) {
				pq.Push(newState, newScore)
			}
		}
		// Turn right
		newState = state{curState.pos, curState.dir.Right()}
		newScore = score + 1000
		if possibleNewState(minCost, newState, newScore) {
			pq.Push(newState, newScore)
		}
		// Or turn left
		newState = state{curState.pos, curState.dir.Left()}
		if possibleNewState(minCost, newState, newScore) {
			pq.Push(newState, newScore)
		}
	}
	panic("failed")
}

func possibleNewState(minCost map[state]int, s state, score int) bool {
	if v, ok := minCost[s]; !ok || score < v {
		minCost[s] = score
		return true
	}
	return false
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
