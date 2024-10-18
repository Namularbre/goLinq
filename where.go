package linq

func Where[T any](input []T, filter func(T) bool) []T {
	var output []T

	for _, inputElem := range input {
		ok := make(chan bool)

		go func() {
			defer close(ok)

			ok <- filter(inputElem)
		}()

		if <-ok {
			output = append(output, inputElem)
		}
	}

	return output
}
