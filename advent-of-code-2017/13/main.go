package p13

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 13[a|b] <filename>")
	}

	switch args[0] {
	case "13a", "13":
		fmt.Print(Severity(args[1]))
	case "13b":
		fmt.Print(Delay(args[1]))
	}
}

type Firewall struct {
	n int
	p []int
	w []int
	d []int
}

func load(s string) (*Firewall, int) {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	max, m := 0, make(map[int]int)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		i := strings.Index(line, ":")
		if i < 0 {
			log.Fatal("unknown format")
		}

		a, err := strconv.Atoi(line[:i])
		if err != nil {
			log.Fatal(err)
		}

		b, err := strconv.Atoi(line[i+2:])
		if err != nil {
			log.Fatal(err)
		}

		if a > max || len(m) == 0 {
			max = a
		}

		m[a] = b
	}

	p := make([]int, max+1)
	w := make([]int, max+1)
	d := make([]int, max+1)

	for i := 0; i <= max; i++ {
		p[i] = 0
		w[i] = m[i]
		d[i] = 1
	}

	return &Firewall{n: max, p: p, w: w, d: d}, max
}

func (f *Firewall) step() {
	for i := 0; i <= f.n; i++ {
		if f.w[i] <= 1 {
			continue
		}

		if f.p[i] == 0 && f.d[i] == -1 {
			f.d[i] = 1
		}

		if f.p[i] == f.w[i]-1 && f.d[i] == 1 {
			f.d[i] = -1
		}

		f.p[i] += f.d[i]
	}
}

func (f *Firewall) reset() {
	for i := 0; i <= f.n; i++ {
		f.p[i] = 0
		f.d[i] = 1
	}
}

func Severity(s string) int {
	f, n := load(s)

	total := 0

	for i := 0; i <= n; i++ {
		if f.p[i] == 0 {
			total += i * f.w[i]
		}

		f.step()
	}

	return total
}

func Delay(s string) int {
	f, n := load(s)
	delay := 0

	for {
		f.reset()

		for i := 0; i < delay; i++ {
			f.step()
		}

		ok := true
		for i := 0; i <= n; i++ {
			if f.w[i] > 1 {
				if ok = f.p[i] != 0; !ok {
					break
				}
			}

			f.step()
		}

		if ok {
			break
		}

		delay++
	}

	return delay
}

// The above sucks.
// It takes forever to get an answer...
// Let's try again.

func Delay2(s string) int {
	f, _ := load(s)

	// Each layer takes n=2w-2 steps to get from 0 back to 0 e.g. n=4 for w=3; n=6 for w=4
	// So a layer is at 0 each %n e.g. for n=4, when t=0,4,8,12...

	m := make(map[int]int)

	for i, w := range f.w {
		if w <= 1 {
			continue
		}

		m[i] = 2*w - 2
	}

	// Then, we can predict that if we wait k steps, we will be at layer i at t=k+i
	// And we will hit 0 only if t%n is 0 or (k+i)%(2w-2) is 0

	k := 0

	for {
		ok := true
		for i, n := range m {
			if (k+i)%n == 0 {
				ok = false
				break
			}
		}

		if ok {
			break
		}

		k++
	}

	return k
}
