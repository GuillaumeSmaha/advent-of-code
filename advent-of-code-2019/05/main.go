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
		fmt.Printf("\t amp %s is stopping\n", a.Name)
		a.IsRunning = false
	default:
		panic(fmt.Sprintf("unknow op: %d", op))	
	}
}


func getAmp(code []int) *Amp {
	cInput := make(chan int)
	cOutput := make(chan int)
	return &Amp{
		Name:   "A",
		Phase:  1,
		Code:   append(code[:0:0], code...),
		Pos:    0,
		Input:  cInput,
		Output: cOutput,
	}
}

func processCode(code []int) {
	fmt.Println("---")
	fmt.Println("process")
	a := getAmp(code)
	a.Start()
	for a.IsRunning {
		select {
		case o := <-a.Output:
			fmt.Println(o)
		default:
		}
	}
	close(a.Input)
	close(a.Output)
}

func main() {
	processCode(parseData("3,9,8,9,10,9,4,9,99,-1,8"))
	processCode(parseData("3,3,1105,-1,9,1101,0,0,12,4,12,99,1"))
	processCode(parseFile("data.txt"))
	// processCode(parseData("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"))
	// processCode(parseData("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"))
}
