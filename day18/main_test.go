// Advent of Code 2024, Day 18
package main

import (
	"testing"

	"github.com/ghonzo/advent2024/common"
)

func Test_part1(t *testing.T) {
	type args struct {
		entries  []string
		gridSize int
		numRows  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 7, 12}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries, tt.args.gridSize, tt.args.numRows); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		entries  []string
		gridSize int
		numRows  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt"), 7, 12}, "6,1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries, tt.args.gridSize, tt.args.numRows); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
