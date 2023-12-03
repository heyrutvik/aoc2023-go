package utils

import "strings"

type Predicate[T any] func(T) bool

func Reverse(s string) string {
	var result string
	for _, c := range s {
		result = string(c) + result
	}
	return result
}

func KeepLeftWhile(pred Predicate[rune], str string) string {
	var buffer strings.Builder
	length := len(str)
	runes := []rune(str)

	for i := length - 1; i >= 0; i-- {
		c := runes[i]
		if pred(c) {
			buffer.WriteRune(c)
		} else {
			break
		}
	}

	return Reverse(buffer.String())
}

func KeepRightWhile(pred Predicate[rune], str string) string {
	var buffer strings.Builder
	for _, c := range str {
		if pred(c) {
			buffer.WriteRune(c)
		} else {
			break
		}
	}
	return buffer.String()
}
