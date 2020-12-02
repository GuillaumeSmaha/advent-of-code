package main

import (
	"fmt"
	"strconv"
)

func processMain1(codes [][]string) {
	r := &Robot{}
	r.Init()
	r.Direction = RIGHT

	for _, c := range codes {
		v, _ := strconv.Atoi(c[1])

		switch c[0] {
		case "N":
			for range seq(v) {
				r.MoveUp(false)
			}
		case "S":
			for range seq(v) {
				r.MoveDown(false)
			}
		case "E":
			for range seq(v) {
				r.MoveRight(false)
			}
		case "W":
			for range seq(v) {
				r.MoveLeft(false)
			}
		case "L":
			for range seq(v / 90) {
				r.TurnLeft()
			}
		case "R":
			for range seq(v / 90) {
				r.TurnRight()
			}
		case "F":
			for range seq(v) {
				r.MoveForward()
			}
		}
	}

	fmt.Println("Ship", r.X, r.Y)
	fmt.Printf("Result: %d\n", AbsInt(r.X)+AbsInt(r.Y))
}

func main1() {
	processMain1(parseFileText2DFirstChar("list.test.txt"))
	processMain1(parseFileText2DFirstChar("list.txt"))
}
