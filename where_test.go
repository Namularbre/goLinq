package linq

import (
	"slices"
	"testing"
)

func TestWhereWithInt(t *testing.T) {
	slice := []int{2, 4, 6, 7, 8}
	assertRes := []int{2, 4, 6, 8}
	res := Where(slice, func(x int) bool {
		return x%2 == 0
	})

	if !slices.Equal(assertRes, res) {
		t.Fatalf("Error. Res should be %v but got %v", assertRes, res)
	}
}

func TestWhereWithStrings(t *testing.T) {
	slice := PrepareTestSource()
	assertRes := []TestSource{
		{
			Username: "Alex",
			Age:      22,
		},
	}
	res := Where(slice, func(source TestSource) bool {
		return source.Username == "Alex"
	})
	if !slices.Equal(res, assertRes) {
		t.Fatalf("Error. Res should be %v but got %v", assertRes, res)
	}
}

func TestWhereWhenNoResult(t *testing.T) {
	slice := PrepareTestSource()
	var assertRes []TestSource
	res := Where(slice, func(source TestSource) bool {
		return source.Age == -2000
	})

	if len(res) != 0 || !slices.Equal(assertRes, res) {
		t.Fatalf("Error. Res should be %v but got %v", assertRes, res)
	}
}
