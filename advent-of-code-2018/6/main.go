package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 6[a|b] 'input'")
	}

	switch args[0] {
	case "6a", "6":
		fmt.Println(Part1(args[1]))
	case "6b":
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

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func manathan(x1 int, y1 int, x2 int, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}

func manathanCoord(c1 *Coord, c2 *Coord) int {
	return Abs(c1.X-c2.X) + Abs(c1.Y-c2.Y)
}

func manathanIntCoord(x1 int, y1 int, c2 *Coord) int {
	return Abs(x1-c2.X) + Abs(y1-c2.Y)
}

func compress(s string) (string, int) {
	res := ""
	prevC := '\n'
	cnt := 0

	for _, c := range s {
		if prevC != '\n' {
			if Abs(int(c)-int(prevC)) == 32 {
				res = res[:len(res)-1]
				prevC = '\n'
				cnt++
				continue
			} else {
				res = fmt.Sprintf("%v%c", res, c)
			}
		} else {
			res = fmt.Sprintf("%v%c", res, c)
		}
		prevC = c
	}

	return res, cnt
}

type Coord struct {
	X int
	Y int
}

func getMark(idx int) string {
	return string(int('A') + idx)
}

func getSubMark(idx int) string {
	return string(int('a') + idx)
}

func loadCoords(s []string) []*Coord {
	coords := make([]*Coord, 0)
	for _, line := range s {
		values := strings.Split(line, ",")

		x, _ := strconv.Atoi(strings.Trim(values[0], " "))
		y, _ := strconv.Atoi(strings.Trim(values[1], " "))
		coords = append(coords, &Coord{
			X: x,
			Y: y,
		})
	}

	return coords
}

func getSize(coords []*Coord) *Coord {
	res := &Coord{
		X: 0,
		Y: 0,
	}
	for _, c := range coords {
		if c.X > res.X {
			res.X = c.X
		}
		if c.Y > res.Y {
			res.Y = c.Y
		}
	}

	res.X += 2
	res.Y++
	return res
}

func createGrid(coords []*Coord) [][]int {
	size := getSize(coords)

	grid := make([][]int, size.Y)
	for i := 0; i < size.Y; i++ {
		grid[i] = make([]int, size.X)
	}

	for ci, c := range coords {
		grid[c.Y][c.X] = ci + 1
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				idx := []int{}
				dist := 100000
				for ci, c := range coords {
					tmpDist := manathanIntCoord(j, i, c)
					if tmpDist < dist {
						dist = tmpDist
						idx = []int{ci}
					} else if dist == tmpDist {
						idx = append(idx, ci)
					}
				}
				if len(idx) == 1 {
					grid[i][j] = -idx[0] - 1
				}
			}
		}
	}

	return grid
}

func getInfinites(grid [][]int) []int {
	coverage := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		val := grid[i][0]
		if val != 0 {
			if val > 0 {
				coverage[val-1] = 1
			} else {
				coverage[-val-1] = 1
			}
		}
		val = grid[i][len(grid[i])-1]
		if val != 0 {
			if val > 0 {
				coverage[val-1] = 1
			} else {
				coverage[-val-1] = 1
			}
		}
	}
	for j := 0; j < len(grid[0]); j++ {
		val := grid[0][j]
		if val != 0 {
			if val > 0 {
				coverage[val-1] = 1
			} else {
				coverage[-val-1] = 1
			}
		}
		val = grid[len(grid)-1][j]
		if val != 0 {
			if val > 0 {
				coverage[val-1] = 1
			} else {
				coverage[-val-1] = 1
			}
		}
	}

	keys := make([]int, 0, len(coverage))
	for k := range coverage {
		keys = append(keys, k)
	}

	return keys
}

func getFinites(grid [][]int, coords []*Coord) []int {
	coverage := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		val := grid[i][0]
		if val != 0 {
			if val > 0 {
				coverage[val-1] = 1
			} else {
				coverage[-val-1] = 1
			}
		}
		val = grid[i][len(grid[i])-1]
		if val != 0 {
			if val > 0 {
				coverage[val-1] = 1
			} else {
				coverage[-val-1] = 1
			}
		}
	}
	for j := 0; j < len(grid[0]); j++ {
		val := grid[0][j]
		if val != 0 {
			if val > 0 {
				coverage[val-1] = 1
			} else {
				coverage[-val-1] = 1
			}
		}
		val = grid[len(grid)-1][j]
		if val != 0 {
			if val > 0 {
				coverage[val-1] = 1
			} else {
				coverage[-val-1] = 1
			}
		}
	}

	keys := make([]int, 0)
	for i := 0; i < len(coords); i++ {
		if _, ok := coverage[i]; !ok {
			keys = append(keys, i)
		}
	}

	return keys
}

func maxCoverage(grid [][]int) (int, int) {
	coverage := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				if grid[i][j] > 0 {
					coverage[grid[i][j]-1]++
				} else {
					coverage[-grid[i][j]-1]++
				}
			}
		}
	}

	idx := 0
	cnt := 0
	fmt.Println(coverage)
	for i, c := range coverage {
		fmt.Println(i, c)
		if c > cnt {
			idx = i
			cnt = c
		}
	}

	return idx, cnt
}

func maxCoverageForFinites(grid [][]int, coords []*Coord) (int, int) {
	finites := getFinites(grid, coords)
	coverage := make(map[int]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				if grid[i][j] > 0 {
					coverage[grid[i][j]-1]++
				} else {
					coverage[-grid[i][j]-1]++
				}
			}
		}
	}

	idx := 0
	cnt := 0
	for i, c := range coverage {
		if intInSlice(i, finites) && c > cnt {
			idx = i
			cnt = c
		}
	}

	return idx, cnt
}

func printGrid(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				if grid[i][j] > 0 {
					fmt.Printf("%v", getMark(grid[i][j]-1))
				} else {
					fmt.Printf("%v", getSubMark(-grid[i][j]-1))
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func Part1(s string) int {
	str := load(s)
	coords := loadCoords(str)
	grid := createGrid(coords)
	// printGrid(grid)
	fmt.Printf("Infinites: %v\n", getInfinites(grid))
	fmt.Printf("Finites: %v\n", getFinites(grid, coords))
	idx, cnt := maxCoverageForFinites(grid, coords)
	fmt.Printf("Most present: %v\n", idx)
	fmt.Printf("Most present mark: %v\n", getMark(idx))
	fmt.Printf("Count most present: %v\n", cnt)
	return cnt
}

func createGridRegion(coords []*Coord, dist int) [][]int {
	size := getSize(coords)

	grid := make([][]int, size.Y)
	for i := 0; i < size.Y; i++ {
		grid[i] = make([]int, size.X)
	}

	for ci, c := range coords {
		grid[c.Y][c.X] = -ci - 2
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			sumDist := 0
			for _, c := range coords {
				sumDist += manathanIntCoord(j, i, c)
				if sumDist >= dist {
					break
				}
			}
			if sumDist < dist {
				if grid[i][j] == 0 {
					grid[i][j] = 1
				} else {
					grid[i][j] = -grid[i][j]
				}
			}
		}
	}

	return grid
}

func printGridRegion(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				fmt.Printf("%v", getSubMark(-grid[i][j]-2))
			} else if grid[i][j] == 1 {
				fmt.Printf("#")
			} else if grid[i][j] > 0 {
				fmt.Printf("%v", getMark(grid[i][j]-2))
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func sizeRegion(grid [][]int) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				cnt++
			}
		}
	}

	return cnt
}

func Part2(s string) int {
	str := load(s)
	coords := loadCoords(str)
	sizeGrid := getSize(coords)
	fmt.Printf("size: %v\n", sizeGrid)
	fmt.Printf("total element: %v\n", sizeGrid.X*sizeGrid.Y)
	// grid := createGridRegion(coords, 32)
	grid := createGridRegion(coords, 10000)
	// printGridRegion(grid)
	size := sizeRegion(grid)
	return size
}
