package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	part := flag.Int("part", 1, "run part 1 or part 2")
	flag.Parse()

	if *part == 1 {
		ans := part1(input)
		fmt.Println("Answear for part 1:", ans)
	} else {
		ans := part2(input, 3)
		fmt.Println("Answear for part 2:", ans)
	}
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	priorities := make([]int, len(lines))

	for i, line := range lines {
		dup := duplicateRune(line[:len(line)/2], line[len(line)/2:])
		if dup == -1 {
			panic("no duplicate")
		}
		priorities[i] += priority(dup)
	}

	sum := 0
	for _, priority := range priorities {
		sum += priority
	}
	return sum
}

func duplicateRune(s1, s2 string) rune {
	for _, r1 := range s1 {
		for _, r2 := range s2 {
			if r1 == r2 {
				return r1
			}
		}
	}
	return -1
}

func priority(r rune) int {
	if unicode.IsLower(r) {
		return (int)(r-'a') + 1
	}
	return (int)(r-'A') + 27
}

func part2(input string, groupSize int) int {
	lines := strings.Split(input, "\n")
	priorities := make([]int, len(lines)/groupSize)

	group := make([]string, groupSize)
	for i, line := range lines {
		group[i%groupSize] = line

		if i%3 == groupSize-1 {
			// process the whole group
			common := commonRune(group...)
			if common == -1 {
				panic("no common rune was found")
			}
			priorities[i/groupSize] += priority(common)
		}
	}

	sum := 0
	for _, priority := range priorities {
		sum += priority
	}
	return sum
}

func commonRune(strs ...string) rune {
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
