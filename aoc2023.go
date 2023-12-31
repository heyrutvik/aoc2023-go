package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/heyrutvik/aoc2023/day1"
	"github.com/heyrutvik/aoc2023/day2"
	"github.com/heyrutvik/aoc2023/day3"
	"github.com/heyrutvik/aoc2023/day4"
	"github.com/heyrutvik/aoc2023/day5"
	"github.com/heyrutvik/aoc2023/day6"
	"github.com/heyrutvik/aoc2023/day7"
	"github.com/heyrutvik/aoc2023/day8"
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

		part := 1
		if len(args) == 4 {
			if args[2] == "--part" || args[2] == "-p" {
				part, err = strconv.Atoi(args[3])
				if err != nil {
					if part < 1 || part > 2 {
						fmt.Println("help: --part expects either 1 or 2. [default: 1]")
						os.Exit(0)
					}
				}
			}
		}

		var puzzle Puzzle

		switch day {
		case 1:
			puzzle = day1.MakeTrebuchet(part)
			solution(puzzle)
		case 2:
			puzzle = day2.MakeCubeConundrum(part)
			solution(puzzle)
		case 3:
			puzzle = day3.MakeGearRatios(part)
			solution(puzzle)
		case 4:
			puzzle = day4.MakeScratchcards(part)
			solution(puzzle)
		case 5:
			puzzle = day5.MakeSeedFertilizer(part)
			solution(puzzle)
		case 6:
			puzzle = day6.MakeWaitForIt(part)
			solution(puzzle)
		case 7:
			puzzle = day7.MakeCamelCards(part)
			solution(puzzle)
		case 8:
			puzzle = day8.MakeHauntedWasteland(part)
			solution(puzzle)
		default:
			fmt.Println("The solution of the day", day, "does not exist yet.")
		}
	}
}
