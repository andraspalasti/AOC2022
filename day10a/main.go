package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	// The sum of the specific signal strengts
	sum := 0

	// In what cycle do we want to know the registers value
	watchCycle := 20

	// This means that from `cycle` X = `register`
	cycle, register := 1, 1
	for _, line := range lines {
		prev := register
		if strings.HasPrefix(line, "addx ") {
			val, err := strconv.Atoi(line[strings.IndexRune(line, ' ')+1:])
			if err != nil {
				panic(err)
			}
			register += val
			cycle += 2
		} else if strings.HasPrefix(line, "noop") {
			cycle += 1
		}

		if watchCycle < cycle {
			sum += prev * watchCycle
			watchCycle += 40
		}
	}

	fmt.Println("Answear:", sum)
}
