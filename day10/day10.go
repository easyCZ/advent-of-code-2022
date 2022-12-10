package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := parse(os.Stdin)

	strengths := []int{
		20 * ExecuteInstructions(instructions, 20),
		60 * ExecuteInstructions(instructions, 60),
		100 * ExecuteInstructions(instructions, 100),
		140 * ExecuteInstructions(instructions, 140),
		180 * ExecuteInstructions(instructions, 180),
		220 * ExecuteInstructions(instructions, 220),
	}
	fmt.Println("Strengths", strengths)
	sum := 0
	for _, s := range strengths {
		sum += s
	}
	fmt.Println("Sum:", sum)

	drawing := ExecuteInstructionsDrawing(instructions)
	fmt.Println(printDrawing(drawing))
}

func ExecuteInstructionsDrawing(instructions []Ins) [][]int {
	register := 1

	drawingIdx := 0
	drawing := [][]int{
		make([]int, 40, 40),
		make([]int, 40, 40),
		make([]int, 40, 40),
		make([]int, 40, 40),
		make([]int, 40, 40),
		make([]int, 40, 40),
	}

	cycle := 1
	instructionIdx := 0
	for {
		if instructionIdx >= len(instructions) {
			return drawing
		}

		instruction := instructions[instructionIdx]

		// draw
		valueToDraw := 0
		if drawingIdx-1 == register || drawingIdx == register || drawingIdx+1 == register {
			valueToDraw = 1
		}

		modIdx := drawingIdx % 40
		if modIdx-1 == register || modIdx == register || modIdx+1 == register {
			valueToDraw = 1
		}

		col := drawingIdx % 40
		row := drawingIdx / 40
		drawing[row][col] = valueToDraw

		if instruction.Name == "noop" {

		}

		if instruction.Name == "addx" {
			register += instruction.Val
		}

		cycle += 1
		instructionIdx += 1
		drawingIdx = (drawingIdx + 1) % (6 * 40)
	}
}

func printDrawing(drawing [][]int) string {
	b := bytes.NewBufferString("")
	for _, r := range drawing {
		b.WriteString(fmt.Sprintf("%v", r))
		b.WriteString("\n")
	}
	return b.String()
}

func ExecuteInstructions(ins []Ins, toCycle int) int {
	register := 1
	insIdx := 0

	i := 1
	for {
		instruction := ins[insIdx]
		if i == toCycle {
			return register
		}

		switch instruction.Name {
		case "addx":
			register += instruction.Val
		}

		insIdx += 1
		i += 1
	}

	return register
}

type Ins struct {
	Name string
	Val  int
}

func parse(r io.Reader) []Ins {
	s := bufio.NewScanner(r)

	var instructions []Ins
	for s.Scan() {
		t := strings.TrimSpace(s.Text())

		parts := strings.Split(t, " ")

		switch parts[0] {
		case "addx":
			instructions = append(instructions, Ins{
				Name: "noop",
				Val:  0,
			}, Ins{
				Name: parts[0],
				Val:  mustParseInt(parts[1]),
			})
		case "noop":
			instructions = append(instructions, Ins{
				Name: parts[0],
			})
		default:
			panic("unknown op" + parts[0])
		}
	}

	return instructions
}

func mustParseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("failed to parse int")
	}
	return int(n)
}
