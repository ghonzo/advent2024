// Advent of Code 2024, Day 3
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghonzo/advent2024/common"
)

// Day 3: Mull It Over
// Part 1 answer: 157621318
// Part 2 answer: 79845780
func main() {
	fmt.Println("Advent of Code 2024, Day 3")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

var mulRegexp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func part1(entries []string) int {
	var total int
	for _, line := range entries {
		for _, match := range mulRegexp.FindAllStringSubmatch(line, -1) {
			total += common.Atoi(match[1]) * common.Atoi(match[2])
		}
	}
	return total
}

func part2(entries []string) int {
	enabled := true
	var enabledEntries []string
	for _, line := range entries {
		for len(line) > 0 {
			if enabled {
				nextDont := strings.Index(line, "don't()")
				if nextDont == -1 {
					enabledEntries = append(enabledEntries, line)
					line = ""
				} else {
					enabledEntries = append(enabledEntries, line[:nextDont])
					line = line[nextDont+7:]
					enabled = false
				}
			}
			if !enabled {
				nextDo := strings.Index(line, "do()")
				if nextDo == -1 {
					line = ""
					continue
				} else {
					line = line[nextDo+4:]
					enabled = true
				}
			}
		}
	}
	return part1(enabledEntries)
}
