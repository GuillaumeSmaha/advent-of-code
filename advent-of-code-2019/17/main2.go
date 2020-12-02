package main

import (
	"fmt"
)

func getMvt(s string) int {
	return int(s[0])
}

func getMvts(s string) []int {
	r := []int{}
	for _, c := range s {
		r = append(r, int(c))
	}
	return r
}

func processCodeAmp2(codes [][]int) ([]int, *Amp) {
	fmt.Println("---")
	fmt.Println("process")

	r := &Robot{}
	r.Init()

	aMap := getAmp(codes[0], []int{})
	aMap.Start()

	r.Direction = UP
	x, y := 0, 0
	for aMap.IsRunning {
		select {
		case o := <-aMap.Output.Chan:
			switch o {
			case 35:
				r.SetXY(x, y, CASE_WALL)
				x++
			case 46:
				r.SetXY(x, y, CASE_EMPTY)
				x++
			case 10:
				x = 0
				y++
			case 94, 118, 60, 62:
				r.X = x
				r.Y = y
				x++
			}
			if aMap.Output.WaitFor {
				aMap.Output.WaitGroup.Done()
			}
		default:
		}
	}
	close(aMap.Output.Chan)

	// minX, minY, maxX, maxY := r.Bounds()
	// for y := minY; y <= maxY; y++ {
	// 	for x := minX; x <= maxX; x++ {
	// 		if v, ok := r.Map[[2]int{x, y}]; ok && v == CASE_WALL {
	// 			c := 0
	// 			for i := -1; i < 2; i++ {
	// 				for j := -1; j < 2; j++ {
	// 					if i == 0 || j == 0 {
	// 						if v, ok := r.Map[[2]int{x + i, y + j}]; ok && v == CASE_WALL {
	// 							c++
	// 						}
	// 					}
	// 				}
	// 			}
	// 			if c == 5 {
	// 				r.SetXY(x, y, CASE_INTERSEC)
	// 			}
	// 		}
	// 	}
	// }

	fmt.Println("Map")
	r.DrawMap(true, false)

	res := ""
	t := 1
	for t != 0 {
		r.TurnRight()
		v := r.GetXY(r.NextForward())
		d := ""
		if v != CASE_WALL && v != CASE_WALL_CROSS {
			r.TurnLeft()
			r.TurnLeft()
			d += "L"
		} else {
			d += "R"
		}

		i := 0
		for r.GetXY(r.NextForward()) == CASE_WALL || r.GetXY(r.NextForward()) == CASE_WALL_CROSS || r.GetXY(r.NextForward()) == CASE_INTERSEC {
			x, y := r.NextForward()
			r.SetXY(x, y, CASE_WALL_CROSS)
			r.MoveForward()
			i++
		}

		if r.GetXY(r.NextForward()) == CASE_INTERSEC {
			r.MoveForward()
			i++
		}

		if i > 0 {
			res += fmt.Sprintf("%s,%d,", d, i)
			fmt.Println(res)
		}
		r.DrawMap(true, false)

		// time.Sleep(500 * time.Millisecond)

		t = 0
		for _, r := range r.Map {
			if r == CASE_WALL {
				t++
			}
		}
	}

	fmt.Println(res)

	codes[0][0] = 2
	a := getAmp(codes[0], []int{})
	a.Start()

	output := []int{}
	// prevX, prevY := 0, 0
	// mode := 0
	x, y = 0, 0
	mapping := true
	returnLine := 0
	a.SetInputs(getMvts("A,B,A,C,A,B,C,B,C,A\nL,12,R,4,R,4,L,6\nL,12,R,4,R,4,R,12\nL,10,L,6,R,4\n\n\n\n\n"), false)
	for a.IsRunning {
		select {
		case o := <-a.Output.Chan:
			if returnLine > 1 {
				mapping = false
			}
			if mapping {
				switch o {
				case 35:
					returnLine = 0
					r.SetXY(x, y, CASE_WALL)
					x++
				case 46:
					returnLine = 0
					r.SetXY(x, y, CASE_EMPTY)
					x++
				case 10:
					x = 0
					y++
					returnLine++
				case 94, 118, 60, 62:
					r.X = x
					r.Y = y
					x++
					fmt.Println(x, y)
				default:
					fmt.Println(string(o))
					fmt.Println(o)
					return output, a
				}
			} else {
				// r.DrawMap(false, false)
				// fmt.Println(o)
				// fmt.Println(string(o))
				if o < 256 {
					fmt.Printf(string(o))
				} else {
					fmt.Printf("Output: %d\n", o)
				}
			}
			// fmt.Println(mapping)
			// fmt.Println(x, y)
			// r.DrawMap(true, false)

			// // output = append(output, o)
			// // fmt.Println(o)
			// switch o {
			// case 0:
			// 	r.SetXY(r.X, r.Y, CASE_WALL)
			// 	r.MoveBackward()
			// 	r.TurnLeft()
			// case 1:
			// 	r.SetXY(r.X, r.Y, CASE_EMPTY)
			// 	r.RecordLength()
			// case 2:
			// 	r.SetXY(r.X, r.Y, CASE_OXYGEN)
			// 	r.RecordLength()
			// 	fmt.Println("Done")
			// 	fmt.Println(r.MapLength[[2]int{prevX, prevY}])
			// 	fmt.Println(r.MapLength[[2]int{r.X, r.Y}])
			// 	return output, a
			// }
			// r.TurnRight()
			// v := r.GetXY(r.NextForward())
			// if v == CASE_WALL {
			// 	r.TurnLeft()
			// }
			// v = r.GetXY(r.NextForward())
			// if v == CASE_WALL {
			// 	r.TurnLeft()
			// }
			// r.MoveForward()
			// fmt.Print("\033[H\033[2J")
			// fmt.Printf("Movement: %d\n", r.Direction)
			// b1, b2, b3, b4 := r.Bounds()
			// fmt.Printf("Bounds: %d/%d %d/%d\n", b1, b2, b3, b4)
			// fmt.Printf("X: %d, Y: %d\n", r.X, r.Y)
			// a.SetInput(r.Direction, false)

			// fmt.Println()
			// r.DrawMap(true, true)

			// prevX, prevY = r.X, r.Y

			if a.Output.WaitFor {
				a.Output.WaitGroup.Done()
			}
			// time.Sleep(1000 * time.Millisecond)
		default:
		}
	}
	// close(a.Input.Chan)
	close(a.Output.Chan)

	return output, a
}

func main2() {
	// processCodeAmp([][]int{parseLine("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")})
	// processCodeAmp([][]int{parseLine("1102,34915192,34915192,7,4,7,99,0")})
	// processCodeAmp([][]int{parseLine("104,1125899906842624,99")})
	processCodeAmp2(parseFile("data.txt"))
	// processCodeChainAmp([][]int{parseLine("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")})
	// processCodeChainAmp([][]int{parseLine("1102,34915192,34915192,7,4,7,99,0")})
	// processCodeChainAmp([][]int{parseLine("104,1125899906842624,99")})
	// processCodeChainAmp(parseFile("data.txt"))
}
