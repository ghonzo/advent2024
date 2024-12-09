// Advent of Code 2024, Day 9
package main

import (
	"fmt"
	"math"

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

type diskmap []int

func (d diskmap) swap(i, j, n int) {
	// This probably doesn't work if the ranges overlap fyi
	for x := 0; x < n; x++ {
		d[i+x], d[j+x] = d[j+x], d[i+x]
	}
}

func (d diskmap) checksum() int {
	var total int
	for pos, fileId := range d {
		if fileId >= 0 {
			total += pos * fileId
		}
	}
	return total
}

func part1(entries []string) int {
	disk := readDiskmap(entries[0])
	leftIndex := 0
	rightIndex := len(disk) - 1
	for {
		// Decrement right index until we find a non empty
		for ; disk[rightIndex] < 0; rightIndex-- {
		}
		// Increment left index until we find empty
		for ; disk[leftIndex] >= 0; leftIndex++ {
		}
		if rightIndex <= leftIndex {
			break
		}
		disk.swap(rightIndex, leftIndex, 1)
	}
	return disk.checksum()
}

func readDiskmap(s string) diskmap {
	var disk diskmap
	fileId := 0
	expectFile := true
	for _, c := range []byte(s) {
		n := int(c - '0')
		blockValue := -1
		if expectFile {
			blockValue = fileId
		}
		for i := 0; i < n; i++ {
			disk = append(disk, blockValue)
		}
		if !expectFile {
			fileId++
		}
		expectFile = !expectFile
	}
	return disk
}

func part2(entries []string) int {
	disk := readDiskmap(entries[0])
	maxId := math.MaxInt
	rightIndex := len(disk) - 1
	for {
		// Find the next file to move
		// Decrement right index until we find a non empty
		for ; disk[rightIndex] < 0; rightIndex-- {
		}
		// Save that fileId and indexPos
		rid := disk[rightIndex]
		if rid == 0 {
			// reached the end
			break
		}
		rightmostIndex := rightIndex
		// Decrement until we get something different
		for rightIndex--; disk[rightIndex] == rid; rightIndex-- {
		}
		// Make sure we aren't trying to move it twice
		if rid < maxId {
			tryToMoveFile(disk, rightIndex+1, rightmostIndex-rightIndex)
			maxId = min(rid, maxId)
		}
	}
	return disk.checksum()
}

func tryToMoveFile(disk diskmap, i, n int) {
	var sizeEmpty int
	for leftIndex := 0; leftIndex < i; leftIndex++ {
		if disk[leftIndex] < 0 {
			sizeEmpty++
			if sizeEmpty == n {
				// We found a blank chunk big enough
				disk.swap(leftIndex-n+1, i, n)
				return
			}
		} else {
			// Keep looking
			sizeEmpty = 0
		}
	}
}
