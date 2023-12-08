package day7

import "strings"

type Part interface {
	Weights() map[byte]int
	Upgrade(hand Hand) Hand
	Compare(hand1 Hand, hand2 Hand) (order int)
}

type Part1 struct{}

func (p *Part1) Weights() map[byte]int {
	weigths := map[byte]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'J': 9,
		'T': 8,
		'9': 7,
		'8': 6,
		'7': 5,
		'6': 4,
		'5': 3,
		'4': 2,
		'3': 1,
		'2': 0,
	}
	return weigths
}

func (p *Part1) Upgrade(hand Hand) Hand {
	return hand
}

func (p *Part1) Compare(hand1 Hand, hand2 Hand) int {
	weightComparison := compareHandsByWeight(hand1, hand2)
	if weightComparison != 0 {
		return weightComparison
	}

	weights := p.Weights()
	return compareCardByCard(hand1, hand2, &weights)
}

type Part2 struct{}

func (p *Part2) Weights() map[byte]int {
	weigths := map[byte]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
	}
	return weigths
}

// upgrade J to make hand more valuable
func (p *Part2) Upgrade(hand Hand) Hand {
	result, max := hand, weight(hand)
	if strings.Contains(hand.cards, "J") {
		for card := range p.Weights() {
			cards := strings.ReplaceAll(hand.cards, "J", string(card))
			phand := Hand{cards}
			pweight := weight(phand)
			if max < pweight {
				result = phand
				max = pweight
			}
		}
	}
	return result
}

func (p *Part2) Compare(hand1 Hand, hand2 Hand) int {
	uhand1, uhand2 := p.Upgrade(hand1), p.Upgrade(hand2)
	weightComparison := compareHandsByWeight(uhand1, uhand2)
	if weightComparison != 0 {
		return weightComparison
	}

	weights := p.Weights()
	return compareCardByCard(hand1, hand2, &weights)
}

// hand weight
// five  of a kind (AAAAA) 6 | 1*1 + 1*2 + 1*3 + 1*4 + 1*5 = 15
// four  of a kind (AA8AA) 5 | 1*1 + 1*2 + 1*3 + 1*4 +   1 = 11
// full house      (23332) 4 | 1*1 + 1*2 + 1*1 + 1*2 + 1*3 =  9
// three of a kind (TTT98) 3 | 1*1 + 1*2 + 1*3 +   1 +   1 =  8
// two  pair       (23432) 2 | 1*1 + 1*2 + 1*1 + 1*2 +   1 =  7
// one  pair       (A23A4) 1 | 1*1 + 1*2 +   1 +   1 +   1 =  6
// high card       (23456) 0 |   1 +   1 +   1 +   1 +   1 =  5
func weight(hand Hand) int {
	cards := make(map[rune]int)
	for _, c := range hand.cards {
		_, exist := cards[c]
		if exist {
			cards[c] += 1
		} else {
			cards[c] = 1
		}
	}

	weight := 0
	for _, val := range cards {
		for i := val; i > 0; i-- {
			weight += i
		}
	}

	return weight
}

func compareCardByCard(hand1 Hand, hand2 Hand, weights *map[byte]int) int {
	for i := 0; i < 5; i++ {
		c1 := (*weights)[hand1.cards[i]]
		c2 := (*weights)[hand2.cards[i]]
		if c1 < c2 {
			return -1
		} else if c1 > c2 {
			return 1
		}
	}
	return 0
}

func compareHandsByWeight(hand1, hand2 Hand) int {
	weight1, weight2 := weight(hand1), weight(hand2)
	if weight1 < weight2 {
		return -1
	} else if weight2 < weight1 {
		return 1
	} else {
		return 0
	}
}
