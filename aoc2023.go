package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/heyrutvik/aoc2023/day1"
	"github.com/heyrutvik/aoc2023/day2"
	"github.com/heyrutvik/aoc2023/day3"
)

type Puzzle interface {
	Desc()
	Solve()
}

func solution(p Puzzle) {
	p.Desc()
	p.Solve()
}

func main() {
	args := os.Args[2:]

	if args[0] == "--day" || args[0] == "-d" {
		day, err := strconv.Atoi(args[1])
		if err != nil || (day < 1 || day > 25) {
			fmt.Println("help: --day expects a number between 1 and 25.")
			os.Exit(0)
		}

		var puzzle Puzzle

		switch day {
		case 1:
			puzzle = day1.MakeTrebuchet()
			solution(puzzle)
		case 2:
			puzzle = day2.MakeCubeConundrum()
			solution(puzzle)
		case 3:
			puzzle = day3.MakeGearRatios()
			solution(puzzle)
		default:
			fmt.Println("The solution of the day", day, "does not exist yet.")
		}
	}
}
