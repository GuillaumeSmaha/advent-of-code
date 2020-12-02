package p22

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 22[a|b] filename")
	}

	switch args[0] {
	case "22a", "22":
		fmt.Print(Infect(args[1]))
	case "22b":
		fmt.Print(Infect2(args[1]))
	}
}

func load(s string) (map[[2]int]struct{}, int, int) {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	m := make(map[[2]int]struct{})

	y := 0
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		for x := range line {
			if line[x] == '#' {
				m[[2]int{x, y}] = struct{}{}
			}
		}

		y++
	}

	return m, y / 2, y / 2
}

func Infect(s string) int {
	m, x, y := load(s)

	count, dir := 0, 0

	for burst := 0; burst < 10000; burst++ {
		if _, ok := m[[2]int{x, y}]; ok {
			dir++
			delete(m, [2]int{x, y})
		} else {
			if dir--; dir < 0 {
				dir += 4
			}

			m[[2]int{x, y}] = struct{}{}
			count++
		}

		switch dir % 4 {
		case 0:
			y--
		case 1:
			x++
		case 2:
			y++
		case 3:
			x--
		}
	}

	return count
}

const (
	Weakened = 1
	Infected = 2
	Flagged  = 3
)

func Infect2(s string) int {
	loaded, x, y := load(s)

	// Initialize the input as being Infected and keep track of the state in the map.
	m := make(map[[2]int]int)
	for k := range loaded {
		m[k] = Infected
	}

	count, dir := 0, 0

	for burst := 0; burst < 10000000; burst++ {
		state, ok := m[[2]int{x, y}]
		if !ok {
			if dir--; dir < 0 {
				dir += 4
			}

			m[[2]int{x, y}] = Weakened
		} else {
			switch state {
			case Weakened:
				m[[2]int{x, y}] = Infected
				count++
			case Infected:
				dir++
				m[[2]int{x, y}] = Flagged
			case Flagged:
				dir += 2
				delete(m, [2]int{x, y})
			}
		}

		switch dir % 4 {
		case 0:
			y--
		case 1:
			x++
		case 2:
			y++
		case 3:
			x--
		}
	}

	return count
}
