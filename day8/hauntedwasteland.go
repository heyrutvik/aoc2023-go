package day8

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/heyrutvik/aoc2023/utils"
)

type HauntedWasteland struct {
	Blocks [][]string
}

func MakeHauntedWasteland(part int) *HauntedWasteland {
	return &HauntedWasteland{
		utils.ReadBlocks("./day8/input.txt"),
	}
}

func (s *HauntedWasteland) Desc() {
	fmt.Println("Puzzle:  ", "Haunted Wasteland")
	fmt.Println("Link:    ", "https://adventofcode.com/2023/day/8")
}

func (s *HauntedWasteland) Solve() {
	result := part1(s)
	fmt.Println("Solution:", result)
}

func part1(s *HauntedWasteland) int {
	tape := MakeTape(s.Blocks[0][0])
	machine, _ := MakeStateMachine(s.Blocks[1])
	machine.init("AAA")

	for {
		head := tape.next()
		currState := machine.next(head)
		if currState == "ZZZ" {
			break
		}
	}

	return machine.step
}

type State string

type Fork struct {
	left  State
	right State
}

// tuple = "(AAA, BBB)"
func MakeFork(tuple string) Fork {
	parse := func(a string) string {
		var sb strings.Builder
		for _, c := range a {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				sb.WriteRune(c)
			}
		}
		return sb.String()
	}

	t := strings.Split(tuple, ",")
	left, right := parse(t[0]), parse(t[1])
	return Fork{State(left), State(right)}
}

type Tape struct {
	index    int
	legnth   int
	commands []int
}

// tape = "LLR"
func MakeTape(tape string) *Tape {
	commands := make([]int, 0)

	for _, c := range tape {
		if c == 'L' {
			commands = append(commands, 0)
		} else {
			commands = append(commands, 1)
		}
	}

	return &Tape{0, len(commands), commands}
}

func (t *Tape) next() int {
	index := t.index % t.legnth
	result := t.commands[index]
	t.index += 1
	return result
}

type StateMachine struct {
	current State
	states  map[State]Fork
	step    int
}

func MakeStateMachine(table []string) (machine *StateMachine, err error) {
	var step int
	var state State
	states := make(map[State]Fork)

	for _, row := range table {
		parts := strings.Split(row, "=")
		state = State(strings.TrimSpace(parts[0]))
		fork := MakeFork(parts[1])
		states[state] = fork
	}

	m := StateMachine{state, states, step}

	m, err = m.init(state) // init with whatever the last state was
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	machine = &m
	return
}

func (m *StateMachine) init(s State) (machine StateMachine, err error) {
	_, exist := m.states[s]
	if !exist {
		err = fmt.Errorf("state `%v` is not available in table", s)
	}
	m.current = s
	m.step = 0
	machine = StateMachine{s, m.states, 0}
	return
}

func (sm *StateMachine) next(head int) State {
	var nextState State
	if head == 0 {
		nextState = State(sm.states[sm.current].left)
	} else {
		nextState = State(sm.states[sm.current].right)
	}
	sm.current = nextState
	sm.step += 1
	return nextState
}
