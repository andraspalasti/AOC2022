package main

import "testing"

func TestPart1(t *testing.T) {
	markerSize := 4
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "example-1", input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 7},
		{name: "example-2", input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
		{name: "example-3", input: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
		{name: "example-4", input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
		{name: "example-5", input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
		{name: "actual", input: input, want: 1034},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, markerSize); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	markerSize := 14
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "example-1", input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 19},
		{name: "example-2", input: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 23},
		{name: "example-3", input: "nppdvjthqldpwncqszvftbrmjlhg", want: 23},
		{name: "example-4", input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 29},
		{name: "example-5", input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 26},
		{name: "actual", input: input, want: 2472},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, markerSize); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
