package utils

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) (ret *Set[T]) {
	ret = &Set[T]{}
	for _, item := range items {
		ret.Add(item)
	}
	return
}

func (s *Set[T]) Add(item T) {
	(*s)[item] = struct{}{}
}

func (s *Set[T]) Has(item T) (ret bool) {
	_, ret = (*s)[item]
	return
}

func (s *Set[T]) Intersection(other *Set[T]) (ret *Set[T]) {
	ret = NewSet[T]()
	for item := range *other {
		if s.Has(item) {
			ret.Add(item)
		}
	}
	return
}

func (s *Set[T]) Size() int {
	return len(*s)
}
