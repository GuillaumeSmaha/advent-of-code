package p19

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 19[a|b] filename")
	}

	fmt.Print(Walk(args[1]))
}

func Walk(s string) (string, int) {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	m := make([]string, 0)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		m = append(m, line)
	}

	x, y, dx, dy := 0, 0, 0, 1

	for i := range m[0] {
		if m[0][i] == '|' {
			x = i
			break
		}
	}

	get := func(x, y int) byte {
		if y < 0 || y >= len(m) {
			return ' '
		}

		p := m[y]

		if x < 0 || x >= len(p) {
			return ' '
		}

		return p[x]
	}

	sum := make([]byte, 0)

	n := 1

	for {
		c := get(x+dx, y+dy)

		if c >= 'A' && c <= 'Z' {
			sum = append(sum, c)
		}

		if dy == 0 && (c != ' ') {
			x += dx
			n++
			continue
		}

		if dx == 0 && (c != ' ') {
			y += dy
			n++
			continue
		}

		if dx == 0 && (c == ' ') {
			dy = 0

			if get(x+1, y) != ' ' {
				dx = 1
				continue
			}

			if get(x-1, y) != ' ' {
				dx = -1
				continue
			}

			break
		}

		if dy == 0 && (c == ' ') {
			dx = 0

			if get(x, y+1) != ' ' {
				dy = 1
				continue
			}

			if get(x, y-1) != ' ' {
				dy = -1
				continue
			}

			break
		}
	}

	return string(sum), n
}
