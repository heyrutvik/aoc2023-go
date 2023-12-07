package day6

import (
	"fmt"

	"github.com/heyrutvik/aoc2023/utils"
)

type WaitForIt struct {
	part int
}

func MakeWaitForIt(part int) *WaitForIt {
	return &WaitForIt{part}
}

func (s *WaitForIt) Desc() {
	fmt.Println("Puzzle:  ", "Wait For It")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/6")
}

func (s *WaitForIt) Solve() {
	var result int
	if s.part == 1 {
		a := len(comp(60, 475))
		b := len(comp(94, 2138))
		c := len(comp(78, 1015))
		d := len(comp(82, 1650))
		result = a * b * c * d
	} else {
		result = len(comp(60947882, 475213810151650))
	}
	fmt.Println("Solution:", result)
}

func comp(time int, record int) (result []int) {
	set := utils.MakeSet[int](result)
	t := time / 2

	offset := 0
	if time%2 == 0 {
		offset = -1
	}

	for i := t; i > 0; i-- {
		offset += 1
		speed := i
		run := time - i
		distance := speed * run
		if distance <= record {
			break
		}
		set.Add(i)
		set.Add(t + offset)
	}

	result = set.Elements()
	return
}
