package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 12[a|b] 'input'")
	}

	switch args[0] {
	case "12a", "12":
		fmt.Println(Part1(args[1]))
	case "12b":
		fmt.Println(Part2(args[1]))
	}
}

func load(s string) []string {
	res := make([]string, 0)
	if strings.HasPrefix(s, "@") {
		f, err := ioutil.ReadFile(s[1:])
		if err != nil {
			log.Fatal(err)
		}

		s = strings.TrimSpace(string(f))

		for i, line := range strings.Split(s, "\n") {
			if len(line) == 0 {
				continue
			}
			fmt.Println(i, line)
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

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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

type Pot struct {
	ID    int
	Plant bool
}

// var Parser *regexp.Regexp = regexp.MustCompile(`Step (?P<from>\w) must be finished before step (?P<to>\w) can begin.`)
var Parser *regexp.Regexp = regexp.MustCompile(`(.+) => (.)`)

func parseString(s string) (string, string) {
	match := Parser.FindStringSubmatch(s)
	from := match[1]
	to := match[2]
	return from, to
}

type Rule struct {
	Regex  *regexp.Regexp
	Match  string
	Result string
}

func loadData(s map[int]string) map[int]*Rule {
	rules := make(map[int]*Rule)
	for i, line := range s {
		from, to := parseString(line)

		rules[i] = &Rule{
			Regex:  regexp.MustCompile("/" + strings.Replace(from, ".", "\\.", -1) + "/"),
			Match:  from,
			Result: to,
		}
	}

	return rules
}

func loadData2(s []string) map[string]string {
	rules := make(map[string]string)
	for _, line := range s {
		from, to := parseString(line)

		rules[from] = to
	}

	return rules
}

// func printData(tree map[int]*Node) {
// 	for _, node := range tree {
// 		fmt.Printf("Node %v:\n", node.ID)
// 		fmt.Printf("\tChildren:\n")
// 		for _, n := range node.Children {
// 			fmt.Printf("\t\t- %v\n", n)
// 		}
// 		fmt.Printf("\tParents:\n")
// 		for _, n := range node.Parents {
// 			fmt.Printf("\t\t- %v\n", n)
// 		}
// 		fmt.Printf("\tMetadata:\n")
// 		for _, n := range node.Metadata {
// 			fmt.Printf("\t\t- %v\n", n)
// 		}
// 	}
// }

func replaceAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func countHashtag(str string) int {
	cnt := 0
	for _, c := range str {
		if c == '#' {
			cnt++
		}
	}
	return cnt
}

func sumPlant(str string, it int) int {
	sum := 0
	center := (len(str) - it) / 2
	for i, c := range str {
		if c == '#' {
			sum += (i - center)
		}
	}
	return sum
}

func Part1(s string) int {
	str := load(s)
	rules := loadData2(str)

	pots := "#..#.#..##......###...###"
	pots = ".##..##..####..#.#.#.###....#...#..#.#.#..#...#....##.#.#.#.#.#..######.##....##.###....##..#.####.#"
	cnt := countHashtag(pots)
	sum := sumPlant(pots, 0)
	fmt.Printf("%02d: %v (cnt:%02d) (sum:%02d)\n", 0, pots, cnt, sum)
	// fmt.Printf("\t%02d: %v\n", 0, pots)
	for i := 0; i < 20; i++ {
		// fmt.Printf("%02d: %v (cnt:%02d) (sum:%02d)\n", i+1, pots, cnt, sum)
		pots = "...." + pots + "...."
		// fmt.Printf("\t%02d: %v\n", i+1, pots)
		newPots := ""
		// fmt.Printf("\tRule %v = %v:\n", r.Match, r.Result)
		for j := 2; j < len(pots)-2; j++ {
			sub := pots[j-2 : j+3]
			res, ok := rules[sub]
			if ok {
				newPots += res
			}
		}
		pots = newPots
		cnt = countHashtag(pots)
		sum = sumPlant(pots, (i+1)*5)
		fmt.Printf("%02d: %v (cnt:%02d) (sum:%02d)\n", i+1, pots, cnt, sum)
		// fmt.Printf("\t%02d: %v\n", i+1, pots)
	}

	return sumPlant(pots, 20*5)
}

func Part2(s string) int {
	str := load(s)
	rules := loadData2(str)

	// pots := loadInitPot("#..#.#..##......###...###")
	// applyS(str)
	// printData(data)

	initPots := "#..#.#..##......###...###"
	initPots = ".##..##..####..#.#.#.###....#...#..#.#.#..#...#....##.#.#.#.#.#..######.##....##.###....##..#.####.#"
	pots := initPots
	cnt := countHashtag(pots)
	sum := sumPlant(pots, 0)
	zero := 0
	prevZero := zero
	fmt.Printf("%02d: %v (cnt:%02d) (sum:%02d) (zero:%d)\n", 0, pots, cnt, sum, zero)
	// fmt.Printf("\t%02d: %v\n", 0, pots)
	// 2000000
	// 50000000000
	prevPots := pots
	for i := 0; i < 50000000000; i++ {
		// fmt.Printf("%02d: %v (cnt:%02d) (sum:%02d)\n", i+1, pots, cnt, sum)
		pots = "...." + pots + "...."
		zero += 5
		// fmt.Printf("\t%02d: %v\n", i+1, pots)
		newPots := ""
		// fmt.Printf("\tRule %v = %v:\n", r.Match, r.Result)
		for j := 2; j < len(pots)-2; j++ {
			sub := pots[j-2 : j+3]
			res, ok := rules[sub]
			if ok {
				newPots += res
			}
		}
		pots = newPots
		for pots[0] == '.' {
			pots = pots[1:]
			zero--
		}
		for pots[len(pots)-1] == '.' {
			pots = pots[:len(pots)-1]
		}

		if prevPots == pots {
			fmt.Printf("End at %v with %v zeros (before %v)\n", i, zero, prevZero)
			fmt.Printf("Sum: %v\n", sumPlant(pots, zero))
			newZero := (50000000000-i+8)*(zero-prevZero) + zero

			fmt.Printf("At %v with %v zeros\n", 50000000000, newZero)
			fmt.Printf("Sum should be %v\n", sumPlant(pots, newZero))
			// return sumPlant(pots, newZero)

			diff := sumPlant(pots, zero) - sumPlant(pots, prevZero)
			fmt.Printf("Sum diff %v\n", diff)
			return sumPlant(pots, zero) + (50000000000-i+8)*diff

			break
		}

		// cnt = countHashtag(pots)
		// sum = sumPlant(pots, (i+1)*5)
		fmt.Printf("%02d: %v (cnt:%02d) (sum:%02d) (zero:%d)\n", i+1, pots, cnt, sum, zero)
		// fmt.Printf("\t%02d: %v\n", i+1, pots)
		prevPots = pots
		prevZero = zero
	}

	return -1
}
