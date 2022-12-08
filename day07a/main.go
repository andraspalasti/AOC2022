package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const (
	treshold  = 100_000
	delimiter = "-"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	pwd := []string{}
	folders := make(map[string]int)

	for _, line := range lines {
		if strings.HasPrefix(line, "$ ") {
			// this is a command
			cmd := strings.SplitN(line, " ", 3)
			if cmd[1] == "cd" {
				if cmd[2] == ".." {
					pwd = pwd[:len(pwd)-1]
				} else {
					pwd = append(pwd, cmd[2])
				}
			}
		} else if !strings.HasPrefix(line, "dir ") {
			// Its a row from ls where the row is a file
			size, err := strconv.Atoi(strings.Split(line, " ")[0])
			if err != nil {
				panic(err)
			}

			fpath := strings.Builder{}
			for i, part := range pwd {
				if i > 0 {
					fpath.WriteString(delimiter)
				}
				fpath.WriteString(part)
				folders[fpath.String()] += size
			}
		}
	}

	sum := 0
	for _, size := range folders {
		if size <= treshold {
			sum += size
		}
	}
	println(folders["/"])

	fmt.Println("Answear:", sum)
}
