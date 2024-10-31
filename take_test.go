package linq

import (
	"slices"
	"testing"
)

func TestLimit(t *testing.T) {
	slice := PrepareTestSource()
	assertRes := []TestSource{
		{
			Username: "Dupont",
			Age:      60,
		},
		{
			Username: "Dupond",
			Age:      44,
		},
	}
	res := Take(slice, 2)

	if len(res) != 2 && !slices.Equal(res, assertRes) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}

func TestWithLimitBiggerThan(t *testing.T) {
	slice := PrepareTestSource()
	assertRes := slice
	res := Take(slice, 3_000_000)

	if len(res) != len(assertRes) && slices.Equal(assertRes, res) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}
