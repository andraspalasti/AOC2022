package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func ParseRange(s string) [2]int {
	i := strings.Index(s, "-")
	if i == -1 {
		panic("Unable to parse range")
	}

	l, err := strconv.Atoi(s[:i])
	if err != nil {
		panic("Unable to parse range")
	}
	h, err := strconv.Atoi(s[i+1:])
	if err != nil {
		panic("Unable to parse range")
	}
	return [2]int{l, h}
}

func main() {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		part1, part2, found := strings.Cut(line, ",")
		if !found {
			panic("Invalid line")
		}
		range1, range2 := ParseRange(part1), ParseRange(part2)

		if range1[0] <= range2[0] && range2[1] <= range1[1] {
			count++
		} else if range2[0] <= range1[0] && range1[1] <= range2[1] {
			count++
		}
	}

	fmt.Println("Answear:", count)
}
