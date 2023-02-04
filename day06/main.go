package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	part := flag.Int("part", 1, "run part 1 or part 2")
	flag.Parse()

	if *part == 1 {
		ans := part1(input, 4)
		fmt.Println("Answear for part 1:", ans)
	} else {
		ans := part2(input, 14)
		fmt.Println("Answear for part 2:", ans)
	}
}

func part1(input string, markerSize int) int {
	startIndex := len(input) - 1
	for i := markerSize; i < len(input); i++ {
		marker := input[i-markerSize : i]

		if !hasDuplicate(marker) {
			startIndex = i
			break
		}
	}
	return startIndex
}

func part2(input string, markerSize int) int {
	startIndex := len(input) - 1
	for i := markerSize; i < len(input); i++ {
		marker := input[i-markerSize : i]

		if !hasDuplicate(marker) {
			startIndex = i
			break
		}
	}
	return startIndex
}

func hasDuplicate(s string) bool {
	runes := make(map[rune]struct{})
	for _, r := range s {
		_, ok := runes[r]
		if ok {
			return true
		} else {
			runes[r] = struct{}{}
		}
	}
	return false
}
