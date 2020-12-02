package p24

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 24[a|b] filename")
	}

	switch args[0] {
	case "24a", "24":
		fmt.Print(Strongest(args[1]))
	case "24b":
		fmt.Print(Longest(args[1]))
	}
}

type Conn [2]int

func load(s string) ([]Conn, []int) {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	conns := make([]Conn, 0)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		i := strings.Index(line, "/")
		if i < 0 {
			log.Fatal()
		}

		a, err := strconv.Atoi(line[:i])
		if err != nil {
			log.Fatal(err)
		}

		b, err := strconv.Atoi(line[i+1:])
		if err != nil {
			log.Fatal(err)
		}

		conns = append(conns, Conn{a, b})
	}

	left := []int{}
	for i := range conns {
		left = append(left, i)
	}

	return conns, left
}

func build(conns []Conn, port int, left []int) [][]int {
	result := make([][]int, 0)

	// Here we build all possible bridges.
	// We keep track of which components are left and search for one that match the end-of-the-bridge.
	// Each step recursively build a smaller bridge with left components.

	for i, n := range left {
		less := func() []int {
			list := append([]int{}, left[:i]...)
			if i != len(left) {
				list = append(list, left[i+1:]...)
			}

			return list
		}

		step := func(p int) [][]int {
			tail := build(conns, conns[n][p], less())
			for i, list := range tail {
				tail[i] = append([]int{n}, list...)
			}

			if len(tail) == 0 {
				tail = [][]int{[]int{n}}
			}

			return tail
		}

		if conns[n][0] == port {
			result = append(result, step(1)...)
			continue
		}

		if conns[n][1] == port {
			result = append(result, step(0)...)
			continue
		}
	}

	return result
}

func Strongest(s string) int {
	conns, left := load(s)

	max := 0

	list := build(conns, 0, left)
	for _, item := range list {
		n := 0
		for _, i := range item {
			n += conns[i][0] + conns[i][1]
		}

		if n > max {
			max = n
		}
	}

	return max
}

func Longest(s string) int {
	conns, left := load(s)

	max, longest := 0, 0

	list := build(conns, 0, left)
	for _, item := range list {
		if len(item) >= longest {
			n := 0
			for _, i := range item {
				n += conns[i][0] + conns[i][1]
			}

			if n > max {
				longest, max = len(item), n
			}
		}
	}

	return max
}
