package linq

func Select[TSource any, TResult any](input []TSource, selector func(TSource) TResult) []TResult {
	var output []TResult

	for _, inputElem := range input {
		output = append(output, selector(inputElem))
	}

	return output
}
