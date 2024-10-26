package linq

import (
	"slices"
	"testing"
)

func TestSelectWhere(t *testing.T) {
	slice := PrepareTestSource()
	assertRes := []int{60, 44, 76, 36}
	res := Where(
		Select[TestSource, int](slice, func(elem TestSource) int {
			return elem.Age
		}),
		func(source int) bool {
			return source > 30
		})

	if !slices.Equal(assertRes, res) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}
