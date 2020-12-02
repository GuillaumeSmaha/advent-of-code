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
	for p < 0 {
		p += len(a.Code)
	}
	if immediate {
		return a.Code[p]
	}
	return a.GetVal(a.Code[p], true)
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
			val[2] == '0',
			val[1] == '0',
		}
		op = op % 100
	}

	switch op {
	case 1:
		a.Code[a.Pos+3] = a.GetVal(a.Pos+1, getImmediate(0, immediates)) + a.GetVal(a.Pos+2, getImmediate(1, immediates))
		a.Pos += 4
	case 2:
		a.Code[a.Pos+3] = a.GetVal(a.Pos+1, getImmediate(0, immediates)) * a.GetVal(a.Pos+2, getImmediate(1, immediates))
		a.Pos += 4
	case 3:
		a.Code[a.Pos+1] = <-a.Input
		fmt.Printf("\t amp %s read value %d\n", a.Name, a.Code[a.Pos+1])
		a.Pos += 2
	case 4:
		fmt.Printf("\t amp %s write value %d\n", a.Name, a.Code[a.Pos+1])
		a.Output <- a.Code[a.Pos+1]
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
		a.Code[a.Pos+3] = 0
		if a.GetVal(a.Pos+1, getImmediate(0, immediates)) < a.GetVal(a.Pos+2, getImmediate(1, immediates)) {
			a.Code[a.Pos+3] = 1
		}
	case 8:
		a.Code[a.Pos+3] = 0
		if a.GetVal(a.Pos+1, getImmediate(0, immediates)) == a.GetVal(a.Pos+2, getImmediate(1, immediates)) {
			a.Code[a.Pos+3] = 1
		}
	case 99:
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

func getAmp(code []int, phases []int) *Amp {
	cInput := make(chan int)
	cOutput := make(chan int)

	ampAcOutput := make(chan int)
	ampA := &Amp{
		Name:   "A",
		Phase:  phases[0],
		Code:   code,
		Pos:    0,
		Input:  cInput,
		Output: ampAcOutput,
	}

	ampBcOutput := make(chan int)
	ampB := &Amp{
		Name:   "B",
		Phase:  phases[1],
		Code:   code,
		Pos:    0,
		Input:  ampAcOutput,
		Output: ampBcOutput,
	}

	ampCcOutput := make(chan int)
	ampC := &Amp{
		Name:   "C",
		Phase:  phases[2],
		Code:   code,
		Pos:    0,
		Input:  ampBcOutput,
		Output: ampCcOutput,
	}

	ampDcOutput := make(chan int)
	ampD := &Amp{
		Name:   "D",
		Phase:  phases[3],
		Code:   code,
		Pos:    0,
		Input:  ampCcOutput,
		Output: ampDcOutput,
	}

	ampE := &Amp{
		Name:   "E",
		Phase:  phases[4],
		Code:   code,
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

func processCode(code []int) {
	fmt.Println("process")
	c := getChainAmps(code, []int{4, 3, 2, 1, 0})
	defer c.Stop()
	c.Start()
	c.Input <- 0
	for c.IsRunning() {
		select {
		case o := <-c.Output:
			fmt.Println("o")
			fmt.Println(o)
		default:
		}
	}
}

func main() {
	// processCode(parseData("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"))
	// processCode(parseData("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"))
	// processCode(parseData("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"))
}
