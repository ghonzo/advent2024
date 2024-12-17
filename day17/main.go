// Advent of Code 2024, Day 17
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2024/common"
)

// Day 17:
// Part 1 answer: 1,0,2,0,5,7,2,1,3
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2024, Day 17")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %s\n", part1(entries))
	//fmt.Printf("Part 2: %d\n", part2(entries))
}

type registerStore [3]int

func (r registerStore) combo(operand int) int {
	if operand < 4 {
		return operand
	}
	return r[operand-4]
}

func part1(entries []string) string {
	registers := registerStore{common.ConvertToInts(entries[0])[0], common.ConvertToInts(entries[1])[0], common.ConvertToInts(entries[2])[0]}
	program := common.ConvertToInts(entries[4])
	var ip int
	var output []int
	for ip < len(program) {
		opcode := program[ip]
		operand := program[ip+1]
		ip += 2
		switch opcode {
		case 0:
			registers[0] = registers[0] / (1 << registers.combo(operand))
		case 1:
			registers[1] = registers[1] ^ operand
		case 2:
			registers[1] = registers.combo(operand) % 8
		case 3:
			if registers[0] != 0 {
				ip = operand
			}
		case 4:
			registers[1] = registers[1] ^ registers[2]
		case 5:
			output = append(output, registers.combo(operand)%8)
		case 6:
			registers[1] = registers[0] / (1 << registers.combo(operand))
		case 7:
			registers[2] = registers[0] / (1 << registers.combo(operand))
		}
	}
	return intsToString(output)
}

func intsToString(ints []int) string {
	s := make([]string, len(ints))
	for i, v := range ints {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
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
