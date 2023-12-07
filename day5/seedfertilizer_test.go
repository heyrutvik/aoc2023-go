package day5

import (
	"testing"
)

func validate(t *testing.T, want int, got int) {
	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}

func TestSeedFertillizer(t *testing.T) {
	m := Map{
		[]Slice{
			Slice{50, 98, 2},
			Slice{52, 50, 48},
		},
	}
	validate(t, 0, m.get(0))
	validate(t, 1, m.get(1))
	validate(t, 48, m.get(48))
	validate(t, 52, m.get(50))
	validate(t, 53, m.get(51))
	validate(t, 98, m.get(96))
	validate(t, 99, m.get(97))
	validate(t, 50, m.get(98))
	validate(t, 51, m.get(99))
}

var blocks [][]string = [][]string{
	[]string{
		"seeds: 79 14 55 13",
	},
	[]string{
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	},
	[]string{
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
	},
	[]string{
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
	},
	[]string{
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
	},
	[]string{
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
	},
	[]string{
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
	},
	[]string{
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	},
}

func TestPart1(t *testing.T) {
	s := instance(&Part1{}, blocks)
	want := 35
	got := solve(s)
	validate(t, want, got)
}

func TestPart2(t *testing.T) {
	s := instance(&Part2{}, blocks)
	want := 46
	got := solve(s)
	validate(t, want, got)
}
