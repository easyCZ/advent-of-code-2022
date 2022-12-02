package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	var capacities [][]int

	var current []int
	for s.Scan() {
		t := s.Text()

		s := strings.TrimSpace(t)
		if s == "" {
			// start a new entry
			capacities = append(capacities, current)
			current = []int{}
			continue
		}

		val, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("invalid val %s", s))
		}

		current = append(current, val)
	}

	var carrying []int
	for _, c := range capacities {
		sum := 0
		for _, i := range c {
			sum += i

		}
		carrying = append(carrying, sum)
	}

	maxCarryingOne := 0
	for _, c := range carrying {
		if c > maxCarryingOne {
			maxCarryingOne = c
		}
	}

	fmt.Println("Max: ", maxCarryingOne)

	sort.Ints(carrying)
	sum3 := carrying[len(carrying)-1] + carrying[len(carrying)-2] + carrying[len(carrying)-3]
	fmt.Println("Max 3: ", sum3)
}
