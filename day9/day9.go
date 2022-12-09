package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	moves := parse(os.Stdin)

	rope := Rope{
		Head: Point{
			X: 0,
			Y: 0,
		},
		Tail: Point{
			X: 0,
			Y: 0,
		},
	}
	visited := traverse(rope, moves)
	fmt.Println("Tail visited: ", len(visited))

	chain := []Point{
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
	}

	visited, _ = TraverseChain(chain, moves)
	visitedLong := dedupePoints(visited)
	fmt.Println("long tail visited", len(visitedLong))

}

func TraverseChain(chain []Point, moves []Move) ([]Point, []Point) {
	var seen []Point

	for _, m := range moves {
		fmt.Println("Move", m)
		for i := 0; i < m.Value; i++ {
			fmt.Println("Step", i)
			head := SingleMoveHead(chain[0], m.Direction)
			chain[0] = head

			for j := 1; j < len(chain); j++ {
				tail := chain[j]

				newTail := SingleMoveTail(head, tail)
				chain[j] = newTail

				head = newTail
			}

			seen = append(seen, chain[len(chain)-1])
		}
		fmt.Println(chain)
	}

	return dedupePoints(seen), chain
}

type Rope struct {
	Head Point
	Tail Point
}

func (r *Rope) Move(m Move) []Point {
	var tailPoints []Point

	tailPoints = append(tailPoints, Point{
		X: r.Tail.X,
		Y: r.Tail.Y,
	})

	for i := 0; i < m.Value; i++ {
		r.move(Move{
			Direction: m.Direction,
			Value:     1,
		})

		tailPoints = append(tailPoints, Point{
			X: r.Tail.X,
			Y: r.Tail.Y,
		})
	}

	return tailPoints
}

func (r *Rope) move(m Move) {
	head, tail := SingleMove(r.Head, r.Tail, m)
	r.Head = head
	r.Tail = tail
}

func SingleMoveHead(head Point, direction string) Point {
	// move the head
	switch direction {
	case "U":
		head.Y += 1
	case "D":
		head.Y -= 1
	case "L":
		head.X -= 1
	case "R":
		head.X += 1
	}
	return head
}

func SingleMoveTail(head, tail Point) Point {
	fmt.Println("Move tail", head, tail)
	// move the tail based on the head
	if manhattan(head, tail) > 1 {
		if head.Y == tail.Y {
			// on the same Y axis
			// move on the X axis

			polarity := head.X - tail.X
			if polarity > 0 {
				tail.X = tail.X + 1
			} else {
				tail.X = tail.X - 1
			}
		}

		if head.X == tail.X {
			polarity := head.Y - tail.Y
			if polarity < 0 {
				tail.Y = tail.Y - 1
			} else {
				tail.Y = tail.Y + 1
			}
		}

		// they are not on the same axis
		if head.X-tail.X == 2 && head.Y-tail.Y == 1 {
			tail.X += 1
			tail.Y += 1
		}

		if head.X-tail.X == 2 && head.Y-tail.Y == -1 {
			tail.X += 1
			tail.Y -= 1
		}

		if head.X-tail.X == 1 && head.Y-tail.Y == -2 {
			tail.X += 1
			tail.Y -= 1
		}

		if head.X-tail.X == -1 && head.Y-tail.Y == -2 {
			tail.X -= 1
			tail.Y -= 1
		}

		if head.X-tail.X == -2 && head.Y-tail.Y == -1 {
			tail.X -= 1
			tail.Y -= 1
		}

		if head.X-tail.X == -2 && head.Y-tail.Y == 1 {
			tail.X -= 1
			tail.Y += 1
		}

		if head.X-tail.X == -1 && head.Y-tail.Y == 2 {
			tail.X -= 1
			tail.Y += 1
		}

		if head.X-tail.X == 1 && head.Y-tail.Y == 2 {
			tail.X += 1
			tail.Y += 1
		}

		if head.X-tail.X == 2 && head.Y-tail.Y == 2 {
			tail.X += 1
			tail.Y += 1
		}

		if head.X-tail.X == -2 && head.Y-tail.Y == 2 {
			tail.X -= 1
			tail.Y += 1
		}

		if head.X-tail.X == -2 && head.Y-tail.Y == -2 {
			tail.X -= 1
			tail.Y -= 1
		}

		if head.X-tail.X == 2 && head.Y-tail.Y == -2 {
			tail.X += 1
			tail.Y -= 1
		}
	}

	return tail
}

func SingleMove(head, tail Point, m Move) (Point, Point) {
	// move the head
	switch m.Direction {
	case "U":
		head.Y += m.Value
	case "D":
		head.Y -= m.Value
	case "L":
		head.X -= m.Value
	case "R":
		head.X += m.Value
	}

	// move the tail based on the head
	if manhattan(head, tail) > 1 {
		if head.Y == tail.Y {
			// on the same Y axis
			// move on the X axis

			polarity := head.X - tail.X
			if polarity > 0 {
				tail.X = tail.X + 1
			} else {
				tail.X = tail.X - 1
			}
		}

		if head.X == tail.X {
			polarity := head.Y - tail.Y
			if polarity < 0 {
				tail.Y = tail.Y - 1
			} else {
				tail.Y = tail.Y + 1
			}
		}

		// they are not on the same axis
		if head.X-tail.X == 2 && head.Y-tail.Y == 1 {
			tail.X += 1
			tail.Y += 1
		}

		if head.X-tail.X == 2 && head.Y-tail.Y == -1 {
			tail.X += 1
			tail.Y -= 1
		}

		if head.X-tail.X == 1 && head.Y-tail.Y == -2 {
			tail.X += 1
			tail.Y -= 1
		}

		if head.X-tail.X == -1 && head.Y-tail.Y == -2 {
			tail.X -= 1
			tail.Y -= 1
		}

		if head.X-tail.X == -2 && head.Y-tail.Y == -1 {
			tail.X -= 1
			tail.Y -= 1
		}

		if head.X-tail.X == -2 && head.Y-tail.Y == 1 {
			tail.X -= 1
			tail.Y += 1
		}

		if head.X-tail.X == -1 && head.Y-tail.Y == 2 {
			tail.X -= 1
			tail.Y += 1
		}

		if head.X-tail.X == 1 && head.Y-tail.Y == 2 {
			tail.X += 1
			tail.Y += 1
		}
	}

	return head, tail
}

type Point struct {
	X int
	Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%d, %d)", p.X, p.Y)
}

func traverse(init Rope, moves []Move) []Point {
	var visited []Point

	for _, m := range moves {
		visited = append(visited, init.Move(m)...)
	}

	return dedupePoints(visited)
}

func dedupePoints(ps []Point) []Point {
	deduped := make(map[Point]struct{})
	for _, v := range ps {
		deduped[v] = struct{}{}
	}

	var collected []Point
	for point := range deduped {
		collected = append(collected, point)
	}

	return collected
}

type Move struct {
	Direction string
	Value     int
}

func parse(r io.Reader) []Move {
	s := bufio.NewScanner(r)

	var moves []Move
	for s.Scan() {
		t := strings.TrimSpace(s.Text())

		moves = append(moves, parseMove(t))
	}

	return moves
}

func parseMove(s string) Move {
	parts := strings.Split(s, " ")

	n, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic("failed to parse input")
	}

	return Move{
		Direction: parts[0],
		Value:     int(n),
	}
}

func manhattan(a, b Point) int {
	return int(math.Abs(float64(b.X-a.X))) + int(math.Abs(float64(b.Y-a.Y)))
}

func distance(a, b int) int {
	return int(math.Abs(float64(b - a)))
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
