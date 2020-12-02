package p15

import (
	"fmt"
	"log"
	"strconv"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 15[a|b] a b")
	}

	a, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	b, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "15a", "15":
		fmt.Print(Count(a, b))
	case "15b":
		fmt.Print(CountSlow(a, b))
	}
}

func Count(a, b int) int {
	total := 0
	p, q := int64(a), int64(b)

	for i := 0; i < 40000000; i++ {
		p = p * 16807
		p %= 2147483647
		q = q * 48271
		q %= 2147483647
		if p&0xffff == q&0xffff {
			total++
		}
	}

	return total
}

func CountSlow(a, b int) int {
	total := 0
	p, q := int64(a), int64(b)

	for i := 0; i < 5000000; i++ {
		for {
			p = p * 16807
			p %= 2147483647
			if p%4 == 0 {
				break
			}
		}

		for {
			q = q * 48271
			q %= 2147483647
			if q%8 == 0 {
				break
			}
		}

		if p&0xffff == q&0xffff {
			total++
		}
	}

	return total
}
