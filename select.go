package linq

// Select is used to select fields in a struct in an array
func Select[TSource any, TResult any](input []TSource, selector func(TSource) TResult) []TResult {
	var output []TResult

	for _, inputElem := range input {
		output = append(output, selector(inputElem))
	}

	return output
}
