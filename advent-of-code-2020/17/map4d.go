package main

import (
	"fmt"
)

type Map4D struct {
	Map       map[[4]int]int
	MapLength map[[4]int]int
	Robots    []*Robot
	Start     [4]int
	InvertY   bool
}

func (m *Map4D) Init() {
	m.Map = map[[4]int]int{}
	m.MapLength = map[[4]int]int{}
	m.Robots = []*Robot{}
}

func (m *Map4D) Clone() *Map4D {
	nm := &Map4D{}
	*nm = *m
	nm.Map = make(map[[4]int]int, len(m.Map))
	nm.MapLength = make(map[[4]int]int, len(m.MapLength))
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

func (m *Map4D) AddRobot(r *Robot) {
	m.Robots = append(m.Robots, r)
}

func (m *Map4D) IsBorderXYZ(x, y, z, w int) bool {
	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW := m.Bounds()
	return minX == x || minY == y || minZ == z || minW == w || maxX == x || maxY == y || maxZ == z || maxW == w
}

func (m *Map4D) ExpandBorder(caseSearch int, caseSet int, alignSize bool, keepCubic bool) {
	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW := m.Bounds()

	m.ForeachXYApply(func(m *Map4D, x, y, z, w int) {
		if m.GetXYZ(x, y, z, w) == caseSearch {
			if x == minX {
				m.SetXYZ(x-1, y, z, w, caseSet)
			} else if x == maxX {
				m.SetXYZ(x+1, y, z, w, caseSet)
			}
			if y == minY {
				m.SetXYZ(x, y-1, z, w, caseSet)
			} else if y == maxY {
				m.SetXYZ(x, y+1, z, w, caseSet)
			}
			if z == minZ {
				m.SetXYZ(x, y, z-1, w, caseSet)
			} else if z == maxZ {
				m.SetXYZ(x, y, z+1, w, caseSet)
			}
			if w == minW {
				m.SetXYZ(x, y, z, w-1, caseSet)
			} else if w == maxW {
				m.SetXYZ(x, y, z, w+1, caseSet)
			}

		}
	})

	if alignSize {
		for i := 0; i < 3; i++ {
			minX, minY, minZ, minW, maxX, maxY, maxZ, maxW = m.Bounds()
			if -minX != maxX {
				if -minX > maxX {
					m.SetXYZ(-minX, minY, minZ, minW, caseSet)
				} else {
					m.SetXYZ(-maxX, minY, minZ, minW, caseSet)
				}
			}
			if -minY != maxY {
				if -minY > maxY {
					m.SetXYZ(minX, -minY, minZ, minW, caseSet)
				} else {
					m.SetXYZ(minX, -maxY, minZ, minW, caseSet)
				}
			}
			if -minZ != maxZ {
				if -minZ > maxZ {
					m.SetXYZ(minX, minY, -minZ, minW, caseSet)
				} else {
					m.SetXYZ(minX, minY, -maxZ, minW, caseSet)
				}
			}
			if -minW != maxW {
				if -minW > maxW {
					m.SetXYZ(minX, minY, minZ, -maxW, caseSet)
				} else {
					m.SetXYZ(minX, minY, minZ, -maxW, caseSet)
				}
			}
			if keepCubic {
				if maxX != maxY {
					if maxX > maxY {
						m.SetXYZ(maxX, maxX, minZ, minW, caseSet)
					} else {
						m.SetXYZ(maxY, maxY, minZ, minW, caseSet)
					}
				}
				if maxX != maxZ {
					if maxX > maxZ {
						m.SetXYZ(maxX, minY, maxX, minW, caseSet)
					} else {
						m.SetXYZ(maxZ, minY, maxZ, minW, caseSet)
					}
				}
				if maxX != maxW {
					if maxX > maxW {
						m.SetXYZ(maxX, minY, minZ, maxX, caseSet)
					} else {
						m.SetXYZ(maxW, minY, minZ, maxW, caseSet)
					}
				}
			}
		}
	}

	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW = m.BoundsWith(caseSearch)
	m.ForeachXYApply(func(m *Map4D, x, y, z, w int) {
		if x < minX-1 || x > maxX+1 || y < minY-1 || y > maxY+1 || z < minZ-1 || z > maxZ+1 || w < minW-1 || w > maxW+1 {
			m.DeleteXYZ(x, y, z, w)
		}
	})
}

func (m *Map4D) GetXYZ(x, y, z, w int) int {
	if v, ok := m.Map[[4]int{x, y, z, w}]; ok {
		return v
	}
	return CASE_UNKNOW
}

func (m *Map4D) GetChar(x, y, z, w int) string {
	return CHARS_CASE[m.GetXYZ(x, y, z, w)]
}

func (m *Map4D) SetXYZ(x, y, z, w int, v int) {
	m.Map[[4]int{x, y, z, w}] = v
}

func (m *Map4D) DeleteXYZ(x, y, z, w int) {
	delete(m.Map, [4]int{x, y, z, w})
}

func (m *Map4D) ForeachXY() [][4]int {
	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW := m.Bounds()

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
	r := [][4]int{}
	for w := minW; w <= maxW; w++ {
		for z := minZ; z <= maxZ; z++ {
			for y := yStart; yCmp(y); y += yInc {
				for x := minX; x <= maxX; x++ {
					r = append(r, [4]int{x, y, z, w})
				}
			}
		}
	}
	return r
}

func (m *Map4D) ForeachXYApply(apply func(m *Map4D, x, y, z, w int)) {
	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW := m.Bounds()

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
	for w := minW; w <= maxW; w++ {
		for z := minZ; z <= maxZ; z++ {
			for y := yStart; yCmp(y); y += yInc {
				for x := minX; x <= maxX; x++ {
					apply(m, x, y, z, w)
				}
			}
		}
	}
}

func (m *Map4D) ForeachNeighbors(mx, my, mz, mw int, withCenter bool, withBound bool, size ...int) [][4]int {
	s := 1
	if len(size) > 0 {
		s = size[0]
	}

	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW := m.Bounds()
	if withBound {
		minX = IntMax(minX, mx-s)
		minY = IntMax(minY, my-s)
		minZ = IntMax(minZ, mz-s)
		minW = IntMax(minW, mz-s)
		maxX = IntMin(maxX, mx+s)
		maxY = IntMax(maxY, my+s)
		maxZ = IntMax(maxZ, mz+s)
		maxW = IntMax(maxW, mz+s)
	} else {
		minX = mx - s
		minY = my - s
		minZ = mz - s
		minW = mw - s
		maxX = mx + s
		maxY = my + s
		maxZ = mz + s
		maxW = mw + s
	}
	r := [][4]int{}
	for w := minW; w <= maxW; w++ {
		for z := minZ; z <= maxZ; z++ {
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxY; x++ {
					if withCenter && x == mx && y == my && z == mz && w == mw || x != mx || y != my || z != mz || w != mw {
						r = append(r, [4]int{x, y, z, w})
					}
				}
			}
		}
	}
	return r
}

func (m *Map4D) ForeachNeighborsApply(apply func(m *Map4D, x, y, z, w int), mx, my, mz, mw int, withCenter, withBound bool, size ...int) {
	s := 1
	if len(size) > 0 {
		s = size[0]
	}

	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW := m.Bounds()
	if withBound {
		minX = IntMax(minX, mx-s)
		minY = IntMax(minY, my-s)
		minZ = IntMax(minZ, mz-s)
		minW = IntMax(minW, mw-s)
		maxX = IntMin(maxX, mx+s)
		maxY = IntMax(maxY, my+s)
		maxZ = IntMax(maxZ, mz+s)
		maxW = IntMax(maxW, mw+s)
	} else {
		minX = mx - s
		minY = my - s
		minZ = mz - s
		minW = mw - s
		maxX = mx + s
		maxY = my + s
		maxZ = mz + s
		maxW = mw + s
	}
	for w := minW; w <= maxW; w++ {
		for z := minZ; z <= maxZ; z++ {
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					if withCenter && x == mx && y == my && z == mz && w == mw || x != mx || y != my || z != mz || w != mw {
						apply(m, x, y, z, w)
					}
				}
			}
		}
	}
}

func (m *Map4D) GetRobotsMap() map[[4]int][]*Robot {
	robots := map[[4]int][]*Robot{}
	// for _, r := range m.Robots {
	// 	xyz := r.GetXYZSlice()
	// 	if _, ok := robots[xyz]; !ok {
	// 		robots[xyz] = []*Robot{}
	// 	}
	// 	robots[xyz] = append(robots[xy], r)
	// }
	return robots
}

func (m *Map4D) BoundsList() [8]int {
	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW := m.Bounds()
	return [8]int{minX, minY, minZ, minW, maxX, maxY, maxZ, maxW}
}

func (m *Map4D) Bounds() (minX, minY, minZ, minW, maxX, maxY, maxZ, maxW int) {
	for xy := range m.Map {
		minX = xy[0]
		minY = xy[1]
		minZ = xy[2]
		minW = xy[3]
		maxX = xy[0]
		maxY = xy[1]
		maxZ = xy[2]
		maxW = xy[3]
		break
	}
	for xyz := range m.Map {
		x := xyz[0]
		y := xyz[1]
		z := xyz[2]
		w := xyz[3]
		if minX > x {
			minX = x
		}
		if minY > y {
			minY = y
		}
		if minZ > z {
			minZ = z
		}
		if minW > w {
			minW = w
		}
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
		if maxZ < z {
			maxZ = z
		}
		if maxW < w {
			maxW = w
		}
	}
	return
}

func (m *Map4D) BoundsWith(caseKind int) (minX, minY, minZ, minW, maxX, maxY, maxZ, maxW int) {
	for xy, v := range m.Map {
		if v != caseKind {
			continue
		}
		minX = xy[0]
		minY = xy[1]
		minZ = xy[2]
		minW = xy[3]
		maxX = xy[0]
		maxY = xy[1]
		maxZ = xy[2]
		maxW = xy[3]
		break
	}
	for xyz, v := range m.Map {
		if v != caseKind {
			continue
		}
		x := xyz[0]
		y := xyz[1]
		z := xyz[2]
		w := xyz[3]
		if minX > x {
			minX = x
		}
		if minY > y {
			minY = y
		}
		if minZ > z {
			minZ = z
		}
		if minW > w {
			minW = w
		}
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
		if maxZ < z {
			maxZ = z
		}
		if maxW < w {
			maxW = w
		}
	}
	return
}

func (m *Map4D) RecordLengthFor(x, y, z, w int) {
	min := 999999999
	for k := -1; k < 2; k++ {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 || j == 0 {
					if v, ok := m.MapLength[[4]int{x + i, y + j, z + k}]; ok {
						if min > v {
							min = v
						}
					}
				}
			}
		}
	}
	m.MapLength[[4]int{x, y, z, w}] = min + 1
}

func (m *Map4D) DrawMap(drawStart bool, drawRobot bool, drawMapLength bool) {
	robots := m.GetRobotsMap()
	minX, minY, minZ, minW, maxX, maxY, maxZ, maxW := m.Bounds()

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
	for w := minW; w <= maxW; w++ {
		for z := minZ; z <= maxZ; z++ {
			fmt.Printf("z = %d\tw = %d\n", z, w)
			for y := yStart; yCmp(y); y += yInc {
				for x := minX; x <= maxX; x++ {
					r, hasRobot := robots[[4]int{x, y, z, w}]
					if drawStart && m.Start[0] == x && m.Start[1] == y && m.Start[2] == z {
						fmt.Printf(CHARS_CASE[CASE_START])
					} else if drawRobot && hasRobot {
						fmt.Printf(r[0].DrawRobot())
					} else {
						c := m.GetChar(x, y, z, w)
						fmt.Printf(c)
					}
				}
				fmt.Printf("\n")
			}
		}
	}

	if drawMapLength {
		fmt.Printf("\n")
		for w := minW; w <= maxW; w++ {
			for z := minZ; z <= maxZ; z++ {
				fmt.Printf("z = %d\tw = %d\n", z, w)
				for y := minY; y <= maxY; y++ {
					for x := minX; x <= maxX; x++ {
						r, hasRobot := robots[[4]int{x, y, z, w}]
						if drawRobot && hasRobot {
							fmt.Printf(CHARS_CASE[CASE_EMPTY])
							fmt.Printf(r[0].DrawRobot())
							fmt.Printf(CHARS_CASE[CASE_EMPTY])
						} else {
							v := m.GetXYZ(x, y, z, w)
							c := m.GetChar(x, y, z, w)
							if v == CASE_WALL {
								fmt.Printf(c)
								fmt.Printf(c)
								fmt.Printf(c)
							} else if v, ok := m.MapLength[[4]int{x, y, z, w}]; ok {
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
	}
}
