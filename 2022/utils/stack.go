package utils

type Stack[T comparable] []T

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (ret T) {
	i := len(*s) - 1
	ret = (*s)[i]
	*s = (*s)[:i]
	return
}
