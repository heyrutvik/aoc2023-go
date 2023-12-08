package day1

import (
	"strconv"
	"strings"
	"unicode"
)

type Part1 struct{}

func MakePart1() *Part1 {
	return &Part1{}
}

func (p *Part1) Clean(line string) string {
	var sb strings.Builder
	for _, c := range line {
		if unicode.IsDigit(c) {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func (p *Part1) Calibrate(line string) (val int, err error) {
	last := len(line) - 1
	val, err = strconv.Atoi(string([]byte{line[0], line[last]}))
	return
}
