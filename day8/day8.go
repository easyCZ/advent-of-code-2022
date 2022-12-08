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
	grid := parse(os.Stdin)

	visible := visibleFromOutside(grid)
	fmt.Println("Visible trees:", visible)

	fmt.Println("Max score:", findScore(grid))
}

func findScore(grid [][]int) int {
	max := 0
	for r := range grid {
		for c := range grid[r] {
			v := treesVisible(r, c, grid)

			if v > max {
				max = v
			}
		}
	}
	return max
}

func treesVisible(row, col int, grid [][]int) int {
	return treesDown(row, col, grid) * treesUp(row, col, grid) * treesLeft(row, col, grid) * treesRight(row, col, grid)
}

func treesUp(row, col int, grid [][]int) int {
	visibleTrees := 0
	for rowIdx := row - 1; rowIdx >= 0; rowIdx-- {
		tree := grid[rowIdx][col]

		if tree < grid[row][col] {
			visibleTrees += 1
		} else {
			visibleTrees += 1
			return visibleTrees
		}
	}

	return visibleTrees
}

func treesDown(row, col int, grid [][]int) int {
	visibleTrees := 0
	for rowIdx := row + 1; rowIdx < len(grid); rowIdx++ {
		tree := grid[rowIdx][col]

		if tree < grid[row][col] {
			visibleTrees += 1
		} else {
			visibleTrees += 1
			return visibleTrees
		}
	}

	return visibleTrees
}

func treesRight(row, col int, grid [][]int) int {
	visibleTrees := 0
	for colIdx := col + 1; colIdx < len(grid[0]); colIdx++ {
		tree := grid[row][colIdx]

		if tree < grid[row][col] {
			visibleTrees += 1
		} else {
			visibleTrees += 1
			return visibleTrees
		}
	}

	return visibleTrees
}

func treesLeft(row, col int, grid [][]int) int {
	visibleTrees := 0
	for colIdx := col - 1; colIdx >= 0; colIdx-- {
		tree := grid[row][colIdx]

		if tree < grid[row][col] {
			visibleTrees += 1
		} else {
			visibleTrees += 1
			return visibleTrees
		}
	}

	return visibleTrees
}

type Visible struct {
	Row       int
	Col       int
	Value     int
	Direction string
}

func visibleFromOutside(grid [][]int) int {
	outside := len(grid)*2 + len(grid[0])*2 - 4

	fmt.Println("outside", outside)

	left := visibleLeft(grid)
	right := visibleRight(grid)
	top := visibleTop(grid)
	bot := visibleBottom(grid)

	var visibles []Visible
	visibles = append(visibles, top...)
	visibles = append(visibles, bot...)
	visibles = append(visibles, left...)
	visibles = append(visibles, right...)

	seen := make(map[string]struct{})
	for _, v := range visibles {
		seen[fmt.Sprintf("%d-%d", v.Row, v.Col)] = struct{}{}
	}

	return len(seen) + outside

}

func visibleLeft(grid [][]int) []Visible {
	var visibles []Visible
	// left
	for rowIdx := 1; rowIdx < len(grid)-1; rowIdx++ {
		visible := grid[rowIdx][0]

		for colIdx := 1; colIdx < len(grid[0])-1; colIdx++ {
			v := grid[rowIdx][colIdx]
			if v > visible {
				visible = v
				visibles = append(visibles, Visible{
					Row:       rowIdx,
					Col:       colIdx,
					Value:     v,
					Direction: "left",
				})
			}
		}
	}
	return visibles
}

func visibleRight(grid [][]int) []Visible {
	var visibles []Visible
	// right
	for rowIdx := 1; rowIdx < len(grid)-1; rowIdx++ {
		visible := grid[rowIdx][len(grid[0])-1]

		for colIdx := len(grid[rowIdx]) - 2; colIdx > 0; colIdx-- {
			v := grid[rowIdx][colIdx]
			if v > visible {
				visible = v
				visibles = append(visibles, Visible{
					Row:       rowIdx,
					Col:       colIdx,
					Value:     v,
					Direction: "right",
				})
			}
		}
	}

	return visibles
}

func visibleTop(grid [][]int) []Visible {
	var visibles []Visible
	// top
	for colIdx := 1; colIdx < len(grid[0])-1; colIdx++ {
		visible := grid[0][colIdx]
		for rowIdx := 1; rowIdx < len(grid)-1; rowIdx++ {
			v := grid[rowIdx][colIdx]
			if v > visible {
				visible = v
				visibles = append(visibles, Visible{
					Row:       rowIdx,
					Col:       colIdx,
					Value:     v,
					Direction: "top",
				})
			}
		}
	}
	return visibles
}

func visibleBottom(grid [][]int) []Visible {
	var visibles []Visible
	// bottom
	for colIdx := len(grid[0]) - 2; colIdx > 0; colIdx-- {
		visible := grid[len(grid)-1][colIdx]
		for rowIdx := len(grid) - 2; rowIdx > 0; rowIdx-- {
			v := grid[rowIdx][colIdx]
			if v > visible {
				visible = v
				visibles = append(visibles, Visible{
					Row:       rowIdx,
					Col:       colIdx,
					Value:     v,
					Direction: "bottom",
				})
			}
		}
	}
	return visibles
}

func parse(r io.Reader) [][]int {
	s := bufio.NewScanner(r)

	var rows [][]int
	for s.Scan() {
		t := strings.TrimSpace(s.Text())
		row := parseRow(t)
		rows = append(rows, row)
	}

	return rows
}

func parseRow(s string) []int {
	var row []int
	for _, c := range s {
		n, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			panic(fmt.Sprintf("failed to parse %s", string(c)))
		}

		row = append(row, int(n))
	}

	return row
}
