package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
)

type AmpInput struct {
	Init       []int
	AfterCode  bool
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
	if a.Input.AfterCode {
		a.Input.Pos = len(a.Code)
	}
	if a.Input.MemorySize < 1 {
		a.Input.MemorySize = 1
	}
	a.Input.Memory = make([]int, a.Input.MemorySize)
	for i := 0; i < a.MemorySize; i++ {
		a.Code = append(a.Code, 0)
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
				if a.Input.AfterCode {
					a.Code[a.Input.Pos+i] = v
				} else {
					a.Input.Memory[i] = v
				}
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
	if a.Input.AfterCode {
		a.Code[a.Input.Pos] = i
	} else {
		a.Input.Memory[a.Input.Pos] = i
	}
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
			if a.Input.AfterCode {
				valueInput = a.Code[a.Input.Pos]
			} else {
				valueInput = a.Input.Memory[a.Input.Pos]
			}
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
