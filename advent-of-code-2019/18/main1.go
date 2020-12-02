package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// A Data Structure for queue used in BFS
type queueNode struct {
	xy   [2]int // The cordinates of a cell
	dist int    // cell's distance of from the source
}

// check whether given cell (row, col) is a valid
// cell or not.
func isValid(r *Robot, y, x int) bool {
	// return true if row number and column number
	// is in range
	minX, minY, maxX, maxY := r.Bounds()
	return (y >= minY) && (y < maxY) &&
		(x >= minX) && (x < maxX)
}

// These arrays are used to get row and column
// numbers of 4 neighbours of a given cell
var yNum = []int{-1, 0, 0, 1}
var xNum = []int{0, -1, 1, 0}

// BFSDoor function to find the shortest path between
// a given source cell to a destination cell.
func BFSDoor(r *Robot, src [2]int, dest [2]int, keys [][2]int, allKeys *map[[2]int]rune, allDoors *map[[2]int]rune) *queueNode {
	// check source and destination cell
	// of the matrix have value 1
	if _, ok := r.Map[src]; !ok {
		return nil
	}
	if _, ok := r.Map[dest]; !ok {
		return nil
	}

	visited := map[[2]int]bool{}

	// Mark the source cell as visited
	visited[src] = true

	// Create a queue for BFS
	q := list.New()

	// Distance of source cell is 0
	s := &queueNode{xy: src, dist: 0}
	q.PushBack(s) // Enqueue source cell

	// Do a BFS starting from source cell
	for q.Len() != 0 {
		curr := q.Front().Value.(*queueNode)
		pt := curr.xy

		// If we have reached the destination cell,
		// we are done
		if pt[0] == dest[0] && pt[1] == dest[1] {
			return curr
		}

		// Otherwise dequeue the front cell in the queue
		// and enqueue its adjacent cells
		q.Remove(q.Front())

		for i := 0; i < 4; i++ {
			x := pt[0] + xNum[i]
			y := pt[1] + yNum[i]
			xy := [2]int{x, y}

			// if adjacent cell is valid, has path and
			// not visited yet, enqueue it.
			v, okMap := r.Map[xy]
			_, okVisited := visited[xy]
			hasKey := false

			// fnd = 0
			// 	for _, kCurrXY := range curr.keys {
			// 		if doors[do] == keys[kCurrXY] {
			// 			fnd++
			// 			break
			// 		}
			// 	}
			// }
			// doors[do] == keys[kCurrXY]
			if v == CASE_DOOR {
				for _, kXY := range keys {
					if (*allKeys)[kXY] == (*allDoors)[xy] {
						hasKey = true
						break
					}
				}
			}
			isCanOpenDoor := v != CASE_DOOR || v == CASE_DOOR && hasKey
			if isValid(r, y, x) && okMap && v != CASE_WALL && isCanOpenDoor && !okVisited {
				// mark cell as visited and enqueue it
				visited[xy] = true
				adjcell := &queueNode{xy: xy, dist: curr.dist + 1}
				q.PushBack(adjcell)
			}
		}
	}

	// Return -1 if destination cannot be reached
	return nil
}

// A Data Structure for queue used in BFS
type queueUserNode struct {
	xy   [2]int // The cordinates of a cell
	dist int    // cell's distance of from the source
	keys [][2]int
}

func process1(filename string) {
	fmt.Println("---")
	fmt.Println("process")

	r := &Robot{}
	keys := map[[2]int]rune{}
	doors := map[[2]int]rune{}
	keysR := map[rune][2]int{}
	doorsR := map[rune][2]int{}

	parseData := func() {
		r.Init()
		file, _ := os.Open(filename)
		fscanner := bufio.NewScanner(file)
		x, y := 0, 0
		for fscanner.Scan() {
			l := fscanner.Text()
			x = 0
			for _, c := range l {
				switch c {
				case '#':
					r.SetXY(x, y, CASE_WALL)
				case '.':
					r.SetXY(x, y, CASE_EMPTY)
				case '@':
					r.SetXY(x, y, CASE_EMPTY)
					// r.SetXY(x-1, y, CASE_WALL)
					r.X = x
					r.Y = y
				default:
					if int(c) < 91 {
						doors[[2]int{x, y}] = c
						doorsR[c] = [2]int{x, y}
						// doors[[2]int{x, y}] = c + ('A' - 'a')
						r.SetXY(x, y, CASE_DOOR)
					} else {
						// keys[[2]int{x, y}] = c
						keys[[2]int{x, y}] = c + ('A' - 'a')
						keysR[c+('A'-'a')] = [2]int{x, y}
						r.SetXY(x, y, CASE_KEY)
					}
				}
				x++
			}
			y++
		}
	}

	parseData()
	r.DrawMap(true, false)

	// graphsDoorKey := map[string][]string{}
	// fmt.Println("Dist doors:")
	// tt := 0
	// for k, v := range keysR {
	// 	p, ok := doorsR[k]
	// 	if ok {
	// 		d := BFS(r, v, p)
	// 		if d != nil {
	// 			tt += d.dist
	// 		}

	// 		s := []string{}
	// 		for do := range d.doors {
	// 			s = append(s, fmt.Sprintf("%c", doors[do]))
	// 		}
	// 		graphsDoorKey[fmt.Sprintf("%c", k)] = s
	// 		fmt.Printf("\t%c: dist: %d\t doors: %s\n", k, d.dist, strings.Join(s, ", "))
	// 	}
	// }
	// fmt.Println("tt = ", tt)
	// fmt.Println("Users keys:")

	// graphsUserKeys := map[string][]string{}
	// tt = 0
	// for k, v := range keysR {
	// 	d := BFS(r, [2]int{r.X, r.Y}, v)
	// 	if d != nil {
	// 		tt += d.dist
	// 	}

	// 	s := []string{}
	// 	for do := range d.doors {
	// 		s = append(s, fmt.Sprintf("%c", doors[do]))
	// 	}
	// 	graphsUserKeys[fmt.Sprintf("%c", k)] = s
	// 	fmt.Printf("\t%c: dist: %d\t doors: %s\n", k, d.dist, strings.Join(s, ", "))
	// }
	// fmt.Println("tt = ", tt)

	src := [2]int{r.X, r.Y}

	cpu, _ := strconv.Atoi(os.Getenv("EAI_CPU_LIMIT"))
	if cpu < 1 {
		cpu = 1
	}

	/*
		ti := time.Now()
		fmt.Printf("Graph key to key with %d workers (%d keys: %d to calcul)\n", cpu, len(keysR), len(keysR)*len(keysR))
		graphsDoorKey := map[[2]int]map[[2]int]*queueNode{}
		graphsDoorKey[src] = map[[2]int]*queueNode{}
		for _, v1 := range keysR {
			graphsDoorKey[v1] = map[[2]int]*queueNode{}
		}

		type trGraph struct {
			src  [2]int
			dest [2]int
		}

		type trGraphRes struct {
			src  [2]int
			dest [2]int
			q    *queueNode
		}

		jobsGraph := make(chan trGraph, 100)
		resultsGraph := make(chan trGraphRes, 100)
		workerGraph := func(r *Robot, jobs chan trGraph, results chan trGraphRes) {
			for j := range jobs {
				results <- trGraphRes{src: j.src, dest: j.dest, q: BFS(r, j.src, j.dest)}
			}
		}

		for w := 0; w < cpu; w++ {
			go workerGraph(r, jobsGraph, resultsGraph)
		}

		wgGraph := sync.WaitGroup{}
		go func() {
			for r := range resultsGraph {
				graphsDoorKey[r.src][r.dest] = r.q
				wgGraph.Done()
			}
		}()

		fmt.Printf("\tUser\n")
		for _, v2 := range keysR {
			wgGraph.Add(1)
			jobsGraph <- trGraph{src, v2}
		}

		for k1, v1 := range keysR {
			fmt.Printf("\t%c\n", k1)
			for k2, v2 := range keysR {
				if k1 == k2 {
					continue
				}
				wgGraph.Add(1)
				jobsGraph <- trGraph{v1, v2}
			}
		}
		close(jobsGraph)
		wgGraph.Wait()
		close(resultsGraph)

		fmt.Printf("Graph key to key: Done in %v\n", time.Now().Sub(ti))
	*/

	type tr struct {
		src  [2]int
		dest [2]int
	}

	type trRes struct {
		src  [2]int
		dest [2]int
		q    *queueNode
	}

	jobs := make(chan *queueUserNode, 1024)
	results := make(chan *queueUserNode, 1024)
	var minDist *queueUserNode
	muxMin := sync.Mutex{}

	wg := sync.WaitGroup{}
	worker := func(r *Robot, jobs chan *queueUserNode, results chan *queueUserNode) {
		for curr := range jobs {
			// var minDistLocal *queueNode
			for _, kxy := range keysR {
				// Pass already taken key
				fnd := 0
				for _, kCurrXY := range curr.keys {
					if kCurrXY == kxy {
						fnd++
						break
					}
				}
				if fnd != 0 {
					continue
				}
				d := BFSDoor(r, curr.xy, kxy, curr.keys, &keys, &doors)

				if d != nil {
					newKeys := [][2]int{}
					for _, v := range curr.keys {
						newKeys = append(newKeys, v)
					}
					newKeys = append(newKeys, d.xy)
					adjcell := &queueUserNode{xy: d.xy, dist: curr.dist + d.dist, keys: newKeys}
					wg.Add(1)
					results <- adjcell
				}
				// if d != nil {
				// 	if minDistLocal == nil {
				// 		minDistLocal = d
				// 	} else if minDistLocal.dist > d.dist {
				// 		minDistLocal = d
				// 	}
				// }
			}

			// if len(curr.keys) == len(keys) {
			// 	muxMin.Lock()
			// 	if minDist == nil {
			// 		minDist = curr
			// 		fmt.Println(minDist)
			// 		fmt.Println(minDist.dist)
			// 	} else if curr.dist < minDist.dist {
			// 		minDist = curr
			// 		fmt.Println(minDist)
			// 		fmt.Println(minDist.dist)
			// 	}
			// 	muxMin.Unlock()
			// 	wg.Done()
			// 	return
			// }

			// if minDistLocal != nil {
			// 	newKeys := [][2]int{}
			// 	for _, v := range curr.keys {
			// 		newKeys = append(newKeys, v)
			// 	}
			// 	newKeys = append(newKeys, minDistLocal.xy)
			// 	adjcell := &queueUserNode{xy: minDistLocal.xy, dist: curr.dist + minDistLocal.dist, keys: newKeys}
			// 	wg.Add(1)
			// 	results <- adjcell
			// }

			wg.Done()
		}
	}

	for w := 0; w < cpu; w++ {
		go worker(r, jobs, results)
	}

	go func() {
		for curr := range results {
			// fmt.Printf("Result: %v keys: %v dist: %v\n", curr.xy, curr.keys, curr.dist)
			if len(curr.keys) == len(keys) {
				muxMin.Lock()
				if minDist == nil {
					minDist = curr
					fmt.Println(minDist)
					fmt.Println(minDist.dist)
				} else if curr.dist < minDist.dist {
					minDist = curr
					fmt.Println(minDist)
					fmt.Println(minDist.dist)
				}
				muxMin.Unlock()
			} else {
				wg.Add(1)
				// jobs <- curr
				go func() {
					jobs <- curr
				}()
			}
			wg.Done()
		}
	}()

	wg.Add(1)
	jobs <- &queueUserNode{xy: src, dist: 0, keys: [][2]int{}}
	wg.Wait()
	close(jobs)
	close(results)

	/*
		// Create a queue for BFS
		// q := list.New()
		qCh := make(chan *queueUserNode, 33)

		// Distance of source cell is 0
		s := &queueUserNode{xy: src, dist: 0, keys: [][2]int{}}
		// q.PushBack(s) // Enqueue source cell
		qCh <- s

		for c := 0; c < cpu; c++ {
			wg.Add(1)
			fmt.Printf("Routine %d: Start\n", c)
			go func() {
				var ok bool
				for curr := range qCh {
					// curr := q.Front().Value.(*queueUserNode)

					// If we have reached the destination cell,
					// we are done
					if len(curr.keys) == len(keys) {
						muxMin.Lock()
						if minDist == nil {
							minDist = curr
							fmt.Println(minDist)
							fmt.Println(minDist.dist)
						} else if curr.dist < minDist.dist {
							minDist = curr
							fmt.Println(minDist)
							fmt.Println(minDist.dist)
						}
						muxMin.Unlock()
						break
					}

					// Otherwise dequeue the front cell in the queue
					// and enqueue its adjacent cells
					// q.Remove(q.Front())

					for _, kxy := range keysR {
						// Pass already taken key
						fnd := 0
						for _, kCurrXY := range curr.keys {
							if kCurrXY == kxy {
								fnd++
								break
							}
						}
						if fnd != 0 {
							continue
						}
						var d *queueNode
						if d, ok = graphsDoorKey[curr.xy][kxy]; !ok {
							d = BFS(r, curr.xy, kxy)
						}
						fnd = 0
						for do := range d.doors {
							for _, kCurrXY := range curr.keys {
								if doors[do] == keys[kCurrXY] {
									fnd++
									break
								}
							}
						}
						if fnd == len(d.doors) {
							newKeys := [][2]int{}
							for _, v := range curr.keys {
								newKeys = append(newKeys, v)
							}
							newKeys = append(newKeys, kxy)
							adjcell := &queueUserNode{xy: kxy, dist: curr.dist + d.dist, keys: newKeys}
							qCh <- adjcell
							// q.PushBack(adjcell)
						}
					}
				}
				wg.Done()
			}()
			fmt.Printf("Routine %d: Done\n", c)
		}

		wg.Wait()
		close(qCh)
		fmt.Println(minDist)
		fmt.Println(minDist.dist)
	*/

	return

	/*
		doorsGet := map[rune]rune{}
		keysGet := map[rune]rune{}
		canOpenDoor := func() bool {
			x, y := r.NextForward()
			v := r.GetXY(x, y)
			if v == CASE_DOOR {
				_, ok := keysGet[doors[[2]int{x, y}]]
				return ok
			}
			return true
		}

		t := 0
		r.Direction = UP
		r.MapLength[[2]int{r.X, r.Y}] = 0
		mlength := map[rune]map[[2]int]int{}
		// rdlength := map[rune]int{}
		for _, k := range keys {
			mlength[k] = map[[2]int]int{}
		}

		fmt.Printf("Keys: %v\n", keysGet)
		fmt.Printf("X/Y: %d/%d\n", r.X, r.Y)
		r.DrawMap(true, true)
		for len(keysGet) != len(keys) {
			for r.GetXY(r.X, r.Y) != CASE_KEY && r.GetXY(r.X, r.Y) != CASE_DOOR {
				r.TurnLeft()
				v := r.GetXY(r.NextForward())
				if v == CASE_WALL || !canOpenDoor() {
					r.TurnRight()
				}
				v = r.GetXY(r.NextForward())
				if v == CASE_WALL || !canOpenDoor() {
					r.TurnRight()
				}
				v = r.GetXY(r.NextForward())
				if v == CASE_WALL || !canOpenDoor() {
					r.TurnRight()
				}
				r.MoveForward()
				if _, ok := r.MapLength[[2]int{r.X, r.Y}]; !ok {
					r.RecordLength()
				}
				time.Sleep(time.Millisecond * 50)
				fmt.Print("\033[H\033[2J")
				fmt.Printf("Keys: %v\n", keysGet)
				fmt.Printf("X/Y: %d/%d\n", r.X, r.Y)
				r.DrawMap(true, false)
			}

			if r.GetXY(r.X, r.Y) == CASE_KEY {
				k := keys[[2]int{r.X, r.Y}]
				keysGet[k] = k
				fmt.Printf("Key found: %c !\n", k)
				t += r.MapLength[[2]int{r.X, r.Y}]
				r.MapLength = map[[2]int]int{}
				r.MapLength[[2]int{r.X, r.Y}] = 0
				r.SetXY(r.X, r.Y, CASE_EMPTY)
			}

			if r.GetXY(r.X, r.Y) == CASE_DOOR {
				d := doors[[2]int{r.X, r.Y}]
				doorsGet[d] = d
				fmt.Printf("Door found: %c !\n", d)
				t += r.MapLength[[2]int{r.X, r.Y}]
				r.MapLength = map[[2]int]int{}
				r.MapLength[[2]int{r.X, r.Y}] = 0
				r.SetXY(r.X, r.Y, CASE_EMPTY)
			}

			fmt.Print("\033[H\033[2J")
			fmt.Printf("Keys: %v\n", keysGet)
			fmt.Printf("Doors: %v\n", doorsGet)
			fmt.Printf("X/Y: %d/%d\n", r.X, r.Y)
			r.DrawMap(true, false)
		}

		fmt.Println()
		fmt.Println()
		fmt.Printf("Total: %d\n", t)
	*/

}

func main1() {
	// processCodeAmp([][]int{parseLine("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")})
	// processCodeAmp([][]int{parseLine("1102,34915192,34915192,7,4,7,99,0")})
	// processCodeAmp([][]int{parseLine("104,1125899906842624,99")})
	// process1("test.txt")
	// process1("test2.txt")
	// process1("test3.txt")
	// process1("test4.txt")
	// process1("test5.txt")
	process1("data.txt")
	// processCodeChainAmp([][]int{parseLine("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")})
	// processCodeChainAmp([][]int{parseLine("1102,34915192,34915192,7,4,7,99,0")})
	// processCodeChainAmp([][]int{parseLine("104,1125899906842624,99")})
	// processCodeChainAmp(parseFile("data.txt"))
}
