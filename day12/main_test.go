// Advent of Code 2024, Day 12
package main

import (
	"testing"

	"github.com/ghonzo/advent2024/common"
)

func Test_part1(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 140},
		{"example2", args{common.ReadStringsFromFile("testdata/example2.txt")}, 772},
		{"example3", args{common.ReadStringsFromFile("testdata/example3.txt")}, 1930},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 80},
		{"example2", args{common.ReadStringsFromFile("testdata/example2.txt")}, 436},
		{"example4", args{common.ReadStringsFromFile("testdata/example4.txt")}, 236},
		{"example5", args{common.ReadStringsFromFile("testdata/example5.txt")}, 368},
		{"example3", args{common.ReadStringsFromFile("testdata/example3.txt")}, 1206},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
