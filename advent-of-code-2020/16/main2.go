package main

import (
	"fmt"
	"strings"
)

func processMain2(data *Data) {
	fmt.Println("----------")
	// fmt.Println("Valid Tickets Idx:")
	m := map[int]map[string]struct{}{}
	wordsIdx := map[string][]int{}
	for i := range data.Ticket {
		// fmt.Printf("\ti = %#v\n", i)
		// fmt.Printf("\tfor idx: %v\n", data.FindNamesForIdx(i))

		vIdx := data.FindNamesForIdx(i)
		m[i] = map[string]struct{}{}
		for k := range vIdx {
			m[i][k] = struct{}{}
			if _, ok := wordsIdx[k]; !ok {
				wordsIdx[k] = []int{}
			}
			wordsIdx[k] = append(wordsIdx[k], i)
		}
	}

	// fmt.Printf("len(wordsIdx) = %#v\n", len(wordsIdx))
	// fmt.Println("wordsIdx:")
	// for k, words := range wordsIdx {
	// 	fmt.Printf("\t%v: %#v\n", k, words)
	// }

	wordsCntIdx := map[int][]string{}
	for k, v := range wordsIdx {
		l := len(v)
		if _, ok := wordsCntIdx[l]; !ok {
			wordsCntIdx[l] = []string{}
		}
		wordsCntIdx[l] = append(wordsCntIdx[l], k)
	}

	// fmt.Println("wordsCntIdx:")
	// for k, words := range wordsCntIdx {
	// 	fmt.Printf("\t%v: %#v\n", k, words)
	// }

	idxUsed := make([]bool, len(data.Ticket))
	idxWords := make([]string, len(data.Ticket))
	for i := 0; i < len(wordsCntIdx); i++ {
		if words, ok := wordsCntIdx[i]; ok {
			for _, w := range words {
				for _, idx := range wordsIdx[w] {
					if !idxUsed[idx] {
						idxWords[idx] = w
						idxUsed[idx] = true
						break
					}
				}
			}
		}
	}

	fmt.Printf("%#v\n", idxUsed)
	fmt.Printf("%#v\n", idxWords)

	res := 1
	for i, w := range idxWords {
		if strings.Contains(w, "departure") {
			fmt.Printf("Found %v : idx=%d\n", w, i)
			res *= data.Ticket[i]
		}
	}
	fmt.Printf("Total: %#v\n", res)
}

func main2() {
	// processMain2(parseFile("list.test.txt"))
	processMain2(parseFile("list.test2.txt"))
	processMain2(parseFile("list.txt"))
}
