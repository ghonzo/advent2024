// Advent of Code 2024, Day 24
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2024/common"
)

// Day 24:
// Part 1 answer: 60614602965288
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 24")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	//fmt.Printf("Part 2: %d\n", part2(entries))
}

type rule struct {
	a, gate, b, c string
}

func (r rule) eval(registers map[string]bool) bool {
	var aVal, bVal bool
	var ok bool
	if aVal, ok = registers[r.a]; !ok {
		return false
	}
	if bVal, ok = registers[r.b]; !ok {
		return false
	}
	switch r.gate {
	case "AND":
		registers[r.c] = aVal && bVal
	case "OR":
		registers[r.c] = aVal || bVal
	case "XOR":
		registers[r.c] = (aVal != bVal)
	}
	return true
}

func part1(entries []string) uint64 {
	registers := make(map[string]bool)
	var rules []rule
	var linenum int
	var line string
	for linenum, line = range entries {
		if len(line) == 0 {
			break
		}
		registers[line[0:3]] = (line[5] == '1')
	}
	// Now read rules
	for _, line = range entries[linenum+1:] {
		parts := strings.Split(line, " ")
		rules = append(rules, rule{parts[0], parts[1], parts[2], parts[4]})
	}
outer:
	for len(rules) > 0 {
		for i, r := range rules {
			if r.eval(registers) {
				rules[i] = rules[len(rules)-1]
				rules = rules[:len(rules)-1]
				continue outer
			}
		}
	}
	var total uint64
	for k, v := range registers {
		if v && k[0] == 'z' {
			total += (1 << common.Atoi(k[1:3]))
		}
	}
	return total
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
