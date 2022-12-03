package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	rucksacks := parse(os.Stdin)

	var cSum int
	for i := 0; i < len(rucksacks)/3; i++ {
		a := rucksacks[i*3]
		b := rucksacks[i*3+1]
		c := rucksacks[i*3+2]

		cSum += commonPriority(a, b, c)
	}
	fmt.Println(cSum)

	// sum := 0
	// for _, r := range rucksacks {
	// 	cmps := r.Compartments()
	// 	dupes := duplicates(string(cmps[0]), string(cmps[1]))

	// 	sum += sumPriority(dupes)
	// }
	// fmt.Println("Compartment sum", sum)

	// bs := badgeSum(groups)
	// fmt.Println("badge sum:", bs)
}

type Rucksack struct {
	Items string
}

type Compartment string

func (r *Rucksack) Compartments() []Compartment {
	l := len(r.Items)
	half := int(l / 2.0)

	return []Compartment{Compartment(r.Items[0:half]), Compartment(r.Items[half:l])}
}

func badgeSum(groups [][]Rucksack) int {
	badgeSum := 0
	for _, g := range groups {
		a := g[0]
		b := g[1]
		c := g[2]

		dupes := duplicates(a.Items, duplicates(b.Items, c.Items))
		fmt.Println(dupes)

		sum := sumPriority(dupes)
		fmt.Println(sum)
		badgeSum += sum
	}

	return badgeSum
}

func parse(r io.Reader) []string {
	s := bufio.NewScanner(r)

	var rucksacks []string
	for s.Scan() {
		t := strings.TrimSpace(s.Text())

		rucksacks = append(rucksacks, t)
	}

	return rucksacks
}

func duplicates(a, b string) string {
	setA := toMap(string(a))
	setB := toMap(string(b))

	var dupes []string
	for i := range setA {
		_, exists := setB[i]
		if exists {
			dupes = append(dupes, i)
		}
	}

	return strings.Join(dupes, "")
}

func toMap(s string) map[string]struct{} {
	m := make(map[string]struct{})

	for _, i := range s {
		m[string(i)] = struct{}{}
	}

	return m
}

func sumPriority(s string) int {
	sum := 0
	for _, c := range s {
		sum += priority(string(c))
	}

	return sum
}

func priority(s string) int {
	v := int(s[0])
	if v >= int('A') && v <= int('Z') {
		return v - int('A') + 27
	}

	if v >= int('a') && v <= int('z') {
		return v - int('a') + 1
	}

	panic("invalid val: " + s)
}

func group(rucksacks []Rucksack) [][]Rucksack {
	var groups [][]Rucksack
	var group []Rucksack

	for _, r := range rucksacks {
		if len(group) == 3 {
			groups = append(groups, group)
			group = []Rucksack{}
		}

		group = append(group, r)
	}

	return groups
}

func commonPriority(a, b, c string) int {
	com := common(a, b, c)
	return priority(com)
}

func common(a, b, c string) string {
	setA := toMap(string(a))
	setB := toMap(string(b))
	setC := toMap(string(c))

	var dupes []string
	for i := range setA {
		_, existsInB := setB[i]
		_, existsInC := setC[i]
		if existsInB && existsInC {
			dupes = append(dupes, i)
		}
	}

	return strings.Join(dupes, "")
}
