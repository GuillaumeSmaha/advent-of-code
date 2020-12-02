package p21

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 21[a|b] 'input' iterations")
	}

	n, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(Pixels(args[1], n))
}

type rule struct {
	s []string
}

func load(s string) map[string]rule {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	m := make(map[string]rule)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		p := strings.Split(line, " => ")
		if len(p) != 2 {
			log.Fatal("unknown rule format")
		}

		a := strings.Split(p[0], "/")
		b := strings.Split(p[1], "/")

		if len(a)+1 != len(b) {
			log.Fatal("wrong size")
		}

		m[p[0]] = rule{s: b}
	}

	p := make(map[string]rule)

	for k, r := range m {
		q := strings.Split(k, "/")
		n := len(q)
		g := make([]byte, n*n)
		h := make([]byte, n*n)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g[i*n+j] = q[i][j]
			}
		}

		add := func() {
			q := make([]string, n)
			for i := range q {
				q[i] = string(h[i*n : i*n+n])
			}

			w := strings.Join(q, "/")
			p[w] = r
			g, h = h, g
		}

		flap := func() {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					h[i*n+j] = g[i*n+(n-j-1)]
				}
			}
			add()
		}

		flip := func() {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					h[i*n+j] = g[(n-i-1)*n+j]
				}
			}
			add()
		}

		rotate := func() {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					h[i*n+j] = g[j*n+i]
				}
			}
			add()
		}

		for i := 0; i < 4; i++ {
			flap()
			flap()
			flip()
			flap()
			flap()
			flip()
			rotate()
		}
	}

	return p
}

func Pixels(s string, loops int) int {
	m := load(s)

	p := []string{".#.", "..#", "###"}

	fill := func(u, v int) {
		n := len(p) / u
		q := make([]string, n*v)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				s := make([]string, u)

				for x := 0; x < u; x++ {
					c := p[i*u+x]
					k := c[j*u : (j+1)*u]
					s[x] = k
				}

				k := strings.Join(s, "/")
				r, ok := m[k]
				if !ok {
					log.Fatal("rule not found")
				}

				for h, y := range r.s {
					q[i*len(r.s)+h] += y
				}
			}
		}

		p = q
	}

	for k := 0; k < loops; k++ {
		n := len(p)

		if n%2 == 0 {
			fill(2, 3)
			continue
		}

		if n%3 == 0 {
			fill(3, 4)
			continue
		}

		log.Fatal("wrong size")
	}

	count := 0
	for _, line := range p {
		count += strings.Count(line, "#")
	}

	return count
}
