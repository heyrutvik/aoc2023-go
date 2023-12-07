package day6

import "testing"

func TestPart1(t *testing.T) {
	a := len(comp(7, 9))
	b := len(comp(15, 40))
	c := len(comp(30, 200))

	want := 288
	got := a * b * c

	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 71503
	got := len(comp(71530, 940200))

	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}
