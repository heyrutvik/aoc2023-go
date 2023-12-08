package day7

import (
	"testing"
)

var input []string = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestPart1(t *testing.T) {
	part := &Part1{}
	c := CamelCards{
		Lines: input,
		Part:  part,
	}

	want := 6440
	got := solve(&c)
	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	part := &Part2{}
	c := CamelCards{
		Lines: input,
		Part:  part,
	}

	want := 5905
	got := solve(&c)
	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}
