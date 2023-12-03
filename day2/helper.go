package day2

type Part interface {
	Solve(comb Comb, total int) (result int)
}
