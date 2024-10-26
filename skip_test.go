package linq

import (
	"slices"
	"testing"
)

func TestSkip(t *testing.T) {
	slice := PrepareTestSource()
	assertRes := []TestSource{
		{
			Username: "Alex",
			Age:      22,
		},
		{
			Username: "Steve",
			Age:      28,
		},
		{
			Username: "Alice",
			Age:      76,
		},
		{
			Username: "Bob",
			Age:      36,
		},
	}
	res := Skip(slice, 2)

	if !slices.Equal(assertRes, res) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}

func TestSkipWithXBiggerThanLenOfSlice(t *testing.T) {
	slice := PrepareTestSource()
	var assertRes []TestSource
	skipLenOfSlice := uint(len(slice))
	res := Skip(slice, skipLenOfSlice)

	if !slices.Equal(assertRes, res) || len(res) != 0 {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}

func TestSkipZero(t *testing.T) {
	slice := PrepareTestSource()
	assertRes := slice
	res := Skip(slice, 0)

	if !slices.Equal(assertRes, res) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}
