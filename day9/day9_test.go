package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMove(t *testing.T) {
	initial := Rope{
		Head: Point{
			X: 0,
			Y: 0,
		},
		Tail: Point{
			X: 0,
			Y: 0,
		},
	}

	initial.Move(Move{
		Direction: "R",
		Value:     4,
	})
	require.Equal(t, Rope{
		Head: Point{
			X: 4,
			Y: 0,
		},
		Tail: Point{
			X: 3,
			Y: 0,
		},
	}, initial)

	initial.Move(Move{
		Direction: "U",
		Value:     4,
	})
	require.Equal(t, Rope{
		Head: Point{
			X: 4,
			Y: 4,
		},
		Tail: Point{
			X: 4,
			Y: 3,
		},
	}, initial)

	initial.Move(Move{
		Direction: "L",
		Value:     3,
	})
	require.Equal(t, Rope{
		Head: Point{
			X: 1,
			Y: 4,
		},
		Tail: Point{
			X: 2,
			Y: 4,
		},
	}, initial)

	initial.Move(Move{
		Direction: "D",
		Value:     1,
	})
	require.Equal(t, Rope{
		Head: Point{
			X: 1,
			Y: 3,
		},
		Tail: Point{
			X: 2,
			Y: 4,
		},
	}, initial)

	initial.Move(Move{
		Direction: "R",
		Value:     4,
	})
	require.Equal(t, Rope{
		Head: Point{
			X: 5,
			Y: 3,
		},
		Tail: Point{
			X: 4,
			Y: 3,
		},
	}, initial)

	initial.Move(Move{
		Direction: "D",
		Value:     1,
	})
	require.Equal(t, Rope{
		Head: Point{
			X: 5,
			Y: 2,
		},
		Tail: Point{
			X: 4,
			Y: 3,
		},
	}, initial)

	initial.Move(Move{
		Direction: "L",
		Value:     5,
	})
	require.Equal(t, Rope{
		Head: Point{
			X: 0,
			Y: 2,
		},
		Tail: Point{
			X: 1,
			Y: 2,
		},
	}, initial)

	initial.Move(Move{
		Direction: "R",
		Value:     2,
	})
	require.Equal(t, Rope{
		Head: Point{
			X: 2,
			Y: 2,
		},
		Tail: Point{
			X: 1,
			Y: 2,
		},
	}, initial)
}

// func TestRopeChainMove(t *testing.T) {
// 	initial := []Point{
// 		{X: 0, Y: 0},
// 		{X: 0, Y: 0},
// 	}

// 	initial.MoveMulti(Move{
// 		Direction: "R",
// 		Value:     4,
// 	})
// 	require.Equal(t, []Point{
// 		{X: 4, Y: 0},
// 		{X: 3, Y: 0},
// 	}, initial.Points)

// 	initial.MoveMulti(Move{
// 		Direction: "U",
// 		Value:     4,
// 	})
// 	require.Equal(t, []Point{
// 		{X: 4, Y: 4},
// 		{X: 4, Y: 3},
// 	}, initial.Points)

// 	initial.MoveMulti(Move{
// 		Direction: "L",
// 		Value:     3,
// 	})
// 	require.Equal(t, []Point{
// 		{X: 1, Y: 4},
// 		{X: 2, Y: 4},
// 	}, initial.Points)

// 	initial.MoveMulti(Move{
// 		Direction: "D",
// 		Value:     1,
// 	})
// 	require.Equal(t, []Point{
// 		{X: 1, Y: 3},
// 		{X: 2, Y: 4},
// 	}, initial.Points)

// 	initial.MoveMulti(Move{
// 		Direction: "R",
// 		Value:     4,
// 	})
// 	require.Equal(t, []Point{
// 		{X: 5, Y: 3},
// 		{X: 4, Y: 3},
// 	}, initial.Points)

// 	initial.MoveMulti(Move{
// 		Direction: "D",
// 		Value:     1,
// 	})
// 	require.Equal(t, []Point{
// 		{X: 5, Y: 2},
// 		{X: 4, Y: 3},
// 	}, initial.Points)

// 	initial.MoveMulti(Move{
// 		Direction: "L",
// 		Value:     5,
// 	})
// 	require.Equal(t, []Point{
// 		{X: 0, Y: 2},
// 		{X: 1, Y: 2},
// 	}, initial.Points)

// 	initial.MoveMulti(Move{
// 		Direction: "R",
// 		Value:     2,
// 	})
// 	require.Equal(t, []Point{
// 		{X: 2, Y: 2},
// 		{X: 1, Y: 2},
// 	}, initial.Points)
// }

func TestLong(t *testing.T) {
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

	moves := []Move{
		{Direction: "R", Value: 5},
		{Direction: "U", Value: 8},
		{Direction: "L", Value: 8},
		{Direction: "D", Value: 3},
		{Direction: "R", Value: 17},
		{Direction: "D", Value: 10},
		{Direction: "L", Value: 25},
		{Direction: "U", Value: 20},
	}

	visisted, _ := TraverseChain(chain, moves)
	require.Equal(t, 36, len(visisted))
}

func TestLongSingle(t *testing.T) {
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

	visited, chain := TraverseChain(chain, []Move{
		{Direction: "R", Value: 5},
	})
	require.Equal(t, []Point{
		{X: 5, Y: 0},
		{X: 4, Y: 0},
		{X: 3, Y: 0},
		{X: 2, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
		{X: 0, Y: 0},
	}, chain)
	require.Equal(t, []Point{
		{X: 0, Y: 0},
	}, visited)

	visited, chain = TraverseChain(chain, []Move{
		{Direction: "U", Value: 8},
	})
	require.Equal(t, []Point{
		{X: 5, Y: 8},
		{X: 5, Y: 7},
		{X: 5, Y: 6},
		{X: 5, Y: 5},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 3},
		{X: 2, Y: 2},
		{X: 1, Y: 1},
		{X: 0, Y: 0},
	}, chain)
	require.Equal(t, []Point{
		{X: 0, Y: 0},
	}, visited)

}
