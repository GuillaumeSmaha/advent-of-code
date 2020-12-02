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
		log.Fatal("usage: advent-of-code-2018 3[a|b] 'input'")
	}

	switch args[0] {
	case "3a", "3":
		fmt.Println(Part1(args[1]))
	case "3b":
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

func getSize(s string) []int {

	size := make([]int, 2)
	r := regexp.MustCompile(`#(?P<Id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)`)
	match := r.FindStringSubmatch(s)
	left, _ := strconv.Atoi(match[2])
	top, _ := strconv.Atoi(match[3])
	width, _ := strconv.Atoi(match[4])
	height, _ := strconv.Atoi(match[5])
	size[0] = left*2 + width
	size[1] = top*2 + height

	return size
}

func applyOnRug(rug [][]int, s string) {

	r := regexp.MustCompile(`#(?P<Id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)`)
	//`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`
	match := r.FindStringSubmatch(s)
	left, _ := strconv.Atoi(match[2])
	top, _ := strconv.Atoi(match[3])
	width, _ := strconv.Atoi(match[4])
	height, _ := strconv.Atoi(match[5])

	for i := left; i < (left + width); i++ {
		for j := top; j < (top + height); j++ {
			rug[i][j] += 1
		}
	}
}

func sum(s string) int {
	res := load(s)

	size := make([]int, 2)
	for _, r := range res {
		csize := getSize(r)
		size[0] = max(size[0], csize[0])
		size[1] = max(size[1], csize[1])
	}
	fmt.Printf("size: %v\n", size)

	rug := make([][]int, size[0])
	for i := 0; i < size[0]; i++ {
		rug[i] = make([]int, size[1])
	}

	for _, r := range res {
		applyOnRug(rug, r)
	}

	count := 0
	for i := 0; i < size[0]; i++ {
		for j := 0; j < size[1]; j++ {
			if rug[i][j] > 1 {
				count += 1
			}
		}
	}
	return count
}

func Part1(s string) int {
	return sum(s)
}

func applyOnRugKeys(rugKeys [][][]int, s string) {

	r := regexp.MustCompile(`#(?P<Id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)`)
	//`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`
	match := r.FindStringSubmatch(s)
	id, _ := strconv.Atoi(match[1])
	left, _ := strconv.Atoi(match[2])
	top, _ := strconv.Atoi(match[3])
	width, _ := strconv.Atoi(match[4])
	height, _ := strconv.Atoi(match[5])

	for i := left; i < (left + width); i++ {
		for j := top; j < (top + height); j++ {
			rugKeys[i][j][id] = 1
		}
	}
}

func diff(s string) int {
	res := load(s)

	size := make([]int, 2)
	for _, r := range res {
		csize := getSize(r)
		size[0] = max(size[0], csize[0])
		size[1] = max(size[1], csize[1])
	}
	fmt.Printf("size: %v\n", size)

	rug := make([][]int, size[0])
	for i := 0; i < size[0]; i++ {
		rug[i] = make([]int, size[1])
	}

	rugKeys := make([][][]int, size[0])
	for i := 0; i < size[0]; i++ {
		rugKeys[i] = make([][]int, size[1])
		for j := 0; j < size[1]; j++ {
			rugKeys[i][j] = make([]int, len(res)+1)
		}
	}

	for _, r := range res {
		applyOnRug(rug, r)
		applyOnRugKeys(rugKeys, r)
	}

	checkOverlay := func(s string) bool {
		r := regexp.MustCompile(`#(?P<Id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)`)
		//`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`
		match := r.FindStringSubmatch(s)
		left, _ := strconv.Atoi(match[2])
		top, _ := strconv.Atoi(match[3])
		width, _ := strconv.Atoi(match[4])
		height, _ := strconv.Atoi(match[5])

		for i := left; i < (left + width); i++ {
			for j := top; j < (top + height); j++ {
				if rug[i][j] > 1 {
					return false
				}
			}
		}
		return true
	}

	cached := make([]int, len(res)+1)

	for i := 0; i < size[0]; i++ {
		for j := 0; j < size[1]; j++ {
			if rug[i][j] == 1 {
				for k := 0; k < len(res); k++ {
					if rugKeys[i][j][k] == 1 {
						if cached[k] == 0 && checkOverlay(res[k]) {
							return k + 1
						}
						cached[k] = 1
						break
					}
				}
			}
		}
	}
	fmt.Println("return 0")
	return 0
}

func Part2(s string) int {
	return diff(s)
}
