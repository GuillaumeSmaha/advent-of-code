package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 7[a|b] 'input'")
	}

	switch args[0] {
	case "7a", "7":
		fmt.Println(Part1(args[1]))
	case "7b":
		fmt.Println(Part2(args[1]))
	}
}

type Node struct {
	Name     string
	Children []string
	Parents  []string
	Printed  bool
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

var Parser *regexp.Regexp = regexp.MustCompile(`Step (?P<from>\w) must be finished before step (?P<to>\w) can begin.`)

func parseString(s string) (string, string) {
	match := Parser.FindStringSubmatch(s)
	from := match[1]
	to := match[2]
	return from, to
}

func loadGraph(s []string) map[string]*Node {
	graph := make(map[string]*Node)
	for _, line := range s {
		from, to := parseString(line)

		if _, ok := graph[from]; !ok {
			graph[from] = &Node{
				Name:     from,
				Children: []string{to},
				Parents:  []string{},
			}
		} else {
			graph[from].Children = append(graph[from].Children, to)
		}

		if _, ok := graph[to]; !ok {
			graph[to] = &Node{
				Name:     to,
				Children: []string{},
				Parents:  []string{from},
			}
		} else {
			graph[to].Parents = append(graph[to].Parents, from)
		}
	}

	return graph
}

func printGraph(graph map[string]*Node) {
	for _, node := range graph {
		fmt.Printf("Node %v:\n", node.Name)
		fmt.Printf("\tChildren:\n")
		for _, n := range node.Children {
			fmt.Printf("\t\t- %v\n", n)
		}
		fmt.Printf("\tParents:\n")
		for _, n := range node.Parents {
			fmt.Printf("\t\t- %v\n", n)
		}
	}
}

func printGraph2(graph map[string]*Node, start string, padding string) {
	node := graph[start]
	fmt.Printf("%v%v:\n", padding, node.Name)
	for _, n := range node.Children {
		printGraph2(graph, n, padding+"\t")
	}
}

func getGraphParentsNode(graph map[string]*Node) []string {
	starts := []string{}
	for _, node := range graph {
		if len(node.Parents) == 0 {
			starts = append(starts, node.Name)
			fmt.Printf("Start possible: %v\n", node.Name)
		}
	}

	return starts
}

func resetGraphPrinted(graph map[string]*Node) {
	for n, _ := range graph {
		graph[n].Printed = false
	}
}

func getGraphNodeString(graph map[string]*Node, name string) string {

	if graph[name].Printed {
		return ""
	}
	for _, p := range graph[name].Parents {
		if !graph[p].Printed {
			return ""
		}
	}
	graph[name].Printed = true

	sort.Slice(graph[name].Children, func(i, j int) bool {
		return graph[name].Children[i] < graph[name].Children[j]
	})

	res := name
	for _, c := range graph[name].Children {
		res += getGraphNodeString(graph, c)
	}

	return res
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

func Part1(s string) string {
	str := load(s)
	graph := loadGraph(str)
	// printGraph(graph)
	// start := getGraphParentNode(graph)
	// printGraph2(graph, start, "")
	// resetGraphPrinted(graph)

	nodes := make([]string, 0)
	nodesAvailable := make([]string, 0)
	for _, g := range graph {
		available := true
		for _, p := range graph[g.Name].Parents {
			if !graph[p].Printed {
				available = false
				break
			}
		}
		if available {
			nodesAvailable = append(nodesAvailable, g.Name)
		} else {
			nodes = append(nodes, g.Name)
		}
	}

	fmt.Printf("Start list: %v\n", nodes)
	fmt.Printf("Start available: %v\n", nodesAvailable)
	res := ""
	for len(nodesAvailable) != 0 {
		sort.Slice(nodesAvailable, func(i, j int) bool {
			return nodesAvailable[i] < nodesAvailable[j]
		})

		element := nodesAvailable[0]
		if !graph[element].Printed {
			fmt.Printf("--\n")
			fmt.Printf("Element to add: %v\n", element)
			res += element
			fmt.Printf("Res: %v\n", res)
			nodesAvailable = deleteFromSlice(nodesAvailable, element)
			fmt.Printf("List available: %v\n", nodesAvailable)
			fmt.Printf("List remaining: %v\n", nodes)
			graph[element].Printed = true
		}

		toAdd := make([]string, 0)
		for _, n := range nodes {
			available := true
			for _, p := range graph[n].Parents {
				if !graph[p].Printed {
					available = false
					break
				}
			}
			if available {
				fmt.Printf("New available: %v\n", n)
				toAdd = append(toAdd, n)
			}
		}

		for _, n := range toAdd {
			nodes = deleteFromSlice(nodes, n)
		}
		nodesAvailable = append(nodesAvailable, toAdd...)
	}

	return res
}

func stepDuration(element string) int {
	return int(element[0]) - int('A') + 61
}

type Worker struct {
	ID            int
	WorkOn        string
	RemainingTime int
}

func Part2(s string) int {
	str := load(s)
	graph := loadGraph(str)
	// printGraph(graph)
	// start := getGraphParentNode(graph)
	// printGraph2(graph, start, "")
	// resetGraphPrinted(graph)

	nodes := make([]string, 0)
	nodesAvailable := make([]string, 0)
	workers := make([]*Worker, 0)
	for i := 0; i < 5; i++ {
		workers = append(workers, &Worker{
			ID: i,
		})
	}

	// for _, g := range graph {
	// 	available := true
	// 	for _, p := range graph[g.Name].Parents {
	// 		if !graph[p].Printed {
	// 			available = false
	// 			break
	// 		}
	// 	}
	// 	if available {
	// 		nodesAvailable = append(nodesAvailable, g.Name)
	// 	} else {
	// 		nodes = append(nodes, g.Name)
	// 	}
	// }
	for _, g := range graph {
		nodes = append(nodes, g.Name)
	}

	fmt.Printf("Start list: %v\n", nodes)
	fmt.Printf("Start available: %v\n", nodesAvailable)

	fmt.Printf("Second")
	for _, w := range workers {
		fmt.Printf("\tWker %v", w.ID)
	}
	fmt.Printf("\tDone\n")

	res := ""
	time := 0
	isWorking := true
	for isWorking || len(nodesAvailable) != 0 {
		sort.Slice(nodesAvailable, func(i, j int) bool {
			return nodesAvailable[i] < nodesAvailable[j]
		})

		toAdd := make([]string, 0)
		for _, n := range nodes {
			available := true
			for _, p := range graph[n].Parents {
				if !graph[p].Printed {
					available = false
					break
				}
			}
			if available {
				// fmt.Printf("New available: %v\n", n)
				toAdd = append(toAdd, n)
			}
		}
		for _, n := range toAdd {
			nodes = deleteFromSlice(nodes, n)
		}
		nodesAvailable = append(nodesAvailable, toAdd...)

		for _, w := range workers {
			if len(nodesAvailable) > 0 {
				element := nodesAvailable[0]
				if !graph[element].Printed {
					if w.WorkOn != "" {
						continue
					}

					w.WorkOn = element
					w.RemainingTime = stepDuration(element)
					nodesAvailable = deleteFromSlice(nodesAvailable, element)
					// fmt.Printf("--\n")
					// fmt.Printf("Element to add: %v\n", element)
					// fmt.Printf("Worker %v during %v seconds\n", worker.ID, worker.RemainingTime)
					// fmt.Printf("Res: %v\n", res)
					// fmt.Printf("List available: %v\n", nodesAvailable)
					// fmt.Printf("List remaining: %v\n", nodes)
				}
			}
		}

		fmt.Printf(" %3d", time)
		for _, w := range workers {
			if w.WorkOn != "" {
				fmt.Printf("\t   %v", w.WorkOn)
			} else {
				fmt.Printf("\t   .")
			}
		}
		fmt.Printf("\t  %v\n", res)

		isWorking = false
		for _, w := range workers {
			if w.WorkOn != "" {
				isWorking = true
				w.RemainingTime--
				if w.RemainingTime == 0 {
					// fmt.Printf("Worker %v finish task %v\n", w.ID, w.WorkOn)
					graph[w.WorkOn].Printed = true
					res += w.WorkOn
					w.WorkOn = ""
				}
			}
		}

		time++
	}

	return time - 1
}
