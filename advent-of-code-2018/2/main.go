package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 2[a|b] 'input'")
	}

	switch args[0] {
	case "2a", "2":
		fmt.Println(Sum(args[1]))
	case "2b":
		fmt.Println(Compare(args[1]))
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

func countAppear(s string) (bool, bool) {
	freq := make([]int, 255)
	for _, c := range s {
		freq[int(c)]++
	}

	twice := false
	third := false
	for _, f := range freq {
		if f == 2 {
			twice = true
		}
		if f == 3 {
			third = true
		}
	}

	return twice, third
}

func sum(s string) int {
	res := load(s)

	freq := make([]int, 4)
	for _, r := range res {
		t, tt := countAppear(r)
		if t {
			freq[2]++
		}
		if tt {
			freq[3]++
		}
	}
	fmt.Printf("%v\n", freq)
	return freq[2] * freq[3]
}

func Sum(s string) int {
	return sum(s)
}

func diff(r string, r2 string) int {
	diff := 0
	for i, _ := range r {
		if r[i] != r2[i] {
			diff += 1
		}
	}

	return diff
}

func Compare(s string) string {
	data := load(s)

	for _, r := range data {
		for _, r2 := range data {
			if r == r2 {
				continue
			}
			if diff(r, r2) == 1 {
				res := ""
				for i, _ := range r {
					if r[i] == r2[i] {
						res = fmt.Sprintf("%s%c", res, r[i])
					}
				}
				return res
			}
		}
	}
	return ""
}
