package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFullyContains(t *testing.T) {

	require.Equal(t, true, Section{Start: 2, End: 8}.FullyContains(Section{Start: 3, End: 7}))
	require.Equal(t, false, Section{Start: 2, End: 3}.FullyContains(Section{Start: 3, End: 7}))
	require.Equal(t, true, Section{Start: 2, End: 3}.FullyContains(Section{Start: 2, End: 3}))
}

func TestOverlaps(t *testing.T) {
	require.Equal(t, true, Section{Start: 2, End: 8}.Overlaps(Section{Start: 3, End: 7}))
	require.Equal(t, true, Section{Start: 2, End: 3}.Overlaps(Section{Start: 3, End: 7}))
	require.Equal(t, true, Section{Start: 2, End: 8}.Overlaps(Section{Start: 3, End: 4}))
	require.Equal(t, true, Section{Start: 2, End: 8}.Overlaps(Section{Start: 1, End: 3}))

	require.Equal(t, false, Section{Start: 1, End: 2}.Overlaps(Section{Start: 3, End: 4}))
	require.Equal(t, false, Section{Start: 1, End: 2}.Overlaps(Section{Start: 3, End: 4}))
}
