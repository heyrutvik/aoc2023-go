package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/heyrutvik/aoc2023/utils"
)

type Scratchcards struct {
	Lines []string
	Part  Part
}

type Game struct {
	id      int
	winning utils.Set[int]
	ours    utils.Set[int]
}

func MakeScratchcards(part int) *Scratchcards {
	var p Part
	p = &Part1{}
	if part == 2 {
		p = &Part2{}
	}
	return &Scratchcards{
		Lines: utils.ReadLines("./day4/input.txt"),
		Part:  p,
	}
}

func (s *Scratchcards) Desc() {
	fmt.Println("Puzzle:  ", "Scratchcards")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/4")
}

func (s *Scratchcards) Solve() {
	total := s.Part.Solve(s)
	fmt.Println("Solution:", total)
}

type Part1 struct{}

func (p *Part1) Solve(s *Scratchcards) int {
	total := 0
	for _, line := range s.Lines {
		cardPoints := part1(line)
		total = cardPoints + total
	}
	return total
}

type Part2 struct{}

func (p *Part2) Solve(s *Scratchcards) int {
	total := 0
	matching := make(map[int]int)
	cards := make(map[int]int)
	stack := stack.New()
	for _, line := range s.Lines {
		game := parse(line)
		set := game.winning.Intersection(&game.ours)
		matching[game.id] = set.Size()
		cards[game.id] = 0
		stack.Push(game.id)
	}
	for stack.Len() > 0 {
		gid := stack.Pop().(int)
		cards[gid] = cards[gid] + 1
		matches := matching[gid]
		for i := 1; i <= matches; i++ {
			stack.Push(gid + i)
		}
	}
	for _, count := range cards {
		total = count + total
	}
	return total
}

func part1(line string) (cardPoints int) {
	game := parse(line)
	set := game.winning.Intersection(&game.ours)
	count := set.Size()
	cardPoints = double(count)
	return
}

func parse(line string) Game {
	start := strings.Index(line, " ")
	numbers := strings.Index(line, ":")
	separator := strings.Index(line, "|")

	id := atoi(strings.TrimSpace(line[start+1 : numbers]))
	wn := utils.MakeSet[int](utils.Map(atoi, strings.Fields(line[numbers+1:separator])))
	on := utils.MakeSet[int](utils.Map(atoi, strings.Fields(line[separator+1:])))

	return Game{id, wn, on}
}

func atoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func double(count int) (result int) {
	switch count {
	case 0:
		result = 0
	case 1:
		result = 1
	default:
		result = int(math.Pow(2, float64(count-1)))
	}
	return
}
