package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule [2]int

func (r *Rule) Validate(v int) bool {
	return v >= r[0] && v <= r[1]
}

type Rules []Rule

func (r *Rules) Validate(v int) bool {
	for _, ru := range *r {
		if ru.Validate(v) {
			return true
		}
	}
	return false
}

type Ticket []int

type Data struct {
	Rules             map[string]Rules
	Ticket            Ticket
	NearbyTicket      []Ticket
	NearbyTicketValid []Ticket
}

func (d *Data) ValidateRule(v int) bool {
	for _, r := range d.Rules {
		if r.Validate(v) {
			return true
		}
	}
	return false
}

func (d *Data) FirstInvalidTicketValue(t Ticket) (int, bool) {
	for _, tv := range t {
		if !d.ValidateRule(tv) {
			return tv, false
		}
	}
	return 0, true
}

func (d *Data) CheckCompleteInvalidTicket(t Ticket) bool {
	c := 0
	for _, tv := range t {
		if !d.ValidateRule(tv) {
			c++
		}
	}
	return len(t) == c
}

func (d *Data) CheckCompleteValidTicket(t Ticket) bool {
	c := 0
	for _, tv := range t {
		if d.ValidateRule(tv) {
			c++
		}
	}
	return len(t) == c
}

func (d *Data) FindFirstRuleName(v int) string {
	for n, r := range d.Rules {
		if r.Validate(v) {
			return n
		}
	}
	return ""
}

func (d *Data) FindRuleNames(v int) map[string]int {
	m := map[string]int{}
	for n, r := range d.Rules {
		if r.Validate(v) {
			if _, ok := m[n]; !ok {
				m[n] = 0
			}
			m[n]++
		}
	}
	return m
}

func (d *Data) FindTicketRuleName(t Ticket) string {
	m := map[string]int{}
	for _, v := range t {
		for n, r := range d.Rules {
			if r.Validate(v) {
				if _, ok := m[n]; !ok {
					m[n] = 0
				}
				m[n]++
			}
		}
	}
	max := 0
	maxN := ""
	for n, c := range m {
		if c > max {
			maxN = n
			max = c
		}
	}
	return maxN
}

func (d *Data) FindNamesForIdx(idx int) map[string]int {
	m := map[string]int{}
	tt := d.FindValidTickets()
	for _, v := range tt {
		mm := map[string]int{}
		for n, r := range d.Rules {
			if r.Validate(v[idx]) {
				if _, ok := m[n]; !ok {
					m[n] = 0
				}
				m[n]++
				if _, ok := mm[n]; !ok {
					mm[n] = 0
				}
				mm[n]++
			}
		}
	}
	for n, c := range m {
		if c != len(tt) {
			delete(m, n)
		}
	}
	return m
}

func (d *Data) FindCompleteInvalidIndex() int {
	for i, t := range d.NearbyTicket {
		if d.CheckCompleteInvalidTicket(t) {
			return i
		}

	}
	return -1
}

func (d *Data) FindCompleteValidIndex() int {
	for i, t := range d.NearbyTicket {
		if d.CheckCompleteValidTicket(t) {
			return i
		}

	}
	return -1
}

func (d *Data) FindValidTickets() []Ticket {
	d.NearbyTicketValid = []Ticket{}
	for _, t := range d.NearbyTicket {
		if _, b := d.FirstInvalidTicketValue(t); b {
			d.NearbyTicketValid = append(d.NearbyTicketValid, t)
		}
	}
	return d.NearbyTicketValid
}

func parseFile(filename string) *Data {
	data := &Data{
		Rules:        map[string]Rules{},
		Ticket:       Ticket{},
		NearbyTicket: []Ticket{},
	}
	prev := ""
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	step := 0
	for fscanner.Scan() {
		s := fscanner.Text()
		// fmt.Println(s)
		if s == "" || s == "\n" {
			step++
			continue
		}
		p := strings.Split(s, ":")
		value := ""
		if len(p) == 1 {
			value = p[0]
		} else if len(p[1]) == 0 {
			prev = p[0]
			continue
		} else {
			prev = p[0]
			value = p[1]
		}

		if step == 0 {
			for _, v := range strings.Split(strings.TrimSpace(value), " or ") {
				pv := strings.Split(v, "-")
				start, _ := strconv.Atoi(pv[0])
				end, _ := strconv.Atoi(pv[1])
				if _, ok := data.Rules[prev]; !ok {
					data.Rules[prev] = Rules{}
				}
				data.Rules[prev] = append(data.Rules[prev], Rule{start, end})
			}
		} else if step == 1 {
			for _, vs := range strings.Split(value, ",") {
				i, _ := strconv.Atoi(vs)
				data.Ticket = append(data.Ticket, i)
			}
		} else {
			v := Ticket{}
			for _, vs := range strings.Split(value, ",") {
				i, _ := strconv.Atoi(vs)
				v = append(v, i)
			}
			data.NearbyTicket = append(data.NearbyTicket, v)
		}
	}

	return data
}

func processMain1(data *Data) {
	fmt.Println("----------")
	fmt.Printf("%#v\n", data)

	s := 0
	for _, t := range data.NearbyTicket {
		v, b := data.FirstInvalidTicketValue(t)
		if b {
			fmt.Printf("%#v: %d\n", t, v)
			s += v
		}
	}
	fmt.Printf("Toal: %d\n", s)
}

func main1() {
	processMain1(parseFile("list.test.txt"))
	processMain1(parseFile("list.txt"))
}
