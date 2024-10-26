package linq

import (
	"encoding/json"
	"fmt"
	"os"
)

// TestSource The goal of this type is to test with custom structs in linq
type TestSource struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func PrepareTestSource() []TestSource {
	return []TestSource{
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
