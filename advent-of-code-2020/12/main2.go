package main

import (
	"fmt"
	"strconv"
)

func processMain2(codes [][]string) {
	ship := &Robot{
		RestingChar: "S",
	}
	ship.Init()

	waypoint := &Robot{
		RestingChar: "W",
	}
	waypoint.Init()
	waypoint.X = 10
	waypoint.Y = 1

	m := &Map{}
	m.Init()
	m.AddRobot(ship)
	m.AddRobot(waypoint)

	fmt.Println("---")
	fmt.Println("Ship", ship.X, ship.Y)
	fmt.Println("Waypoint", waypoint.X, waypoint.Y)

	for _, c := range codes {
		v, _ := strconv.Atoi(c[1])

		switch c[0] {
		case "N":
			waypoint.MoveUp(false, v)
		case "S":
			waypoint.MoveDown(false, v)
		case "E":
			waypoint.MoveRight(false, v)
		case "W":
			waypoint.MoveLeft(false, v)
		case "L":
			for range seq(v / 90) {
				waypoint.TurnLeftAround(0, 0)
			}
		case "R":
			for range seq(v / 90) {
				waypoint.TurnRightAround(0, 0)
			}
		case "F":
			ship.X += v * waypoint.X
			ship.Y += v * waypoint.Y
		}

		m.SetXY(ship.X, ship.Y, CASE_UNKNOW)
		m.SetXY(waypoint.X, waypoint.Y, CASE_UNKNOW)
	}
	fmt.Printf("Bounds: %v\n", m.BoundsList())
	fmt.Println("Ship", ship.X, ship.Y)
	fmt.Println("Waypoint", waypoint.X, waypoint.Y)
	fmt.Printf("Result: %d\n", AbsInt(ship.X)+AbsInt(ship.Y))
}

func main2() {
	processMain2(parseFileText2DFirstChar("list.test.txt"))
	processMain2(parseFileText2DFirstChar("list.txt"))
}
