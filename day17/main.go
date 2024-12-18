// Advent of Code 2024, Day 17
package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ghonzo/advent2024/common"
)

// Day 17: Chronospatial Computer
// Part 1 answer: 1,0,2,0,5,7,2,1,3
// Part 2 answer: 265652340990875
func main() {
	fmt.Println("Advent of Code 2024, Day 17")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %s\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type registerStore [3]uint64

func (r registerStore) combo(operand int) uint64 {
	if operand < 4 {
		return uint64(operand)
	}
	return r[operand-4]
}

func part1(entries []string) string {
	// Every example has b and c as 0, so we just care about a
	a := common.ConvertToInts(entries[0])[0]
	program := common.ConvertToInts(entries[4])
	return intsToString(runProgram(uint64(a), program))
}

func runProgram(a uint64, program []int) []int {
	registers := registerStore{a, 0, 0}
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
			registers[1] = registers[1] ^ uint64(operand)
		case 2:
			registers[1] = registers.combo(operand) % 8
		case 3:
			if registers[0] != 0 {
				ip = operand
			}
		case 4:
			registers[1] = registers[1] ^ registers[2]
		case 5:
			output = append(output, int(registers.combo(operand)%8))
		case 6:
			registers[1] = registers[0] / (1 << registers.combo(operand))
		case 7:
			registers[2] = registers[0] / (1 << registers.combo(operand))
		}
	}
	return output
}

func intsToString(ints []int) string {
	s := make([]string, len(ints))
	for i, v := range ints {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}

// Got this by looking at the output ... so it might be specific to my input
func part2(entries []string) uint64 {
	program := common.ConvertToInts(entries[4])
	candidateSolutions := []uint64{0}
	// Think in octal! We need to find the digits in reverse order
	for i := len(program) - 1; i >= 0; i-- {
		shift := uint64(1) << (i * 3)
		var newCandidateSolutions []uint64
		for _, a := range candidateSolutions {
			for digit := 0; digit < 8; digit++ {
				output := runProgram(a, program)
				if len(output) == len(program) && slices.Equal(output[i:], program[i:]) {
					newCandidateSolutions = append(newCandidateSolutions, a)
				}
				a += shift
			}
		}
		candidateSolutions = newCandidateSolutions
	}
	return slices.Min(candidateSolutions)
}
