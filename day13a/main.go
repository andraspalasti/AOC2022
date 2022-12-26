package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Packet struct {
	N    int
	List []Packet
}

func (p Packet) IsInt() bool {
	return p.List == nil
}

// Return values:
//   - 1  Left less than right
//   - -1 Left greater than right
//   - 0  Left equals right
func ComparePackets(l, r Packet) int {
	lInt, rInt := l.IsInt(), r.IsInt()
	if lInt && rInt {
		return CompareInts(l.N, r.N)
	}

	if lInt && !rInt {
		return ComparePackets(Packet{List: []Packet{l}}, r)
	}

	if !lInt && rInt {
		return ComparePackets(l, Packet{List: []Packet{r}})
	}

	length := min(len(l.List), len(r.List))
	for i := 0; i < length; i++ {
		result := ComparePackets(l.List[i], r.List[i])
		if result != 0 {
			return result
		}
	}

	if len(l.List) < len(r.List) {
		return 1
	} else if len(l.List) > len(r.List) {
		return -1
	}
	return 0
}

// Return values:
//   - -1 Left less than right
//   - 1  Left greater than right
//   - 0  Left equals right
func CompareInts(l, r int) int {
	if l < r {
		return 1
	} else if l > r {
		return -1
	}
	return 0
}

func MustParsePacket(packet string) Packet {
	if !strings.HasPrefix(packet, "[") {
		n, err := strconv.Atoi(packet)
		if err != nil {
			panic(err)
		}
		return Packet{N: n, List: nil}
	}

	depth, l := 0, 1
	p := Packet{List: []Packet{}}
	for i, r := range packet {
		if r == '[' {
			depth++
		}

		if depth == 1 && l < i && r == ',' {
			s := packet[l:i]
			p.List = append(p.List, MustParsePacket(s))
			l = i + 1
		}

		if depth == 1 && l < i && r == ']' {
			s := packet[l : i+1]
			if s[0] != '[' {
				s = s[:len(s)-1]
			}
			p.List = append(p.List, MustParsePacket(s))
			l = i + 1
		}

		if r == ']' {
			depth--
		}
	}
	return p
}

// Returns the smaller value
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//go:embed input.txt
var input string

func main() {
	sum := 0

	pairs := strings.Split(input, "\n\n")
	for i, pair := range pairs {
		left, right, found := strings.Cut(pair, "\n")
		if !found {
			panic("No new line separator was found")
		}

		result := ComparePackets(MustParsePacket(left), MustParsePacket(right))
		if result > 0 {
			sum += i + 1
		}
	}

	fmt.Println("Answear:", sum)
}
