package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPruneEmptyStrings(t *testing.T) {
	test := []string{"a", "", "", "b"}
	// converts back
	res := PruneEmptyStrings(test)
	require.Equal(t, []string{"a", "b"}, res, "strings not pruned correctly")
}

func TestPruneEqual(t *testing.T) {
	test := []string{"a", "", "", "b"}
	// converts back
	res := PruneEqual(test, "b")
	require.Equal(t, []string{"a", "", ""}, res, "strings not pruned correctly")
}

func TestSeen(t *testing.T) {
	test := []string{"a", "a", "b", "b"}
	// converts back
	res := Dedupe(test)
	require.Equal(t, []string{"a", "b"}, res, "strings not deduped correctly")
}
