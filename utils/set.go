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

func (s *Set[T]) Add(t T) {
	if !s.coll[t] {
		s.coll[t] = true
	}
}

func (s *Set[T]) Elements() (result []T) {
	for key := range s.coll {
		result = append(result, key)
	}
	return
}

func (s *Set[T]) Size() int {
	return len(s.Elements())
}

func (s *Set[T]) Contains(e T) (exist bool) {
	for elem := range s.coll {
		if elem == e {
			exist = true
			break
		}
	}
	return
}

func (s *Set[T]) Intersection(other *Set[T]) Set[T] {
	result := MakeSet[T]([]T{})
	for elem := range s.coll {
		if other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}
