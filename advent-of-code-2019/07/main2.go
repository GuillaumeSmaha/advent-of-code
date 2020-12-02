package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func parseLine(data string) []int {
	ds := strings.Split(data, ",")
	r := []int{}
	for _, d := range ds {
		i, _ := strconv.Atoi(d)
		r = append(r, i)
	}
	return r
}

func parseFile(filename string) [][]int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	codes := [][]int{}
	for fscanner.Scan() {
		codes = append(codes, parseLine(fscanner.Text()))
	}
	return codes
}

type Amp struct {
	Name       string
	Phase      int
	InputInit  []int
	Code       []int
	Pos        int
	Input      chan int
	Output     chan int
	IsRunning  bool
	WaitGroup  *sync.WaitGroup
	stackTrace []string
}

func (a *Amp) Start() {
	a.stackTrace = []string{}
	a.IsRunning = true
	a.Input <- a.Phase
	if a.InputInit != nil {
		for _, i := range a.InputInit {
			a.Input <- i
		}
	}
	go func() {
		for a.IsRunning {
			a.ExecOp()
		}
	}()
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

func (a *Amp) Dump() {
	fmt.Println(strings.Join(a.stackTrace, "\n"))
}

func getBoolStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func (a *Amp) getTraceVal(p int, immediate bool) string {
	if immediate {
		return fmt.Sprintf("%04d   ", a.Code[p])
	}
	return fmt.Sprintf("c[%04d]", a.Code[p])
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

	i1 := getImmediate(0, immediates)
	i2 := getImmediate(1, immediates)

	switch op {
	case 1:
		v1 := a.GetVal(a.Pos+1, i1)
		v2 := a.GetVal(a.Pos+2, i2)
		a.Code[a.Code[a.Pos+3]] = v1 + v2
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: ADD   c[%04d] = %s + %s = %d + %d = %d", a.Name, a.Pos, a.Code[a.Pos+3], a.getTraceVal(a.Pos+1, i1), a.getTraceVal(a.Pos+2, i2), v1, v2, a.Code[a.Code[a.Pos+3]]))
		a.Pos += 4
	case 2:
		v1 := a.GetVal(a.Pos+1, i1)
		v2 := a.GetVal(a.Pos+2, i2)
		a.Code[a.Code[a.Pos+3]] = v1 * v2
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: MUL   c[%04d] = %s * %s = %d * %d = %d", a.Name, a.Pos, a.Code[a.Pos+3], a.getTraceVal(a.Pos+1, i1), a.getTraceVal(a.Pos+2, i2), v1, v2, a.Code[a.Code[a.Pos+3]]))
		a.Pos += 4
	case 3:
		a.Code[a.Code[a.Pos+1]] = <-a.Input
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: IN    c[%04d] = %d", a.Name, a.Pos, a.Code[a.Pos+1], a.Code[a.Code[a.Pos+1]]))
		a.Pos += 2
	case 4:
		a.Output <- a.Code[a.Code[a.Pos+1]]
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: OUT   c[%04d] : %d", a.Name, a.Pos, a.Code[a.Pos+1], a.Code[a.Code[a.Pos+1]]))
		a.Pos += 2
	case 5:
		v1 := a.GetVal(a.Pos+1, i1)
		v2 := a.GetVal(a.Pos+2, i2)
		r := v1 != 0
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: JMPNE POS = %s = %d   IF %s != 0 : %d != 0 : %v", a.Name, a.Pos, a.getTraceVal(a.Pos+2, i2), v2, a.getTraceVal(a.Pos+1, i1), v1, r))
		if r {
			a.Pos = v2
		} else {
			a.Pos += 3
		}
	case 6:
		v1 := a.GetVal(a.Pos+1, i1)
		v2 := a.GetVal(a.Pos+2, i2)
		r := v1 == 0
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: JMPEQ POS = %s = %d   IF %s == 0 : %d == 0 : %v", a.Name, a.Pos, a.getTraceVal(a.Pos+2, i2), v2, a.getTraceVal(a.Pos+1, i1), v1, r))
		if r {
			a.Pos = v2
		} else {
			a.Pos += 3
		}
	case 7:
		v1 := a.GetVal(a.Pos+1, i1)
		v2 := a.GetVal(a.Pos+2, i2)
		r := v1 < v2
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: LT    c[%04d] = %s == %s : %d == %d : %v", a.Name, a.Pos, a.Code[a.Pos+3], a.getTraceVal(a.Pos+1, i1), a.getTraceVal(a.Pos+2, i2), v1, v2, r))
		if r {
			a.Code[a.Code[a.Pos+3]] = 1
		} else {
			a.Code[a.Code[a.Pos+3]] = 0
		}
		a.Pos += 4
	case 8:
		v1 := a.GetVal(a.Pos+1, i1)
		v2 := a.GetVal(a.Pos+2, i2)
		r := v1 == v2
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: EQ    c[%04d] = %s == %s : %d == %d : %v", a.Name, a.Pos, a.Code[a.Pos+3], a.getTraceVal(a.Pos+1, i1), a.getTraceVal(a.Pos+2, i2), v1, v2, r))
		if r {
			a.Code[a.Code[a.Pos+3]] = 1
		} else {
			a.Code[a.Code[a.Pos+3]] = 0
		}
		a.Pos += 4
	case 99:
		a.stackTrace = append(a.stackTrace, fmt.Sprintf("%s: pos %04d: STOP", a.Name, a.Pos))
		a.IsRunning = false
		if a.WaitGroup != nil {
			a.WaitGroup.Done()
		}
	default:
		a.Dump()
		panic(fmt.Sprintf("unknow op: %d at pos %d", op, a.Pos))
	}
}

type ChainAmps struct {
	Amps   []*Amp
	Input  chan int
	Output chan int
	done   chan struct{}
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
	var waitGroup sync.WaitGroup
	c.done = make(chan struct{})
	for _, a := range c.Amps {
		a.WaitGroup = &waitGroup
		waitGroup.Add(1)
		a.Start()
	}

	go func() {
		waitGroup.Wait()
		c.close()
	}()
}

func (c *ChainAmps) close() {
	find := func(slice []chan int, val chan int) (int, bool) {
		for i, item := range slice {
			if item == val {
				return i, true
			}
		}
		return -1, false
	}

	closed := []chan int{}
	for _, a := range c.Amps {
		if _, in := find(closed, a.Output); !in {
			closed = append(closed, a.Output)
			close(a.Output)
		}
	}
	<-c.done
	if _, in := find(closed, c.Output); !in {
		closed = append(closed, c.Output)
		close(c.Output)
	}

	if _, in := find(closed, c.Input); !in {
		closed = append(closed, c.Input)
		close(c.Input)
	}
	for _, a := range c.Amps {
		if _, in := find(closed, a.Input); !in {
			closed = append(closed, a.Input)
			close(a.Input)
		}
	}
	close(c.done)
}

func (c *ChainAmps) Dump() {
	for _, a := range c.Amps {
		a.Dump()
	}
}

func getChainAmps(codes [][]int, phases []int) *ChainAmps {
	cInput := make(chan int, 128)
	cOutput := make(chan int, 128)

	amps := []*Amp{}
	initInput := []int{0}
	var cInputPrev chan int = cInput
	for i, c := range codes {
		cAmpOutput := make(chan int, 128)
		amp := &Amp{
			Name:      string(65 + i),
			Phase:     phases[i],
			InputInit: initInput,
			Code:      append(c[:0:0], c...),
			Input:     cInputPrev,
			Output:    cAmpOutput,
		}
		cInputPrev = cAmpOutput
		initInput = nil
		amps = append(amps, amp)
	}

	c := &ChainAmps{
		Amps:   amps,
		Input:  cInput,
		Output: cOutput,
	}

	go func() {
		for o := range amps[len(amps)-1].Output {
			cOutput <- o
			cInput <- o
		}
		c.done <- struct{}{}
	}()

	return c
}

func processCodePhase(code [][]int, phase []int) ([]int, *ChainAmps) {
	c := getChainAmps(code, phase)
	c.Start()
	output := []int{}
	for o := range c.Output {
		output = append(output, o)
	}
	return output, c
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

func processCode(codes [][]int) {
	fmt.Println("---")
	fmt.Println("process")
	var m int
	var mPhase []int
	var mChain *ChainAmps
	codes = append(codes, codes[0])
	codes = append(codes, codes[0])
	codes = append(codes, codes[0])
	codes = append(codes, codes[0])
	// cmb := getCombinaisonPhases([]int{0, 1, 2, 3, 4})
	cmb := getCombinaisonPhases([]int{5, 6, 7, 8, 9})
	for _, p := range cmb {
		o, c := processCodePhase(codes, p)
		v := o[len(o)-1]
		if m < v {
			m = v
			mPhase = p
			mChain = c
		}
	}
	mChain.Dump()
	fmt.Printf("Max: %d\n", m)
	fmt.Printf("MaxPhase: %v\n", mPhase)
}

func main() {
	// processCode([][]int{parseLine("3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5")})
	// processCode([][]int{parseLine("3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10")})
	processCode(parseFile("data.txt"))
}
