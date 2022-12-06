package main

import (
	_ "embed"
	"fmt"
)

const markerSize = 4

//go:embed input.txt
var input string

func HasDuplicate(s string) bool {
	runes := make(map[rune]struct{})
	for _, r := range s {
		_, ok := runes[r]
		if ok {
			return true
		} else {
			runes[r] = struct{}{}
		}
	}
	return false
}

func main() {
	startIndex := len(input) - 1
	for i := markerSize; i < len(input); i++ {
		marker := input[i-markerSize : i]

		if !HasDuplicate(marker) {
			startIndex = i
			break
		}
	}

	fmt.Println("Answear:", startIndex)
}
