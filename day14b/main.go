package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func MoveDown(cave map[Point]bool, p Point) Point {
	new := Point{x: p.x, y: p.y + 1}
	if !cave[new] {
		return new
	}

	new.x = p.x - 1
	if !cave[new] {
		return new
	}

	new.x = p.x + 1
	if !cave[new] {
		return new
	}
	return p
}

func MustParsePoint(point string) Point {
	xStr, yStr, found := strings.Cut(point, ",")
	if !found {
		panic("Could not find ',' separator")
	}

	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		panic(err)
	}

	return Point{x, y}
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

//go:embed input.txt
var input string

func main() {
	cave := make(map[Point]bool)

	highestY := 0
	for _, line := range strings.Split(input, "\n") {
		points := strings.Split(line, " -> ")

		start := MustParsePoint(points[0])
		for i := 1; i < len(points); i++ {
			end := MustParsePoint(points[i])

			dv := Point{x: end.x - start.x, y: end.y - start.y}
			if dv.x != 0 {
				dv.x = dv.x / Abs(dv.x)
			} else {
				dv.y = dv.y / Abs(dv.y)
			}

			tmp := start
			for tmp != end {
				cave[tmp] = true
				tmp.x += dv.x
				tmp.y += dv.y
			}

			if highestY < end.y {
				highestY = end.y
			}

			cave[end] = true
			start = end
		}
	}

	count := 0
	floor := highestY + 2
	for {
		sand, prev := Point{x: 500, y: 0}, Point{x: 500, y: -1}

		// Check if the origin of the sand is blocked
		if sand == MoveDown(cave, sand) {
			count++
			break
		}

		for prev != sand {
			prev = sand
			sand = MoveDown(cave, sand)

			if sand.y == floor-1 {
				break
			}
		}

		// Successfully dropped sand
		cave[sand] = true
		count++
	}

	fmt.Println("Answear:", count)
}
