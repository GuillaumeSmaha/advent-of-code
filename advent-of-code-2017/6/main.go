package p6

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 6[a|b] '1,2,3,...'")
	}

	ints := make([]int, 0)

	for _, k := range strings.Fields(args[1]) {
		i, err := strconv.Atoi(k)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, i)
	}

	switch args[0] {
	case "6a", "6":
		fmt.Print(Run(ints))
	case "6b":
		fmt.Print(Loop(ints))
	}
}

type Runner interface {
	Test(banks []int) bool
}

type HitRunner struct {
	m map[string]struct{}
}

func (r *HitRunner) Test(banks []int) bool {
	i := fmt.Sprintf("%v", banks)
	if _, ok := r.m[i]; ok {
		return true
	}

	r.m[i] = struct{}{}
	return false
}

func run(banks []int, r Runner) {
	max := func() (int, int) {
		k, n := 0, banks[0]

		for i, p := range banks {
			if p > n {
				k, n = i, p
			}
		}

		return k, n
	}

	for {
		k, n := max()

		banks[k] = 0

		for i := 0; i != n; i++ {
			j := k + i + 1
			banks[j%len(banks)]++
		}

		if r.Test(banks) {
			break
		}
	}
}

func Run(banks []int) int {
	r := HitRunner{m: make(map[string]struct{})}
	run(banks, &r)
	return len(r.m) + 1
}

type LoopRunner struct {
	p []int
	n int
}

func (r *LoopRunner) Test(banks []int) bool {
	r.n++

	for i := range banks {
		if banks[i] != r.p[i] {
			return false
		}
	}

	return true
}

func Loop(banks []int) int {
	Run(banks)
	r := LoopRunner{p: append([]int{}, banks...)}
	run(banks, &r)
	return r.n
}
