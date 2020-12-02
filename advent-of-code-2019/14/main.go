package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
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

func getBoolStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

type AmpInput struct {
	Init       []int
	MemorySize int
	Memory     []int
	Pos        int
	UseChannel bool
	Chan       chan int
}

type AmpOutput struct {
	WaitGroup sync.WaitGroup
	WaitFor   bool
	Chan      chan int
}

type Amp struct {
	Name             string
	MemorySize       int
	Input            AmpInput
	Output           AmpOutput
	Code             []int
	Pos              int
	RelativeBase     int
	IsRunning        bool
	WaitGroup        *sync.WaitGroup
	stackTrace       []string
	stackTraceWriter *tabwriter.Writer
	Debug            bool
}

func (a *Amp) Start() {
	if a.Input.MemorySize < 1 {
		a.Input.MemorySize = 1
	}
	a.Input.Memory = make([]int, a.Input.MemorySize)
	for i := 0; i < a.MemorySize; i++ {
		a.Code =
			append(a.Code, 0)
	}
	if a.Input.UseChannel {
		if a.Input.Init != nil {
			go func() {
				for _, i := range a.Input.Init {
					a.Input.Chan <- i
				}
			}()
		}
	} else {
		if a.Input.Init != nil {
			for i, v := range a.Input.Init {
				a.Code[a.Input.Pos+i] = v
				a.Input.Memory[i] = v
			}
		}
	}

	a.stackTrace = []string{}
	a.IsRunning = true
	go func() {
		for a.IsRunning {
			a.ExecOp()
		}
	}()
}

func (a *Amp) GetPos(p int, mode int) int {
	switch mode {
	case 1:
		return p
	case 2:
		return a.RelativeBase + a.Code[p]
	}
	return a.Code[p]
}
func (a *Amp) GetVal(p int, mode int) int {
	return a.Code[a.GetPos(p, mode)]
}

func (a *Amp) getTraceVal(p int, mode int) string {
	switch mode {
	case 1:
		return fmt.Sprintf("%d", a.Code[p])
	case 2:
		return fmt.Sprintf("c[%04d]", a.RelativeBase+a.Code[p])
	}
	return fmt.Sprintf("c[%04d]", a.Code[p])
}

func (a *Amp) addStack(line string) {
	a.stackTrace = append(a.stackTrace, line)
	if a.Debug {
		fmt.Println(strings.ReplaceAll(line, "\t", " "))
	}
}

func (a *Amp) Dump() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	w.Write([]byte(strings.Join(a.stackTrace, "\n")))
	w.Flush()
}

func (a *Amp) SetInput(i int, goroutine bool) {
	a.Input.Memory[a.Input.Pos] = i
	if a.Input.UseChannel && goroutine {
		go func() {
			a.SetInput(i, false)
		}()
		return
	} else if a.Input.UseChannel {
		a.Input.Chan <- i
	}
}

func OpName(op int) string {
	opStr := "UNKNOW"
	switch op {
	case 1:
		opStr = "ADD"
	case 2:
		opStr = "MUL"
	case 3:
		opStr = "INPUT"
	case 4:
		opStr = "OUTPUT"
	case 5:
		opStr = "JMPNE"
	case 6:
		opStr = "JMPEQ"
	case 7:
		opStr = "LT"
	case 8:
		opStr = "EQ"
	case 9:
		opStr = "RLTVBS"
	case 99:
		opStr = "STOP"
	}

	return opStr
}

func (a *Amp) ExecOp() {
	op := a.Code[a.Pos]

	modes := make([]int, 7)
	if op > 99 {
		val := fmt.Sprintf("%09d", op)
		for i := 6; i >= 0; i-- {
			m, _ := strconv.Atoi(string(val[i]))
			modes[6-i] = m
		}
		op = op % 100
	}

	// Params management
	paramsCount := 0
	paramsValues := []int{}
	switch op {
	case 1:
		paramsCount = 3
		paramsValues = append(paramsValues, []int{0, 1}...)
	case 2:
		paramsCount = 3
		paramsValues = append(paramsValues, []int{0, 1}...)
	case 3:
		paramsCount = 1
	case 4:
		paramsCount = 1
		paramsValues = append(paramsValues, []int{0}...)
	case 5:
		paramsCount = 2
		paramsValues = append(paramsValues, []int{0, 1}...)
	case 6:
		paramsCount = 2
		paramsValues = append(paramsValues, []int{00, 1}...)
	case 7:
		paramsCount = 3
		paramsValues = append(paramsValues, []int{0, 1}...)
	case 8:
		paramsCount = 3
		paramsValues = append(paramsValues, []int{0, 1}...)
	case 9:
		paramsCount = 1
		paramsValues = append(paramsValues, []int{0}...)
	case 99:
		paramsCount = 0
	default:
		a.Dump()
		panic(fmt.Sprintf("unknow op: %d at pos %d", op, a.Pos))
	}

	poses := make([]int, paramsCount)
	for i := 0; i < paramsCount; i++ {
		poses[i] = a.GetPos(a.Pos+1+i, modes[i])
	}

	values := make([]int, paramsCount)
	for _, v := range paramsValues {
		values[v] = a.Code[poses[v]]
	}

	baseStack := func() string {
		s := fmt.Sprintf("%s\tpos %04d\t%-6s\t%d", a.Name, a.Pos, OpName(op), a.Code[a.Pos])
		for i := 0; i < paramsCount; i++ {
			s += fmt.Sprintf(",%d", a.Code[a.Pos+i+1])
		}
		s += "\t"
		return s
	}

	// Exec operation
	switch op {
	case 1:
		r := values[0] + values[1]
		a.addStack(fmt.Sprintf("%s: c[%04d] =\t%s + %s =\t%d + %d =\t%d", baseStack(), poses[2], a.getTraceVal(a.Pos+1, modes[0]), a.getTraceVal(a.Pos+2, modes[1]), values[0], values[1], r))
		a.Code[poses[2]] = r
		a.Pos += 4
	case 2:
		r := values[0] * values[1]
		a.addStack(fmt.Sprintf("%s: c[%04d] =\t%s * %s =\t%d * %d =\t%d", baseStack(), poses[2], a.getTraceVal(a.Pos+1, modes[0]), a.getTraceVal(a.Pos+2, modes[1]), values[0], values[1], r))
		a.Code[poses[2]] = r
		a.Pos += paramsCount + 1
	case 3:
		var valueInput int
		if a.Input.UseChannel {
			valueInput = <-a.Input.Chan
		} else {
			valueInput = a.Input.Memory[a.Input.Pos]
		}
		a.addStack(fmt.Sprintf("%s: %s =\t%d", baseStack(), a.getTraceVal(a.Pos+1, modes[0]), valueInput))
		a.Code[poses[0]] = valueInput
		a.Pos += paramsCount + 1
	case 4:
		a.addStack(fmt.Sprintf("%s: c[%04d] :\t%d", baseStack(), a.Code[a.Pos+1], values[0]))
		if a.Output.WaitFor {
			a.Output.WaitGroup.Add(1)
		}
		a.Output.Chan <- values[0]
		if a.Output.WaitFor {
			a.Output.WaitGroup.Wait()
		}
		a.Pos += paramsCount + 1
	case 5:
		r := values[0] != 0
		a.addStack(fmt.Sprintf("%s: POS = %s = %d\tIF %s != 0 :\t%d != 0 :\t%v", baseStack(), a.getTraceVal(a.Pos+2, modes[1]), values[1], a.getTraceVal(a.Pos+1, modes[0]), values[0], r))
		if r {
			a.Pos = values[1]
		} else {
			a.Pos += paramsCount + 1
		}
	case 6:
		r := values[0] == 0
		a.addStack(fmt.Sprintf("%s: POS = %s = %d\tIF %s == 0 :\t%d == 0 :\t%v", baseStack(), a.getTraceVal(a.Pos+2, modes[1]), values[1], a.getTraceVal(a.Pos+1, modes[0]), values[0], r))
		if r {
			a.Pos = values[1]
		} else {
			a.Pos += paramsCount + 1
		}
	case 7:
		r := values[0] < values[1]
		a.addStack(fmt.Sprintf("%s: c[%04d] =\t%s < %s\t:\t%d < %d :\t%v", baseStack(), poses[2], a.getTraceVal(a.Pos+1, modes[0]), a.getTraceVal(a.Pos+2, modes[1]), values[0], values[1], r))
		if r {
			a.Code[poses[2]] = 1
		} else {
			a.Code[poses[2]] = 0
		}
		a.Pos += paramsCount + 1
	case 8:
		r := values[0] == values[1]
		a.addStack(fmt.Sprintf("%s: c[%04d] =\t%s == %s :\t%d == %d :\t%v", baseStack(), poses[2], a.getTraceVal(a.Pos+1, modes[0]), a.getTraceVal(a.Pos+2, modes[1]), values[0], values[1], r))
		if r {
			a.Code[poses[2]] = 1
		} else {
			a.Code[poses[2]] = 0
		}
		a.Pos += paramsCount + 1
	case 9:
		posNew := a.RelativeBase + values[0]
		a.addStack(fmt.Sprintf("%s: SET FROM %d TO %d + %s =\t%d + %d =\t%d", baseStack(), a.RelativeBase, a.RelativeBase, a.getTraceVal(a.Pos+1, modes[0]), a.RelativeBase, values[0], posNew))
		a.RelativeBase = posNew
		a.Pos += paramsCount + 1
	case 99:
		a.addStack(fmt.Sprintf("%s: STOP", baseStack()))
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
	Amps                []*Amp
	Input               chan int
	Output              chan int
	InjectOutputToInput bool
	done                chan struct{}
}

func (c *ChainAmps) IsRunning() bool {
	for _, a := range c.Amps {
		if a.IsRunning {
			return true
		}
	}
	return false
}

func (c *ChainAmps) None() {
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
		if _, in := find(closed, a.Output.Chan); !in {
			closed = append(closed, a.Output.Chan)
			close(a.Output.Chan)
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
		if _, in := find(closed, a.Input.Chan); !in {
			closed = append(closed, a.Input.Chan)
			close(a.Input.Chan)
		}
	}
	close(c.done)
}

func (c *ChainAmps) Dump() {
	for _, a := range c.Amps {
		a.Dump()
	}
}

func getChainAmps(codes [][]int, init []int) *ChainAmps {
	cInput := make(chan int, 128)
	cOutput := make(chan int, 128)

	amps := []*Amp{}
	var cInputPrev chan int = cInput
	for i, c := range codes {
		cAmpOutput := make(chan int, 128)
		amp := &Amp{
			Name:       string(65 + i),
			MemorySize: 2048,
			Input: AmpInput{
				Init:       init,
				UseChannel: false,
				Chan:       cInputPrev,
			},
			Output: AmpOutput{
				Chan:    cAmpOutput,
				WaitFor: true,
			},
			Code:  append(c[:0:0], c...),
			Debug: true,
		}
		cInputPrev = cAmpOutput
		init = nil
		amps = append(amps, amp)
	}

	c := &ChainAmps{
		Amps:   amps,
		Input:  cInput,
		Output: cOutput,
	}

	go func() {
		for o := range amps[len(amps)-1].Output.Chan {
			cOutput <- o
			if c.InjectOutputToInput {
				cInput <- o
			}
		}
		c.done <- struct{}{}
	}()

	return c
}

func processChainAmps(code [][]int, phase []int) ([]int, *ChainAmps) {
	c := getChainAmps(code, phase)
	c.Start()
	output := []int{}
	for o := range c.Output {
		output = append(output, o)
	}
	return output, c
}

func processCodeChainAmp(codes [][]int) {
	fmt.Println("---")
	fmt.Println("process")
	o, c := processChainAmps(codes, nil)
	for _, oo := range o {
		fmt.Println(oo)
	}
	c.None()
	// c.Dump()
}

func getAmp(code []int, init []int) *Amp {
	cInput := make(chan int)
	cOutput := make(chan int)
	return &Amp{
		Name:       "A",
		MemorySize: 2048,
		Input: AmpInput{
			Init:       init,
			UseChannel: false,
			Chan:       cInput,
		},
		Output: AmpOutput{
			Chan:    cOutput,
			WaitFor: true,
		},
		Code:  append(code[:0:0], code...),
		Debug: false,
	}
}

type Robot struct {
	Map       map[string]int
	MapEdit   map[string]int
	PosX      int
	PosY      int
	Direction int // up = 0, -1 == left, 1 = right, down = 2
	Color     int // black = 0, white = 1
}

func processCodeAmp(codes [][]int) ([]int, *Amp) {
	fmt.Println("---")
	fmt.Println("process")
	// r := Robot{
	// 	Map: map[string]int{},
	// }
	// r.Map["0/0"] = 1
	a := getAmp(codes[0], []int{})
	a.Start()

	output := []int{}
	minX := 0
	minY := 0
	maxX := 0
	maxY := 0

	carte := map[string]int{}

	mode := 0

	drawMap := map[int]string{
		0: " ", // vide
		1: "█", // mural
		2: "░", // bloc
		3: "═", // palette horizontal
		4: "Θ", // boule rebondissante
	}

	x := 0
	y := 0
	for a.IsRunning {
		select {
		case o := <-a.Output.Chan:
			// output = append(output, o)
			// fmt.Println(o)
			if mode == 0 {
				x = o
			} else if mode == 1 {
				y = o
			} else if mode == 2 {
				mode = -1
				carte[fmt.Sprintf("%d/%d", x, y)] = o
				if minX > x {
					minX = x
				}
				if minY > y {
					minY = y
				}
				if maxX < x {
					maxX = x
				}
				if maxY < y {
					maxY = y
				}
			}
			mode++
			if a.Output.WaitFor {
				a.Output.WaitGroup.Done()
			}
		default:
		}
	}
	// close(a.Input.Chan)
	close(a.Output.Chan)

	fmt.Println()
	fmt.Printf("Mmap: %v\n", carte)
	fmt.Println()

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if a, ok := carte[fmt.Sprintf("%d/%d", x, y)]; ok {
				fmt.Printf(drawMap[a])
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	t := 0
	for _, v := range carte {
		if v == 2 {
			t++
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Printf("Total: %d\n", t)
	fmt.Printf("Count map: %d\n", len(carte))

	// a.Dump()
	return output, a
}

func main() {
	// processCodeAmp([][]int{parseLine("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")})
	// processCodeAmp([][]int{parseLine("1102,34915192,34915192,7,4,7,99,0")})
	// processCodeAmp([][]int{parseLine("104,1125899906842624,99")})
	processCodeAmp(parseFile("data.txt"))
	// processCodeChainAmp([][]int{parseLine("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")})
	// processCodeChainAmp([][]int{parseLine("1102,34915192,34915192,7,4,7,99,0")})
	// processCodeChainAmp([][]int{parseLine("104,1125899906842624,99")})
	// processCodeChainAmp(parseFile("data.txt"))
}
