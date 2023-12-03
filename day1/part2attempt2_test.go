package day1

import "testing"

func TestCalibrationValue2(t *testing.T) {
	p := MakePart2Attempt2()
	val, _ := p.Calibrate("2")
	validate(t, 22, val)
	val, _ = p.Calibrate("12")
	validate(t, 12, val)
	val, _ = p.Calibrate("2427")
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

func TestSanitizeInput2(t *testing.T) {
	p := MakePart2Attempt2()
	validate(t, "", p.Clean("abcd"))
	validate(t, "123456789", p.Clean("onetwothreefourfivesixseveneightnine"))
}
