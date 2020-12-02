package p10

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 10[a|b] 1,2,3,4... [a,b,c,d...]")
	}

	list := ""
	if len(args) == 3 {
		list = args[2]
	}

	switch args[0] {
	case "10a", "10":
		fmt.Print(Multiply(list, args[1]))
	case "10b":
		fmt.Print(Hash(args[1]))
	}
}

func ints(s string) []int {
	p := make([]int, 0)

	for _, item := range strings.Split(s, ",") {
		if item == "" {
			continue
		}

		i, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}

		p = append(p, i)
	}

	return p
}

func Multiply(list, length string) int {
	p := ints(list)
	q := ints(length)

	if len(p) == 0 {
		for i := 0; i < 256; i++ {
			p = append(p, i)
		}
	}

	reverse := func(a, b int) {
		n := len(p)
		for i, j := a, b-1; i < j; i, j = i+1, j-1 {
			p[i%n], p[j%n] = p[j%n], p[i%n]
		}
	}

	i, skip := 0, 0

	for _, size := range q {
		reverse(i, i+size)
		i += size + skip
		skip++
	}

	return p[0] * p[1]
}

func round(p, q []byte, i, skip int) (int, int) {
	reverse := func(a, b int) {
		n := len(p)
		for i, j := a, b-1; i < j; i, j = i+1, j-1 {
			p[i%n], p[j%n] = p[j%n], p[i%n]
		}
	}

	for _, size := range q {
		n := int(size)
		reverse(i, i+n)
		i += n + skip
		skip++
	}

	return i, skip
}

func join(p []byte) byte {
	x := p[0]
	for _, k := range p[1:] {
		x ^= k
	}
	return x
}

func Hash(length string) string {
	q := append([]byte(length), 17, 31, 73, 47, 23)

	p := make([]byte, 256)
	for i := range p {
		p[i] = byte(i)
	}

	i, skip := 0, 0

	for n := 0; n < 64; n++ {
		i, skip = round(p, q, i, skip)
	}

	dense := make([]byte, 16)

	for i := range dense {
		n := 16 * i
		dense[i] = join(p[n : n+16])
	}

	return hex.EncodeToString(dense)
}
