package p9

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 9[a|b] 'input'")
	}

	switch args[0] {
	case "9a", "9":
		fmt.Print(Score(args[1]))
	case "9b":
		fmt.Print(Garbage(args[1]))
	}
}

func load(s string) string {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	return s
}

func group(s string, score int) (int, int, int) {
	s = load(s)

	n := score
	k := 0
	p := 0

	for {
		ignore := false

		for i, c := range s[k:] {
			if ignore {
				ignore = false
				continue
			}

			if c == '}' {
				return k + i + 1, n, p
			}

			if c == '{' {
				k += i + 1
				j, m, q := group(s[k:], score+1)
				k += j
				n += m
				p += q
				break
			}

			if c == '<' {
				k += i + 1
				j, q := garbage(s[k:])
				k += j
				p += q
				break
			}

			if c == '!' {
				ignore = true
			}
		}

		if k >= len(s)-1 {
			break
		}
	}

	return k, n, p
}

func garbage(s string) (int, int) {
	ignore := false

	n := 0
	for i, c := range s {
		if ignore {
			ignore = false
			continue
		}

		if c == '>' {
			return i, n
		}

		if c == '!' {
			ignore = true
			continue
		}

		n++
	}

	log.Println("garbage until EOS")
	return -1, n
}

func Score(s string) int {
	_, n, _ := group(s, 0)
	return n
}

func Garbage(s string) int {
	_, _, p := group(s, 0)
	return p
}
