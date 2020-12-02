package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 8[a|b] 'input'")
	}

	switch args[0] {
	case "8a", "8":
		fmt.Println(Part1(args[1]))
	case "8b":
		fmt.Println(Part2(args[1]))
	}
}

type Node struct {
	ID         int
	Children   map[int]int
	Parents    map[int]int
	Metadata   []int
	Value      int
	Calculated bool
}

func load(s string) []string {
	res := make([]string, 0)
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = strings.TrimSpace(string(f))

		for _, line := range strings.Split(s, "\n") {
			if len(line) == 0 {
				continue
			}
			res = append(res, line)
		}
	}

	return res
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

var GlobalIdx int

func parseNode(tree map[int]*Node, values []string, start int) (int, int) {
	newIdx := GlobalIdx
	GlobalIdx++
	// fmt.Printf("---\n")
	// fmt.Printf("Node %v:\n", newIdx)
	// fmt.Printf("[%v]Start %v:\n", newIdx, start)

	tree[newIdx] = &Node{
		ID:       newIdx,
		Children: map[int]int{},
		Parents:  map[int]int{},
	}

	childrenCount, _ := strconv.Atoi(values[start])
	metadataCount, _ := strconv.Atoi(values[start+1])
	// fmt.Printf("[%v]childrenCount %v:\n", newIdx, childrenCount)
	// fmt.Printf("[%v]metadataCount %v:\n", newIdx, metadataCount)

	ptr := start + 2
	for i := 0; i < childrenCount; i++ {
		newPtr, idx := parseNode(tree, values, ptr)
		ptr = newPtr
		tree[newIdx].Children[len(tree[newIdx].Children)] = idx
		tree[idx].Parents[len(tree[idx].Parents)] = newIdx
	}
	// fmt.Printf("[%v]Children %v:\n", newIdx, tree[newIdx].Children)

	for j := ptr; j < (ptr + metadataCount); j++ {
		val, _ := strconv.Atoi(values[j])
		tree[newIdx].Metadata = append(tree[newIdx].Metadata, val)
	}
	// fmt.Printf("[%v]Metadata %v:\n", newIdx, tree[newIdx].Metadata)

	return ptr + metadataCount, newIdx
}

func loadTree(s string) map[int]*Node {

	tree := make(map[int]*Node)
	values := strings.Split(s, " ")
	GlobalIdx = 0
	parseNode(tree, values, 0)
	return tree
}

func printTree(tree map[int]*Node) {
	for _, node := range tree {
		fmt.Printf("Node %v:\n", node.ID)
		fmt.Printf("\tChildren:\n")
		for _, n := range node.Children {
			fmt.Printf("\t\t- %v\n", n)
		}
		fmt.Printf("\tParents:\n")
		for _, n := range node.Parents {
			fmt.Printf("\t\t- %v\n", n)
		}
		fmt.Printf("\tMetadata:\n")
		for _, n := range node.Metadata {
			fmt.Printf("\t\t- %v\n", n)
		}
	}
}

func sumMetadata(tree map[int]*Node) int {
	sum := 0
	for _, node := range tree {
		for _, n := range node.Metadata {
			sum += n
		}
	}
	return sum
}

func Part1(s string) int {
	str := load(s)
	tree := loadTree(str[0])
	printTree(tree)

	return sumMetadata(tree)
}

func calculateNode(tree map[int]*Node, idx int) int {
	node := tree[idx]
	if node.Calculated {
		return node.Value
	}

	if len(node.Children) == 0 {
		sum := 0
		for _, n := range node.Metadata {
			sum += n
		}
		node.Value = sum
		node.Calculated = true
		return sum
	} else {
		sum := 0
		for _, n := range node.Metadata {
			childIdx, ok := node.Children[n-1]
			if ok {
				sum += calculateNode(tree, childIdx)
			}
		}
		node.Value = sum
		node.Calculated = true
		return sum
	}
}

func Part2(s string) int {
	str := load(s)
	tree := loadTree(str[0])
	// printTree(tree)
	// for i, _ := range tree {
	// 	fmt.Printf("val %v = %v\n", i, calculateNode(tree, i))
	// }

	return calculateNode(tree, 0)
}
