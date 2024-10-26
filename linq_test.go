package linq

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"testing"
)

// TestSource The goal of this type is to test custom structs in linq
type TestSource struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func loadTestsSource() []TestSource {
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	var tests []TestSource

	err = json.Unmarshal(data, &tests)
	if err != nil {
		panic(err)
	}

	fmt.Printf("test set len %d \n", len(tests))

	return tests
}

func TestWhere(t *testing.T) {
	slice := []int{2, 4, 6, 7, 8}
	assertRes := []int{2, 4, 6, 8}
	res := Where(slice, func(x int) bool {
		return x%2 == 0
	})

	if !slices.Equal(assertRes, res) || len(assertRes) != len(res) {
		t.Fatalf("Error. Res should be %v but got %v", assertRes, res)
	}
}

func TestSelect(t *testing.T) {
	slice := []TestSource{
		{
			Username: "Alex",
			Age:      22,
		},
		{
			Username: "Steve",
			Age:      28,
		},
	}
	assertRes := []string{
		"Alex",
		"Steve",
	}

	res := Select[TestSource, string](slice, func(source TestSource) string {
		return source.Username
	})

	if !slices.Equal(assertRes, res) {
		t.Fatalf("Error, Res should be %v but got %v", assertRes, res)
	}
}

func TestLimit(t *testing.T) {
	slice := []TestSource{
		{
			Username: "Dupont",
			Age:      60,
		},
		{
			Username: "Dupond",
			Age:      44,
		},
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
	assertOutput := []TestSource{
		{
			Username: "Dupont",
			Age:      60,
		},
		{
			Username: "Dupond",
			Age:      44,
		},
	}
	output := Limit(slice, 2)

	if len(output) != 2 && !slices.Equal(output, assertOutput) {
		t.Fatalf("Error, Res should be %v but got %v", assertOutput, output)
	}
}

func BenchmarkSelect(b *testing.B) {
	data := loadTestsSource()

	var ages []int

	for i := 0; i < b.N; i++ {
		ages = Select[TestSource, int](data, func(source TestSource) int {
			return source.Age
		})
	}

	fmt.Printf("result size: %d\n", len(ages))
}

func BenchmarkWhere(b *testing.B) {
	data := loadTestsSource()

	var res []TestSource

	for i := 0; i < b.N; i++ {
		res = Where[TestSource](data, func(elem TestSource) bool {
			return elem.Age > 20 && elem.Age < 25
		})
	}

	fmt.Printf("result size: %d\n", len(res))
}
