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
	heights := parseInput(input)
	visibleTrees := map[int]struct{}{}
	size := len(heights)

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
	return sides + len(visibleTrees)
}

func part2(input string) int {
	heights := parseInput(input)
	size := len(heights)

	max := 0
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			score := scenicScore(heights, i, j)
			if max < score {
				max = score
			}
		}
	}
	return max
}

func scenicScore(heights [][]int, row, col int) int {
	height := heights[row][col]
	size := len(heights)

	// The view distances from the specific direction
	var (
		wdTop    = 0
		wdLeft   = 0
		wdRight  = 0
		wdBottom = 0
	)
	for i := row - 1; 0 <= i; i-- {
		wdTop++
		if height <= heights[i][col] {
			break
		}
	}
	for i := col - 1; 0 <= i; i-- {
		wdLeft++
		if height <= heights[row][i] {
			break
		}
	}
	for i := row + 1; i < size; i++ {
		wdBottom++
		if height <= heights[i][col] {
			break
		}
	}
	for i := col + 1; i < size; i++ {
		wdRight++
		if height <= heights[row][i] {
			break
		}
	}
	return wdTop * wdLeft * wdRight * wdBottom
}

func parseInput(input string) [][]int {
	heights := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		heights = append(heights, make([]int, len(line)))
		for col, r := range line {
			heights[len(heights)-1][col] = int(r - '0')
		}
	}
	return heights
}
