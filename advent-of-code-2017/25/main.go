package p25

import (
	"fmt"
	"log"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 25")
	}

	fmt.Print(Turing())
}

type State [2]Action

type Action struct {
	set  int
	step int
	next int
}

func Turing() int {
	states := []State{
		State{
			Action{set: 1, step: 1, next: 1},
			Action{set: 0, step: -1, next: 3},
		},
		State{
			Action{set: 1, step: 1, next: 2},
			Action{set: 0, step: 1, next: 5},
		},
		State{
			Action{set: 1, step: -1, next: 2},
			Action{set: 1, step: -1, next: 0},
		},
		State{
			Action{set: 0, step: -1, next: 4},
			Action{set: 1, step: 1, next: 0},
		},
		State{
			Action{set: 1, step: -1, next: 0},
			Action{set: 0, step: 1, next: 1},
		},
		State{
			Action{set: 0, step: 1, next: 2},
			Action{set: 0, step: 1, next: 4},
		},
	}

	m := make(map[int]int)
	p := 0
	s := 0

	for i := 0; i < 12302209; i++ {
		x := m[p]
		state := states[s][x]
		m[p] = state.set
		p, s = p+state.step, state.next
	}

	n := 0
	for _, v := range m {
		n += v
	}

	return n
}
