package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 1[a|b] 'input'")
	}

	switch args[0] {
	case "1a", "1":
		fmt.Println(Sum(args[1]))
	case "1b":
		fmt.Println(Half(args[1]))
	}
}

func load(s string) string {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = strings.TrimSpace(string(f))
	}

	return s
}

func sum(s string, steps int) int {
	s = load(s)

	get := func(i int) (result rune, next int) {
		// let's not assume that the string is ASCII; should work with any utf-8 string
		for k, c := range s[i:] {
			if k != 0 {
				next = i + k
				break
			}

			result = c
		}

		return
	}

	match := func(i int) bool {
		a, j := get(i)
		b, _ := get((j + steps - 1) % len(s))
		return a == b
	}

	sum := 0
	for i, c := range s {
		if match(i) {
			sum += int(c) - int('0')
		}
	}

	return sum
}

func Sum(s string) int {
	return sum(s, 1)
}

func Half(s string) int {
	s = load(s)

	// make it work with any utf-8 string
	half := 0
	for range s {
		half++
	}

	return sum(s, half/2)
}
