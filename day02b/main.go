package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Hand int

const (
	Rock Hand = iota
	Paper
	Scissors
)

var win = map[Hand]Hand{Rock: Scissors, Paper: Rock, Scissors: Paper}
var lose = map[Hand]Hand{Scissors: Rock, Rock: Paper, Paper: Scissors}

func Score(me, opponent Hand) int {
	score := int(me) + 1
	if me == opponent {
		score += 3
	} else if win[me] == opponent {
		score += 6
	}
	return score
}

func HandFromString(str string) Hand {
	var stringToHand = map[string]Hand{"A": Rock, "B": Paper, "C": Scissors}
	hand, ok := stringToHand[strings.TrimSpace(str)]
	if !ok {
		panic("Can't convert '" + str + "' to Hand")
	}
	return hand
}

func main() {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		opponent, outcome, found := strings.Cut(line, " ")
		if !found {
			panic("Invalid input")
		}

		{
			var me Hand
			opponent := HandFromString(opponent)
			if outcome == "X" {
				// need to lose
				me = win[opponent]
			} else if outcome == "Y" {
				// need to draw
				me = opponent
			} else if outcome == "Z" {
				// need to win
				me = lose[opponent]
			}
			total += Score(me, opponent)
		}
	}

	fmt.Println("Answear:", total)
}
