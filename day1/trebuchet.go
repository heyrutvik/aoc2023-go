package day1

import (
	"fmt"

	"github.com/heyrutvik/aoc2023/utils"
)

type Trebuchet struct {
	Lines []string
	Part  Part
}

func MakeTrebuchet(part int) *Trebuchet {
	var p Part
	p = MakePart1()
	if part == 2 {
		p = MakePart2Attempt2()
	}
	return &Trebuchet{
		Lines: utils.ReadLines("./day1/input.txt"),
		Part:  p,
	}
}

func (t *Trebuchet) Desc() {
	fmt.Println("Puzzle:  ", "Trebuchet")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/1")
}

func (t *Trebuchet) Solve() {
	total := 0
	for _, line := range t.Lines {
		s := t.Part.Clean(line)
		val, err := t.Part.Calibrate(s)
		if err != nil {
			panic(err.Error())
		}
		total = val + total
	}
	fmt.Println("Solution:", total)
}
