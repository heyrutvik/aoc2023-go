package day5

type Part interface {
	ParseSeeds(block []string) (seeds []int)
}
