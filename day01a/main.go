package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	elfs := strings.Split(input, "\n\n")
	cals := make([]int, len(elfs), len(elfs))

	for i, items := range elfs {
		for _, cal := range strings.Split(items, "\n") {
			if cal, err := strconv.Atoi(cal); err == nil {
				cals[i] += cal
			}
		}
	}

	max := 0
	for _, cal := range cals {
		if max < cal {
			max = cal
		}
	}
	fmt.Println("Answear:", max)
}
