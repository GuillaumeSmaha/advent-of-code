package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseData(data string) []int {
	ds := strings.Split(data, ",")
	r := []int{}
	for _, d := range ds {
		i, _ := strconv.Atoi(d)
		r = append(r, i)
	}
	return r
}

func parseFile(filename string) []int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	fscanner.Scan()
	return parseData(fscanner.Text())
}

type Amp struct {
	Name      string
	Phase     int
	Code      []int
	Pos       int
	Input     chan int
	Output    chan int
	IsRunning bool
}

func (a *Amp) Start() {
	a.IsRunning = true
	go func() {
		for a.IsRunning {
			a.ExecOp()
		}
	}()
	a.Input <- a.Phase
}

func (a *Amp) GetVal(p int, immediate bool) int {
	if immediate {
		return a.Code[p]
	}
	return a.Code[a.Code[p]]
}

func getImmediate(i int, immediates []bool) bool {
	if immediates == nil {
		return false
	}
	return immediates[i]
}

func (a *Amp) ExecOp() {
	op := a.Code[a.Pos]

	var immediates []bool
	if op > 99 {
		val := fmt.Sprintf("%05d", op)
		immediates = []bool{
			val[2] == '1',
			val[1] == '1',
		}
		op = op % 100
	}

	switch op {
	case 1:
		a.Code[a.Code[a.Pos+3]] = a.GetVal(a.Pos+1, getImmediate(0, immediates)) + a.GetVal(a.Pos+2, getImmediate(1, immediates))
		a.Pos += 4
	case 2:
		a.Code[a.Code[a.Pos+3]] = a.GetVal(a.Pos+1, getImmediate(0, immediates)) * a.GetVal(a.Pos+2, getImmediate(1, immediates))
		a.Pos += 4
	case 3:
		a.Code[a.Code[a.Pos+1]] = <-a.Input
		// fmt.Printf("\t amp %s read value %d\n", a.Name, a.Code[a.Pos+1])
		a.Pos += 2
	case 4:
		// fmt.Printf("\t amp %s write value %d\n", a.Name, a.Code[a.Pos+1])
		a.Output <- a.Code[a.Code[a.Pos+1]]
		a.Pos += 2
	case 5:
		if a.GetVal(a.Pos+1, getImmediate(0, immediates)) != 0 {
			a.Pos = a.GetVal(a.Pos+2, getImmediate(1, immediates))
		} else {
			a.Pos += 3
		}
	case 6:
		if a.GetVal(a.Pos+1, getImmediate(0, immediates)) == 0 {
			a.Pos = a.GetVal(a.Pos+2, getImmediate(1, immediates))
		} else {
			a.Pos += 3
		}
	case 7:
		if a.GetVal(a.Pos+1, getImmediate(0, immediates)) < a.GetVal(a.Pos+2, getImmediate(1, immediates)) {
			a.Code[a.Code[a.Pos+3]] = 1
		} else {
			a.Code[a.Code[a.Pos+3]] = 0
		}
		a.Pos += 4
	case 8:
		if a.GetVal(a.Pos+1, getImmediate(0, immediates)) == a.GetVal(a.Pos+2, getImmediate(1, immediates)) {
			a.Code[a.Code[a.Pos+3]] = 1
		} else {
			a.Code[a.Code[a.Pos+3]] = 0
		}
		a.Pos += 4
	case 99:
		// fmt.Printf("\t amp %s is stopping\n", a.Name)
		a.IsRunning = false
	default:
		panic(fmt.Sprintf("unknow op: %d", op))
	}
}

type ChainAmps struct {
	Amps   []*Amp
	Input  chan int
	Output chan int
}

func (c *ChainAmps) IsRunning() bool {
	for _, a := range c.Amps {
		if a.IsRunning {
			return true
		}
	}
	return false
}

func (c *ChainAmps) Start() {
	for _, a := range c.Amps {
		a.Start()
	}
}

func (c *ChainAmps) Stop() {
	for _, a := range c.Amps {
		a.IsRunning = false
		close(a.Input)
	}
	close(c.Output)
}

func getChainAmps(code []int, phases []int) *ChainAmps {
	cInput := make(chan int)
	cOutput := make(chan int)

	ampAcOutput := make(chan int)
	ampA := &Amp{
		Name:   "A",
		Phase:  phases[0],
		Code:   append(code[:0:0], code...),
		Pos:    0,
		Input:  cInput,
		Output: ampAcOutput,
	}

	ampBcOutput := make(chan int)
	ampB := &Amp{
		Name:   "B",
		Phase:  phases[1],
		Code:   append(code[:0:0], code...),
		Pos:    0,
		Input:  ampAcOutput,
		Output: ampBcOutput,
	}

	ampCcOutput := make(chan int)
	ampC := &Amp{
		Name:   "C",
		Phase:  phases[2],
		Code:   append(code[:0:0], code...),
		Pos:    0,
		Input:  ampBcOutput,
		Output: ampCcOutput,
	}

	ampDcOutput := make(chan int)
	ampD := &Amp{
		Name:   "D",
		Phase:  phases[3],
		Code:   append(code[:0:0], code...),
		Pos:    0,
		Input:  ampCcOutput,
		Output: ampDcOutput,
	}

	ampE := &Amp{
		Name:   "E",
		Phase:  phases[4],
		Code:   append(code[:0:0], code...),
		Pos:    0,
		Input:  ampDcOutput,
		Output: cOutput,
	}

	return &ChainAmps{
		Amps: []*Amp{
			ampA,
			ampB,
			ampC,
			ampD,
			ampE,
		},
		Input:  cInput,
		Output: cOutput,
	}
}

func processCodePhase(code []int, phase []int) []int {
	c := getChainAmps(code, phase)
	defer c.Stop()
	c.Start()
	c.Input <- 0
	output := []int{}
	for c.IsRunning() {
		select {
		case o := <-c.Output:
			output = append(output, o)
		default:
		}
	}
	return output
}

func getCombinaisonPhases(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func processCode(code []int) {
	fmt.Println("---")
	fmt.Println("process")
	var m int
	var mPhase []int
	for _, p := range getCombinaisonPhases([]int{0, 1, 2, 3, 4}) {
		o := processCodePhase(code, p)
		if m < o[0] {
			m = o[0]
			mPhase = p
		}
	}
	fmt.Printf("Max: %d\n", m)
	fmt.Printf("MaxPhase: %v\n", mPhase)
}

func main() {
	processCode(parseData("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"))
	processCode(parseData("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"))
	processCode(parseFile("data.txt"))
}
