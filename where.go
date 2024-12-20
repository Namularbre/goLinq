package linq

// Where is used to make a projection of a new slice where all elements match the filter
func Where[T any](input []T, predicate func(T) bool) []T {
	var output []T

	for _, inputElem := range input {
		if predicate(inputElem) {
			output = append(output, inputElem)
		}
	}

	return output
}
