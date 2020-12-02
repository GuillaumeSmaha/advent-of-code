package main

import (
	"fmt"
)

func processCodeAmp1(codes [][]int) ([]int, *Amp) {
	fmt.Println("---")
	fmt.Println("process")

	r := &Robot{}
	r.Init()

	a := getAmp(codes[0], []int{})
	a.Start()

	output := []int{}

	r.MoveUp(true)
	// prevX, prevY := 0, 0
	x, y := 0, 0
	for a.IsRunning {
		select {
		case o := <-a.Output.Chan:
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
			}
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

	r.DrawMap(false, false)

	t := 0
	minX, minY, maxX, maxY := r.Bounds()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if v, ok := r.Map[[2]int{x, y}]; ok && v == CASE_WALL {
				c := 0
				for i := -1; i < 2; i++ {
					for j := -1; j < 2; j++ {
						if i == 0 || j == 0 {
							if v, ok := r.Map[[2]int{x + i, y + j}]; ok && v == CASE_WALL {
								c++
							}
						}
					}
				}
				if c == 5 {
					fmt.Println(x, y, c)
					t += x * y
				}
			}
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Printf("Total: %d\n", t)

	return output, a
}

func main1() {
	// processCodeAmp([][]int{parseLine("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")})
	// processCodeAmp([][]int{parseLine("1102,34915192,34915192,7,4,7,99,0")})
	// processCodeAmp([][]int{parseLine("104,1125899906842624,99")})
	processCodeAmp1(parseFile("data.txt"))
	// processCodeChainAmp([][]int{parseLine("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")})
	// processCodeChainAmp([][]int{parseLine("1102,34915192,34915192,7,4,7,99,0")})
	// processCodeChainAmp([][]int{parseLine("104,1125899906842624,99")})
	// processCodeChainAmp(parseFile("data.txt"))
}
