package linq

// Skip skips x element in a slice
func Skip[T any](input []T, x uint) []T {
	if x >= uint(len(input)) {
		return []T{}
	}
	return input[x:]
}
