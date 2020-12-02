package main

import "fmt"

func processMain1(codes []string) {
	currentMap := &Map3D{
		InvertY: true,
	}
	currentMap.Init()

	lenZ := (len(codes) - 1) / 2
	for y, l := range codes {
		for x, c := range l {
			for z := -lenZ; z <= lenZ; z++ {
				currentMap.SetXYZ(x-lenZ, y-lenZ, z, CASE_EMPTY)
			}
			if c == '#' {
				currentMap.SetXYZ(x-lenZ, y-lenZ, 0, CASE_WALL)
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

		bak.ForeachXYApply(func(m *Map3D, x, y, z int) {
			cn := 0
			m.ForeachNeighborsApply(func(m *Map3D, x, y, z int) {
				v := m.GetXYZ(x, y, z)
				if v == CASE_WALL {
					cn++
				}

			}, x, y, z, false, false)
			// fmt.Println("x, y, z: ", x, y, z, ": cn: ", cn)
			v := m.GetXYZ(x, y, z)
			if v == CASE_WALL {
				if cn == 2 || cn == 3 {
					currentMap.SetXYZ(x, y, z, CASE_WALL)
				} else {
					currentMap.SetXYZ(x, y, z, CASE_EMPTY)
				}
			} else {
				if cn == 3 {
					currentMap.SetXYZ(x, y, z, CASE_WALL)
				} else {
					currentMap.SetXYZ(x, y, z, CASE_EMPTY)
				}
			}
		})

		// currentMap.DrawMap(false, false, false)
		fmt.Println("-%%%%%%%%%%%%%%%%%%%%%%%%%ù")
		currentMap.ExpandBorder(CASE_WALL, CASE_EMPTY, false, false)
		currentMap.DrawMap(false, false, false)

		cnt := 0
		currentMap.ForeachXYApply(func(m *Map3D, x, y, z int) {
			v := m.GetXYZ(x, y, z)
			if v == CASE_WALL {
				cnt++
			}
		})
		fmt.Println("Total: ", cnt)
	}

}

func main1() {
	processMain1(parseFileText1D("list.test.txt"))
	processMain1(parseFileText1D("list.txt"))
}
