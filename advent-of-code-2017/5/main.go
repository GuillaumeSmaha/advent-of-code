package p5

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 5 'input'")
	}

	s := args[1]

	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	ints := make([]int, 0)

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		ints = append(ints, i)
	}

	switch args[0] {
	case "5a", "5":
		fmt.Print(Run(ints))
	case "5b":
		fmt.Print(Run2(ints))
	}
}

func Run(jumps []int) int {
	n, i := 0, 0

	for {
		k := jumps[i]
		jumps[i]++
		n++
		i += k

		if i < 0 || i >= len(jumps) {
			break
		}
	}

	return n
}

func Run2(jumps []int) int {
	n, i := 0, 0

	for {
		k := jumps[i]

		if k < 3 {
			jumps[i]++
		} else {
			jumps[i]--
		}

		n++
		i += k

		if i < 0 || i >= len(jumps) {
			break
		}
	}

	return n
}
