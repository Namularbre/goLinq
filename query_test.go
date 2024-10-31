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

func TestQuery_First(t *testing.T) {
	input := PrepareTestSource()
	query := NewQuery[TestSource, TestSource](input)

	output, err := query.First()

	if err != nil {
		t.Fatalf("Error, the test shouldn't return an error: %v", err)
	}

	if *output == input[0] {
		t.Fatalf("Error, ouput should be %v but is %v", input[0], output)
	}
}

func TestQuery_FirstWithError(t *testing.T) {
	input := PrepareTestSource()
	query := NewQuery[TestSource, string](input)
	query.Where(func(user TestSource) bool {
		return user.Username == "é'gfdv"
	})

	_, err := query.First()

	if err == nil || err.Error() != "no result in the query output" {
		t.Fatalf("Error, we should have an error here but got %v", err)
	}
}

func TestQuery_Where(t *testing.T) {
	input := PrepareTestSource()
	query := NewQuery[TestSource, string](input)
	query.Where(func(user TestSource) bool {
		return user.Username == "Dupont"
	})

	output := query.FirstOrNil()

	if output == nil {
		t.Fatalf("Error, we should have a value here")
	}
}

func TestQuery_WhereNoResult(t *testing.T) {
	input := PrepareTestSource()
	query := NewQuery[TestSource, string](input)
	query.Where(func(user TestSource) bool {
		return user.Username == "69"
	})

	output := query.FirstOrNil()

	if output != nil {
		t.Fatalf("Error, we should have an error here")
	}
}

func TestQuery_FirstOrDefaultForDefault(t *testing.T) {
	input := PrepareTestSource()
	query := NewQuery[TestSource, TestSource](input)
	query.Where(func(user TestSource) bool {
		return user.Username == "69"
	})
	assert := &TestSource{Username: "Michel", Age: 69}

	output := query.FirstOrDefault(assert)

	if output != assert {
		t.Fatalf("Error, ouput should be %v but is %v", input[0], output)
	}
}

func TestQuery_FirstOrDefault(t *testing.T) {
	input := PrepareTestSource()
	query := NewQuery[TestSource, TestSource](input)
	query.Where(func(user TestSource) bool {
		return user.Username == "Dupont"
	})

	assert := &TestSource{Username: "Michel", Age: 69}

	output := query.FirstOrDefault(assert)

	if output == assert {
		t.Fatalf("Error, ouput should not be %v but is %v", assert, output)
	}
}
