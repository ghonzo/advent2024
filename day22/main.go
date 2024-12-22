// Advent of Code 2024, Day 22
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
)

// Day 22: Monkey Market
// Part 1 answer: 17960270302
// Part 2 answer: 2042
func main() {
	fmt.Println("Advent of Code 2024, Day 22")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var total int
	for _, line := range entries {
		n := common.Atoi(line)
		for i := 0; i < 2000; i++ {
			n = nextSecret(n)
		}
		total += n
	}
	return total
}

func nextSecret(n int) int {
	n = ((64 * n) ^ n) % 16777216
	n = ((n / 32) ^ n) % 16777216
	n = ((n * 2048) ^ n) % 16777216
	return n
}

type diffHistory [4]int

func (dh diffHistory) addDiff(d int) diffHistory {
	return diffHistory{dh[1], dh[2], dh[3], d}
}

func part2(entries []string) int {
	totalMap := make(map[diffHistory]int)
	for _, line := range entries {
		n := common.Atoi(line)
		var dh diffHistory
		found := make(map[diffHistory]bool)
		for i := 0; i < 2000; i++ {
			lastPrice := n % 10
			n = nextSecret(n)
			price := n % 10
			dh = dh.addDiff(price - lastPrice)
			if i > 2 {
				if !found[dh] {
					found[dh] = true
					totalMap[dh] += price
				}
			}
		}
	}
	var maxBananas int
	for _, v := range totalMap {
		maxBananas = max(maxBananas, v)
	}
	return maxBananas
}
