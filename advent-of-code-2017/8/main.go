package p7

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 8[a|b] 'filename'")
	}

	switch args[0] {
	case "8a", "8":
		fmt.Print(Maximum(args[1]))
	case "8b":
		fmt.Print(Highest(args[1]))
	}
}

type CPU map[string]int

func (cpu CPU) execute(s string) {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	num := func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		return i
	}

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		p := strings.Fields(line)
		cpu.run(p[0], p[1], num(p[2]), p[4], p[5], num(p[6]))
	}
}

func (cpu CPU) run(a, op string, i int, b, cond string, j int) {
	test := func(x int) bool {
		switch cond {
		case "<":
			return x < j
		case "<=":
			return x <= j
		case "==":
			return x == j
		case "!=":
			return x != j
		case ">":
			return x > j
		case ">=":
			return x >= j
		}
		log.Fatal("unknown condition")
		return false
	}

	if !test(cpu[b]) {
		return
	}

	switch op {
	case "inc":
		cpu[a] += i
	case "dec":
		cpu[a] -= i
	}

	if k := cpu[a]; k > cpu["_max"] {
		cpu["_max"] = k
	}
}

func Maximum(s string) int {
	cpu := CPU{}
	cpu.execute(s)

	first, max := true, 0
	for k, v := range cpu {
		if k[0] != '_' && (first || v > max) {
			max = v
			first = false
		}
	}

	return max
}

func Highest(s string) int {
	cpu := CPU{}
	cpu.execute(s)
	return cpu["_max"]
}
