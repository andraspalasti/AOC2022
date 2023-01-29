package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type point struct {
	x, y int
}

type sensor struct {
	pos    point
	radius int
}

func (s *sensor) contains(p point) bool {
	return ManhattanDistance(s.pos, p) <= s.radius
}

const bound = 4_000_000

func main() {
	sensors := []sensor{}
	for _, line := range strings.Split(input, "\n") {
		sensor, beacon := sensor{}, point{}
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.pos.x, &sensor.pos.y, &beacon.x, &beacon.y)
		if err != nil {
			panic(err)
		}
		sensor.radius = ManhattanDistance(sensor.pos, beacon)
		sensors = append(sensors, sensor)
	}

	cur := point{x: 0, y: 0}
mainLoop:
	for y := 0; y < bound; y++ {
		for x := 0; x < bound; x++ {
			cur.x = x
			cur.y = y

			// Find which sensor does the current position collide with
			var collide *sensor = nil
			for _, s := range sensors {
				if s.contains(cur) {
					collide = &s
					break
				}
			}

			// If there is no collision than we found the point
			if collide == nil {
				break mainLoop
			}

			// Jump to the next possible location
			x = collide.pos.x + (collide.radius - Abs(y-collide.pos.y))
		}
	}

	ans := bound*cur.x + cur.y
	fmt.Println("Answear:", ans)
}

func ManhattanDistance(p1, p2 point) int {
	return Abs(p1.x-p2.x) + Abs(p1.y-p2.y)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
