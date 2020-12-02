package p11

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 11[a|b] 'input'")
	}

	switch args[0] {
	case "11a", "11":
		fmt.Print(Distance(args[1]))
	case "11b":
		fmt.Print(Furthest(args[1]))
	}
}

func walk(s string) (int, int) {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	x := 0
	y := 0
	z := 0

	abs := func(x int) int {
		if x >= 0 {
			return x
		}
		return -x
	}

	max := func(x ...int) int {
		p := x[0]
		for _, i := range x {
			if i > p {
				p = i
			}
		}
		return p
	}

	distance := func() int {
		return max(abs(x), abs(y), abs(z))
	}

	far := 0

	for _, d := range strings.Split(s, ",") {
		switch d {
		case "n":
			x--
			y++
		case "s":
			x++
			y--
		case "ne":
			y++
			z--
		case "sw":
			y--
			z++
		case "nw":
			x--
			z++
		case "se":
			x++
			z--
		}

		d := max(abs(x), abs(y), abs(z))
		if d > far {
			far = d
		}
	}

	return distance(), far
}

func Distance(s string) int {
	d, _ := walk(s)
	return d
}

func Furthest(s string) int {
	_, d := walk(s)
	return d
}
