package p3

import (
	"fmt"
	"log"
	"strconv"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 3[a|b] number")
	}

	i, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "3a", "3":
		fmt.Println(Steps(i))
	case "3b":
		fmt.Println(Sums(i))
	}
}

func cell(k int) (int, int, int) {
	if k == 0 || k == 1 {
		return 0, 0, 0
	}

	// here is my thinking:
	// if we label each ring with r=0, r=1, r=2... they end with d=(2r+1)^2
	// so r=0;d=1 r=1;d=9 r=2;d=25
	// each ring goes from p+1=d(r-1)+1=(2r-1)^2+1 to d i.e. from 2 to 9, from 10 to 25
	// and thus if q=2r+1, the number of cells is n=4(q-1)
	// let's label them from 1 to n by substracting their cell number k with p
	// and if we % with n, we got a nice 0 based indexing starting at the lower right corner
	// i=(k-p)%n

	square := func(x int) int { return x * x }
	r := 1
	for square(r*2+1) < k {
		r++
	}

	q := r*2 + 1
	p := square(r*2 - 1)
	n := 4 * (q - 1)
	i := (k - p) % n

	// the first q are vertical up and the q from h=n/2=2q-2 are vertical down
	// so for q=5 we have h=8 and i=
	// 0 | 1,2,3 | 4 | 5,6,7 | 8 | 9,10,11 | 12 | 13-14-15
	// if we % h, the first q=5 are vertical and the last q=5 are horizontal including corners
	// thus dx and dy have a range of +/- u=q/2
	// dy varies in the 1st half of h while dx varies on the 2nd half of h
	// thus, dy is centered on u and dx on h-u
	// if we say that j=i%h and dx=j-(h-u) and dy=j-u, we will get the following for q=5:
	//   dx: -6 | -5,-4,-3 | -2 | -1,0,1 | -6 | -5,-4,-3 | -2 | -1,0,1
	//   dy: -2 | -1, 0, 1 |  2 |  3,4,5 | -2 | -1, 0, 1 |  2 |  3,4,5
	// for dx, we must invert the sign of the 1st half
	// idem for the 2nd half of dy
	// if we want only the number of steps, we can use abs anyway

	h := n / 2
	u := q / 2
	j := i % h
	dx := j - (h - u)
	dy := j - u

	if i < h {
		dx = -dx
	} else {
		dy = -dy
	}

	return dx, dy, u
}

func Steps(k int) int {
	dx, dy, u := cell(k)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	max := func(x int) int {
		if x > u {
			return u
		}
		return x
	}

	dx = max(abs(dx))
	dy = max(abs(dy))

	return dx + dy
}

func Sums(stop int) int {
	m := map[string]int{
		"0,0": 1,
	}

	// my thinking here is that:
	// we will compute the x,y position of each cell and keep that in a map
	// with dx +/- 1 and dy +/- 1 we can lookup and sum

	look := func(x, y int) int {
		i := fmt.Sprintf("%d,%d", x, y)
		return m[i]
	}

	around := func(k int) int {
		x, y, u := cell(k)

		max := func(x int) int {
			if x < -u {
				return -u
			}
			if x > u {
				return u
			}
			return x
		}

		x = max(x)
		y = max(y)
		value := 0
		value += look(x-1, y-1)
		value += look(x-1, y)
		value += look(x-1, y+1)
		value += look(x, y-1)
		value += look(x, y+1)
		value += look(x+1, y-1)
		value += look(x+1, y)
		value += look(x+1, y+1)
		m[fmt.Sprintf("%d,%d", x, y)] = value
		return value
	}

	sum := 1
	for k := 2; k <= stop; k++ {
		sum = around(k)
	}

	return sum
}
