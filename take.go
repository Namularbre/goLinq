package linq

// Take limits the number of element in a slice. It begins at the start of the slice and give the next x elements
func Take[T any](input []T, x uint) []T {
	if x >= uint(len(input)) {
		return input
	}
	return input[:x]
}
