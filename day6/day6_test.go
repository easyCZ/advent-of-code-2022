package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindPacketMarker(t *testing.T) {
	require.Equal(t, 7, findMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	require.Equal(t, 5, findMarker("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	require.Equal(t, 6, findMarker("nppdvjthqldpwncqszvftbrmjlhg"))
	require.Equal(t, 10, findMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	require.Equal(t, 11, findMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}

func TestFindMessageMarker(t *testing.T) {
	require.Equal(t, 19, findMessageMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	require.Equal(t, 23, findMessageMarker("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	require.Equal(t, 23, findMessageMarker("nppdvjthqldpwncqszvftbrmjlhg"))
	require.Equal(t, 29, findMessageMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	require.Equal(t, 26, findMessageMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
}
