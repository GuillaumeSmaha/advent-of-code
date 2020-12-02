package p1

import (
	"fmt"
	"log"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 9[a|b] 'input'")
	}

	switch args[0] {
	case "9a", "9":
		fmt.Println(Part1(args[1]))
	case "9b":
		fmt.Println(Part2(args[1]))
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func deleteFromSlice(list []string, element string) []string {
	for i, v := range list {
		if v == element {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}

	return list
}

func play(l List, marble int) {

}

func Part1(s string) int {
	// l := new(List)

	return 0
}

func Part2(s string) int {
	// printTree(tree)
	// for i, _ := range tree {
	// 	fmt.Printf("val %v = %v\n", i, calculateNode(tree, i))
	// }

	return 0
}
