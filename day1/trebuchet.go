package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var dict = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var keys = func() []string {
	ks := make([]string, 0, len(dict))
	for k := range dict {
		ks = append(ks, k)
	}
	return ks
}()

type Trebuchet struct {
	Lines []string
}

func MakeTrebuchet(filePath string) *Trebuchet {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return &Trebuchet{
		Lines: lines,
	}
}

func (t *Trebuchet) Desc() {
	fmt.Println("Puzzle:  ", "Trebuchet")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/1")
}

func (t *Trebuchet) Solve() {
	total := 0
	for _, line := range t.Lines {
		s := sensitize(line)
		val, err := calibrationValue(s)
		if err != nil {
			panic(err.Error())
		}
		total = val + total
	}
	fmt.Println("Solution:", total)
}

func sensitize(s string) (r string) {
	r = loop(s, "", false)
	return
}

func loop(rest string, acc string, wasDigitInWords bool) string {
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

func calibrationValue(s string) (val int, err error) {
	n1, err := firstNumber(s)
	n2, _ := firstNumber(reverse(s)) // no need to check error here
	val, _ = strconv.Atoi(fmt.Sprint(n1) + fmt.Sprint(n2))
	return
}

func reverse(s string) string {
	var result string
	for _, c := range s {
		result = string(c) + result
	}
	return result
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
