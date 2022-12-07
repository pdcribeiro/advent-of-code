package utils

func Map[T, U any](input []T, callback func(T) U) (ret []U) {
	ret = make([]U, len(input))
	for i, v := range input {
		ret[i] = callback(v)
	}
	return
}
