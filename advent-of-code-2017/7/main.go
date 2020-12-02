package p7

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 7[a|b] 'filename'")
	}

	switch args[0] {
	case "7a", "7":
		fmt.Print(Root(args[1]))
	case "7b":
		fmt.Print(FirstWrong(args[1]))
	}
}

func load(s string) Nodes {
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = string(f)
	}

	nodes := Nodes{}
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		nodes.Add(line)
	}

	return nodes
}

type Nodes map[string]*Node

type Node struct {
	Name     string
	Disk     int
	Parent   string
	Children []string
}

func (m Nodes) Add(line string) {
	p := strings.Split(line, " ")

	get := func(name string) *Node {
		node, ok := m[name]
		if !ok {
			node = &Node{Name: name}
			m[name] = node
		}

		return node
	}

	node := get(p[0])

	k, err := strconv.Atoi(p[1][1 : len(p[1])-1])
	if err != nil {
		log.Fatal(err)
	}

	node.Disk = k

	if len(p) == 2 || len(p) == 3 {
		return
	}

	for _, i := range p[3:] {
		i = strings.Trim(i, ",")
		node.Children = append(node.Children, i)
		child := get(i)
		if child.Parent != "" {
			log.Fatal("more than 1 parent")
		}

		child.Parent = node.Name
	}
}

func (m Nodes) FirstWrong(name string) (int, bool) {
	node := m[name]

	if len(node.Children) == 0 {
		return node.Disk, false
	}

	items := []int{}

	for _, i := range node.Children {
		n, ok := m.FirstWrong(i)
		if ok {
			return n, ok
		}

		items = append(items, n)
	}

	n := make(map[int]int)

	for _, i := range items {
		n[i]++
	}

	if len(n) == 1 {
		for w, k := range n {
			return node.Disk + w*k, false
		}
	}

	if len(n) != 2 {
		log.Fatal("should only have 2 values")
	}

	one, all := []int{}, []int{}

	for k, v := range n {
		if v == 1 {
			one = append(one, k)
		} else {
			all = append(all, k)
		}
	}

	if len(one) != 1 || len(all) != 1 {
		log.Fatal("cannot figure out good and bad weight")
	}

	bad := -1
	for i, v := range items {
		if v == one[0] {
			bad = i
		}
	}

	wrong := m[node.Children[bad]]
	value := wrong.Disk + all[0] - items[bad]
	return value, true
}

func (m Nodes) Root() string {
	root := ""
	for k, v := range m {
		if v.Parent == "" {
			if root != "" {
				log.Fatal("more than 1 root")
			}

			root = k
		}
	}

	return root
}

func Root(s string) string {
	nodes := load(s)
	return nodes.Root()
}

func FirstWrong(s string) int {
	nodes := load(s)

	root := nodes.Root()
	n, ok := nodes.FirstWrong(root)
	if !ok {
		log.Fatal("nothing is unbalanced")
	}

	return n
}
