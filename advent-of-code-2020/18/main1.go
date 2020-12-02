package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Value int

type Op struct {
	kind int
	v    int
}

type Section struct {
	lvl    int
	start  int
	end    int
	value  int
	solved bool
	str    string
}

type SolverFct func(all []*Section, se *Section)

func solveSection(all []*Section, se *Section) {
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
	// fmt.Printf("calcul rep: %s\n", s)
	v := 0
	st := ""
	prevOp := '0'
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
	vp, err := strconv.Atoi(st)
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
	}

	se.value = v
	se.solved = true
	// fmt.Printf("value = %#v\n", v)
}

func parseOp(s string, solveSection SolverFct) int {
	s = strings.ReplaceAll(s, " ", "")
	// fmt.Printf("--------------------\n")
	// fmt.Printf("%#v\n", s)
	parenthesis := make([]int, len(s))
	maxParenthesis := 0
	p := 0
	for i, r := range s {
		switch r {
		case '(':
			p++
			if p > maxParenthesis {
				maxParenthesis = p
			}
		case ')':
			p--
		}
		parenthesis[i] = p
	}

	sections := []*Section{}
	sectionsMap := map[int][]*Section{}
	for i := range seq(maxParenthesis + 1) {
		sectionsMap[i] = []*Section{}
	}

	pPrev := -1
	sections = append(sections, &Section{
		start: 0,
		end:   -1,
		lvl:   0,
	})
	for i, p := range parenthesis {
		if p < pPrev {
			// fmt.Println("look for lvl", pPrev)
			for _, sec := range sections {
				// fmt.Println(" -- check :", sec)
				if sec.end == -1 && sec.lvl == pPrev {
					v := s[sec.start : i+1]
					sec.str = v
					sec.end = i + 1
					sectionsMap[sec.lvl] = append(sectionsMap[sec.lvl], sec)
					break
				}
			}
		} else if p > pPrev {
			sections = append(sections, &Section{
				start: i,
				end:   -1,
				lvl:   p,
			})
		}
		pPrev = p
	}
	sec := sections[0]
	if sec.end == -1 {
		v := s[sec.start:]
		sec.str = v
		sec.end = len(s)
		sectionsMap[sec.lvl] = append(sectionsMap[sec.lvl], sec)
	}

	// fmt.Printf("len s: %d\n", len(s))
	// fmt.Printf("%#v\n", parenthesis)
	// fmt.Printf("Sections:\n")
	// for _, sec := range sections {
	// 	fmt.Printf("\t%#v\n", sec)
	// }
	// fmt.Printf("%#v\n", sectionsMap)
	// for lvl, secs := range sectionsMap {
	// 	fmt.Printf("\tLvl: %#v\n", lvl)
	// 	for _, sec := range secs {
	// 		fmt.Printf("\t\t%#v\n", sec)
	// 	}
	// }

	for _, i := range seqReverse(maxParenthesis + 1) {
		// fmt.Printf("%#v\n", i)
		for _, sec := range sectionsMap[i] {
			solveSection(sections, sec)
		}
	}

	return sections[0].value
}

func processMain1(lines []string) {
	s := 0
	for _, l := range lines {
		// fmt.Printf("%s ?..\n", l)
		v := parseOp(l, solveSection)
		s += v
		fmt.Printf("%s = %d\n", l, v)
	}
	fmt.Printf("Result: %d\n", s)
}

func main1() {
	// processMain1(parseFileText1D("list.test.txt"))
	// processMain1(parseFileText1D("list.test2.txt"))
	// processMain1(parseFileText1D("list.test3.txt"))
	// processMain1(parseFileText1D("list.test4.txt"))
	processMain1(parseFileText1D("list.txt"))
}
