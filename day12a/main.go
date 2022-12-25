package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

func main() {
	lines := strings.Split(input, "\n")

	start, end := Point{-1, -1}, Point{1, -1}
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		heightMap[i] = make([]int, len(line))
		for j, r := range line {
			switch r {
			case 'S':
				start.x, start.y = j, i
				heightMap[i][j] = 0
			case 'E':
				end.x, end.y = j, i
				heightMap[i][j] = int('z' - 'a')
			default:
				heightMap[i][j] = int(r - 'a')
			}
		}
	}

	neighbours := func(n Point) []Move[Point] {
		height := heightMap[n.y][n.x]
		nbs := make([]Move[Point], 0, 4) // Neighbours

		// Down
		if n.y+1 < len(heightMap) && heightMap[n.y+1][n.x]-1 <= height {
			nbs = append(nbs, Move[Point]{Node: Point{n.x, n.y + 1}, Cost: 1})
		}

		// Right
		if n.x+1 < len(heightMap[0]) && heightMap[n.y][n.x+1]-1 <= height {
			nbs = append(nbs, Move[Point]{Node: Point{n.x + 1, n.y}, Cost: 1})
		}

		// Up
		if 0 <= n.y-1 && heightMap[n.y-1][n.x]-1 <= height {
			nbs = append(nbs, Move[Point]{Node: Point{n.x, n.y - 1}, Cost: 1})
		}

		// Left
		if 0 <= n.x-1 && heightMap[n.y][n.x-1]-1 <= height {
			nbs = append(nbs, Move[Point]{Node: Point{n.x - 1, n.y}, Cost: 1})
		}
		return nbs
	}

	isEnd := func(n Point) bool {
		return n.x == end.x && n.y == end.y
	}

	length := ShortestPath(start, neighbours, isEnd)
	fmt.Println("Answear:", length)
}
