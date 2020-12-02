package p12

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 12[a|b] 'input'")
	}

	switch args[0] {
	case "12a", "12":
		fmt.Print(Count(args[1]))
	case "12b":
		fmt.Print(Total(args[1]))
	}
}

type Group map[int]struct{}
type Graph map[int][]int

func load(s string) Graph {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	m := Graph{}

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		i := strings.Index(line, "<->")
		if i < 0 {
			log.Fatal("unknown line format")
		}

		a, err := strconv.Atoi(line[:i-1])
		if err != nil {
			log.Fatal(err)
		}

		ints := make([]int, 0)

		for _, j := range strings.Split(line[i+4:], ", ") {
			b, err := strconv.Atoi(j)
			if err != nil {
				log.Fatal(err)
			}

			ints = append(ints, b)
		}

		m[a] = ints
	}

	return m
}

func (m Graph) Group(q int) Group {
	g := Group{q: struct{}{}}
	w := []int{q}

	for {
		if len(w) == 0 {
			break
		}

		n := len(w) - 1
		i := w[n]
		w = w[:n]

		for _, k := range m[i] {
			if _, ok := g[k]; ok {
				continue
			}

			g[k] = struct{}{}
			w = append(w, k)
		}
	}

	return g
}

func Count(s string) int {
	m := load(s)
	g := m.Group(0)
	return len(g)
}

func Total(s string) int {
	m := load(s)
	w := make(map[int]struct{})
	h := []Group{}

	for k := range m {
		if _, ok := w[k]; ok {
			continue
		}

		g := m.Group(k)

		for u := range g {
			w[u] = struct{}{}
		}

		h = append(h, g)
	}

	return len(h)
}
