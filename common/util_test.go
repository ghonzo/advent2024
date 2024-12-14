package common

import (
	"reflect"
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pos", args{60}, 60},
		{"zero", args{0}, 0},
		{"neg", args{-999}, 999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pos", args{"225"}, 225},
		{"zero", args{"0"}, 0},
		{"neg", args{"-10"}, -10},
		{"empty", args{""}, 0},
		{"invalid", args{"pickle"}, 0},
		{"spaces", args{" 33 "}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atoi(tt.args.s); got != tt.want {
				t.Errorf("Atoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSgn(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pos", args{60}, 1},
		{"zero", args{0}, 0},
		{"neg", args{-999}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sgn(tt.args.a); got != tt.want {
				t.Errorf("Sgn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, ""},
		{"one char", args{"A"}, "A"},
		{"two char", args{"AB"}, "BA"},
		{"three char", args{"ABC"}, "CBA"},
		{"four char", args{"ABCD"}, "DCBA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToInts(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{""}, []int{}},
		{"empty2", args{"empty"}, []int{}},
		{"one", args{"123"}, []int{123}},
		{"one negative", args{"-345"}, []int{345}},
		{"one with junk", args{"  -325*"}, []int{325}},
		{"comma separated", args{"2,4,6,8"}, []int{2, 4, 6, 8}},
		{"comma and space separated", args{"2, 4, 6, 18"}, []int{2, 4, 6, 18}},
		{"space separated", args{"112233     445566"}, []int{112233, 445566}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToInts(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMod(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pos pos", args{11, 5}, 1},
		{"neg pos", args{-11, 5}, 4},
		{"pos neg", args{11, -5}, -4},
		{"neg neg", args{-11, -5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mod(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Mod() = %v, want %v", got, tt.want)
			}
		})
	}
}
