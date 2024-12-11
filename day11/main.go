// Advent of Code 2024, Day 11
package main

import (
	"fmt"
	"strconv"

	"github.com/ghonzo/advent2024/common"
)

// Day 11: Plutonian Pebbles
// Part 1 answer: 182081
// Part 2 answer: 216318908621637
func main() {
	fmt.Println("Advent of Code 2024, Day 11")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	stones := common.ConvertToInts(entries[0])
	for i := 0; i < 25; i++ {
		var newStones []int
		for _, stone := range stones {
			newStones = append(newStones, blink(stone)...)
		}
		stones = newStones
	}
	return len(stones)
}

func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if s := strconv.Itoa(stone); len(s)%2 == 0 {
		return []int{common.Atoi(s[:len(s)/2]), common.Atoi(s[len(s)/2:])}
	} else {
		return []int{stone * 2024}
	}
}

func part2(entries []string) int {
	// Let's use a map of stone -> number instead
	stoneMap := make(map[int]int)
	for _, v := range common.ConvertToInts(entries[0]) {
		stoneMap[v]++
	}
	for i := 0; i < 75; i++ {
		newStoneMap := make(map[int]int)
		for stone, num := range stoneMap {
			for _, nextStone := range blink(stone) {
				newStoneMap[nextStone] += num
			}
		}
		stoneMap = newStoneMap
	}
	var total int
	for _, v := range stoneMap {
		total += v
	}
	return total
}
