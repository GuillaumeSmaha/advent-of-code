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
		log.Fatal("usage: advent-of-code-2018 11[a|b] 'input'")
	}

	switch args[0] {
	case "11a", "11":
		fmt.Println(Part1(args[1]))
	case "11b":
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

func createGrid(size int, input int) [][]int {
	size++
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]int, size)
	}

	getFuel := func(x int, y int) int {
		rackID := x + 10
		tmp := (rackID*y + input) * rackID
		return (tmp/100)%10 - 5
	}

	for y := 1; y < len(grid); y++ {
		for x := 1; x < len(grid[y]); x++ {
			grid[y][x] = getFuel(x, y)
		}
	}

	return grid
}

func findMaxSum(grid [][]int, maskSize int) (int, int, int) {

	maxMaskValue := maskSize * maskSize * 4

	getMaskValue := func(x int, y int) int {
		sum := 0
		// for j := y - (maskSize-1)/2; j <= y+(maskSize-1)/2; j++ {
		// 	for i := x - (maskSize-1)/2; i <= x+(maskSize-1)/2; i++ {
		// 		sum += grid[j][i]
		// 	}
		// }
		for j := y; j < y+maskSize; j++ {
			for i := x; i < x+maskSize; i++ {
				sum += grid[j][i]
			}
		}
		return sum
	}

	maxX := 1
	maxY := 1
	maxValue := -maxMaskValue
	max := len(grid) - maskSize + 1
	for y := 1; y < max; y++ {
		for x := 1; x < max; x++ {
			val := getMaskValue(x, y)
			if val > maxValue {
				maxX = x
				maxY = y
				maxValue = val
				if val == maxMaskValue {
					return maxX, maxY, maxValue
				}
			}
		}
	}

	return maxX, maxY, maxValue
}

func Part1(s string) string {
	size := 300
	input := 18
	input = 9110
	grid := createGrid(size, input)

	x, y, val := findMaxSum(grid, 3)
	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
	fmt.Printf("val = %v\n", val)

	return fmt.Sprintf("%v,%v", x, y)
}

func Part2(s string) string {
	size := 300
	input := 18
	input = 42
	input = 9110
	grid := createGrid(size, input)

	maxX := 1
	maxY := 1
	maxMaskSize := 1
	maxValue := 0
	for i := 1; i <= size; i++ {
		x, y, val := findMaxSum(grid, i)
		fmt.Printf("%03d: %v,%v,%v\n", i, x, y, val)
		if val > maxValue {
			maxMaskSize = i
			maxValue = val
			maxX = x
			maxY = y
		}
	}

	return fmt.Sprintf("%v,%v,%v", maxX, maxY, maxMaskSize)
}
