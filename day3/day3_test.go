package main

import (
	"fmt"
	"testing"
)

func TestBadgeSum(t *testing.T) {
	s := badgeSum([][]Rucksack{
		[]Rucksack{
			Rucksack{
				Items: "vJrwpWtwJgWrhcsFMMfFFhFp",
			},
			Rucksack{
				Items: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			},
			Rucksack{
				Items: "PmmdzqPrVvPwwTWBwg",
			},
		},
		[]Rucksack{
			Rucksack{
				Items: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			},
			Rucksack{
				Items: "ttgJtRGJQctTZtZT",
			},
			Rucksack{
				Items: "CrZsJsPPZsGzwwsLwLmpwMDw",
			},
		},
	})

	fmt.Println(s)
}
