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

func main() {
	const ROW = 2_000_000
	taken := make(map[point]struct{})

	beacons := []point{}
	for _, line := range strings.Split(input, "\n") {
		sensor, beacon := point{}, point{}
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.x, &sensor.y, &beacon.x, &beacon.y)
		if err != nil {
			panic(err)
		}
		beacons = append(beacons, beacon)

		radius := ManhattanDistance(sensor, beacon)
		if radius < Abs(ROW-sensor.y) {
			continue
		}

		dx := radius - Abs(ROW-sensor.y)
		for i := 0; i <= dx; i++ {
			taken[point{x: sensor.x + i, y: ROW}] = struct{}{}
			taken[point{x: sensor.x - i, y: ROW}] = struct{}{}
		}
	}

	for _, beacon := range beacons {
		delete(taken, beacon)
	}

	fmt.Println("Answear:", len(taken))
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
