package p23

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 23[a|b] filename")
	}

	switch args[0] {
	case "23a", "23":
		fmt.Print(Muls(args[1]))
	case "23b":
		fmt.Print(Debug())
	}
}

func load(s string) [][]string {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	m := make([][]string, 0)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		m = append(m, strings.Fields(line))
	}

	return m
}

func Muls(s string) int {
	m := load(s)

	ip, regs := 0, make([]int, 26)

	value := func(s string) int {
		if n := s[0]; n >= 'a' && n <= 'z' {
			return regs[n-'a']
		}

		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		return i
	}

	set := func(s string, i int) {
		regs[s[0]-'a'] = i
	}

	count := 0

	for ip >= 0 && ip < len(m) {
		op := m[ip]
		ip++

		switch op[0] {
		case "set":
			set(op[1], value(op[2]))
		case "sub":
			set(op[1], value(op[1])-value(op[2]))
		case "mul":
			set(op[1], value(op[1])*value(op[2]))
			count++
		case "jnz":
			if value(op[1]) != 0 {
				ip += value(op[2]) - 1
			}
		default:
			log.Fatal("unknown op: " + strings.Join(op, " "))
		}
	}

	return count
}

// Here we took the assembly dump and converted it to Go code manually.
// It's clear it's looking for primes.
// So we added 3 simple optimisation that gives equivalent code:
//
//   1. exit loops when f == 1
//   2. use a division to replace the inner loop
//   3. early exit on sqrt

func Debug() int {
	b, c, f, h := 0, 0, 0, 0
	d, e := 0, 0
	//set b 57
	//set c b
	//jnz a 2
	//jnz 1 5
	//mul b 100
	//sub b -100000
	//set c b
	//sub c -17000
	q := 57*100 + 100000
	i := 0
	for b, c = q, q+17000; b-17 != c; i++ {
		//set f 1
		f = 1
		//set d 2
		for d = 2; d != b && f == 1; d++ { // added && f == 1
			// optimize: replace inner loop
			if e = b / d; d*e == b {
				f = 0
			}
			// optimize: exit early
			if d*d > b {
				break
			}

			//set e 2
			for e = 2; e != b && f == 1; e++ { // added && f == 1
				//set g d
				//mul g e
				//sub g b
				//jnz g 2
				//set f 0
				if d*e == b {
					f = 0
				}

				//sub e -1
				//set g e
				//sub g b
				//jnz g -8
			}
			//sub d -1
			//set g d
			//sub g b
			//jnz g -13
		}
		//jnz f 2
		//sub h -1
		if f == 0 {
			h++
		}
		//set g b
		//sub g c
		//jnz g 2
		//jnz 1 3
		//sub b -17 // tricky thing here... increment AFTER the test
		//jnz 1 -23
		//log.Println(i, (c-b)/17, b, f, h)
		b += 17
	}

	return h
}
