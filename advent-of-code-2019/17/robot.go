package main

import "fmt"

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
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
	CASE_WALL:       "█",
	CASE_EMPTY:      " ",
	CASE_START:      "X",
	CASE_INTERSEC:   "O",
	CASE_WALL_CROSS: "░",
}

type Robot struct {
	Map       map[[2]int]int
	MapLength map[[2]int]int
	X         int
	Y         int
	Direction int // up = 1, down = 2, left = 3, right = 4
	Color     int // black = 0, white = 1
}

func (r *Robot) Init() {
	r.Map = map[[2]int]int{}
	r.MapLength = map[[2]int]int{}
}

func (r *Robot) RecordLength() {
	min := 999999999
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 || j == 0 {
				if v, ok := r.MapLength[[2]int{r.X + i, r.Y + j}]; ok {
					if min > v {
						min = v
					}
				}
			}
		}
	}
	r.MapLength[[2]int{r.X, r.Y}] = min + 1
}

func (r *Robot) NextForward() (int, int) {
	switch r.Direction {
	case UP:
		return r.X, r.Y - 1
	case DOWN:
		return r.X, r.Y + 1
	case LEFT:
		return r.X - 1, r.Y
	case RIGHT:
		return r.X + 1, r.Y
	}
	return r.X, r.Y
}

func (r *Robot) MoveForward() {
	switch r.Direction {
	case UP:
		r.Y--
	case DOWN:
		r.Y++
	case LEFT:
		r.X--
	case RIGHT:
		r.X++
	}
}

func (r *Robot) MoveBackward() {
	switch r.Direction {
	case UP:
		r.Y++
	case DOWN:
		r.Y--
	case LEFT:
		r.X++
	case RIGHT:
		r.X--
	}
}

func (r *Robot) TurnRight() {
	switch r.Direction {
	case UP:
		r.Direction = RIGHT
	case RIGHT:
		r.Direction = DOWN
	case DOWN:
		r.Direction = LEFT
	case LEFT:
		r.Direction = UP
	}
}

func (r *Robot) TurnLeft() {
	switch r.Direction {
	case UP:
		r.Direction = LEFT
	case LEFT:
		r.Direction = DOWN
	case DOWN:
		r.Direction = RIGHT
	case RIGHT:
		r.Direction = UP
	}
}

func (r *Robot) MoveLeft(direction bool) {
	r.Y--
	if direction {
		r.Direction = LEFT
	}
}

func (r *Robot) MoveRight(direction bool) {
	r.Y++
	if direction {
		r.Direction = RIGHT
	}
}

func (r *Robot) MoveUp(direction bool) {
	r.Y--
	if direction {
		r.Direction = UP
	}
}

func (r *Robot) MoveDown(direction bool) {
	r.Y++
	if direction {
		r.Direction = DOWN
	}
}

func (r *Robot) DrawRobot() string {
	switch r.Direction {
	case UP:
		return "^"
	case LEFT:
		return "<"
	case DOWN:
		return "V"
	case RIGHT:
		return ">"
	}
	return "Θ"
}

func (r *Robot) GetChar(x, y int) string {
	return CHARS_CASE[r.GetXY(x, y)]
}

func (r *Robot) GetXY(x, y int) int {
	if v, ok := r.Map[[2]int{x, y}]; ok {
		return v
	}
	return CASE_UNKNOW
}

func (r *Robot) Bounds() (minX, minY, maxX, maxY int) {
	for xy, _ := range r.Map {
		minX = xy[0]
		minY = xy[1]
		maxX = xy[0]
		maxY = xy[1]
		break
	}
	for xy, _ := range r.Map {
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

func (r *Robot) DrawMap(drawRobot bool, drawMapLength bool) {
	minX, minY, maxX, maxY := r.Bounds()
	// for y := maxY; y >= minY; y-- {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if x == 0 && y == 0 {
				fmt.Printf(CHARS_CASE[CASE_START])
			} else if drawRobot && r.X == x && r.Y == y {
				fmt.Printf(r.DrawRobot())
			} else {
				c := r.GetChar(x, y)
				fmt.Printf(c)
			}
		}
		fmt.Printf("\n")
	}

	if drawMapLength {
		fmt.Printf("\n")
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if drawRobot && r.X == x && r.Y == y {
					fmt.Printf(CHARS_CASE[CASE_EMPTY])
					fmt.Printf(r.DrawRobot())
					fmt.Printf(CHARS_CASE[CASE_EMPTY])
				} else {
					v := r.GetXY(x, y)
					c := r.GetChar(x, y)
					if v == CASE_WALL {
						fmt.Printf(c)
						fmt.Printf(c)
						fmt.Printf(c)
					} else if v, ok := r.MapLength[[2]int{x, y}]; ok {
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

func (r *Robot) SetXY(x, y int, v int) {
	r.Map[[2]int{x, y}] = v
}
