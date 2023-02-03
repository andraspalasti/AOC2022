package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/andraspalasti/aoc2022/mymath"
)

//go:embed input.txt
var input string

var offsets = [...]cube{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

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
	lava := parseInput(input)

	surfaceArea := 0
	for c := range lava {
		for _, offset := range offsets {
			neighbour := cube{c.x + offset.x, c.y + offset.y, c.z + offset.z}
			if !lava[neighbour] {
				surfaceArea++
			}
		}
	}

	return surfaceArea
}

func part2(input string) int {
	lava := parseInput(input)

	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	minZ, maxZ := math.MaxInt, math.MinInt

	for c := range lava {
		minX, maxX = mymath.Min(minX, c.x), mymath.Max(maxX, c.x)
		minY, maxY = mymath.Min(minY, c.y), mymath.Max(maxY, c.y)
		minZ, maxZ = mymath.Min(minZ, c.z), mymath.Max(maxZ, c.z)
	}

	minX, maxX = minX-1, maxX+1
	minY, maxY = minY-1, maxY+1
	minZ, maxZ = minZ-1, maxZ+1

	// This contains the cubes that wrap the lava droplet
	visited := make(map[cube]bool)

	// Unvisited cubes that wrap the droplet
	queue := []cube{{maxX, maxY, maxZ}}

	surfaceArea := 0
	for 0 < len(queue) {
		c := queue[0]
		queue = queue[1:]

		if visited[c] {
			continue
		}
		visited[c] = true

		for _, offset := range offsets {
			neighbour := cube{c.x + offset.x, c.y + offset.y, c.z + offset.z}
			if lava[neighbour] {
				surfaceArea++
			} else if !visited[neighbour] {
				// Check if cube is in the area of the lava
				inArea := minX <= neighbour.x && neighbour.x <= maxX &&
					minY <= neighbour.y && neighbour.y <= maxY &&
					minZ <= neighbour.z && neighbour.z <= maxZ
				if inArea {
					queue = append(queue, neighbour)
				}
			}
		}
	}

	return surfaceArea
}

type cube struct {
	x, y, z int
}

func parseInput(input string) map[cube]bool {
	cubes := make(map[cube]bool)
	for _, line := range strings.Split(input, "\n") {
		cube := cube{}
		_, err := fmt.Sscanf(line, "%d,%d,%d", &cube.x, &cube.y, &cube.z)
		if err != nil {
			panic(err)
		}
		cubes[cube] = true
	}
	return cubes
}
