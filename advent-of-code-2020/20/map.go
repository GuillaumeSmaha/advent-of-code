package main

import (
	"fmt"
	"strconv"
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
	FlippedX  bool
	FlippedY  bool
	Rotation  int

	_cacheBounds []int
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
	nm._cacheBounds = nil
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

	m.ForeachXYApply(func(mm *Map, x, y, v int) {
		if v == caseSearch {
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
	// fmt.Println(m.GetStatePretty())
	// fmt.Println(m.BoundsList())
	// m.DrawMap(false, false, false)
	// panic(fmt.Sprintln("unknow x,y", x, y))
	return CASE_UNKNOW
}

func (m *Map) GetChar(x, y int) string {
	return CHARS_CASE[m.GetXY(x, y)]
}

func (m *Map) Fill(lenX, lenY int, v int) {
	for y := 0; y <= lenY; y++ {
		for x := 0; x <= lenX; x++ {
			m.SetXY(x, y, v)
		}
	}
}

func (m *Map) SetXY(x, y int, v int) {
	m._cacheBounds = nil
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

func (m *Map) ForeachXYApply(apply func(m *Map, x, y, v int)) {
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
			apply(m, x, y, m.GetXY(x, y))
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

func (m *Map) FlipX() {
	newM := map[[2]int]int{}
	_, _, _, maxY := m.Bounds()
	for _, xy := range m.ForeachXY() {
		v := m.Map[xy]
		xy[1] = maxY - xy[1]
		newM[xy] = v
	}
	m.Map = newM
	m.FlippedX = !m.FlippedX
	m._cacheBounds = nil
}

func (m *Map) FlipY() {
	newM := map[[2]int]int{}
	_, _, maxX, _ := m.Bounds()
	for _, xy := range m.ForeachXY() {
		v := m.Map[xy]
		xy[0] = maxX - xy[0]
		newM[xy] = v
	}
	m.Map = newM
	m.FlippedY = !m.FlippedY
	m._cacheBounds = nil
}

func (m *Map) RotateLeft() {
	newM := map[[2]int]int{}
	for _, xy := range m.ForeachXY() {
		newM[[2]int{-xy[1], xy[0]}] = m.Map[xy]
	}
	m.Map = newM
	m.Rotation = (m.Rotation - 90 + 360) % 360
	m._cacheBounds = nil
}

func (m *Map) RotateRight() {
	newM := map[[2]int]int{}
	for _, xy := range m.ForeachXY() {
		newM[[2]int{xy[1], -xy[0]}] = m.Map[xy]
	}
	m.Map = newM
	m.Rotation = (m.Rotation + 90) % 360
	m._cacheBounds = nil
}

func (m *Map) GetStatePretty() string {
	s := fmt.Sprintf("%d°", m.Rotation)
	if m.FlippedX {
		s += " flippedX"
	}
	if m.FlippedY {
		s += " flippedY"
	}
	return s
}

func (m *Map) GetState() string {
	s := fmt.Sprintf("%d", m.Rotation)
	if m.FlippedX {
		s += "X"
	}
	if m.FlippedY {
		s += "Y"
	}
	return s
}

func (m *Map) UndoState(clone bool) *Map {
	mm := m
	if clone {
		mm = m.Clone()
	}
	if m.FlippedX {
		mm.FlipX()
	}
	if m.FlippedY {
		mm.FlipY()
	}
	for mm.Rotation != 0 {
		mm.RotateLeft()
	}
	return mm
}

func (m *Map) ApplyState(state string, clone bool) *Map {
	mm := m
	if clone {
		mm = m.Clone()
	}
	n := len(state) - 1
	if state[n] == 'Y' && !mm.FlippedY {
		mm.FlipY()
		n--
	} else if state[n] != 'Y' && mm.FlippedY {
		mm.FlipY()
	}

	if state[n] == 'X' && !mm.FlippedX {
		mm.FlipX()
		n--
	} else if state[n] != 'X' && mm.FlippedX {
		mm.FlipX()
	}

	rot, err := strconv.Atoi(state[0 : n+1])
	if err != nil {
		panic(err)
	}
	for mm.Rotation != rot {
		mm.RotateLeft()
	}
	return mm
}

func (m *Map) AllRotateFlips() []*Map {
	maps := []*Map{}
	var mm *Map

	mm = m.Clone()
	maps = append(maps, mm)

	// Rotate 90
	mm = mm.Clone()
	mm.RotateRight()
	maps = append(maps, mm)

	// Rotate 180
	mm = mm.Clone()
	mm.RotateRight()
	maps = append(maps, mm)

	// Rotate 270
	mm = mm.Clone()
	mm.RotateRight()
	maps = append(maps, mm)

	// FlipX
	mm = m.Clone()
	mm.FlipX()
	maps = append(maps, mm)

	// FlipX & Rotate 90
	mm = mm.Clone()
	mm.RotateRight()
	maps = append(maps, mm)

	// FlipX & Rotate 180
	mm = mm.Clone()
	mm.RotateRight()
	maps = append(maps, mm)

	// FlipX & Rotate 270
	mm = mm.Clone()
	mm.RotateRight()
	maps = append(maps, mm)

	return maps
}

func (m *Map) CheckBorder(mo *Map) [][2]int {
	m1minX, m1minY, m1maxX, m1maxY := m.Bounds()
	m2minX, m2minY, m2maxX, m2maxY := mo.Bounds()
	lenX := m.LenX()
	lenY := m.LenY()
	if lenX != mo.LenX() {
		return [][2]int{}
	}
	if lenY != mo.LenY() {
		return [][2]int{}
	}

	borders := [][2]int{}

	x1 := 0
	x2 := 0
	y1 := m1maxY
	y2 := m2minY
	v := true
	for x := 0; x <= lenX; x++ {
		x1 = x + m1minX
		x2 = x + m2minX
		v = v && m.GetXY(x1, y1) == mo.GetXY(x2, y2)
		if !v {
			break
		}
	}
	if v {
		borders = append(borders, [2]int{0, -1})
	}

	x1 = 0
	x2 = 0
	y1 = m1minY
	y2 = m2maxY
	v = true
	for x := 0; x <= lenX; x++ {
		x1 = x + m1minX
		x2 = x + m2minX
		v = v && m.GetXY(x1, y1) == mo.GetXY(x2, y2)
		if !v {
			break
		}
	}
	if v {
		borders = append(borders, [2]int{0, 1})
	}

	x1 = m1minX
	x2 = m2maxX
	y1 = 0
	y2 = 0
	v = true
	for y := 0; y <= lenY; y++ {
		y1 = y + m1minY
		y2 = y + m2minY
		v = v && m.GetXY(x1, y1) == mo.GetXY(x2, y2)
		if !v {
			break
		}
	}
	if v {
		borders = append(borders, [2]int{-1, 0})
	}

	x1 = m1maxX
	x2 = m2minX
	y1 = 0
	y2 = 0
	v = true
	for y := 0; y <= lenY; y++ {
		y1 = y + m1minY
		y2 = y + m2minY
		v = v && m.GetXY(x1, y1) == mo.GetXY(x2, y2)
		if !v {
			break
		}
	}
	if v {
		borders = append(borders, [2]int{1, 0})
	}

	return borders
}

func (m *Map) AllRotateFlipsCheckBorder(mo *Map) []*MatchMap {
	r := []*MatchMap{}
	for _, m2 := range mo.AllRotateFlips() {
		vs := m.CheckBorder(m2)
		for _, v := range vs {
			r = append(r, &MatchMap{
				pos: v,
				m1:  m,
				m2:  m2,
			})
		}
	}
	return r
}

func (m *Map) AllRotateFlipsCheckBorderFull(mo *Map) []*MatchMap {
	r := []*MatchMap{}
	for _, m1 := range m.AllRotateFlips() {
		for _, m2 := range mo.AllRotateFlips() {
			vs := m1.CheckBorder(m2)
			for _, v := range vs {
				r = append(r, &MatchMap{
					pos: v,
					m1:  m1,
					m2:  m2,
				})
			}
		}
	}
	return r
}

func (m *Map) LenX() int {
	m1minX, _, m1maxX, _ := m.Bounds()
	return IntAbs(m1maxX-m1minX) + 1
}

func (m *Map) LenY() int {
	_, m1minY, _, m1maxY := m.Bounds()
	return IntAbs(m1maxY-m1minY) + 1
}

func (m *Map) AlignCoordsToZero() {
	minX, minY, _, _ := m.Bounds()
	mm := map[[2]int]int{}
	m.ForeachXYApply(func(m *Map, x, y, v int) {
		mm[[2]int{x - minX, y - minY}] = v
	})
	m.Map = mm
	m._cacheBounds = nil
}

func (m *Map) SearchMask(mask *Map) [][2]int {
	lenX := m.LenX()
	lenY := m.LenY()
	maskLenX := mask.LenX()
	maskLenY := mask.LenY()
	if lenX < maskLenX {
		return [][2]int{}
	}
	if lenY < maskLenY {
		return [][2]int{}
	}

	_, _, mapMaxX, mapMaxY := m.Bounds()
	maskMinX, maskMinY, _, _ := mask.Bounds()
	maxX := mapMaxX - maskLenX
	maxY := mapMaxY - maskLenY

	result := [][2]int{}
	m.ForeachXYApply(func(m *Map, x, y, _ int) {
		if x <= maxX && y <= maxY {
			fnd := true
			for j := 0; j < maskLenY && fnd; j++ {
				mY := y + j
				maskY := maskMinY + j
				for i := 0; i < maskLenX; i++ {
					mX := x + i
					maskX := maskMinX + i
					v := mask.GetXY(maskX, maskY)
					if v != CASE_UNKNOW && m.GetXY(mX, mY) != v {
						fnd = false
						break
					}
				}
			}
			if fnd {
				result = append(result, [2]int{x, y})
			}
		}
	})

	return result
}

func (m *Map) Bounds() (minX, minY, maxX, maxY int) {
	if m._cacheBounds != nil {
		minX = m._cacheBounds[0]
		minY = m._cacheBounds[1]
		maxX = m._cacheBounds[2]
		maxY = m._cacheBounds[3]
		return
	}

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
	m._cacheBounds = []int{minX, minY, maxX, maxY}
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

func (m *Map) DrawMap(drawStart bool, drawRobot bool, drawMapLength bool) {
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
			if drawStart && m.Start[0] == x && m.Start[1] == y {
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
