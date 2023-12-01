package day1

import "testing"

func validate(t *testing.T, want any, got any) {
	if got != want {
		t.Fatalf("Got: %v; Want: %v", got, want)
	}
}

func TestReverseString(t *testing.T) {
	validate(t, "dcba", reverse("abcd"))
}

func TestFirstNumber(t *testing.T) {
	num, _ := firstNumber("abc2def")
	validate(t, 2, num)
	num, _ = firstNumber("2abcdef")
	validate(t, 2, num)
	num, _ = firstNumber("abcdef2")
	validate(t, 2, num)
	num, _ = firstNumber(sensitize("abcninedef2"))
	validate(t, 9, num)
}

func TestFirstNumberFailed(t *testing.T) {
	_, err := firstNumber("abcdef")
	validate(t, "number not found in `abcdef`", err.Error())
}

func TestCalibrationValue(t *testing.T) {
	val, _ := calibrationValue("abc2def")
	validate(t, 22, val)
	val, _ = calibrationValue("a1bc2def")
	validate(t, 12, val)
	val, _ = calibrationValue("a2b4c2d7ef")
	validate(t, 27, val)
	val, _ = calibrationValue(sensitize("a2b4c2d7efone"))
	validate(t, 21, val)
	val, _ = calibrationValue(sensitize("two1nine"))
	validate(t, 29, val)
	val, _ = calibrationValue(sensitize("eightwothree"))
	validate(t, 83, val)
	val, _ = calibrationValue(sensitize("abcone2threexyz"))
	validate(t, 13, val)
	val, _ = calibrationValue(sensitize("xtwone3four"))
	validate(t, 24, val)
	val, _ = calibrationValue(sensitize("4nineeightseven2"))
	validate(t, 42, val)
	val, _ = calibrationValue(sensitize("zoneight234"))
	validate(t, 14, val)
	val, _ = calibrationValue(sensitize("7pqrstsixteen"))
	validate(t, 76, val)
	val, _ = calibrationValue(sensitize("eighthree"))
	validate(t, 83, val)
	val, _ = calibrationValue(sensitize("sevenine"))
	validate(t, 79, val)
}

func TestSensitizeInput(t *testing.T) {
	validate(t, "abcd", sensitize("abcd"))
	validate(t, "123456789", sensitize("onetwothreefourfivesixseveneightnine"))
}
