package utils

type Set[T comparable] struct {
	coll map[T]bool
}

func MakeSet[T comparable](elements []T) Set[T] {
	coll := make(map[T]bool)
	for _, e := range elements {
		if !coll[e] {
			coll[e] = true
		}
	}
	return Set[T]{coll}
}

func (s Set[T]) Add(t T) {
	if !s.coll[t] {
		s.coll[t] = true
	}
}

func (s Set[T]) Elements() (result []T) {
	for key := range s.coll {
		result = append(result, key)
	}
	return
}
