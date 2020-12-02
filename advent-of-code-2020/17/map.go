package main

import (
	"fmt"
)

const (
	CASE_UNKNOW     = 0
	CASE_WALL       = 1
	CASE_EMPTY      = 2
	CASE_START      = 3
	CASE_INTERSEC   = 4
	CASE_WALL_CROSS = 5
)

var CHARS_CASE = map[int]string{
	CASE_UNKNOW:     "░",
	CASE_WALL:       "#",
	CASE_EMPTY:      ".",
	CASE_START:      "X",
	CASE_INTERSEC:   "O",
	CASE_WALL_CROSS: "░",
}

type Map struct {
	Map       map[[2]int]int
	MapLength map[[2]int]int
	Robots    []*Robot
	Start     [2]int
	InvertY   bool
}

func (m *Map) Init() {
	m.Map = map[[2]int]int{}
	m.MapLength = map[[2]int]int{}
	m.Robots = []*Robot{}
}

func (m *Map) Clone() *Map {
	nm := &Map{}
	*nm = *m
	nm.Map = make(map[[2]int]int, len(m.Map))
	nm.MapLength = make(map[[2]int]int, len(m.MapLength))
	nm.Robots = make([]*Robot, len(m.Robots))
	for k, v := range m.Map {
		kk := k
		nm.Map[kk] = v
	}
	for k, v := range m.MapLength {
		kk := k
		nm.MapLength[kk] = v
	}
	for k, v := range m.Robots {
		nm.Robots[k] = v
	}
	return nm
}

func (m *Map) AddRobot(r *Robot) {
	m.Robots = append(m.Robots, r)
}

func (m *Map) IsBorderXY(x, y, z int) bool {
	minX, minY, maxX, maxY := m.Bounds()
	return minX == x || minY == y || maxX == x || maxY == y
}

func (m *Map) ExpandBorder(caseSearch int, caseSet int, alignSize bool, keepSquare bool) {
	minX, minY, maxX, maxY := m.Bounds()

	m.ForeachXYApply(func(mm *Map, x, y int) {
		if m.GetXY(x, y) == caseSearch {
			if x == minX {
				m.SetXY(x-1, y, caseSet)
			} else if x == maxX {
				m.SetXY(x+1, y, caseSet)
			}
			if y == minY {
				m.SetXY(x, y-1, caseSet)
			} else if y == maxY {
				m.SetXY(x, y+1, caseSet)
			}

		}
	})

	if alignSize {
		for i := 0; i < 3; i++ {
			minX, minY, maxX, maxY = m.Bounds()
			if -minX != maxX {
				if -minX > maxX {
					m.SetXY(-minX, minY, caseSet)
				} else {
					m.SetXY(-maxX, minY, caseSet)
				}
			}
			if -minY != maxY {
				if -minY > maxY {
					m.SetXY(minX, -minY, caseSet)
				} else {
					m.SetXY(minX, -maxY, caseSet)
				}
			}
			if keepSquare {
				if maxX != maxY {
					if maxX > maxY {
						m.SetXY(maxX, maxX, caseSet)
					} else {
						m.SetXY(maxY, maxY, caseSet)
					}
				}
			}
		}
	}
}

func (m *Map) GetXY(x, y int) int {
	if v, ok := m.Map[[2]int{x, y}]; ok {
		return v
	}
	return CASE_UNKNOW
}

func (m *Map) GetChar(x, y int) string {
	return CHARS_CASE[m.GetXY(x, y)]
}

func (m *Map) SetXY(x, y int, v int) {
	m.Map[[2]int{x, y}] = v
}

func (m *Map) ForeachXY() [][2]int {
	minX, minY, maxX, maxY := m.Bounds()

	yStart := maxY
	yInc := -1
	yCmp := func(y int) bool {
		return y >= minY
	}
	if m.InvertY {
		yStart = minY
		yInc = 1
		yCmp = func(y int) bool {
			return y <= maxY
		}
	}
	r := [][2]int{}
	for y := yStart; yCmp(y); y += yInc {
		for x := minX; x <= maxX; x++ {
			r = append(r, [2]int{x, y})
		}
	}
	return r
}

func (m *Map) ForeachXYApply(apply func(m *Map, x, y int)) {
	minX, minY, maxX, maxY := m.Bounds()

	yStart := maxY
	yInc := -1
	yCmp := func(y int) bool {
		return y >= minY
	}
	if m.InvertY {
		yStart = minY
		yInc = 1
		yCmp = func(y int) bool {
			return y <= maxY
		}
	}
	for y := yStart; yCmp(y); y += yInc {
		for x := minX; x <= maxX; x++ {
			apply(m, x, y)
		}
	}
}

func (m *Map) ForeachNeighbors(mx, my int, withCenter bool, size ...int) [][2]int {
	s := 1
	if len(size) > 0 {
		s = size[0]
	}

	minX, minY, maxX, maxY := m.Bounds()
	minX = IntMax(minX, mx-s)
	minY = IntMax(minY, my-s)
	maxX = IntMin(maxX, mx+s)
	maxY = IntMax(maxY, my+s)
	r := [][2]int{}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxY; x++ {
			if withCenter && x == mx && y == my || (x != mx || y != my) {
				r = append(r, [2]int{x, y})
			}
		}
	}
	return r
}

func (m *Map) ForeachNeighborsApply(apply func(m *Map, x, y int), mx, my int, withCenter bool, size ...int) {
	s := 1
	if len(size) > 0 {
		s = size[0]
	}

	minX, minY, maxX, maxY := m.Bounds()
	minX = IntMax(minX, mx-s)
	minY = IntMax(minY, my-s)
	maxX = IntMin(maxX, mx+s)
	maxY = IntMin(maxY, my+s)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if withCenter && x == mx && y == my || (x != mx || y != my) {
				apply(m, x, y)
			}
		}
	}
}

func (m *Map) GetRobotsMap() map[[2]int][]*Robot {
	robots := map[[2]int][]*Robot{}
	for _, r := range m.Robots {
		xy := r.GetXYSlice()
		if _, ok := robots[xy]; !ok {
			robots[xy] = []*Robot{}
		}
		robots[xy] = append(robots[xy], r)
	}
	return robots
}

func (m *Map) BoundsList() [4]int {
	minX, minY, maxX, maxY := m.Bounds()
	return [4]int{minX, minY, maxX, maxY}
}

func (m *Map) Bounds() (minX, minY, maxX, maxY int) {
	for xy := range m.Map {
		minX = xy[0]
		minY = xy[1]
		maxX = xy[0]
		maxY = xy[1]
		break
	}
	for xy := range m.Map {
		x := xy[0]
		y := xy[1]
		if minX > x {
			minX = x
		}
		if minY > y {
			minY = y
		}
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
	}
	return
}

func (m *Map) RecordLengthFor(x, y int) {
	min := 999999999
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 || j == 0 {
				if v, ok := m.MapLength[[2]int{x + i, y + j}]; ok {
					if min > v {
						min = v
					}
				}
			}
		}
	}
	m.MapLength[[2]int{x, y}] = min + 1
}

func (m *Map) DrawMap(drawRobot bool, drawMapLength bool) {
	robots := m.GetRobotsMap()
	minX, minY, maxX, maxY := m.Bounds()

	yStart := maxY
	yInc := -1
	yCmp := func(y int) bool {
		return y >= minY
	}
	if m.InvertY {
		yStart = minY
		yInc = 1
		yCmp = func(y int) bool {
			return y <= maxY
		}
	}
	for y := yStart; yCmp(y); y += yInc {
		for x := minX; x <= maxX; x++ {
			r, hasRobot := robots[[2]int{x, y}]
			if m.Start[0] == x && m.Start[1] == y {
				fmt.Printf(CHARS_CASE[CASE_START])
			} else if drawRobot && hasRobot {
				fmt.Printf(r[0].DrawRobot())
			} else {
				c := m.GetChar(x, y)
				fmt.Printf(c)
			}
		}
		fmt.Printf("\n")
	}

	if drawMapLength {
		fmt.Printf("\n")
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				r, hasRobot := robots[[2]int{x, y}]
				if drawRobot && hasRobot {
					fmt.Printf(CHARS_CASE[CASE_EMPTY])
					fmt.Printf(r[0].DrawRobot())
					fmt.Printf(CHARS_CASE[CASE_EMPTY])
				} else {
					v := m.GetXY(x, y)
					c := m.GetChar(x, y)
					if v == CASE_WALL {
						fmt.Printf(c)
						fmt.Printf(c)
						fmt.Printf(c)
					} else if v, ok := m.MapLength[[2]int{x, y}]; ok {
						fmt.Printf("%03d", v)
					} else {
						fmt.Printf(CHARS_CASE[CASE_UNKNOW])
						fmt.Printf(CHARS_CASE[CASE_UNKNOW])
						fmt.Printf(CHARS_CASE[CASE_UNKNOW])
					}
				}
			}
			fmt.Printf("\n")
		}
	}
}
