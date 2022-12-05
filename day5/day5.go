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

	stacks := []string{
		reverse("PGRN"),
		reverse("CDGFLBTJ"),
		reverse("VSM"),
		reverse("PZCRSL"),
		reverse("QDWCVLSP"),
		reverse("SMDWNTC"),
		reverse("PWGDH"),
		reverse("VMCSHPLZ"),
		reverse("ZGWLFPR"),
	}

	moves := parse(os.Stdin)

	for _, m := range moves {
		stacks = Rearrange(stacks, m)
	}

	var tops string
	for _, s := range stacks {
		tops += Top(s)
	}

	fmt.Println("CrateMover 9000:", tops)

	stacks = []string{
		reverse("PGRN"),
		reverse("CDGFLBTJ"),
		reverse("VSM"),
		reverse("PZCRSL"),
		reverse("QDWCVLSP"),
		reverse("SMDWNTC"),
		reverse("PWGDH"),
		reverse("VMCSHPLZ"),
		reverse("ZGWLFPR"),
	}

	for _, m := range moves {
		stacks = RearrangeMany(stacks, m)
	}

	tops = ""
	for _, s := range stacks {
		tops += Top(s)
	}

	fmt.Println("CrateMover 9001:", tops)
}

func Take(stack string, n int) (taken string, updatedStack string) {
	toTake := stack[len(stack)-n:]
	return reverse(toTake), stack[0 : len(stack)-n]
}

func TakeMany(stack string, n int) (taken string, updatedStack string) {
	toTake := stack[len(stack)-n:]
	return toTake, stack[0 : len(stack)-n]
}

func Add(stack string, items string) (newStack string) {
	return stack + items
}

func Top(stack string) string {
	return string(stack[len(stack)-1])
}

func Rearrange(stacks []string, move Move) []string {
	from := stacks[move.From]
	to := stacks[move.To]

	var lifted string
	lifted, stacks[move.From] = Take(from, move.Count)

	stacks[move.To] = Add(to, lifted)

	return stacks
}

func RearrangeMany(stacks []string, move Move) []string {
	from := stacks[move.From]
	to := stacks[move.To]

	var lifted string
	lifted, stacks[move.From] = TakeMany(from, move.Count)

	stacks[move.To] = Add(to, lifted)

	return stacks
}

type Move struct {
	Count int
	From  int
	To    int
}

func ParseMove(s string) Move {
	s = strings.ReplaceAll(s, "move ", "")
	s = strings.ReplaceAll(s, " from ", ",")
	s = strings.ReplaceAll(s, " to ", ",")

	parts := strings.Split(s, ",")
	count := parts[0]
	start := parts[1]
	end := parts[2]

	return Move{Count: mustConvertToInt(count), From: mustConvertToInt(start) - 1, To: mustConvertToInt(end) - 1}
}

func mustConvertToInt(s string) int {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse %s", s))
	}
	return int(v)
}

func parse(r io.Reader) []Move {
	s := bufio.NewScanner(r)

	var moves []Move
	for s.Scan() {
		t := strings.TrimSpace(s.Text())
		m := ParseMove(t)
		moves = append(moves, m)
	}

	return moves
}

func reverse(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
