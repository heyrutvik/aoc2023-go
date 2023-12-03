package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/heyrutvik/aoc2023/utils"
)

type Part2Attempt1 struct{}

func MakePart2Attempt1() *Part2Attempt1 {
	return &Part2Attempt1{}
}

// rest: one2sevenine acc:      wasDigitInWords: false
// rest: e2sevenine   acc: 1    wasDigitInWords: true
// rest: sevenine     acc: 12   wasDigitInWords: false
// rest: nine         acc: 127  wasDigitInWords: true
// rest: nine         acc: 127  wasDigitInWords: false
// rest: e            acc: 1279 wasDigitInWords: true

func (p *Part2Attempt1) Clean(line string) string {
	var loop func(string, string, bool) string
	loop = func(rest string, acc string, wasDigitInWords bool) string {
		// if the length of the rest of the string is empty OR
		// if the last call found a digit (in words) and the rest of the string has only one character
		// return accumulator
		if len(rest) == 0 || (wasDigitInWords && len(rest) == 1) {
			return acc
		}

		// if last call found a digit in words
		// loop if the prefix of the rest of the string is one of the keys again
		// by setting the wasDigitInWords to false
		// otherwise, remove the first character (as it's been already processed)
		if wasDigitInWords {
			for _, key := range keys {
				if strings.HasPrefix(rest, key) {
					return loop(rest, acc, false)
				}
			}
			rest = rest[1:]
		}

		// if rest of the string has one of the keys as its prefix
		// append the corresponding digit to accumulator
		// drop one less length of the key and loop with wasDigitInWords true
		// (for overlapping digits in words e.g. eighthree = 83)
		for _, key := range keys {
			if strings.HasPrefix(rest, key) {
				drop := len(key) - 1
				acc = acc + (dict)[key]
				return loop(rest[drop:], acc, true)
			}
		}

		// otherwise, just remove first character, append it in accumlator and loop
		return loop(rest[1:], acc+string(rest[0]), false)
	}

	return loop(line, "", false)
}

func (p *Part2Attempt1) Calibrate(line string) (val int, err error) {
	n1, err := firstNumber(line)
	n2, _ := firstNumber(utils.Reverse(line)) // no need to check error here
	val, _ = strconv.Atoi(fmt.Sprint(n1) + fmt.Sprint(n2))
	return
}

func firstNumber(s string) (num int, err error) {
	for _, c := range s {
		if num = int(c - '0'); unicode.IsDigit(c) {
			return
		}
	}
	err = fmt.Errorf("number not found in `%v`", s)
	return
}
