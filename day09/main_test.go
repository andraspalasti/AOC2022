package main

import "testing"

var example string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var examplePart2 string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "example", input: example, want: 13},
		{name: "actual", input: input, want: 6067},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "example", input: examplePart2, want: 36},
		{name: "actual", input: input, want: 2471},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
