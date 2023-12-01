package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/heyrutvik/aoc2023/day1"
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
	fmt.Println(`
     ___       __                 __     ____  ____   ______          __        ___  ____ ___  _____
    /   | ____/ _   _____  ____  / /_   / __ \/ __/  / ________  ____/ ___     |__ \/ __ |__ \|__  /
   / /| |/ __  | | / / _ \/ __ \/ __/  / / / / /_   / /   / __ \/ __  / _ \    __/ / / / __/ / /_ < 
  / ___ / /_/ /| |/ /  __/ / / / /_   / /_/ / __/  / /___/ /_/ / /_/ /  __/   / __/ /_/ / __/___/ / 
 /_/  |_\__,_/ |___/\___/_/ /_/\__/   \____/_/     \____/\____/\__,_/\___/   /____\____/____/____/  																									
	`)
	fmt.Println("help: type `quit` to exit.")

	var input string

	for {
		fmt.Print("Check solution of day [1, 25] >> ")
		fmt.Scan(&input)

		if input == "quit" {
			fmt.Print("\nSo long, and thanks for all the fish!\n\n")
			os.Exit(0)
		}

		day, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("You sure have entered a number between 1 and 25!?\n\n")
			continue
		}

		if day < 1 || day > 25 {
			fmt.Print("You are good for nothing! Try again.\n\n")
			continue
		}

		var puzzle Puzzle

		switch day {
		case 1:
			puzzle = day1.MakeTrebuchet("./day1/input")
			solution(puzzle)
		default:
			fmt.Print("The solution of the day", day, "does not exist yet.\n\n")
		}
	}
}
