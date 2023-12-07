package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/heyrutvik/aoc2023/utils"
)

type SeedFertilizer struct {
	seeds                []int
	seed2soil            Map
	soil2fertilizer      Map
	fertilizer2water     Map
	water2light          Map
	light2temperature    Map
	temperature2humidity Map
	humidity2location    Map
	part                 Part
}

func MakeSeedFertilizer(part int) *SeedFertilizer {
	var p Part
	p = &Part1{}
	if part == 2 {
		p = &Part2{}
	}

	blocks := utils.ReadBlocks("./day5/input.txt")

	return instance(p, blocks)
}

func instance(p Part, blocks [][]string) *SeedFertilizer {
	seeds := p.ParseSeeds(blocks[0])
	seed2soil, _ := parseMap("seed-to-soil", blocks[1])
	soil2fertilizer, _ := parseMap("soil-to-fertilizer", blocks[2])
	fertilizer2water, _ := parseMap("fertilizer-to-water", blocks[3])
	water2light, _ := parseMap("water-to-light", blocks[4])
	light2temperature, _ := parseMap("light-to-temperature", blocks[5])
	temperature2humidity, _ := parseMap("temperature-to-humidity", blocks[6])
	humidity2location, _ := parseMap("humidity-to-location", blocks[7])

	return &SeedFertilizer{
		seeds,
		seed2soil,
		soil2fertilizer,
		fertilizer2water,
		water2light,
		light2temperature,
		temperature2humidity,
		humidity2location,
		p,
	}
}

func (s *SeedFertilizer) Desc() {
	fmt.Println("Puzzle:  ", "If You Give A Seed A Fertilizer")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/5")
}

func (s *SeedFertilizer) Solve() {
	result := solve(s)
	fmt.Println("Solution:", result)
}

func solve(s *SeedFertilizer) int {
	result := int(^uint(0) >> 1)
	for _, seed := range s.seeds {
		soil := s.seed2soil.get(seed)
		fertilizer := s.soil2fertilizer.get(soil)
		water := s.fertilizer2water.get(fertilizer)
		light := s.water2light.get(water)
		temperature := s.light2temperature.get(light)
		humidity := s.temperature2humidity.get(temperature)
		location := s.humidity2location.get(humidity)

		if location < result {
			result = location
		}
	}
	return result
}

type Part1 struct{}

func (p *Part1) ParseSeeds(block []string) (seeds []int) {
	seeds, _ = parseSeedsPart1(block)
	return
}

type Part2 struct{}

func (p *Part2) ParseSeeds(block []string) (seeds []int) {
	seeds, _ = parseSeedsPart2(block)
	return
}

type Slice struct {
	dst int
	src int
	rng int
}

type Map struct {
	slices []Slice
}

func (m *Map) get(key int) (value int) {
	value = key
	for _, slice := range m.slices {
		start := slice.src
		end := slice.src + slice.rng
		if key >= start && key < end {
			offset := key - start
			value = slice.dst + offset
			break
		}
	}
	return
}

func parseSeedsPart1(block []string) (seeds []int, err error) {
	if len(block) == 1 && len(block[0]) > 0 {
		idx := strings.Index(block[0], " ")
		for _, seed := range strings.Fields(block[0][idx:]) {
			s, _ := strconv.Atoi(strings.TrimSpace(seed))
			seeds = append(seeds, s)
		}
	} else {
		err = fmt.Errorf("block doesn't contain seeds")
	}
	return
}

func parseSeedsPart2(block []string) (seeds []int, err error) {
	var input []int
	input, err = parseSeedsPart1(block)
	if err != nil {
		return
	}

	mid := make([][2]int, 0)
	for i := 1; i < len(input); i += 2 {
		mid = append(mid, [2]int{input[i-1], input[i]})
	}

	for _, tuple := range mid {
		start, till := tuple[0], tuple[1]
		for i := 0; i < till; i++ {
			seeds = append(seeds, start+i)
		}
	}
	return
}

func parseMap(prefix string, block []string) (m Map, err error) {
	hasHeader := strings.HasPrefix(block[0], prefix)
	if hasHeader {
		slices := []Slice{}
		for _, line := range block[1:] {
			nums := [3]int{}
			for idx, number := range strings.Fields(line) {
				num, _ := strconv.Atoi(strings.TrimSpace(number))
				nums[idx] = num
			}
			slices = append(slices, Slice{nums[0], nums[1], nums[2]})
		}
		m = Map{slices}
	} else {
		err = fmt.Errorf("block doesn't start with %v prefix", prefix)
	}
	return
}
