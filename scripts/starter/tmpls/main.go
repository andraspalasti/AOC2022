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
		ans := part1(input)
		fmt.Println("Answear for part 1:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Answear for part 2:", ans)
	}
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}
