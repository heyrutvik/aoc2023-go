package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Processor2 struct{}

func MakeProcessor2() *Processor2 {
	return &Processor2{}
}

// rest: one2sevenine acc:
// rest: ne2sevenine  acc: 1
// rest: e2sevenine   acc: 1
// rest: 2sevenine    acc: 1
// rest: sevenine     acc: 12
// rest: evenine      acc: 127
// rest: venine       acc: 127
// rest: enine        acc: 127
// rest: nine         acc: 127
// rest: ine          acc: 1279
// rest: ne           acc: 1279
// rest: e            acc: 1279
// rest:              acc: 1279

func (p *Processor2) Clean(line string) string {
	var loop func(string, string) string
	loop = func(rest string, acc string) string {
		if len(rest) == 0 {
			return acc
		}
		val, isDigit := isPrefixDigit(rest)
		if isDigit {
			return loop(rest[1:], acc+fmt.Sprint(val))
		} else {
			return loop(rest[1:], acc)
		}
	}
	return loop(line, "")
}

func isPrefixDigit(line string) (val int, isDigit bool) {
	c := line[0]
	if val = int(c - '0'); unicode.IsDigit(rune(c)) {
		isDigit = true
		return
	}

	for _, key := range keys {
		if strings.HasPrefix(line, key) {
			val, _ = strconv.Atoi(dict[key])
			isDigit = true
			return
		}
	}
	return
}

func (p *Processor2) Calibrate(line string) (val int, err error) {
	l := len(line)
	a, b := line[0], line[l-1]
	val, _ = strconv.Atoi(string([]byte{a, b}))
	return
}
