package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(f string) map[string]*Map {
	lines := parseFileText1D(f)
	maps := map[string]*Map{}
	var m *Map
	t := ""
	x, y := 0, 0
	lenY := -1
	for _, l := range lines[1:] {
		if l == "" {
			continue
		}
		i := strings.Index(l, "Tile")
		if i != -1 {
			break
		}
		lenY++
	}
	for _, l := range lines {
		if l == "" {
			continue
		}
		i := strings.Index(l, "Tile")
		if i != -1 {
			if m != nil && t != "" {
				maps[t] = m
			}
			t = l[5 : len(l)-1]
			m = &Map{
				// InvertY: true,
			}
			m.Init()
			x = 0
			y = 0
		} else {
			for _, c := range l {
				switch c {
				case '#':
					m.SetXY(x, lenY-y, CASE_WALL)
				case '.':
					m.SetXY(x, lenY-y, CASE_EMPTY)
				}
				x++
			}
			x = 0
			y++
		}
	}
	if m != nil && t != "" {
		maps[t] = m
	}
	return maps
}

type MatchMap struct {
	pos [2]int
	m1  *Map
	m2  *Map
}

func processMain1(f string) {
	maps := parse(f)

	match := map[string][]*MatchMap{}
	tileCnt := map[string]int{}
	matchDone := map[string]struct{}{}
	for t1, m1 := range maps {
		for t2, m2 := range maps {
			if t1 != t2 {
				_, ok1 := matchDone[t1+"-"+t2]
				_, ok2 := matchDone[t2+"-"+t1]

				if !ok1 && !ok2 {
					r := m1.AllRotateFlipsCheckBorderFull(m2)
					matchDone[t1+"-"+t2] = struct{}{}
					matchDone[t2+"-"+t1] = struct{}{}

					if len(r) > 0 {
						match[t1+"-"+t2] = r
						match[t2+"-"+t1] = r
						tileCnt[t1] += len(r)
						tileCnt[t2] += len(r)
					}
				}
			}
		}
	}

	for m, mm := range match {
		fmt.Printf("%v: %v\n", m, len(mm))
		for _, ma := range mm {
			fmt.Printf("\t pos: %#v, m1:%v, m2:%v\n", ma.pos, ma.m1.GetStatePretty(), ma.m2.GetStatePretty())
		}
	}

	// 2: corners
	// 3: borders
	// 4: middles
	res := 1
	fmt.Println("--")
	fmt.Println("Corners:")
	for m, c := range tileCnt {
		fmt.Printf("%v: %v: %v\n", m, c, c/8)
		if c/8 == 2 {
			fmt.Println("\t=> ok")
			// fmt.Printf("%v: %v\n", m, c/8)
			v, _ := strconv.Atoi(m)
			res *= v
		}
	}
	fmt.Printf("Result: %v\n", res)
}

func main1() {
	// processMain1("list.test.txt")
	processMain1("list.txt")
}
