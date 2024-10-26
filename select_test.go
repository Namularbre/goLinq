package linq

import (
	"slices"
	"testing"
)

func TestSelect(t *testing.T) {
	slice := PrepareTestSource()
	assertRes := []string{
		"Dupont",
		"Dupond",
		"Alex",
		"Steve",
		"Alice",
		"Bob",
	}

	res := Select[TestSource, string](slice, func(source TestSource) string {
		return source.Username
	})

	if !slices.Equal(assertRes, res) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}

func TestSelectWhereNoResult(t *testing.T) {
	var slice []TestSource
	var assertRes []string
	res := Select[TestSource, string](slice, func(source TestSource) string {
		return source.Username
	})

	if len(res) != 0 || !slices.Equal(assertRes, res) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}
