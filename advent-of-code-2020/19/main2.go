package main

import (
	"fmt"
	"regexp"
	"strings"
)

func processMain2(f string) {
	rules, checks := parse(f)

	// update rule 8
	rules["8"] = &Rule{
		id:       "8",
		original: "42 | 42 8",
		// original: "42 | 14 8",
	}
	rules["8"].parse()

	// update rule 11
	rules["11"] = &Rule{
		id:       "11",
		original: "42 31 | 42 11 31",
	}
	rules["11"].parse()

	for i, r := range rules {
		e := r.expand(rules)
		fmt.Printf("%s: %s: %s\n", i, r.original, e)
	}

	maxLen := 0
	for _, c := range checks {
		if maxLen < len(c) {
			maxLen = len(c)
		}
	}

	// fmt.Println("rules[8].expanded")
	// fmt.Println(rules["8"].expanded)

	// fmt.Println("rules[11].expanded")
	// fmt.Println(rules["11"].expanded)

	// fmt.Println("rules[0].expanded")
	// fmt.Println(rules["0"].expanded)

	rr := []*regexp.Regexp{}
	for i := 1; i < maxLen; i++ {
		s := strings.ReplaceAll(rules["0"].expanded, "{LOOP11}", fmt.Sprintf("{%d}", i))
		// fmt.Println(s)
		r := regexp.MustCompile("^" + s + "$")
		rr = append(rr, r)
	}

	cnt := 0
	for _, c := range checks {
		// fmt.Println(c)
		for _, r := range rr {
			if r.MatchString(c) {
				// fmt.Println("\tmatch")
				cnt++
				break
			}
		}
	}
	fmt.Printf("cnt = %d\n", cnt)

}
func main2() {
	// processMain2("list.test2.txt")
	processMain2("list.txt")
}
