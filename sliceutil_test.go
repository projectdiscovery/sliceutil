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

func TestDedupeStrings(t *testing.T) {
	test := []string{"a", "a", "b", "b"}
	// converts back
	res := Dedupe(test)
	require.Equal(t, []string{"a", "b"}, res, "strings not deduped correctly")
}

func TestDedupeInt(t *testing.T) {
	test := []int{1, 2, 2, 3}
	// converts back
	res := DedupeInt(test)
	require.Equal(t, []int{1, 2, 3}, res, "ints not deduped correctly")
}

func TestPickRandom(t *testing.T) {
	test := []string{"a", "b"}
	// converts back
	res := PickRandom(test)
	require.Contains(t, test, res, "element was not picked correctly")
}

func TestContains(t *testing.T) {
	test1 := []string{"a", "b"}
	test2 := "a"
	// converts back
	res := Contains(test1, test2)
	require.True(t, res, "unexptected result")
}

func TestContainsItems(t *testing.T) {
	test1 := []string{"a", "b", "c"}
	test2 := []string{"a", "b"}
	// converts back
	res := ContainsItems(test1, test2)
	require.True(t, res, "unexptected result")
}
