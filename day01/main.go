package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
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
	elfs := strings.Split(input, "\n\n")
	cals := make([]int, len(elfs))

	for i, items := range elfs {
		for _, cal := range strings.Split(items, "\n") {
			cal, err := strconv.Atoi(cal)
			if err != nil {
				panic(err)
			}
			cals[i] += cal
		}
	}

	max := 0
	for _, cal := range cals {
		if max < cal {
			max = cal
		}
	}
	return max
}

func part2(input string) int {
	elfs := strings.Split(input, "\n\n")
	cals := make([]int, len(elfs))

	for i, items := range elfs {
		for _, cal := range strings.Split(items, "\n") {
			cal, err := strconv.Atoi(cal)
			if err != nil {
				panic(err)
			}
			cals[i] += cal
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cals)))

	sum := 0
	for _, cal := range cals[:3] {
		sum += cal
	}
	return sum
}
