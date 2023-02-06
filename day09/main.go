package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/andraspalasti/aoc2022/mymath"
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
	visited := map[knot]struct{}{}
	head, tail := knot{0, 0}, knot{0, 0}

	// Put inital tail position in visited positions
	visited[tail] = struct{}{}

	for _, line := range strings.Split(input, "\n") {
		dir, amount := "", 0
		_, err := fmt.Sscanf(line, "%s %d", &dir, &amount)
		if err != nil {
			panic("parsing line: " + err.Error())
		}

		for i := 0; i < amount; i++ {
			head.move(dir)
			tail.follow(head)
			visited[tail] = struct{}{}
		}
	}

	return len(visited)
}

func part2(input string) int {
	visited := map[knot]struct{}{}

	rope := make([]knot, 10)
	head, tail := &rope[0], &rope[len(rope)-1]

	// Put inital tail position in visited positions
	visited[*tail] = struct{}{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		dir, amount := "", 0
		_, err := fmt.Sscanf(line, "%s %d", &dir, &amount)
		if err != nil {
			panic("parsing line: " + err.Error())
		}

		for i := 0; i < amount; i++ {
			head.move(dir)
			for j := 1; j < len(rope); j++ {
				rope[j].follow(rope[j-1])
			}

			visited[*tail] = struct{}{}
		}
	}

	return len(visited)
}

type knot struct {
	x, y int
}

func (k *knot) move(dir string) {
	switch dir {
	case "L":
		k.x--
	case "R":
		k.x++
	case "U":
		k.y++
	case "D":
		k.y--
	default:
		panic("invalid direction: " + dir)
	}
}

func (tail *knot) follow(head knot) {
	xDist, yDist := mymath.Abs(head.x-tail.x), mymath.Abs(head.y-tail.y)
	if xDist <= 1 && yDist <= 1 {
		// we are in range so we dont have to do anything
		return
	}

	if 1 < xDist {
		if head.x < tail.x {
			tail.x = head.x + 1
		} else {
			tail.x = head.x - 1
		}
	} else {
		tail.x = head.x
	}

	if 1 < yDist {
		if head.y < tail.y {
			tail.y = head.y + 1
		} else {
			tail.y = head.y - 1
		}
	} else {
		tail.y = head.y
	}
}
