package main

import (
	"fmt"
)

func processMain2(codes []string) {
	currentMap := &Map4D{
		InvertY: true,
	}
	currentMap.Init()

	lenZ := (len(codes) - 1) / 2
	for y, l := range codes {
		for x, c := range l {
			for z := -lenZ; z <= lenZ; z++ {
				for w := -lenZ; w <= lenZ; w++ {
					currentMap.SetXYZ(x-lenZ, y-lenZ, z, w, CASE_EMPTY)
				}
				if c == '#' {
					currentMap.SetXYZ(x-lenZ, y-lenZ, 0, 0, CASE_WALL)
				}
			}
		}
	}

	currentMap.ExpandBorder(CASE_WALL, CASE_EMPTY, false, false)
	currentMap.DrawMap(false, false, false)

	for i := 1; i <= 6; i++ {
		fmt.Println("------")
		fmt.Println("------")
		fmt.Println("------")
		fmt.Println("------")
		fmt.Println("step =", i)
		bak := currentMap.Clone()

		bak.ForeachXYApply(func(m *Map4D, x, y, z, w int) {
			cn := 0
			m.ForeachNeighborsApply(func(m *Map4D, x, y, z, w int) {
				v := m.GetXYZ(x, y, z, w)
				if v == CASE_WALL {
					cn++
				}

			}, x, y, z, w, false, false)
			// fmt.Println("x, y, z: ", x, y, z, ": cn: ", cn)
			v := m.GetXYZ(x, y, z, w)
			if v == CASE_WALL {
				if cn == 2 || cn == 3 {
					currentMap.SetXYZ(x, y, z, w, CASE_WALL)
				} else {
					currentMap.SetXYZ(x, y, z, w, CASE_EMPTY)
				}
			} else {
				if cn == 3 {
					currentMap.SetXYZ(x, y, z, w, CASE_WALL)
				} else {
					currentMap.SetXYZ(x, y, z, w, CASE_EMPTY)
				}
			}
		})

		// currentMap.DrawMap(false, false, false)
		fmt.Println("-%%%%%%%%%%%%%%%%%%%%%%%%%ù")
		currentMap.ExpandBorder(CASE_WALL, CASE_EMPTY, false, false)
		currentMap.DrawMap(false, false, false)

		cnt := 0
		currentMap.ForeachXYApply(func(m *Map4D, x, y, z, w int) {
			v := m.GetXYZ(x, y, z, w)
			if v == CASE_WALL {
				cnt++
			}
		})
		fmt.Println("Total: ", cnt)
	}
}

func main2() {
	// processMain2(parseFileText1D("list.test.txt"))
	processMain2(parseFileText1D("list.txt"))
}
