package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 1[a|b] 'input'")
	}

	switch args[0] {
	case "1a", "1":
		fmt.Println(Sum(args[1]))
	case "1b":
		fmt.Println(Twice(args[1]))
	}
}

func load(s string) []string {
	res := make([]string, 0)
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = strings.TrimSpace(string(f))

		for _, line := range strings.Split(s, "\n") {
			if len(line) == 0 {
				continue
			}
			res = append(res, line)
		}
	}

	return res
}

func sum(s string, steps int) int {
	res := load(s)

	sum := 0
	for _, r := range res {
		v, _ := strconv.Atoi(r[1:len(r)])
		if r[0] == '-' {
			sum -= v
		} else {
			sum += v
		}
	}

	return sum
}

func Sum(s string) int {
	return sum(s, 1)
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Twice(s string) int {
	res := load(s)

	freq := make([]int, 0)
	sum := 0
	i := 1
	for {
		fmt.Printf("Loop %v\n", i)
		for _, r := range res {
			v, _ := strconv.Atoi(r[1:len(r)])
			if r[0] == '-' {
				sum -= v
			} else {
				sum += v
			}
			if intInSlice(sum, freq) {
				return sum
			}
			freq = append(freq, sum)
		}
		i += 1
	}

	return 0
}
