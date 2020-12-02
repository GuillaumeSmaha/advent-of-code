package main

import (
	"fmt"
)

type Map3D struct {
	Map       map[[3]int]int
	MapLength map[[3]int]int
	Robots    []*Robot
	Start     [3]int
	InvertY   bool
}

func (m *Map3D) Init() {
	m.Map = map[[3]int]int{}
	m.MapLength = map[[3]int]int{}
	m.Robots = []*Robot{}
}

func (m *Map3D) Clone() *Map3D {
	nm := &Map3D{}
	*nm = *m
	nm.Map = make(map[[3]int]int, len(m.Map))
	nm.MapLength = make(map[[3]int]int, len(m.MapLength))
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

func (m *Map3D) AddRobot(r *Robot) {
	m.Robots = append(m.Robots, r)
}

func (m *Map3D) IsBorderXYZ(x, y, z int) bool {
	minX, minY, minZ, maxX, maxY, maxZ := m.Bounds()
	return minX == x || minY == y || minZ == z || maxX == x || maxY == y || maxZ == z
}

func (m *Map3D) ExpandBorder(caseSearch int, caseSet int, alignSize bool, keepCubic bool) {
	minX, minY, minZ, maxX, maxY, maxZ := m.Bounds()

	m.ForeachXYApply(func(m *Map3D, x, y, z int) {
		if m.GetXYZ(x, y, z) == caseSearch {
			if x == minX {
				m.SetXYZ(x-1, y, z, caseSet)
			} else if x == maxX {
				m.SetXYZ(x+1, y, z, caseSet)
			}
			if y == minY {
				m.SetXYZ(x, y-1, z, caseSet)
			} else if y == maxY {
				m.SetXYZ(x, y+1, z, caseSet)
			}
			if z == minZ {
				m.SetXYZ(x, y, z-1, caseSet)
			} else if z == maxZ {
				m.SetXYZ(x, y, z+1, caseSet)
			}

		}
	})

	if alignSize {
		for i := 0; i < 3; i++ {
			minX, minY, minZ, maxX, maxY, maxZ = m.Bounds()
			if -minX != maxX {
				if -minX > maxX {
					m.SetXYZ(-minX, minY, minZ, caseSet)
				} else {
					m.SetXYZ(-maxX, minY, minZ, caseSet)
				}
			}
			if -minY != maxY {
				if -minY > maxY {
					m.SetXYZ(minX, -minY, minZ, caseSet)
				} else {
					m.SetXYZ(minX, -maxY, minZ, caseSet)
				}
			}
			if -minZ != maxZ {
				if -minZ > maxZ {
					m.SetXYZ(minX, minY, -minZ, caseSet)
				} else {
					m.SetXYZ(minX, minY, -maxZ, caseSet)
				}
			}
			if keepCubic {
				if maxX != maxY {
					if maxX > maxY {
						m.SetXYZ(maxX, maxX, minZ, caseSet)
					} else {
						m.SetXYZ(maxY, maxY, minZ, caseSet)
					}
				}
				if maxX != maxZ {
					if maxX > maxZ {
						m.SetXYZ(maxX, minY, maxX, caseSet)
					} else {
						m.SetXYZ(maxZ, minY, maxZ, caseSet)
					}
				}
			}
		}
	}

	minX, minY, minZ, maxX, maxY, maxZ = m.BoundsWith(caseSearch)
	m.ForeachXYApply(func(m *Map3D, x, y, z int) {
		if x < minX-1 || x > maxX+1 || y < minY-1 || y > maxY+1 || z < minZ-1 || z > maxZ+1 {
			m.DeleteXYZ(x, y, z)
		}
	})
}

func (m *Map3D) GetXYZ(x, y, z int) int {
	if v, ok := m.Map[[3]int{x, y, z}]; ok {
		return v
	}
	return CASE_UNKNOW
}

func (m *Map3D) GetChar(x, y, z int) string {
	return CHARS_CASE[m.GetXYZ(x, y, z)]
}

func (m *Map3D) SetXYZ(x, y, z int, v int) {
	m.Map[[3]int{x, y, z}] = v
}

func (m *Map3D) DeleteXYZ(x, y, z int) {
	delete(m.Map, [3]int{x, y, z})
}

func (m *Map3D) ForeachXY() [][3]int {
	minX, minY, minZ, maxX, maxY, maxZ := m.Bounds()

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
	r := [][3]int{}
	for z := minZ; z <= maxZ; z++ {
		for y := yStart; yCmp(y); y += yInc {
			for x := minX; x <= maxX; x++ {
				r = append(r, [3]int{x, y, z})
			}
		}
	}
	return r
}

func (m *Map3D) ForeachXYApply(apply func(m *Map3D, x, y, z int)) {
	minX, minY, minZ, maxX, maxY, maxZ := m.Bounds()

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
	for z := minZ; z <= maxZ; z++ {
		for y := yStart; yCmp(y); y += yInc {
			for x := minX; x <= maxX; x++ {
				apply(m, x, y, z)
			}
		}
	}
}

func (m *Map3D) ForeachNeighbors(mx, my, mz int, withCenter bool, withBound bool, size ...int) [][3]int {
	s := 1
	if len(size) > 0 {
		s = size[0]
	}

	minX, minY, minZ, maxX, maxY, maxZ := m.Bounds()
	if withBound {
		minX = IntMax(minX, mx-s)
		minY = IntMax(minY, my-s)
		minZ = IntMax(minZ, mz-s)
		maxX = IntMin(maxX, mx+s)
		maxY = IntMax(maxY, my+s)
		maxZ = IntMax(maxZ, mz+s)
	} else {
		minX = mx - s
		minY = my - s
		minZ = mz - s
		maxX = mx + s
		maxY = my + s
		maxZ = mz + s
	}
	r := [][3]int{}
	for z := minZ; z <= maxZ; z++ {
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxY; x++ {
				if withCenter && x == mx && y == my && z == mz || x != mx || y != my || z != mz {
					r = append(r, [3]int{x, y, z})
				}
			}
		}
	}
	return r
}

func (m *Map3D) ForeachNeighborsApply(apply func(m *Map3D, x, y, z int), mx, my, mz int, withCenter, withBound bool, size ...int) {
	s := 1
	if len(size) > 0 {
		s = size[0]
	}

	minX, minY, minZ, maxX, maxY, maxZ := m.Bounds()
	if withBound {
		minX = IntMax(minX, mx-s)
		minY = IntMax(minY, my-s)
		minZ = IntMax(minZ, mz-s)
		maxX = IntMin(maxX, mx+s)
		maxY = IntMax(maxY, my+s)
		maxZ = IntMax(maxZ, mz+s)
	} else {
		minX = mx - s
		minY = my - s
		minZ = mz - s
		maxX = mx + s
		maxY = my + s
		maxZ = mz + s
	}
	for z := minZ; z <= maxZ; z++ {
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if withCenter && x == mx && y == my && z == mz || x != mx || y != my || z != mz {
					apply(m, x, y, z)
				}
			}
		}
	}
}

func (m *Map3D) GetRobotsMap() map[[3]int][]*Robot {
	robots := map[[3]int][]*Robot{}
	// for _, r := range m.Robots {
	// 	xyz := r.GetXYZSlice()
	// 	if _, ok := robots[xyz]; !ok {
	// 		robots[xyz] = []*Robot{}
	// 	}
	// 	robots[xyz] = append(robots[xy], r)
	// }
	return robots
}

func (m *Map3D) BoundsList() [6]int {
	minX, minY, minZ, maxX, maxY, maxZ := m.Bounds()
	return [6]int{minX, minY, minZ, maxX, maxY, maxZ}
}

func (m *Map3D) Bounds() (minX, minY, minZ, maxX, maxY, maxZ int) {
	for xy := range m.Map {
		minX = xy[0]
		minY = xy[1]
		minZ = xy[2]
		maxX = xy[0]
		maxY = xy[1]
		maxZ = xy[2]
		break
	}
	for xyz := range m.Map {
		x := xyz[0]
		y := xyz[1]
		z := xyz[2]
		if minX > x {
			minX = x
		}
		if minY > y {
			minY = y
		}
		if minZ > z {
			minZ = z
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
	}
	return
}

func (m *Map3D) BoundsWith(caseKind int) (minX, minY, minZ, maxX, maxY, maxZ int) {
	for xy, v := range m.Map {
		if v != caseKind {
			continue
		}
		minX = xy[0]
		minY = xy[1]
		minZ = xy[2]
		maxX = xy[0]
		maxY = xy[1]
		maxZ = xy[2]
		break
	}
	for xyz, v := range m.Map {
		if v != caseKind {
			continue
		}
		x := xyz[0]
		y := xyz[1]
		z := xyz[2]
		if minX > x {
			minX = x
		}
		if minY > y {
			minY = y
		}
		if minZ > z {
			minZ = z
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
	}
	return
}

func (m *Map3D) RecordLengthFor(x, y, z int) {
	min := 999999999
	for k := -1; k < 2; k++ {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 || j == 0 {
					if v, ok := m.MapLength[[3]int{x + i, y + j, z + k}]; ok {
						if min > v {
							min = v
						}
					}
				}
			}
		}
	}
	m.MapLength[[3]int{x, y, z}] = min + 1
}

func (m *Map3D) DrawMap(drawStart bool, drawRobot bool, drawMapLength bool) {
	robots := m.GetRobotsMap()
	minX, minY, minZ, maxX, maxY, maxZ := m.Bounds()

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
	for z := minZ; z <= maxZ; z++ {
		fmt.Printf("z = %d\n", z)
		for y := yStart; yCmp(y); y += yInc {
			for x := minX; x <= maxX; x++ {
				r, hasRobot := robots[[3]int{x, y, z}]
				if drawStart && m.Start[0] == x && m.Start[1] == y && m.Start[2] == z {
					fmt.Printf(CHARS_CASE[CASE_START])
				} else if drawRobot && hasRobot {
					fmt.Printf(r[0].DrawRobot())
				} else {
					c := m.GetChar(x, y, z)
					fmt.Printf(c)
				}
			}
			fmt.Printf("\n")
		}
	}

	if drawMapLength {
		fmt.Printf("\n")
		for z := minZ; z <= maxZ; z++ {
			fmt.Printf("z = %d\n", z)
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					r, hasRobot := robots[[3]int{x, y, z}]
					if drawRobot && hasRobot {
						fmt.Printf(CHARS_CASE[CASE_EMPTY])
						fmt.Printf(r[0].DrawRobot())
						fmt.Printf(CHARS_CASE[CASE_EMPTY])
					} else {
						v := m.GetXYZ(x, y, z)
						c := m.GetChar(x, y, z)
						if v == CASE_WALL {
							fmt.Printf(c)
							fmt.Printf(c)
							fmt.Printf(c)
						} else if v, ok := m.MapLength[[3]int{x, y, z}]; ok {
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
