// Advent of Code 2024, Day 5
package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ghonzo/advent2024/common"
)

// Day 5: Print Queue
// Part 1 answer: 4135
// Part 2 answer: 5285
func main() {
	fmt.Println("Advent of Code 2024, Day 5")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	// Order map ... values must be not be after key
	orderMap := make(map[int]mapset.Set[int])
	var i int
	var line string
	for i, line = range entries {
		if len(line) == 0 {
			break
		}
		rule := common.ConvertToInts(line)
		s, ok := orderMap[rule[1]]
		if !ok {
			s = mapset.NewSet[int]()
			orderMap[rule[1]] = s
		}
		s.Add(rule[0])
	}
	// Now read each update
outer:
	for _, line = range entries[i+1:] {
		update := common.ConvertToInts(line)
		for n, pageNum := range update[:len(update)-1] {
			if vals, ok := orderMap[pageNum]; ok && vals.ContainsAny(update[n+1:]...) {
				continue outer
			}
		}
		total += update[len(update)/2]
	}
	return total
}

func insertInt(array []int, value int, index int) []int {
	return append(array[:index], append([]int{value}, array[index:]...)...)
}

func removeInt(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func moveInt(array []int, srcIndex int, dstIndex int) []int {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

func part2(entries []string) int {
	var total int
	// Order map ... values must be not be after key
	orderMap := make(map[int]mapset.Set[int])
	var i int
	var line string
	for i, line = range entries {
		if len(line) == 0 {
			break
		}
		rule := common.ConvertToInts(line)
		s, ok := orderMap[rule[1]]
		if !ok {
			s = mapset.NewSet[int]()
			orderMap[rule[1]] = s
		}
		s.Add(rule[0])
	}
	// Now read each update
	for _, line = range entries[i+1:] {
		update := common.ConvertToInts(line)
		swapped := false
	inner:
		for {
			for left, pageNum := range update[:len(update)-1] {
				for right, pageNum2 := range update[left+1:] {
					if vals, ok := orderMap[pageNum]; ok && vals.Contains(pageNum2) {
						swapped = true
						update = moveInt(update, right+left+1, left)
						continue inner
					}
				}
			}
			if swapped {
				total += update[len(update)/2]
			}
			break
		}
	}
	return total
}
