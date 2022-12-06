package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

const groupSize = 3

//go:embed input.txt
var input string

func Priority(r rune) int {
	if unicode.IsLower(r) {
		return (int)(r-'a') + 1
	}
	return (int)(r-'A') + 27
}

func CommonRune(strs ...string) rune {
	runes := make(map[rune]int)
	lastIdx := len(strs) - 1
	for i, s := range strs {
		for _, r := range s {
			count := runes[r]
			if count == i {
				runes[r] = count + 1

				if i == lastIdx && count == lastIdx {
					// we are at the last string and
					// the count of the rune is equal to the number of strings
					return r
				}
			}
		}
	}
	return -1
}

func main() {
	lines := strings.Split(input, "\n")
	priorities := make([]int, len(lines)/groupSize, len(lines)/groupSize)

	group := make([]string, groupSize, groupSize)
	for i, line := range lines {
		group[i%groupSize] = line

		if i%3 == groupSize-1 {
			// process the whole group
			common := CommonRune(group...)
			if common == -1 {
				panic("No common rune was found")
			}
			priorities[i/groupSize] += Priority(common)
		}
	}

	sum := 0
	for _, priority := range priorities {
		sum += priority
	}
	fmt.Println("Answear:", sum)
}
