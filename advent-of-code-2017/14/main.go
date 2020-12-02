package p14

import (
	"fmt"
	"log"
	"math/bits"
	"strconv"

	p10 "../10"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 14[a|b] code")
	}

	switch args[0] {
	case "14a", "14":
		fmt.Print(Used(args[1]))
	case "14b":
		fmt.Print(Regions(args[1]))
	}
}

func Used(key string) int {
	total := 0

	for i := 0; i < 128; i++ {
		h := p10.Hash(fmt.Sprintf("%s-%d", key, i))
		n := len(h) / 8
		for j := 0; j < n; j++ {
			u, err := strconv.ParseUint(h[:8], 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			h = h[8:]
			n := bits.OnesCount64(u)
			total += n
		}
	}

	return total
}

func Regions(key string) int {
	grid := make([]string, 128)

	for i := range grid {
		h := p10.Hash(fmt.Sprintf("%s-%d", key, i))
		n := len(h) / 8
		for j := 0; j < n; j++ {
			u, err := strconv.ParseUint(h[:8], 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			h = h[8:]
			grid[i] = fmt.Sprintf("%s%032b", grid[i], u)
		}

		if len(grid[i]) != 128 {
			log.Fatal("len should be exactly 128")
		}
	}

	labels, m := [128][128]int{}, make(map[int]map[int]struct{})

	for i, row := range grid {
		for j, c := range row {
			if c == '0' {
				continue
			}

			p, q := 0, 0

			if j != 0 {
				p = labels[i][j-1]
			}
			if i != 0 {
				q = labels[i-1][j]
			}

			if p != 0 {
				if q != 0 {
					if p != q {
						if p < q {
							labels[i][j] = p
							m[q][p] = struct{}{}
						} else {
							labels[i][j] = q
							m[p][q] = struct{}{}
						}
					} else {
						labels[i][j] = p
					}
				} else {
					labels[i][j] = p
				}
			} else {
				if q != 0 {
					labels[i][j] = q
				} else {
					n := len(m) + 1
					labels[i][j] = n
					m[n] = make(map[int]struct{})
				}
			}
		}
	}

	for i := len(m); i > 0; i-- {
		list := m[i]

		if len(list) == 0 {
			continue
		}

		min := len(m)

		for k := range list {
			if k < min {
				min = k
			}
		}

		for k := range list {
			if k != min {
				m[k][min] = struct{}{}
			}
		}
	}

	total := 0
	for _, list := range m {
		if len(list) == 0 {
			total++
		}
	}

	return total
}
