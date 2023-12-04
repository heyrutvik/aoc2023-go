package utils

func Map[A any, B any](f func(A) B, as []A) (result []B) {
	for _, a := range as {
		result = append(result, f(a))
	}
	return
}
