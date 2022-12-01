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
			if cal, err := strconv.Atoi(cal); err == nil {
				cals[i] += cal
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cals)))
	fmt.Println("Answear:", Sum(cals[:3]))
}

type Number interface {
	int
}

func Sum[T Number](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
