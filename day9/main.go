// Advent of Code 2024, Day 9
package main

import (
	"fmt"

	"github.com/ghonzo/advent2024/common"
)

// Day 9: Disk Fragmenter
// Part 1 answer: 6291146824486
// Part 2 answer: 6307279963620
func main() {
	fmt.Println("Advent of Code 2024, Day 9")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var memory []int
	var id int
	file := true
	for _, c := range []byte(entries[0]) {
		n := int(c - '0')
		var memoryValue int
		if file {
			memoryValue = id
		} else {
			// Empty
			memoryValue = -1
		}
		for i := 0; i < n; i++ {
			memory = append(memory, memoryValue)
		}
		if file {
			file = false
		} else {
			id++
			file = true
		}
	}
	leftIndex := 0
	rightIndex := len(memory) - 1
	for {
		//fmt.Println(memory)
		// Decrement right index until we find a non empty
		for ; memory[rightIndex] < 0; rightIndex-- {
		}
		// Increment left index until we find empty
		for ; memory[leftIndex] >= 0; leftIndex++ {
		}
		if rightIndex <= leftIndex {
			break
		}
		memory[rightIndex], memory[leftIndex] = memory[leftIndex], memory[rightIndex]
	}
	// Now checksum
	var total int
	for pos, m := range memory {
		if m < 0 {
			break
		}
		total += pos * m
	}
	return total
}

func part2(entries []string) int {
	var memory []int
	var id int
	file := true
	for _, c := range []byte(entries[0]) {
		n := int(c - '0')
		var memoryValue int
		if file {
			memoryValue = id
		} else {
			// Empty
			memoryValue = -1
		}
		for i := 0; i < n; i++ {
			memory = append(memory, memoryValue)
		}
		if file {
			file = false
		} else {
			id++
			file = true
		}
	}
	maxId := id + 1
	rightIndex := len(memory) - 1
	for {
		// Find the next memory chunk to move
		// Decrement right index until we find a non empty
		for ; memory[rightIndex] < 0; rightIndex-- {
		}
		// Save that id and indexPos
		rid := memory[rightIndex]
		if rid == 0 {
			break
		}
		rightmostIndex := rightIndex
		// Decrement until we get something different
		for rightIndex--; memory[rightIndex] == rid; rightIndex-- {
		}
		// Make sure we aren't trying to move it twice
		if rid < maxId {
			moveFile(memory, rightIndex+1, rightmostIndex-rightIndex)
			maxId = min(rid, maxId)
		}
	}
	// Now checksum
	var total int
	for pos, m := range memory {
		if m >= 0 {
			total += pos * m
		}
	}
	return total
}

// Return false if there is no way any more will work
func moveFile(memory []int, index, size int) {
	var sizeEmpty int
	for leftIndex := 0; leftIndex < index; leftIndex++ {
		if memory[leftIndex] < 0 {
			sizeEmpty++
			if sizeEmpty == size {
				for i := 0; i < size; i++ {
					memory[leftIndex-i], memory[index+i] = memory[index+i], memory[leftIndex-i]
				}
				return
			}
		} else {
			sizeEmpty = 0
		}
	}
}
