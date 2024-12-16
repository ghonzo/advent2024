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

// Day 16: Reindeer Maze
// Part 1 answer: 95444
// Part 2 answer: 513
func main() {
	fmt.Println("Advent of Code 2024, Day 16")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

// This is all the state we need for part 1
type posAndDir struct {
	pos, dir common.Point
}

// This helps us keep track of the minimum score for each position and orientation
type pathFinder struct {
	minScore     map[posAndDir]int
	includeEqual bool
}

func newPathFinder(includeEqual bool) *pathFinder {
	return &pathFinder{make(map[posAndDir]int), includeEqual}
}

func (pf *pathFinder) isPotentialPath(pad posAndDir, score int) bool {
	if v, ok := pf.minScore[pad]; !ok || score < v || (pf.includeEqual && score == v) {
		pf.minScore[pad] = score
		return true
	}
	return false
}

func part1(entries []string) int {
	maze, start, end := readMaze(entries)
	pf := newPathFinder(false)
	pq := lane.NewMinPriorityQueue[posAndDir, int]()
	pq.Push(posAndDir{start, common.E}, 0)
	for !pq.Empty() {
		curPad, score, _ := pq.Pop()
		// Finish state
		if curPad.pos == end {
			return score
		}
		// Find all the possible new states
		// Move forward
		p := curPad.pos.Add(curPad.dir)
		if v, ok := maze.CheckedGet(p); ok && v != '#' {
			newPad := posAndDir{p, curPad.dir}
			newScore := score + 1
			if pf.isPotentialPath(newPad, newScore) {
				pq.Push(newPad, newScore)
			}
		}
		// Turn right
		newPad := posAndDir{curPad.pos, curPad.dir.Right()}
		newScore := score + 1000
		if pf.isPotentialPath(newPad, newScore) {
			pq.Push(newPad, newScore)
		}
		// Turn left
		newPad = posAndDir{curPad.pos, curPad.dir.Left()}
		if pf.isPotentialPath(newPad, newScore) {
			pq.Push(newPad, newScore)
		}
	}
	panic("failed")
}

// Returns the grid, start point, and end point
func readMaze(entries []string) (maze common.Grid, start, end common.Point) {
	maze = common.ArraysGridFromLines(entries)
	for p := range maze.AllPoints() {
		switch maze.Get(p) {
		case 'S':
			start = p
		case 'E':
			end = p
		}
	}
	return
}

// For part2, we need to retain the entire path
type state struct {
	path []common.Point
	dir  common.Point
}

func (s state) pos() common.Point {
	return s.path[len(s.path)-1]
}

func (s state) asPosAndDir() posAndDir {
	return posAndDir{s.pos(), s.dir}
}

func part2(entries []string) int {
	maze, start, end := readMaze(entries)
	pf := newPathFinder(true)
	bestPathCost := math.MaxInt
	// All the points that are part of the optimal paths are in this set
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
		p := curState.pos().Add(curState.dir)
		if v, ok := maze.CheckedGet(p); ok && v != '#' {
			newScore := score + 1
			if pf.isPotentialPath(posAndDir{p, curState.dir}, newScore) {
				newPath := slices.Clone(curState.path)
				newPath = append(newPath, p)
				pq.Push(state{newPath, curState.dir}, newScore)
			}
		}
		// Turn right
		newState := state{curState.path, curState.dir.Right()}
		newScore := score + 1000
		if pf.isPotentialPath(newState.asPosAndDir(), newScore) {
			pq.Push(newState, newScore)
		}
		// Or turn left
		newState = state{curState.path, curState.dir.Left()}
		if pf.isPotentialPath(newState.asPosAndDir(), newScore) {
			pq.Push(newState, newScore)
		}
	}
	panic("failed")
}
