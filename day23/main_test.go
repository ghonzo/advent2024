// Advent of Code 2024, Day 23
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
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 7},
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
		want string
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, "co,de,ka,ta"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}