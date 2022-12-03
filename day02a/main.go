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

func Score(me, opponent Hand) int {
	rules := map[Hand]Hand{Rock: Scissors, Paper: Rock, Scissors: Paper}
	score := int(me) + 1
	if me == opponent {
		score += 3
	} else if rules[me] == opponent {
		score += 6
	}
	return score
}

func HandFromString(str string) Hand {
	var stringToHand = map[string]Hand{"A": Rock, "B": Paper, "C": Scissors, "X": Rock, "Y": Paper, "Z": Scissors}
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
		opponent, me, found := strings.Cut(line, " ")
		if !found {
			panic("Invalid input")
		}
		total += Score(HandFromString(me), HandFromString(opponent))
	}

	fmt.Println("Answear:", total)
}
