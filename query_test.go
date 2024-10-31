package linq

import (
	"slices"
	"testing"
)

func TestQuery_Select(t *testing.T) {
	input := PrepareTestSource()
	query := NewQuery[TestSource, string](input)

	query.Select(func(user TestSource) string {
		return user.Username
	})

	assertRes := []string{
		"Dupont",
		"Dupond",
		"Alex",
		"Steve",
		"Alice",
		"Bob",
	}
	res := query.ToSlice()

	if !slices.Equal(assertRes, res) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}
