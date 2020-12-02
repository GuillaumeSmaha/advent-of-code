package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 5[a|b] 'input'")
	}

	switch args[0] {
	case "5a", "5":
		fmt.Println(Part1(args[1]))
	case "5b":
		fmt.Println(Part2(args[1]))
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

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getSize(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func compress(s string) (string, int) {
	res := ""
	prevC := '\n'
	cnt := 0

	for _, c := range s {
		if prevC != '\n' {
			if Abs(int(c)-int(prevC)) == 32 {
				res = res[:len(res)-1]
				prevC = '\n'
				cnt++
				continue
			} else {
				res = fmt.Sprintf("%v%c", res, c)
			}
		} else {
			res = fmt.Sprintf("%v%c", res, c)
		}
		prevC = c
	}

	return res, cnt
}

func Part1(s string) string {
	str := load(s)
	res := str[0]
	cnt := 1
	for cnt > 0 {
		res, cnt = compress(res)
	}
	return res
}

func Part2(s string) int {
	str := load(s)
	minCnt := -1
	for z := 97; z < 123; z++ {
		newS := ""
		for _, c := range str[0] {
			if int(c) != z && int(c) != z-32 {
				newS = fmt.Sprintf("%v%c", newS, c)
			}
		}
		cnt := 1
		res := newS
		for cnt > 0 {
			res, cnt = compress(res)
		}
		if minCnt == -1 {
			minCnt = len(res)
		}
		fmt.Printf("%c: %v\n", z, len(res))
		minCnt = min(minCnt, len(res))
	}
	return minCnt
}
