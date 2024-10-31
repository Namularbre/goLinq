package linq

// Take limits the number of element in a slice. It begins at the start of the slice and give the next x elements
func Take[T any](input []T, x uint) []T {
	var output []T

	for iter, elem := range input {
		if uint(iter) >= x {
			return output
		}

		output = append(output, elem)
	}

	return output
}
