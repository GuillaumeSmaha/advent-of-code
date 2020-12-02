package p17

import (
	"fmt"
	"log"
	"strconv"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2017 17[a|b] input")
	}

	k, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "17a", "17":
		fmt.Print(Spin(k))
	case "17b":
		fmt.Print(SpinAngry(k))
	}
}

func Spin(k int) int {
	w := make([]int, 1)

	p := 0
	for n := 0; n < 2017; n++ {
		i := (p+k)%len(w) + 1
		w = append(w[:i], append([]int{n + 1}, w[i:]...)...)
		p = i
	}

	return w[p+1]
}

func SpinAngry(k int) int {
	last := 0

	w := 1
	p := 0
	for n := 0; n < 50000000; n++ {
		i := (p+k)%w + 1
		if i == 1 {
			last = n + 1
		}

		w++
		p = i
	}

	return last
}
