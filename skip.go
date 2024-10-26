package linq

// Skip skips x element in a slice
func Skip[T any](input []T, x uint) []T {
	var output []T

	skip := true

	for iter, elem := range input {
		if uint(iter) == x {
			skip = false
		}

		if !skip {
			output = append(output, elem)
		}
	}

	return output
}
