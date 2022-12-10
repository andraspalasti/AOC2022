package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func ScenicScore(treeMap [][]int, row, col int) int {
	height := treeMap[row][col]
	size := len(treeMap)

	// The view distances from the specific direction
	var (
		wdTop    = 0
		wdLeft   = 0
		wdRight  = 0
		wdBottom = 0
	)
	for i := row - 1; 0 <= i; i-- {
		wdTop++
		if height <= treeMap[i][col] {
			break
		}
	}
	for i := col - 1; 0 <= i; i-- {
		wdLeft++
		if height <= treeMap[row][i] {
			break
		}
	}
	for i := row + 1; i < size; i++ {
		wdBottom++
		if height <= treeMap[i][col] {
			break
		}
	}
	for i := col + 1; i < size; i++ {
		wdRight++
		if height <= treeMap[row][i] {
			break
		}
	}
	return wdTop * wdLeft * wdRight * wdBottom
}

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

	max := 0
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			score := ScenicScore(heights, i, j)
			if max < score {
				max = score
			}
		}
	}

	fmt.Println("Answear:", max)
}
