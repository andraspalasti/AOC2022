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

	// This means that from `cycle` X = `register`
	cycle, register := 1, 1

	for _, line := range lines {
		from, X := cycle, register
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

		for i := from; i < cycle; i++ {
			// The index of the pixel that we are printing now
			cur := (i % 40) - 1

			// Check if we are printing the place where the sprite is at
			if X-1 <= cur && cur <= X+1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}

			if i%40 == 0 {
				fmt.Println()
			}
		}
	}
}
