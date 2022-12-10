package main

import (
	"os"
)

func main() {
	instructions := parse(os.StdIn)
}

type Instruction struct {
	Op string
	Value int
}

func parse(r io.Reader) []Instruction {
	s := bufio.NewScanner(r)
}
