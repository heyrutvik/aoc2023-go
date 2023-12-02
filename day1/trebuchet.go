package day1

import (
	"fmt"
)

type Trebuchet struct {
	Lines []string
	Proc  Processor
}

func MakeTrebuchet() *Trebuchet {
	return &Trebuchet{
		Lines: readLines("./day1/input.txt"),
		Proc:  MakeProcessor2(),
	}
}

func (t *Trebuchet) Desc() {
	fmt.Println("Puzzle:  ", "Trebuchet")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/1")
}

func (t *Trebuchet) Solve() {
	total := 0
	for _, line := range t.Lines {
		s := t.Proc.Clean(line)
		val, err := t.Proc.Calibrate(s)
		if err != nil {
			panic(err.Error())
		}
		total = val + total
	}
	fmt.Println("Solution:", total)
}
