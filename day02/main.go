package main

import (
	_ "embed"
	"flag"
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
	total := 0
	for _, line := range strings.Split(input, "\n") {
		opponent, me, found := strings.Cut(line, " ")
		if !found {
			panic("invalid input: " + line)
		}
		total += score(handFromString(me), handFromString(opponent))
	}
	return total
}

func part2(input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		opponent, outcome, found := strings.Cut(line, " ")
		if !found {
			panic("invalid input: " + line)
		}

		var me Hand
		opponentHand := handFromString(opponent)
		if outcome == "X" {
			// need to lose
			me = win[opponentHand]
		} else if outcome == "Y" {
			// need to draw
			me = opponentHand
		} else if outcome == "Z" {
			// need to win
			me = lose[opponentHand]
		}
		total += score(me, opponentHand)
	}

	return total
}

func score(me, opponent Hand) int {
	score := int(me) + 1
	if me == opponent {
		score += 3
	} else if win[me] == opponent {
		score += 6
	}
	return score
}

func handFromString(str string) Hand {
	var stringToHand = map[string]Hand{"A": Rock, "B": Paper, "C": Scissors, "X": Rock, "Y": Paper, "Z": Scissors}
	hand, ok := stringToHand[strings.TrimSpace(str)]
	if !ok {
		panic("can't convert '" + str + "' to Hand")
	}
	return hand
}
