package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/trees/avltree"
	"github.com/heyrutvik/aoc2023/utils"
)

type CamelCards struct {
	Lines []string
	Part  Part
}

func MakeCamelCards(part int) *CamelCards {
	var p Part
	p = &Part1{}
	if part == 2 {
		p = &Part2{}
	}
	return &CamelCards{
		Lines: utils.ReadLines("./day7/input.txt"),
		Part:  p,
	}
}

func (s *CamelCards) Desc() {
	fmt.Println("Puzzle:  ", "Camel Cards")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/7")
}

func (s *CamelCards) Solve() {
	result := solve(s)
	fmt.Println("Solution:", result)
}

type Hand struct {
	cards string
}

func solve(s *CamelCards) int {
	avl := avltree.NewWith(func(a, b interface{}) int {
		return s.Part.Compare(a.(Hand), b.(Hand))
	})

	for _, line := range s.Lines {
		parsed := strings.Split(line, " ")
		cards, bidStr := parsed[0], parsed[1]
		bid, _ := strconv.Atoi(bidStr)
		avl.Put(Hand{cards}, bid)
	}

	result := 0
	rank := 1
	for _, bid := range avl.Values() {
		result += bid.(int) * rank
		rank += 1
	}
	return result
}
