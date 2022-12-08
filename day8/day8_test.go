package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVisibleTrees(t *testing.T) {
	grid := [][]int{
		parseRow("30373"),
		parseRow("25512"),
		parseRow("65332"),
		parseRow("33549"),
		parseRow("35390"),
	}

	require.Equal(t, 21, visibleFromOutside(grid))

}

func TestVisibleLeft(t *testing.T) {
	grid := [][]int{
		parseRow("30373"),
		parseRow("25512"),
		parseRow("65332"),
		parseRow("33549"),
		parseRow("35390"),
	}
	visibles := visibleLeft(grid)
	require.Equal(t, []Visible{
		{
			Row:       1,
			Col:       1,
			Value:     5,
			Direction: "left",
		},
		{
			Row:       3,
			Col:       2,
			Value:     5,
			Direction: "left",
		},
	}, visibles)
}

func TestVisibleRight(t *testing.T) {
	grid := [][]int{
		parseRow("30373"),
		parseRow("25512"),
		parseRow("65332"),
		parseRow("33549"),
		parseRow("35390"),
	}
	visibles := visibleRight(grid)
	require.Equal(t, []Visible{
		{
			Row:       1,
			Col:       2,
			Value:     5,
			Direction: "right",
		},
		{
			Row:       2,
			Col:       3,
			Value:     3,
			Direction: "right",
		},
		{
			Row:       2,
			Col:       1,
			Value:     5,
			Direction: "right",
		},
	}, visibles)
}

func TestVisibleTop(t *testing.T) {
	grid := [][]int{
		parseRow("30373"),
		parseRow("25512"),
		parseRow("65332"),
		parseRow("33549"),
		parseRow("35390"),
	}
	visibles := visibleTop(grid)
	require.Equal(t, []Visible{
		{
			Row:       1,
			Col:       1,
			Value:     5,
			Direction: "top",
		},
		{
			Row:       1,
			Col:       2,
			Value:     5,
			Direction: "top",
		},
	}, visibles)
}

func TestVisibleBot(t *testing.T) {
	grid := [][]int{
		parseRow("30373"),
		parseRow("25512"),
		parseRow("65332"),
		parseRow("33549"),
		parseRow("35390"),
	}
	visibles := visibleBottom(grid)
	require.Equal(t, []Visible{
		{
			Row:       3,
			Col:       2,
			Value:     5,
			Direction: "bottom",
		},
	}, visibles)
}
