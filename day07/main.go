package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"
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
	const (
		maxDirSize = 100_000
	)

	folders := parseInput(input)

	sum := 0
	for _, size := range folders {
		if size <= maxDirSize {
			sum += size
		}
	}
	return sum
}

func part2(input string) int {
	const (
		totalSize  = 70_000_000
		neededSize = 30_000_000
	)

	folders := parseInput(input)

	min := math.MaxInt
	unused := totalSize - folders["/"]
	for _, size := range folders {
		if neededSize <= unused+size && size < min {
			min = size
		}
	}

	return min
}

// Returns the size of each folder in a map[folder path]size
func parseInput(input string) map[string]int {
	pwd := []string{}
	folders := map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$ ") {
			cmd := strings.SplitN(line, " ", 3)
			if cmd[1] == "cd" {
				if cmd[2] == ".." {
					pwd = pwd[:len(pwd)-1]
				} else {
					pwd = append(pwd, cmd[2])
				}
			}
		} else if !strings.HasPrefix(line, "dir ") {
			size, name := 0, ""
			_, err := fmt.Sscanf(line, "%d %s", &size, &name)
			if err != nil {
				panic("parsing line: " + err.Error())
			}

			fpath := strings.Builder{}
			for i, part := range pwd {
				if i > 0 {
					fpath.WriteRune('/')
				}
				fpath.WriteString(part)
				folders[fpath.String()] += size
			}
		}
	}
	return folders
}
