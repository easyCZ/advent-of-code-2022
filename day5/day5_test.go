package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRearrange(t *testing.T) {
	initial := []string{
		reverse("NZ"),
		reverse("DCM"),
		reverse("P"),
	}
	stack := Rearrange(initial, Move{
		Count: 1,
		From:  1,
		To:    0,
	})
	require.Equal(t, []string{
		reverse("DNZ"),
		reverse("CM"),
		reverse("P"),
	}, stack)

	stack = Rearrange(stack, Move{
		Count: 3,
		From:  0,
		To:    2,
	})
	require.Equal(t, []string{
		reverse(""),
		reverse("CM"),
		reverse("ZNDP"),
	}, stack)

	stack = Rearrange(stack, Move{
		Count: 2,
		From:  1,
		To:    0,
	})
	require.Equal(t, []string{
		reverse("MC"),
		reverse(""),
		reverse("ZNDP"),
	}, stack)

	stack = Rearrange(stack, Move{
		Count: 1,
		From:  0,
		To:    1,
	})
	require.Equal(t, []string{
		reverse("C"),
		reverse("M"),
		reverse("ZNDP"),
	}, stack)
}
