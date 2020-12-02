package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 10[a|b] 'input'")
	}

	switch args[0] {
	case "10a", "10":
		fmt.Println(Part1(args[1]))
	case "10b":
		fmt.Println(Part2(args[1]))
	}
}

func load(s string) []string {
	res := make([]string, 0)
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = strings.TrimSpace(string(f))

		for _, line := range strings.Split(s, "\n") {
			if len(line) == 0 {
				continue
			}
			res = append(res, line)
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func deleteFromSlice(list []string, element string) []string {
	for i, v := range list {
		if v == element {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}

	return list
}

type Point struct {
	X  int
	Y  int
	VX int
	VY int
}

var Parser *regexp.Regexp = regexp.MustCompile(`position=<\s*(?P<x>[0-9-]+),\s*(?P<y>[0-9-]+)> velocity=<\s*(?P<vx>[0-9-]+),\s*(?P<vy>[0-9-]+)>`)

func parseString(s string) *Point {
	match := Parser.FindStringSubmatch(s)
	x, _ := strconv.Atoi(match[1])
	y, _ := strconv.Atoi(match[2])
	vx, _ := strconv.Atoi(match[3])
	vy, _ := strconv.Atoi(match[4])
	return &Point{
		X:  x,
		Y:  y,
		VX: vx,
		VY: vy,
	}
}

func loadData(s []string) []*Point {
	data := make([]*Point, 0)
	for _, line := range s {
		point := parseString(line)
		data = append(data, point)
	}
	return data
}

func getSize(data []*Point) *Point {
	minMax := &Point{
		X:  1000000000,
		Y:  1000000000,
		VX: -1000000000,
		VY: -1000000000,
	}
	for _, p := range data {
		if p.X < minMax.X {
			minMax.X = p.X
		} else if p.X > minMax.VX {
			minMax.VX = p.X
		}
		if p.Y < minMax.Y {
			minMax.Y = p.Y
		} else if p.Y > minMax.VY {
			minMax.VY = p.Y
		}
	}
	return minMax
}

func printData(data []*Point) {
	size := getSize(data)
	for y := size.Y; y <= size.VY; y++ {
		for x := size.X; x <= size.VX; x++ {
			fnd := false
			for _, p := range data {
				if p.X == x && p.Y == y {
					fnd = true
					break
				}
			}

			if fnd {
				fmt.Printf("X")
			} else {
				fmt.Printf(".")
			}

		}
		fmt.Println()
	}
}

func moveStep(data []*Point) {
	for _, p := range data {
		p.X += p.VX
		p.Y += p.VY
	}
}

func manathan(x1 int, y1 int, x2 int, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}

func detectAlignmentX(data []*Point, minAligned int) bool {
	for _, p1 := range data {
		cnt := 0
		for _, p2 := range data {
			if p1.X == p2.X && Abs(p1.Y-p2.Y) < minAligned {
				cnt++
			}
		}
		if cnt > minAligned {
			return true
		}
	}
	return false
}

func Part1(s string) int {
	str := load(s)
	data := loadData(str)

	for !detectAlignmentX(data, 20) {
		moveStep(data)
	}
	printData(data)

	return 0
}

func Part2(s string) int {
	str := load(s)
	data := loadData(str)

	nbMove := 0
	for !detectAlignmentX(data, 20) {
		moveStep(data)
		nbMove++
	}
	printData(data)
	fmt.Printf("Nb move: %v\n", nbMove)

	return 0
}
