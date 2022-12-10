package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	size := len(lines)
	heights := make([][]int, size)
	for row, line := range lines {
		heights[row] = make([]int, size)
		for col, r := range line {
			heights[row][col] = int(r - '0')
		}
	}

	visibleTrees := map[int]struct{}{}

	for i := 1; i < size-1; i++ {
		var (
			maxTop    = heights[0][i]
			maxBottom = heights[size-1][i]
			maxLeft   = heights[i][0]
			maxRight  = heights[i][size-1]
		)

		for j := 1; j < size-1; j++ {
			// going from top to bottom
			if maxTop < heights[j][i] {
				visibleTrees[j*size+i] = struct{}{}
				maxTop = heights[j][i]
			}

			// going from left to right
			if maxLeft < heights[i][j] {
				visibleTrees[i*size+j] = struct{}{}
				maxLeft = heights[i][j]
			}

			// going from bottom to top
			if maxBottom < heights[size-j-1][i] {
				visibleTrees[(size-j-1)*size+i] = struct{}{}
				maxBottom = heights[size-j-1][i]
			}

			// going from right to left
			if maxRight < heights[i][size-j-1] {
				visibleTrees[i*size+size-j-1] = struct{}{}
				maxRight = heights[i][size-j-1]
			}
		}
	}

	sides := 4 * (size - 1)
	fmt.Println("Answear:", sides+len(visibleTrees))
}
