package main

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func Priority(r rune) int {
	if unicode.IsLower(r) {
		return (int)(r-'a') + 1
	}
	return (int)(r-'A') + 27
}

func DuplicateRune(s1, s2 string) rune {
	for _, r1 := range s1 {
		for _, r2 := range s2 {
			if r1 == r2 {
				return r1
			}
		}
	}
	return -1
}

func main() {
	lines := strings.Split(input, "\n")
	priorities := make([]int, len(lines), len(lines))

	for i, line := range lines {
		dup := DuplicateRune(line[:len(line)/2], line[len(line)/2:])
		if dup == -1 {
			panic("No duplicate")
		}
		priorities[i] += Priority(dup)
	}

	sum := 0
	for _, priority := range priorities {
		sum += priority
	}
	fmt.Println("Answear:", sum)
}
