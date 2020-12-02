package p2

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 2[a|b] 'input'")
	}

	n := 0

	switch args[0] {
	case "2a", "2":
		n = DiffSum(args[1])
	case "2b":
		n = EvenlyDivisibleSum(args[1])
	}

	fmt.Println(n)
}

func load(s string) [][]int {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	r := make([][]int, 0)

	for _, line := range strings.Split(s, "\n") {
		if len(line) == 0 {
			continue
		}

		ints := make([]int, 0)
		for i, f := range strings.Split(line, "\t") {
			n, err := strconv.Atoi(f)
			if err != nil {
				log.Fatalf("line %d: token %s: %s", i, f, err)
			}

			ints = append(ints, n)
		}

		r = append(r, ints)
	}

	return r
}

func DiffSum(s string) int {
	rows := load(s)

	diff := func(row []int) int {
		if len(row) == 0 {
			return 0
		}

		min, max := row[0], row[0]

		for _, k := range row {
			if k > max {
				max = k
			}

			if k < min {
				min = k
			}
		}

		return max - min
	}

	sum := 0

	for _, row := range rows {
		sum += diff(row)
	}

	return sum
}

func EvenlyDivisibleSum(s string) int {
	rows := load(s)

	check := func(row []int) int {
		sort.Ints(row)

		total := 0

		for i, x := range row {
			for j := 0; j != i; j++ {
				y := row[j]
				if k := x / y; k*y == x {
					total += k
				}

				if y*2 >= x {
					break
				}
			}
		}

		return total
	}

	sum := 0

	for _, row := range rows {
		sum += check(row)
	}

	return sum
}
