package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

type valve struct {
	name     string
	flowRate int
	tunnels  []string
	open     bool
}

func main() {
	network := make(map[string]*valve)
	for _, line := range strings.Split(input, "\n") {
		v := &valve{}

		valve, tunnels, found := strings.Cut(line, "; tunnels lead to valves ")
		if !found {
			valve, tunnels, found = strings.Cut(line, "; tunnel leads to valve ")
		}
		if !found {
			panic("parsing: could not find separator")
		}

		_, err := fmt.Sscanf(valve, "Valve %s has flow rate=%d", &v.name, &v.flowRate)
		if err != nil {
			panic("parsing: " + err.Error())
		}
		v.tunnels = strings.Split(tunnels, ", ")

		network[v.name] = v
	}

	// If valve is in here than it is opened
	// open := make(map[string]struct{})

	ans := maxReleased("AA", network, 30, 0, make(map[string]int))
	fmt.Println("Answear:", ans)
}

func maxReleased(pos string, network map[string]*valve, minutesLeft, released int, memo map[string]int) int {
	if minutesLeft <= 1 {
		return released
	}

	key := hash(pos, network, minutesLeft, released)
	if v, ok := memo[key]; ok {
		return v
	}
	currentValve := network[pos]

	// Check the max pressure without opening the valve
	max := released
	for _, neighbour := range currentValve.tunnels {
		tmp := maxReleased(neighbour, network, minutesLeft-1, released, memo)
		if max < tmp {
			max = tmp
		}
	}

	// See what is the max released pressure if we open the valve
	if !currentValve.open && currentValve.flowRate > 0 {
		currentValve.open = true

		released = released + network[pos].flowRate*(minutesLeft-1)
		maxOpened := released
		for _, neighbour := range currentValve.tunnels {
			tmp := maxReleased(neighbour, network, minutesLeft-2, released, memo)
			if maxOpened < tmp {
				maxOpened = tmp
			}
		}

		if max < maxOpened {
			max = maxOpened
		}

		currentValve.open = false
	}

	memo[key] = max
	return max
}

func hash(pos string, network map[string]*valve, minutesLeft, released int) string {
	opened := make([]string, 0, len(network))
	for _, valve := range network {
		if valve.open {
			opened = append(opened, valve.name)
		}
	}
	sort.Strings(opened)
	return fmt.Sprintf("%s %v %d %d", pos, opened, minutesLeft, released)
}
