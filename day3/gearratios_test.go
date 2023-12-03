package day3

import (
	"testing"
)

func MakeGearRatiosTest(part Part) *GearRatios {
	return &GearRatios{
		Lines: []string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		},
		Part: part,
	}
}

func TestPart1(t *testing.T) {
	g := MakeGearRatiosTest(&Part1{})
	numbers := make(map[Location]int)
	for row, line := range g.Lines {
		for col, c := range line {
			g.Part.Solve(&g.Lines, &numbers, row, col, c)
		}
	}
	want := 4361
	got := Reduce(&numbers)
	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	g := MakeGearRatiosTest(&Part2{})
	numbers := make(map[Location]int)
	for row, line := range g.Lines {
		for col, c := range line {
			g.Part.Solve(&g.Lines, &numbers, row, col, c)
		}
	}
	want := 467835
	got := Reduce(&numbers)
	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}
