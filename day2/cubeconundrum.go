package day2

import (
	"fmt"

	"github.com/heyrutvik/aoc2023/utils"
)

type CubeConundrum struct {
	Lines []string
	Part  Part
}

func MakeCubeConundrum(part int) *CubeConundrum {
	var p Part
	p = &Part1{}
	if part == 2 {
		p = &Part2{}
	}
	return &CubeConundrum{
		Lines: utils.ReadLines("./day2/input.txt"),
		Part:  p,
	}
}

func (c *CubeConundrum) Desc() {
	fmt.Println("Puzzle:  ", "Cube Conundrum")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/2")
}

func (c *CubeConundrum) Solve() {
	total := 0
	for _, line := range c.Lines {
		comb := Parse(line)
		total = c.Part.Solve(comb, total)
	}
	fmt.Println("Solution:", total)
}

type Part1 struct{}

func (p *Part1) Solve(comb Comb, total int) (result int) {
	result = total
	exlcude := false
	for _, set := range comb.Sets {
		if set.Red > 12 || set.Green > 13 || set.Blue > 14 {
			exlcude = true
			break
		}
	}
	if !exlcude {
		result = comb.Game + total
	}
	return
}

type Part2 struct{}

func (p *Part2) Solve(comb Comb, total int) int {
	red, green, blue := 1, 1, 1
	for _, set := range comb.Sets {
		if red < set.Red {
			red = set.Red
		}
		if green < set.Green {
			green = set.Green
		}
		if blue < set.Blue {
			blue = set.Blue
		}
	}
	return (red * green * blue) + total
}
