package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveSectionPlusPrecedence(all []*Section, se *Section) {
	s := se.str
	// fmt.Printf("----\n")
	// fmt.Printf("calcul: %s\n", s)
	for _, sec := range all {
		if sec.solved && sec.lvl > se.lvl {
			s = strings.Replace(s, sec.str, fmt.Sprintf("%d", sec.value), 1)
		}
	}
	if s[0] == '(' && s[len(s)-1] == ')' {
		s = s[1 : len(s)-1]
	}
	// fmt.Printf("calcul rep +: %s\n", s)

	vsave := 0
	st := ""
	snew := ""
	prevOp := '0'
	for _, c := range s {
		switch c {
		case '+', '-', '*', '/':
			vp, err := strconv.Atoi(st)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("\t---\n")
			// fmt.Printf("\t c = %c\n", c)
			// fmt.Printf("\t prevOp = %c\n", prevOp)
			// fmt.Printf("\t vsave = %v\n", vsave)
			// fmt.Printf("\t snew = %v\n", snew)
			// fmt.Printf("\t vp = %v\n", vp)
			// fmt.Printf("%v op(%c) %v\n", vp, prevOp, v)
			switch prevOp {
			case '+':
				vsave = vsave + vp
			case '-':
				vsave = vsave - vp
			default:
				if prevOp != '0' {
					snew += fmt.Sprintf("%d%c", vsave, prevOp)
				}
				prevOp = '0'
				vsave = vp
			}
			// fmt.Printf("\t vsave after = %v\n", vsave)
			// fmt.Printf("\t snew after = %v\n", snew)
			prevOp = c
			st = ""
		default:
			st += fmt.Sprintf("%c", c)
		}
	}
	vp, err := strconv.Atoi(st)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v op(%c) %v\n", vp, prevOp, v)
	switch prevOp {
	case '+':
		vsave = vsave + vp
	case '-':
		vsave = vsave - vp
	default:
		snew += fmt.Sprintf("%d%c", vsave, prevOp)
		vsave = vp
	}
	snew += fmt.Sprintf("%d", vsave)

	// fmt.Printf("calcul rep after +: %s\n", snew)
	s = snew
	v := 0
	st = ""
	prevOp = '0'
	for _, c := range s {
		switch c {
		case '+', '-', '*', '/':
			vp, err := strconv.Atoi(st)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("%v op(%c) %v\n", vp, prevOp, v)
			switch prevOp {
			case '+':
				v = v + vp
			case '-':
				v = v - vp
			case '*':
				v = v * vp
			case '/':
				v = v / vp
			default:
				v = vp
			}
			prevOp = c
			st = ""
		default:
			st += fmt.Sprintf("%c", c)
		}
	}
	vp, err = strconv.Atoi(st)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v op(%c) %v (end)\n", vp, prevOp, v)
	switch prevOp {
	case '+':
		v = v + vp
	case '-':
		v = v - vp
	case '*':
		v = v * vp
	case '/':
		v = v / vp
	default:
		v = vp
	}

	se.value = v
	se.solved = true
	// fmt.Printf("value = %#v\n", v)
}

func processMain2(lines []string) {
	s := 0
	for _, l := range lines {
		// fmt.Printf("%s ?..\n", l)
		v := parseOp(l, solveSectionPlusPrecedence)
		s += v
		fmt.Printf("%s = %d\n", l, v)
	}
	fmt.Printf("Result: %d\n", s)
}

func main2() {
	// processMain2(parseFileText1D("list.test.txt"))
	// processMain2(parseFileText1D("list.test2.txt"))
	// processMain2(parseFileText1D("list.test3.txt"))
	// processMain2(parseFileText1D("list.test4.txt"))
	processMain2(parseFileText1D("list.txt"))
}
