package day1

import (
	"bufio"
	"log"
	"os"
)

type Processor interface {
	Clean(line string) string
	Calibrate(line string) (val int, err error)
}

func readLines(filePath string) []string {
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

	return lines
}

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
