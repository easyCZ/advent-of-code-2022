package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	pairs := parse(os.Stdin)

	fullyContainsCount := 0
	for _, p := range pairs {
		if p.A.FullyContains(p.B) || p.B.FullyContains(p.A) {
			fullyContainsCount += 1
		}
	}

	fmt.Println("Fully contains: ", fullyContainsCount)

	overlapsCount := 0
	for _, p := range pairs {
		if p.A.Overlaps(p.B) || p.B.Overlaps(p.A) {
			overlapsCount += 1
		}
	}
	fmt.Println("Overlaps count: ", overlapsCount)
}

type Pair struct {
	A Section
	B Section
}

type Section struct {
	Start int
	End   int
}

func (s Section) FullyContains(another Section) bool {
	if another.Start >= s.Start && another.Start <= s.End {
		if another.End >= s.Start && another.End <= s.End {
			return true
		}
	}

	return false
}

func (s Section) Overlaps(another Section) bool {
	if s.FullyContains(another) {
		return true
	}

	// left
	if another.End >= s.Start && another.End <= s.End {
		return true
	}

	// right
	if another.Start <= s.End && another.End >= s.Start {
		return true
	}

	return false
}

func parse(r io.Reader) []Pair {
	s := bufio.NewScanner(r)

	var pairs []Pair
	for s.Scan() {
		t := strings.TrimSpace(s.Text())

		parts := strings.Split(t, ",")

		a := parseSection(parts[0])
		b := parseSection(parts[1])

		pairs = append(pairs, Pair{
			A: a,
			B: b,
		})
	}

	return pairs
}

func parseSection(s string) Section {
	parts := strings.Split(s, "-")

	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid start %s", s))
	}

	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid end %s", s))
	}

	return Section{
		Start: int(start),
		End:   int(end),
	}
}
