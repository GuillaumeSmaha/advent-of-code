package main

import (
	"fmt"
	"sync"
)

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
