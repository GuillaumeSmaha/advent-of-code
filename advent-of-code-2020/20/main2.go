package main

import (
	"fmt"
)

type Tile struct {
	n string
	m *Map
}

func processMain2(f string) {
	maps := parse(f)

	match := map[string][]*MatchMap{}
	tileCnt := map[string]int{}
	matchDone := map[string]struct{}{}
	neighbors := map[string][]string{}
	for t1, _ := range maps {
		neighbors[t1] = []string{}
	}
	for t1, m1 := range maps {
		for t2, m2 := range maps {
			if t1 != t2 {
				_, ok1 := matchDone[t1+"-"+t2]
				if !ok1 {
					r := m1.AllRotateFlipsCheckBorderFull(m2)
					matchDone[t1+"-"+t2] = struct{}{}

					if len(r) > 0 {
						neighbors[t1] = append(neighbors[t1], t2)
						match[t1+"-"+t2] = r
						tileCnt[t1] += len(r)
					}
				}
			}
		}
	}

	// for m, mm := range match {
	// 	fmt.Printf("%v: %v\n", m, len(mm))
	// 	for _, ma := range mm {
	// 		fmt.Printf("\t pos: %#v, m1:%v, m2:%v\n", ma.pos, ma.m1.GetStatePretty(), ma.m2.GetStatePretty())
	// 	}
	// }

	// 2: corners
	// 3: borders
	// 4: middles
	corners := []string{}
	borders := []string{}
	middles := []string{}
	for m, c := range tileCnt {
		fmt.Printf("%v: %v: %v\n", m, c, c/8)
		switch c / 8 {
		case 2:
			corners = append(corners, m)
		case 3:
			borders = append(borders, m)
		case 4:
			middles = append(middles, m)
		}
	}

	fmt.Println("--")
	fmt.Println("Neighbors:")
	for t1, lt := range neighbors {
		fmt.Println(t1)
		for _, t2 := range lt {
			fmt.Printf("\t%v\n", t2)
		}
	}
	if len(corners) != 4 {
		fmt.Printf("Found less or more than 4 corners\n")
		fmt.Printf("----\n")
		fmt.Printf("Corners:\n")
		for _, m := range corners {
			fmt.Printf("\t- %v\n", m)
			for _, n := range neighbors[m] {
				fmt.Printf("\t\t- %v\n", n)
			}
		}
		return
	}

	start := ""
	end := ""
	fmt.Printf("----\n")
	fmt.Printf("Corners:\n")
	stateStart := ""
	i := 0
	for start == "" && end == "" {
		fmt.Printf("Look for corner with state: %v\n", i)
		for _, m := range corners {
			c := 0
			fmt.Printf("\t%v:\n", m)
			s := ""
			for _, n := range neighbors[m] {
				mm := match[fmt.Sprintf("%s-%s", m, n)]
				if mm[i].pos == [2]int{1, 0} || mm[i].pos == [2]int{0, -1} {
					fmt.Printf("\t\t- %v\n", mm[i].pos)
					c++
					s = mm[i].m1.GetState()
				}
			}
			fmt.Printf("\t%v: %v\n", m, c)
			if c == 2 && start == "" {
				start = m
				stateStart = s
			} else if c == 2 && end == "" {
				end = m
				break
			}
		}
		i++
	}
	// if start == "" {
	// 	fmt.Printf("----\n")
	// 	fmt.Printf("No start found, try with one:\n")
	// 	fmt.Printf("Corners:\n")
	// 	start = "1543"
	// 	for _, m := range corners {
	// 		c := 0
	// 		fmt.Printf("\t%v:\n", m)
	// 		for _, n := range neighbors[m] {
	// 			mm := match[fmt.Sprintf("%s-%s", m, n)]
	// 			if mm[0].pos == [2]int{1, 0} || mm[0].pos == [2]int{0, -1} {
	// 				fmt.Printf("\t\t- %v:\n", mm[0].pos)
	// 				c++
	// 			}
	// 		}
	// 		fmt.Printf("\t%v: %v\n", m, c)
	// 		if c == 1 && start == "" {
	// 			start = m
	// 		} else if c == 1 && end == "" {
	// 			end = m
	// 			break
	// 		}
	// 	}
	// }
	fmt.Printf("start: %v\n", start)
	fmt.Printf("end: %v\n", end)

	tileUsed := map[string]bool{}
	tiles := map[[2]int]*Tile{}

	x, y := 0, 0
	tileCurrent := start
	lastTile := start
	tileCurrentPos := [2]int{0, 0}
	tileUsed[start] = true
	tiles[tileCurrentPos] = &Tile{
		n: start,
		m: maps[start],
	}
	tiles[tileCurrentPos].m.ApplyState(stateStart, false)
	x = 1

	crossArray := func(a, b []string) []string {
		r := []string{}
		for _, i := range a {
			for _, j := range b {
				if i == j {
					r = append(r, i)
				}
			}
		}
		return r
	}

	fmt.Printf("---\nBuild big map:\n")
	for len(tileUsed) != len(tileCnt) {
		fnd := false
		state := tiles[tileCurrentPos].m.GetState()
		fmt.Printf("\tLook for %v/%v from tile %v\n", x, y, tileCurrent)
		neighborsTiles := neighbors[tileCurrent]
		fmt.Printf("\t\t neighbors: %#v\n", neighborsTiles)
		if x > 0 && y > 0 {
			upperTile := [2]int{x, y - 1}
			neighborsTilesUpper := neighbors[tiles[upperTile].n]
			fmt.Printf("\t\t neighbors from upper %#v\n", neighborsTilesUpper)
			neighborsTiles = crossArray(neighborsTiles, neighborsTilesUpper)
			fmt.Printf("\t\t crossed neighbors: %#v\n", neighborsTiles)
		}
		for _, n := range neighborsTiles {
			if b, _ := tileUsed[n]; !b {
				fmt.Printf("\t\tneighbor %v\n", n)
				mm := match[fmt.Sprintf("%s-%s", tileCurrent, n)]
				for _, mmm := range mm {
					// After x
					// checkPos := mmm.pos == [2]int{-1, 0} || mmm.pos == [2]int{1, 0}
					checkPos := mmm.pos == [2]int{1, 0}
					// checkPos := mmm.pos == [2]int{-1, 0}
					if x == 0 {
						// checkPos = mmm.pos == [2]int{0, -1} || mmm.pos == [2]int{0, 1}
						checkPos = mmm.pos == [2]int{0, -1}
					}
					fmt.Printf("\t\t\tmmm.pos: %v => checkPos: %v & state == mmm.m1.GetState(): %v == %v :%v\n", mmm.pos, checkPos, state, mmm.m1.GetState(), state == mmm.m1.GetState())
					if state == mmm.m1.GetState() && checkPos {
						fmt.Printf("\t\t\t -> save\n")
						tiles[[2]int{x, y}] = &Tile{
							n: n,
							m: mmm.m2,
						}

						tileCurrent = n
						tileUsed[n] = true
						lastTile = n
						tileCurrentPos = [2]int{x, y}

						x++
						s := tiles[[2]int{0, y}]
						if s.n != n && len(neighbors[n]) == len(neighbors[s.n]) {
							tileCurrent = s.n
							tileCurrentPos = [2]int{0, y}
							y++
							x = 0
						}
						fnd = true
						break
					}
				}
				if fnd {
					break
				}
			}
		}
	}
	end = lastTile
	fmt.Printf("end: %v\n", end)

	cntTileX := 0
	cntTileY := 0
	for xy := range tiles {
		cntTileX = IntMax(cntTileX, xy[0])
		cntTileY = IntMax(cntTileY, xy[1])
	}
	cntTileX++
	cntTileY++

	fmt.Printf("---\nTiles Map:\n")
	fmt.Println(cntTileY)
	fmt.Println(cntTileX)
	for y := 0; y < cntTileY; y++ {
		for x := 0; x < cntTileX; x++ {
			fmt.Printf("%s\t", tiles[[2]int{x, y}].n)
		}
		fmt.Println()
	}
	fmt.Printf("---\nTiles Map state:\n")
	for y := 0; y < cntTileY; y++ {
		for x := 0; x < cntTileX; x++ {
			fmt.Printf("%s\t", tiles[[2]int{x, y}].m.GetState())
		}
		fmt.Println()
	}

	fmt.Printf("---\nExport big map with seperation:\n")
	bigMap := &Map{
		InvertY: true,
	}
	bigMap.Init()
	lenTileX := tiles[[2]int{0, 0}].m.LenX()
	lenTileY := tiles[[2]int{0, 0}].m.LenY()
	bigMaxX := cntTileX * lenTileX
	bigMaxY := cntTileY * lenTileY
	fmt.Println("lenTileX:", lenTileX)
	fmt.Println("lenTileY:", lenTileY)
	fmt.Println("bigMaxX:", bigMaxX)
	fmt.Println("bigMaxY:", bigMaxY)
	for xy, m := range tiles {
		minX, minY, _, _ := m.m.Bounds()
		lenX := lenTileX + 1
		lenY := lenTileY + 1
		xbase := xy[0] * lenX
		ybase := xy[1] * lenY
		m.m.ForeachXYApply(func(m *Map, x, y, v int) {
			nx := xbase + x - minX
			ny := ybase + y - minY
			bigMap.SetXY(nx, ny, v)
		})
	}
	fmt.Println(bigMap.BoundsList())
	bigMap.DrawMap(false, false, false)

	fmt.Printf("---\nExport big map without seperation:\n")
	bigMap = &Map{
		InvertY: true,
	}
	bigMap.Init()
	for xy, m := range tiles {
		minX, minY, maxX, maxY := m.m.Bounds()
		lenX := lenTileX - 2
		lenY := lenTileY - 2
		xbase := xy[0] * lenX
		ybase := xy[1] * lenY
		// fmt.Printf("---: %v\n", xy)
		// fmt.Printf("\t bounds: %v\n", m.m.BoundsList())
		// fmt.Printf("\t lenX: %v\n", lenX)
		// fmt.Printf("\t lenY: %v\n", lenY)
		// fmt.Printf("\t xbase: %v\n", xbase)
		// fmt.Printf("\t ybase: %v\n", ybase)
		m.m.ForeachXYApply(func(m *Map, x, y, v int) {
			nx := xbase + x - minX
			ny := ybase + y - minY
			if x > minX && y > minY && x < maxX && y < maxY {
				// fmt.Printf("\t\t set %v/%v to %v/%v\n", nx, ny, x, y)
				bigMap.SetXY(nx, ny, v)
			}
		})
	}
	bigMap.AlignCoordsToZero()
	fmt.Println(bigMap.BoundsList())
	fmt.Println(bigMap.GetXY(1, 0))
	bigMap.DrawMap(false, false, false)

	fmt.Printf("---\nMask:\n")
	maskMap := &Map{
		// InvertY: true,
	}
	maskMap.Init()
	//                   #
	// #    ##    ##    ###
	//  #  #  #  #  #  #
	maskMap.SetXY(18, 0, CASE_WALL)
	maskMap.SetXY(0, 1, CASE_WALL)
	maskMap.SetXY(5, 1, CASE_WALL)
	maskMap.SetXY(6, 1, CASE_WALL)
	maskMap.SetXY(11, 1, CASE_WALL)
	maskMap.SetXY(12, 1, CASE_WALL)
	maskMap.SetXY(17, 1, CASE_WALL)
	maskMap.SetXY(18, 1, CASE_WALL)
	maskMap.SetXY(19, 1, CASE_WALL)
	maskMap.SetXY(1, 2, CASE_WALL)
	maskMap.SetXY(4, 2, CASE_WALL)
	maskMap.SetXY(7, 2, CASE_WALL)
	maskMap.SetXY(10, 2, CASE_WALL)
	maskMap.SetXY(13, 2, CASE_WALL)
	maskMap.SetXY(16, 2, CASE_WALL)
	// maskMap.Fill(1, 1, CASE_EMPTY)
	// maskMap.SetXY(0, 1, CASE_WALL)
	// maskMap.SetXY(1, 0, CASE_WALL)
	// maskMap.SetXY(1, 1, CASE_WALL)
	maskMap.DrawMap(false, false, false)

	fmt.Printf("---\nMasks found:\n")
	masks := make([][][2]int, 8)
	bigMaps := bigMap.AllRotateFlips()
	for i, m := range bigMaps {
		masks[i] = m.SearchMask(maskMap)
		fmt.Printf("\t%#v\n", masks[i])
	}

	fmt.Printf("---\nFinal big map:\n")
	finalBigMap := bigMap.Clone()
	for i, m := range bigMaps {
		fmt.Printf("\t mask on map %v\n", m.GetStatePretty())
		if len(masks[i]) > 0 {
			finalBigMap.ApplyState(m.GetState(), false)
			for _, xy := range masks[i] {
				fmt.Printf("\t\tMask: %#v\n", xy)
				maskMap.ForeachXYApply(func(m *Map, x, y, v int) {
					if v == CASE_WALL {
						finalBigMap.SetXY(xy[0]+x, xy[1]+y, CASE_INTERSEC)
					}
				})
			}
		}
	}
	finalBigMap.UndoState(false)
	finalBigMap.DrawMap(false, false, false)

	fmt.Printf("---\nCount:\n")
	cnt := 0
	finalBigMap.ForeachXYApply(func(m *Map, x, y, v int) {
		if v == CASE_WALL {
			cnt++
		}
	})

	fmt.Printf("Result: %v\n", cnt)
}
func main2() {
	// processMain2("list.test.txt")
	processMain2("list.txt")
}
