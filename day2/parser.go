package day2

import (
	"strconv"

	parsec "github.com/prataprc/goparsec"
)

type Game struct{ Value int }

type Red struct{ Value int }

type Green struct{ Value int }

type Blue struct{ Value int }

type Set struct {
	Red   int
	Green int
	Blue  int
}

type Sets struct {
	Values []Set
}

type Comb struct {
	Game int
	Sets []Set
}

var num = parsec.Int()

var game = parsec.And(gameNode, parsec.Atom("Game", "GAME"), num)
var red = parsec.And(redNode, num, parsec.Atom("red", "RED"))
var green = parsec.And(greenNode, num, parsec.Atom("green", "GREEN"))
var blue = parsec.And(blueNode, num, parsec.Atom("blue", "BLUE"))
var color = parsec.OrdChoice(oneOfTheThreeColorNode, red, green, blue)
var set = parsec.Many(setNode, color, parsec.Atom(",", "COLOR SEPARATOR"))
var sets = parsec.Many(setsNode, set, parsec.Atom(";", "SET SEPARATOR"))
var input = parsec.And(inputNode, game, parsec.Atom(":", "GAME SEPARATOR"), sets)

func gameNode(pn []parsec.ParsecNode) parsec.ParsecNode {
	value, _ := strconv.Atoi(pn[1].(*parsec.Terminal).Value)
	return Game{value}
}

type Color interface {
	Red | Green | Blue
}

func colorNode[T Color](create func(int) T) func([]parsec.ParsecNode) parsec.ParsecNode {
	return func(pn []parsec.ParsecNode) parsec.ParsecNode {
		value, _ := strconv.Atoi(pn[0].(*parsec.Terminal).Value)
		return create(value)
	}
}

var redNode = colorNode[Red](func(v int) Red { return Red{v} })
var greenNode = colorNode[Green](func(v int) Green { return Green{v} })
var blueNode = colorNode[Blue](func(v int) Blue { return Blue{v} })

func oneOfTheThreeColorNode(pn []parsec.ParsecNode) parsec.ParsecNode {
	return pn[0]
}

func setNode(pn []parsec.ParsecNode) parsec.ParsecNode {
	var red, green, blue int
	for _, p := range pn {
		switch p := p.(type) {
		case Red:
			red = p.Value
		case Green:
			green = p.Value
		case Blue:
			blue = p.Value
		}
	}
	return Set{red, green, blue}
}

func setsNode(pn []parsec.ParsecNode) parsec.ParsecNode {
	var sets []Set
	for _, p := range pn {
		switch p := p.(type) {
		case Set:
			sets = append(sets, p)
		}
	}
	return Sets{sets}
}

func inputNode(pn []parsec.ParsecNode) parsec.ParsecNode {
	var game int
	var sets []Set
	for _, p := range pn {
		switch p := p.(type) {
		case Game:
			game = p.Value
		case Sets:
			sets = p.Values
		}
	}
	return Comb{game, sets}
}

func Parse(line string) Comb {
	scanner := parsec.NewScanner([]byte(line))
	parsed, _ := input(scanner)
	return parsed.(Comb)
}
