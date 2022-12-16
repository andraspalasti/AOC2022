package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type knot struct {
	x, y int
}

func (tail *knot) follow(head knot) {
	xDist, yDist := abs(head.x-tail.x), abs(head.y-tail.y)
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

func (p *knot) move(dir string) {
	if dir == "L" {
		p.x--
	} else if dir == "U" {
		p.y++
	} else if dir == "R" {
		p.x++
	} else if dir == "D" {
		p.y--
	}
}

func abs(n int) int {
	if 0 < n {
		return n
	}
	return -n
}

func main() {
	visited := make(map[knot]struct{})
	rope := make([]knot, 10)
	head, tail := &rope[0], &rope[len(rope)-1]

	// Put inital tail position in visited positions
	visited[*tail] = struct{}{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		dir, amountStr, found := strings.Cut(line, " ")
		if !found {
			panic("Invalid line")
		}

		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			panic(err)
		}

		for i := 0; i < amount; i++ {
			head.move(dir)
			for j := 1; j < len(rope); j++ {
				rope[j].follow(rope[j-1])
			}
			visited[*tail] = struct{}{}
		}
	}

	fmt.Println("Answear:", len(visited))
}
