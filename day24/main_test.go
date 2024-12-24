// Advent of Code 2024, Day 24
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
		want uint64
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
