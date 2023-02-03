package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const tunnelWidth = 7
const rockCount = 2022

type point struct {
	x, y int
}

type rock struct {
	pos    point   // The bottom left corner of the rock
	blocks []point // The blocks in the rock relative to its position
}

func (r *rock) canPlace(rocks map[point]bool, newPos point) bool {
	for _, b := range r.blocks {
		p := point{newPos.x + b.x, newPos.y + b.y}
		if p.x < 0 || tunnelWidth <= p.x || rocks[p] || p.y < 0 {
			return false
		}
	}
	return true
}

var rockTypes = [...]rock{
	{blocks: []point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},         // -
	{blocks: []point{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}}, // +
	{blocks: []point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}}, // ⅃
	{blocks: []point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},         // |
	{blocks: []point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}},         // #
}

func main() {
	jetPattern := strings.TrimSpace(input)

	solidRocks := make(map[point]bool)

	patternPos := 0
	highest := 0

	for i := 0; i < rockCount; i++ {
		// The currently falling rock
		falling := rockTypes[i%len(rockTypes)]
		falling.pos.y = highest + 3
		falling.pos.x = 2

		for {
			moveDir := jetPattern[patternPos%len(jetPattern)]

			var newPos point
			if moveDir == '>' {
				// Move right
				newPos = point{falling.pos.x + 1, falling.pos.y}
			} else {
				// Move left
				newPos = point{falling.pos.x - 1, falling.pos.y}
			}

			if falling.canPlace(solidRocks, newPos) {
				falling.pos = newPos
			}

			patternPos++

			reachedBottom := !falling.canPlace(solidRocks, point{falling.pos.x, falling.pos.y - 1})
			if reachedBottom {
				for _, block := range falling.blocks {
					p := point{falling.pos.x + block.x, falling.pos.y + block.y}
					solidRocks[p] = true
					if highest < p.y+1 {
						highest = p.y + 1
					}
				}
				break
			} else {
				falling.pos.y -= 1
			}
		}
	}

	fmt.Println("Answear:", highest)
}
