package day3

type Part interface {
	Solve(lines *[]string, numbers *map[Location]int, row int, col int, symbol rune)
}
