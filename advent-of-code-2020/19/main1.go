package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Rule struct {
	id           string
	original     string
	expanded     string
	expandedDone bool
}

func (r *Rule) parse() {
	if r.original[0] == '"' {
		r.expanded = r.original[1:2]
		r.expandedDone = true
		return
	}
}

func (r *Rule) expand(rules Rules) string {
	if r.expandedDone {
		return r.expanded
	}

	partsOr := strings.Split(r.original, "|")
	if r.id == "8" || r.id == "11" {
		fmt.Println()
		fmt.Println()
		fmt.Println("rules")
	}
	for io, part := range partsOr {
		part = strings.TrimSpace(part)
		parts := strings.Split(part, " ")
		partsSolved := make([]bool, len(parts))

		if r.id == "8" || r.id == "11" {
			fmt.Println("parts")
			fmt.Println(parts)
		}

		for i, p := range parts {
			if p == r.id {
				partsSolved[i] = true
				parts[i] = ""
				rep := ")+"
				if i > 0 && i < len(p) {
					rep = fmt.Sprintf("){LOOP%s}", r.id)
				}
				if i > 0 {
					parts[0] = "(" + strings.TrimSpace(parts[0])
					parts[i] += rep
				}
				if i < len(p) {
					fmt.Println("add after")
					parts[i] += "("
					parts = append(parts, rep)
					partsSolved = append(partsSolved, true)
				}
			} else if !partsSolved[i] {
				parts[i] = rules[p].expand(rules)
				partsSolved[i] = true
			}
		}

		if r.id == "8" || r.id == "11" {
			fmt.Println(parts)
		}

		partsOr[io] = strings.Join(parts, "")
	}

	r.expandedDone = true
	r.expanded = "(" + strings.Join(partsOr, "|") + ")"
	return r.expanded
}

type Rules map[string]*Rule

func parse(f string) (Rules, []string) {
	lines := parseFileText1D(f)
	rules := Rules{}
	checks := []string{}
	for _, l := range lines {
		if l == "" {
			continue
		}
		i := strings.Index(l, ":")
		if i == -1 {
			checks = append(checks, l)
		} else {
			n := l[:i]
			r := &Rule{
				id:       n,
				original: l[i+2:],
			}
			r.parse()
			rules[n] = r
		}
	}
	return rules, checks
}

func processMain1(f string) {
	rules, checks := parse(f)

	for i, r := range rules {
		e := r.expand(rules)
		fmt.Printf("%s: %s: %s\n", i, r.original, e)
	}

	r0 := regexp.MustCompile("^" + rules["0"].expanded + "$")

	cnt := 0
	for _, c := range checks {
		// fmt.Println(c)
		if r0.MatchString(c) {
			// fmt.Println("\tmatch")
			cnt++
		}
	}
	fmt.Printf("cnt = %d\n", cnt)
}

func main1() {
	// processMain1("list.test.txt")
	processMain1("list.txt")
}
