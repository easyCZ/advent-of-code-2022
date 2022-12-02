package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	rows := parse(os.Stdin)

	var sum int
	for _, r := range rows {
		round := Round{Opponent: sToHand(r.Opponent), Me: sToHand(r.Me)}
		sum += round.Score()
	}

	fmt.Println("Total score:", sum)

	var sumChosen int
	for _, r := range rows {
		choiced := r.ToReachOutcome()
		sumChosen += choiced.Score()
	}
	fmt.Println("To reach outcome:", sumChosen)
}

const (
	Rock     = "rock"
	Paper    = "paper"
	Scissors = "scissors"
)

type Row struct {
	Opponent string
	Me       string
}

type Round struct {
	Opponent string
	Me       string
}

func (r *Round) Score() int {
	var me string
	switch r.Me {
	case "X", Rock:
		me = Rock
	case "Y", Paper:
		me = Paper
	case "Z", Scissors:
		me = Scissors
	}

	if r.Opponent == me {
		return 3 + r.myHandScore(me)
	}

	if r.Opponent == Rock {
		if me == Paper {
			return 6 + r.myHandScore(me)
		}
		if me == Scissors {
			return 0 + r.myHandScore(me)
		}
	}

	if r.Opponent == Paper {
		if me == Rock {
			return 0 + r.myHandScore(me)
		}
		if me == Scissors {
			return 6 + r.myHandScore(me)
		}
	}

	if r.Opponent == Scissors {
		if me == Rock {
			return 6 + r.myHandScore(me)
		}
		if me == Paper {
			return 0 + r.myHandScore(me)
		}
	}

	panic(fmt.Sprintf("unknown score for %v", r))
}

func (r *Row) ToReachOutcome() Round {
	opponent := sToHand(r.Opponent)

	switch r.Me {
	case "Y":
		// draw
		return Round{Opponent: opponent, Me: opponent}
	case "Z":
		// win
		switch opponent {
		case Rock:
			return Round{Opponent: opponent, Me: Paper}
		case Paper:
			return Round{Opponent: opponent, Me: Scissors}
		case Scissors:
			return Round{Opponent: opponent, Me: Rock}
		}
	case "X":
		// lose
		switch opponent {
		case Rock:
			return Round{Opponent: opponent, Me: Scissors}
		case Paper:
			return Round{Opponent: opponent, Me: Rock}
		case Scissors:
			return Round{Opponent: opponent, Me: Paper}
		}
	}

	panic("unknown outcome")
}

func (r *Round) myHandScore(s string) int {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}

	panic("unknown my hand value")
}

func sToHand(s string) string {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}

	panic("unknwon s " + s)
}

func parse(r io.Reader) []Row {
	s := bufio.NewScanner(r)

	var rows []Row
	for s.Scan() {
		s := strings.TrimSpace(s.Text())
		vals := strings.Split(s, " ")
		if len(vals) != 2 {
			panic(fmt.Sprintf("failed to parse row %s", s))
		}

		r := Row{Opponent: vals[0], Me: vals[1]}
		rows = append(rows, r)
	}

	return rows
}
