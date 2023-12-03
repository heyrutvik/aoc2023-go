package day1

import (
	"testing"

	"github.com/heyrutvik/aoc2023/utils"
)

func validate(t *testing.T, want any, got any) {
	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}

func TestReverseString(t *testing.T) {
	validate(t, "dcba", utils.Reverse("abcd"))
}

func TestFirstNumber(t *testing.T) {
	p := MakeProcessor1()
	num, _ := firstNumber("abc2def")
	validate(t, 2, num)
	num, _ = firstNumber("2abcdef")
	validate(t, 2, num)
	num, _ = firstNumber("abcdef2")
	validate(t, 2, num)
	num, _ = firstNumber(p.Clean("abcninedef2"))
	validate(t, 9, num)
}

func TestFirstNumberFailed(t *testing.T) {
	_, err := firstNumber("abcdef")
	validate(t, "number not found in `abcdef`", err.Error())
}

func TestCalibrationValue(t *testing.T) {
	p := MakeProcessor1()
	val, _ := p.Calibrate("abc2def")
	validate(t, 22, val)
	val, _ = p.Calibrate("a1bc2def")
	validate(t, 12, val)
	val, _ = p.Calibrate("a2b4c2d7ef")
	validate(t, 27, val)
	val, _ = p.Calibrate(p.Clean("a2b4c2d7efone"))
	validate(t, 21, val)
	val, _ = p.Calibrate(p.Clean("two1nine"))
	validate(t, 29, val)
	val, _ = p.Calibrate(p.Clean("eightwothree"))
	validate(t, 83, val)
	val, _ = p.Calibrate(p.Clean("abcone2threexyz"))
	validate(t, 13, val)
	val, _ = p.Calibrate(p.Clean("xtwone3four"))
	validate(t, 24, val)
	val, _ = p.Calibrate(p.Clean("4nineeightseven2"))
	validate(t, 42, val)
	val, _ = p.Calibrate(p.Clean("zoneight234"))
	validate(t, 14, val)
	val, _ = p.Calibrate(p.Clean("7pqrstsixteen"))
	validate(t, 76, val)
	val, _ = p.Calibrate(p.Clean("eighthree"))
	validate(t, 83, val)
	val, _ = p.Calibrate(p.Clean("sevenine"))
	validate(t, 79, val)
}

func TestSanitizeInput(t *testing.T) {
	p := MakeProcessor1()
	validate(t, "abcd", p.Clean("abcd"))
	validate(t, "123456789", p.Clean("onetwothreefourfivesixseveneightnine"))
}
