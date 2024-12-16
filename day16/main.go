// Advent of Code 2024, Day 16
package main

import (
	"fmt"
	"math"
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
	"github.com/oleiade/lane/v2"
)

// Day 16:
// Part 1 answer: 95444
// Part 2 answer: 513
func main() {
	fmt.Println("Advent of Code 2024, Day 16")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type state struct {
	path []common.Point
	dir  common.Point
}

func (s state) pos() common.Point {
	return s.path[len(s.path)-1]
}

type posAndDir struct {
	pos, dir common.Point
}

func (s state) asPosAndDir() posAndDir {
	return posAndDir{s.pos(), s.dir}
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
	minCost := make(map[posAndDir]int)
	pq := lane.NewMinPriorityQueue[state, int]()
	pq.Push(state{[]common.Point{start}, common.E}, 0)
	for !pq.Empty() {
		curState, score, _ := pq.Pop()
		// Finish state
		if curState.pos() == end {
			return score
		}
		// Find all the possible new states
		// Move forward
		newPath := slices.Clone(curState.path)
		newPath = append(newPath, curState.pos().Add(curState.dir))
		newState := state{newPath, curState.dir}
		newScore := score + 1
		if v, ok := maze.CheckedGet(newState.pos()); ok && v != '#' {
			if possibleNewState(minCost, newState, newScore) {
				pq.Push(newState, newScore)
			}
		}
		// Turn right
		newState = state{curState.path, curState.dir.Right()}
		newScore = score + 1000
		if possibleNewState(minCost, newState, newScore) {
			pq.Push(newState, newScore)
		}
		// Or turn left
		newState = state{curState.path, curState.dir.Left()}
		if possibleNewState(minCost, newState, newScore) {
			pq.Push(newState, newScore)
		}
	}
	panic("failed")
}

func possibleNewState(minCost map[posAndDir]int, s state, score int) bool {
	pad := s.asPosAndDir()
	if v, ok := minCost[pad]; !ok || score < v {
		minCost[pad] = score
		return true
	}
	return false
}

func part2(entries []string) int {
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
	minCost := make(map[posAndDir]int)
	bestPathCost := math.MaxInt
	allBestPathsPoints := mapset.NewThreadUnsafeSet[common.Point]()
	pq := lane.NewMinPriorityQueue[state, int]()
	pq.Push(state{[]common.Point{start}, common.E}, 0)
	for !pq.Empty() {
		curState, score, _ := pq.Pop()
		// Finish state
		if curState.pos() == end {
			if score <= bestPathCost {
				allBestPathsPoints.Append(curState.path...)
				bestPathCost = score
			} else {
				// We're done
				return allBestPathsPoints.Cardinality()
			}
		}
		// Find all the possible new states
		// Move forward
		newPath := slices.Clone(curState.path)
		newPath = append(newPath, curState.pos().Add(curState.dir))
		newState := state{newPath, curState.dir}
		newScore := score + 1
		if v, ok := maze.CheckedGet(newState.pos()); ok && v != '#' {
			if possibleNewState2(minCost, newState, newScore) {
				pq.Push(newState, newScore)
			}
		}
		// Turn right
		newState = state{curState.path, curState.dir.Right()}
		newScore = score + 1000
		if possibleNewState2(minCost, newState, newScore) {
			pq.Push(newState, newScore)
		}
		// Or turn left
		newState = state{curState.path, curState.dir.Left()}
		if possibleNewState2(minCost, newState, newScore) {
			pq.Push(newState, newScore)
		}
	}
	panic("failed")
}

func possibleNewState2(minCost map[posAndDir]int, s state, score int) bool {
	pad := s.asPosAndDir()
	if v, ok := minCost[pad]; !ok || score <= v {
		minCost[pad] = score
		return true
	}
	return false
}
