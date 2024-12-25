// Advent of Code 2024, Day 25
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
)

// Day 25: Code Chronicle
// Part 1 answer: 2618
// Part 2 answer: Done!
func main() {
	fmt.Println("Advent of Code 2024, Day 25")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	//fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var locks [][5]int
	var keys [][5]int
	for i := 0; i < len(entries); i += 8 {
		var thing [5]int
		if entries[i] == "#####" {
			for j := 1; j < 6; j++ {
				for k := 0; k < 5; k++ {
					if entries[i+j][k] == '#' {
						thing[k] = j
					}
				}
			}
			locks = append(locks, thing)
		} else {
			for j := 1; j < 6; j++ {
				for k := 0; k < 5; k++ {
					if entries[i+6-j][k] == '#' {
						thing[k] = j
					}
				}
			}
			keys = append(keys, thing)
		}
	}
	var total int
	for _, lock := range locks {
		for _, key := range keys {
			if fit(lock, key) {
				total++
			}
		}
	}
	return total
}

func fit(lock, key [5]int) bool {
	for i := 0; i < 5; i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}
