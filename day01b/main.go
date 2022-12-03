package main

import (
	_ "embed"
	"fmt"
	"sort"
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
			cal, err := strconv.Atoi(cal)
			if err != nil {
				panic(err)
			}
			cals[i] += cal
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cals)))

	sum := 0
	for _, cal := range cals[:3] {
		sum += cal
	}
	fmt.Println("Answear:", sum)
}
