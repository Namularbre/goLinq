package linq

// Limit limits the number of element in a slice. It begins at the start of the slice and give the next x elements
func Limit[T any](input []T, x int) []T {
	var output []T

	for iter, elem := range input {
		if iter >= x {
			return output
		}

		output = append(output, elem)
	}

	return output
}
