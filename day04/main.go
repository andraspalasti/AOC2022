package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
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
		ans := part2(input)
		fmt.Println("Answear for part 2:", ans)
	}
}

func part1(input string) int {
	pairs := parseInput(input)
	count := 0
	for _, p := range pairs {
		if p.left[0] <= p.right[0] && p.right[1] <= p.left[1] {
			count++
		} else if p.right[0] <= p.left[0] && p.left[1] <= p.right[1] {
			count++
		}
	}
	return count
}

func part2(input string) int {
	pairs := parseInput(input)
	count := 0
	for _, p := range pairs {
		if p.left[0] <= p.right[1] && p.right[1] <= p.left[1] {
			count++
		} else if p.right[0] <= p.left[1] && p.left[1] <= p.right[1] {
			count++
		}
	}
	return count
}

type pair struct {
	left, right [2]int
}

func parseInput(input string) []pair {
	pairs := []pair{}
	for _, line := range strings.Split(input, "\n") {
		p := pair{}
		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &p.left[0], &p.left[1], &p.right[0], &p.right[1])
		if err != nil {
			panic("parsing line: " + err.Error())
		}
		pairs = append(pairs, p)
	}
	return pairs
}
