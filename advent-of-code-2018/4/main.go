package p1

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-2018 4[a|b] 'input'")
	}

	switch args[0] {
	case "4a", "4":
		fmt.Println(Part1(args[1]))
	case "4b":
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

type Event struct {
	Date     time.Time
	EventStr string
}

type GuardEvent struct {
	Guard int
	Date  time.Time
	Awake bool
}

type GuardRange struct {
	Guard int
	From  time.Time
	To    time.Time
}

var Parser *regexp.Regexp = regexp.MustCompile(`\[(?P<year>\d+)-(?P<month>\d+)-(?P<day>\d+) (?P<hour>\d+):(?P<minute>\d+)\] (?P<event>.+)`)
var ParserGuard *regexp.Regexp = regexp.MustCompile(`Guard #(?P<id>\d+) begins shift`)

func parseString(s string) (time.Time, string) {
	match := Parser.FindStringSubmatch(s)
	year, _ := strconv.Atoi(match[1])
	month, _ := strconv.Atoi(match[2])
	day, _ := strconv.Atoi(match[3])
	hour, _ := strconv.Atoi(match[4])
	minute, _ := strconv.Atoi(match[5])
	event := match[6]
	date := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)
	return date, event
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func parseData(strs []string) []*GuardEvent {

	var events []*Event
	for _, s := range strs {
		date, eventStr := parseString(s)
		event := &Event{
			Date:     date,
			EventStr: eventStr,
		}
		events = append(events, event)
	}

	// fmt.Println("Events")
	// for _, e := range events {
	// 	fmt.Printf("%v: %v\n", e.Date, e.EventStr)
	// }

	sort.Slice(events, func(i, j int) bool {
		return events[j].Date.Sub(events[i].Date) > 0
	})

	// fmt.Println("SORT")
	// for _, e := range events {
	// 	fmt.Printf("%v: %v\n", e.Date, e.EventStr)
	// }

	var guardEvents []*GuardEvent
	lastGuardId := 0
	var guards map[int]bool = make(map[int]bool)
	lastDayMonth := fmt.Sprintf("%v-%v", events[0].Date.Day(), events[0].Date.Month())
	for _, e := range events {
		if e.Date.Hour() > 1 {
			e.Date = e.Date.Add(time.Duration(int64(time.Hour*24) - int64(time.Hour)*int64(e.Date.Hour()) - int64(time.Minute)*int64(e.Date.Minute())))
		}

		if fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month()) != lastDayMonth {
			for g, a := range guards {
				if a {
					date := e.Date.Add(time.Duration(int64(time.Hour*-24) - int64(time.Hour)*int64(e.Date.Hour()-1) - int64(time.Minute)*int64(e.Date.Minute())))
					event := &GuardEvent{
						Guard: g,
						Date:  date,
						Awake: false,
					}
					guardEvents = append(guardEvents, event)
				}
			}
			guards = make(map[int]bool)
			lastDayMonth = fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month())
		}

		if _, ok := guards[lastGuardId]; !ok {
			e.Date = e.Date.Add(time.Duration(-int64(time.Hour)*int64(e.Date.Hour()) - int64(time.Minute)*int64(e.Date.Minute())))
		}
		awake := true
		if e.EventStr == "falls asleep" {
			guards[lastGuardId] = false
			awake = false
		} else if e.EventStr == "wakes up" {
			guards[lastGuardId] = true
			awake = true
		} else {
			match := ParserGuard.FindStringSubmatch(e.EventStr)
			lastGuardId, _ = strconv.Atoi(match[1])
			if e.Date.Hour() < 23 {
				guards[lastGuardId] = true
			}
		}

		event := &GuardEvent{
			Guard: lastGuardId,
			Date:  e.Date,
			Awake: awake,
		}
		guardEvents = append(guardEvents, event)
	}

	e := events[len(events)-1]
	for g, a := range guards {
		if a {
			date := e.Date.Add(time.Duration(-int64(time.Hour)*int64(e.Date.Hour()-1) - int64(time.Minute)*int64(e.Date.Minute())))
			event := &GuardEvent{
				Guard: g,
				Date:  date,
				Awake: false,
			}
			guardEvents = append(guardEvents, event)
		}
	}

	return guardEvents
}

// func convertToRange(events []*GuardEvent) []*AwakeRange {
// 	var awakeRanges []*AwakeRange
// 	sort.Slice(events, func(i, j int) bool {
// 		if events[i].Guard < events[j].Guard {
// 			return true
// 		}
// 		if events[i].Guard > events[j].Guard {
// 			return false
// 		}
// 		return events[j].Date.Sub(events[i].Date) > 0
// 	})
// 	fmt.Println("SORT")
// 	for _, e := range events {
// 		fmt.Printf("%v: %v: %v\n", e.Guard, e.Date, e.Awake)
// 	}

// 	for _, e := range events {
// 	event := &GuardEvent{
// 		Guard: e.Guard,
// 		Date:  e.Date,
// 		Awake: awake,
// 	}
// 	events = append(events, event)
// 	}

// 	return awakeRanges
// }

func getMinuteMostAwake(events []*GuardEvent, guard int) (int, int) {
	var minuteCnt map[int]int = make(map[int]int)
	for i, e := range events {
		if i == 0 {
			continue
		}
		if e.Guard != guard {
			continue
		}

		if fmt.Sprintf("%v-%v", events[i-1].Date.Day(), events[i-1].Date.Month()) == fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month()) {
			if events[i-1].Awake {
				for m := events[i-1].Date.Minute(); m < e.Date.Minute(); m++ {
					minuteCnt[m]++
				}
			}
		}
	}

	maxC := 0
	maxM := 0
	for m, c := range minuteCnt {
		if c > maxC {
			maxC = c
			maxM = m
		}
	}

	return maxM, maxC
}

func getMinuteMostAsleep(events []*GuardEvent, guard int) (int, int) {
	var minuteCnt map[int]int = make(map[int]int)
	for i, e := range events {
		if i == 0 {
			continue
		}
		if e.Guard != guard {
			continue
		}

		if fmt.Sprintf("%v-%v", events[i-1].Date.Day(), events[i-1].Date.Month()) == fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month()) {
			if !events[i-1].Awake {
				for m := events[i-1].Date.Minute(); m < e.Date.Minute(); m++ {
					minuteCnt[m]++
				}
			}
		}
	}

	maxC := 0
	maxM := 0
	for m, c := range minuteCnt {
		if c > maxC {
			maxC = c
			maxM = m
		}
	}

	return maxM, maxC
}

func printTable(events []*GuardEvent, guard int) {

	sort.Slice(events, func(i, j int) bool {
		if events[i].Guard < events[j].Guard {
			return true
		}
		if events[i].Guard > events[j].Guard {
			return false
		}
		return events[j].Date.Sub(events[i].Date) > 0
	})

	fmt.Println()
	fmt.Printf("Guard: %v\n", guard)
	fmt.Printf("Date\tMinute\n")
	fmt.Printf("\t")
	for i := 0; i < 6; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%v", i)
		}
	}
	fmt.Println()
	fmt.Printf("\t")
	for i := 0; i < 6; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%v", j)
		}
	}

	cnt := 0
	for i, e := range events {
		if e.Guard != guard {
			continue
		}
		if i == 0 {
			fmt.Println()
			fmt.Printf("%02d-%02d\t", e.Date.Month(), e.Date.Day())
			continue
		}
		if fmt.Sprintf("%v-%v", events[i-1].Date.Day(), events[i-1].Date.Month()) == fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month()) {
			for m := events[i-1].Date.Minute(); m < e.Date.Minute(); m++ {
				if events[i-1].Awake {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
				cnt++
			}
		} else {
			for m := cnt; m < 60; m++ {
				fmt.Print(".")
			}
			cnt = 0
			fmt.Println()
			fmt.Printf("%02d-%02d\t", e.Date.Month(), e.Date.Day())
		}
	}
	for m := cnt; m < 60; m++ {
		fmt.Print(".")
	}
	fmt.Println()
}

func Part1(s string) int {
	data := load(s)
	events := parseData(data)
	sort.Slice(events, func(i, j int) bool {
		return events[j].Date.Sub(events[i].Date) > 0
	})

	// fmt.Println("SORT")
	// for _, e := range events {
	// 	fmt.Printf("%v: %v: %v\n", e.Guard, e.Date, e.Awake)
	// }

	sort.Slice(events, func(i, j int) bool {
		if events[i].Guard < events[j].Guard {
			return true
		}
		if events[i].Guard > events[j].Guard {
			return false
		}
		return events[j].Date.Sub(events[i].Date) > 0
	})

	fmt.Println("SORT")
	for _, e := range events {
		fmt.Printf("%v: %v: %v\n", e.Guard, e.Date, e.Awake)
	}

	var guardsDays map[int]int = make(map[int]int)
	for i, e := range events {
		if i == 0 {
			guardsDays[e.Guard]++
			continue
		}
		if fmt.Sprintf("%v-%v", events[i-1].Date.Day(), events[i-1].Date.Month()) != fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month()) {
			guardsDays[e.Guard]++
		}
	}
	var guardsAwake map[int]time.Duration = make(map[int]time.Duration)
	var guardsAwakeMax map[int]time.Duration = make(map[int]time.Duration)
	var guardsAsleep map[int]time.Duration = make(map[int]time.Duration)
	var guardsAsleepMax map[int]time.Duration = make(map[int]time.Duration)

	for i, e := range events {
		if i == 0 {
			continue
		}
		if fmt.Sprintf("%v-%v", events[i-1].Date.Day(), events[i-1].Date.Month()) == fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month()) {
			d := e.Date.Sub(events[i-1].Date)
			if events[i-1].Awake {
				guardsAwake[e.Guard] += d
				if guardsAwakeMax[e.Guard] < d {
					guardsAwakeMax[e.Guard] = d
				}
			} else {
				guardsAsleep[e.Guard] += d
				if guardsAsleepMax[e.Guard] < d {
					guardsAsleepMax[e.Guard] = d
				}
			}
		}
	}

	fmt.Println("Guards nb days:")
	for g, d := range guardsDays {
		fmt.Printf("\t %v: %v\n", g, d)
	}

	fmt.Println("Guards awake:")
	for g, d := range guardsAwake {
		fmt.Printf("\t %v: %v (max: %v)\n", g, d, guardsAwakeMax[g])
	}

	fmt.Println("Guards asleep:")
	for g, d := range guardsAsleep {
		fmt.Printf("\t %v: %v (max: %v)\n", g, d, guardsAsleepMax[g])
	}

	mostAwake := -1
	tmp := time.Duration(0)
	for g, d := range guardsAwake {
		if tmp < d {
			mostAwake = g
			tmp = d
		}
	}
	fmt.Printf("Guards most awake: %v\n", mostAwake)
	fmt.Printf("Guards most awake during: %v\n", guardsAwakeMax[mostAwake])

	mostAsleep := -1
	tmp = time.Duration(0)
	for g, d := range guardsAsleep {
		if tmp < d {
			mostAsleep = g
			tmp = d
		}
	}
	fmt.Printf("Guards most asleep: %v\n", mostAsleep)
	fmt.Printf("Guards most asleep during: %v\n", guardsAsleepMax[mostAsleep])

	mostAwakeMax := -1
	tmp = time.Duration(0)
	for g, d := range guardsAwakeMax {
		if tmp < d {
			mostAwakeMax = g
			tmp = d
		}
	}
	fmt.Printf("Guards most awake max: %v\n", mostAwakeMax)
	fmt.Printf("Guards most awake max during: %v\n", guardsAwakeMax[mostAwakeMax])

	mostAsleepMax := -1
	tmp = time.Duration(0)
	for g, d := range guardsAsleepMax {
		if tmp < d {
			mostAsleepMax = g
			tmp = d
		}
	}
	fmt.Printf("Guards most asleep max: %v\n", mostAsleepMax)
	fmt.Printf("Guards most asleep max during: %v\n", guardsAsleepMax[mostAsleepMax])

	minuteMostAwake, _ := getMinuteMostAwake(events, mostAwake)
	fmt.Printf("Minute when guards most awake: %v\n", minuteMostAwake)
	minuteMostAsleep, _ := getMinuteMostAsleep(events, mostAsleep)
	fmt.Printf("Minute when guards most asleep: %v\n", minuteMostAsleep)

	printTable(events, mostAsleep)

	maxC := 0
	guard := 0
	minute := 0
	for g, _ := range guardsDays {
		m, c := getMinuteMostAsleep(events, g)
		fmt.Printf("%v: min: %v for %v time (check: %v)\n", g, m, c, g*m)
		if c > maxC {
			guard = g
			maxC = c
			minute = m
		}
	}
	fmt.Printf("Guard: %v\n", guard)
	fmt.Printf("Minute: %v\n", minute)
	fmt.Printf("Count: %v\n", maxC)

	m, c := getMinuteMostAsleep(events, mostAsleep)
	fmt.Printf("mostAsleep: %v\n", mostAsleep)
	fmt.Printf("Minute: %v\n", m)
	fmt.Printf("Count: %v\n", c)

	return mostAsleep * m
}

func diff(s string) int {
	return 0
}

func Part2(s string) int {
	data := load(s)
	events := parseData(data)
	sort.Slice(events, func(i, j int) bool {
		return events[j].Date.Sub(events[i].Date) > 0
	})

	// fmt.Println("SORT")
	// for _, e := range events {
	// 	fmt.Printf("%v: %v: %v\n", e.Guard, e.Date, e.Awake)
	// }

	sort.Slice(events, func(i, j int) bool {
		if events[i].Guard < events[j].Guard {
			return true
		}
		if events[i].Guard > events[j].Guard {
			return false
		}
		return events[j].Date.Sub(events[i].Date) > 0
	})

	fmt.Println("SORT")
	for _, e := range events {
		fmt.Printf("%v: %v: %v\n", e.Guard, e.Date, e.Awake)
	}

	var guardsDays map[int]int = make(map[int]int)
	for i, e := range events {
		if i == 0 {
			guardsDays[e.Guard]++
			continue
		}
		if fmt.Sprintf("%v-%v", events[i-1].Date.Day(), events[i-1].Date.Month()) != fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month()) {
			guardsDays[e.Guard]++
		}
	}
	var guardsAwake map[int]time.Duration = make(map[int]time.Duration)
	var guardsAwakeMax map[int]time.Duration = make(map[int]time.Duration)
	var guardsAsleep map[int]time.Duration = make(map[int]time.Duration)
	var guardsAsleepMax map[int]time.Duration = make(map[int]time.Duration)

	for i, e := range events {
		if i == 0 {
			continue
		}
		if fmt.Sprintf("%v-%v", events[i-1].Date.Day(), events[i-1].Date.Month()) == fmt.Sprintf("%v-%v", e.Date.Day(), e.Date.Month()) {
			d := e.Date.Sub(events[i-1].Date)
			if events[i-1].Awake {
				guardsAwake[e.Guard] += d
				if guardsAwakeMax[e.Guard] < d {
					guardsAwakeMax[e.Guard] = d
				}
			} else {
				guardsAsleep[e.Guard] += d
				if guardsAsleepMax[e.Guard] < d {
					guardsAsleepMax[e.Guard] = d
				}
			}
		}
	}

	fmt.Println("Guards nb days:")
	for g, d := range guardsDays {
		fmt.Printf("\t %v: %v\n", g, d)
	}

	fmt.Println("Guards awake:")
	for g, d := range guardsAwake {
		fmt.Printf("\t %v: %v (max: %v)\n", g, d, guardsAwakeMax[g])
	}

	fmt.Println("Guards asleep:")
	for g, d := range guardsAsleep {
		fmt.Printf("\t %v: %v (max: %v)\n", g, d, guardsAsleepMax[g])
	}

	mostAwake := -1
	tmp := time.Duration(0)
	for g, d := range guardsAwake {
		if tmp < d {
			mostAwake = g
			tmp = d
		}
	}
	fmt.Printf("Guards most awake: %v\n", mostAwake)
	fmt.Printf("Guards most awake during: %v\n", guardsAwakeMax[mostAwake])

	mostAsleep := -1
	tmp = time.Duration(0)
	for g, d := range guardsAsleep {
		if tmp < d {
			mostAsleep = g
			tmp = d
		}
	}
	fmt.Printf("Guards most asleep: %v\n", mostAsleep)
	fmt.Printf("Guards most asleep during: %v\n", guardsAsleepMax[mostAsleep])

	mostAwakeMax := -1
	tmp = time.Duration(0)
	for g, d := range guardsAwakeMax {
		if tmp < d {
			mostAwakeMax = g
			tmp = d
		}
	}
	fmt.Printf("Guards most awake max: %v\n", mostAwakeMax)
	fmt.Printf("Guards most awake max during: %v\n", guardsAwakeMax[mostAwakeMax])

	mostAsleepMax := -1
	tmp = time.Duration(0)
	for g, d := range guardsAsleepMax {
		if tmp < d {
			mostAsleepMax = g
			tmp = d
		}
	}
	fmt.Printf("Guards most asleep max: %v\n", mostAsleepMax)
	fmt.Printf("Guards most asleep max during: %v\n", guardsAsleepMax[mostAsleepMax])

	minuteMostAwake, _ := getMinuteMostAwake(events, mostAwake)
	fmt.Printf("Minute when guards most awake: %v\n", minuteMostAwake)
	minuteMostAsleep, _ := getMinuteMostAsleep(events, mostAsleep)
	fmt.Printf("Minute when guards most asleep: %v\n", minuteMostAsleep)

	printTable(events, mostAsleep)

	maxC := 0
	guard := 0
	minute := 0
	for g, _ := range guardsDays {
		m, c := getMinuteMostAsleep(events, g)
		fmt.Printf("%v: min: %v for %v time\n", g, m, c)
		if c > maxC {
			guard = g
			maxC = c
			minute = m
		}
	}

	fmt.Printf("Guard: %v\n", guard)
	fmt.Printf("Minute: %v\n", minute)
	fmt.Printf("Count: %v\n", maxC)

	return guard * minute
}
