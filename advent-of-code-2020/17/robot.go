package main

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

type Robot struct {
	X           int
	Y           int
	Direction   int // up = 1, down = 2, left = 3, right = 4
	Color       int // black = 0, white = 1
	InvertY     bool
	RestingChar string
}

func (r *Robot) Init() {
	if r.RestingChar == "" {
		r.RestingChar = "0"
	}
}

func (r *Robot) Clone() *Robot {
	nr := &Robot{}
	*nr = *r
	return nr
}

func (r *Robot) SetXY(x, y int) {
	r.X = x
	r.Y = y
}

func (r *Robot) GetXY() (int, int) {
	return r.X, r.Y
}

func (r *Robot) GetXYSlice() [2]int {
	return [2]int{r.X, r.Y}
}

func (r *Robot) NextForward() (int, int) {
	switch r.Direction {
	case UP:
		if r.InvertY {
			return r.X, r.Y - 1
		} else {
			return r.X, r.Y + 1
		}
	case DOWN:
		if r.InvertY {
			return r.X, r.Y + 1
		} else {
			return r.X, r.Y - 1
		}
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
		if r.InvertY {
			r.Y--
		} else {
			r.Y++
		}
	case DOWN:
		if r.InvertY {
			r.Y++
		} else {
			r.Y--
		}
	case LEFT:
		r.X--
	case RIGHT:
		r.X++
	}
}

func (r *Robot) MoveBackward() {
	switch r.Direction {
	case UP:
		if r.InvertY {
			r.Y++
		} else {
			r.Y--
		}
	case DOWN:
		if r.InvertY {
			r.Y--
		} else {
			r.Y++
		}
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
func (r *Robot) TurnLeftAround(x, y int) {
	tx := r.X
	if r.InvertY {
		r.X = x + (r.Y - y)
		r.Y = y + -1*(tx-x)
	} else {
		r.X = x + -1*(r.Y-y)
		r.Y = y + (tx - x)
	}
}

func (r *Robot) TurnRightAround(x, y int) {
	tx := r.X
	if r.InvertY {
		r.X = x + -1*(r.Y-y)
		r.Y = y + (tx - x)
	} else {
		r.X = x + (r.Y - y)
		r.Y = y + -1*(tx-x)
	}
}

func (r *Robot) MoveLeft(setDirection bool, c ...int) {
	cnt := 1
	if len(c) > 0 {
		cnt = c[0]
	}
	r.X += -1 * cnt
	if setDirection {
		r.Direction = LEFT
	}
}

func (r *Robot) MoveRight(setDirection bool, c ...int) {
	cnt := 1
	if len(c) > 0 {
		cnt = c[0]
	}
	r.X += cnt
	if setDirection {
		r.Direction = RIGHT
	}
}

func (r *Robot) MoveUp(setDirection bool, c ...int) {
	cnt := 1
	if len(c) > 0 {
		cnt = c[0]
	}
	if r.InvertY {
		r.Y += -1 * cnt
	} else {
		r.Y += cnt
	}
	if setDirection {
		r.Direction = UP
	}
}

func (r *Robot) MoveDown(setDirection bool, c ...int) {
	cnt := 1
	if len(c) > 0 {
		cnt = c[0]
	}
	if r.InvertY {
		r.Y += cnt
	} else {
		r.Y += -1 * cnt
	}
	if setDirection {
		r.Direction = DOWN
	}
}

func (r *Robot) Move(direction int, setDirection bool, c ...int) {
	switch direction {
	case UP:
		r.MoveUp(setDirection, c...)
	case DOWN:
		r.MoveDown(setDirection, c...)
	case LEFT:
		r.MoveLeft(setDirection, c...)
	case RIGHT:
		r.MoveRight(setDirection, c...)
	}
}

func (r *Robot) DrawRobot() string {
	switch r.Direction {
	case UP:
		return "^"
	case DOWN:
		return "V"
	case LEFT:
		return "<"
	case RIGHT:
		return ">"
	}
	return r.RestingChar
}
