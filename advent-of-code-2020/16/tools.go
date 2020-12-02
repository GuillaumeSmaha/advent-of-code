package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseLineInt2D(data string) []int {
	ds := strings.Split(data, ",")
	r := []int{}
	for _, d := range ds {
		i, _ := strconv.Atoi(d)
		r = append(r, i)
	}
	return r
}

func parseFileInt2D(filename string) [][]int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	codes := [][]int{}
	for fscanner.Scan() {
		codes = append(codes, parseLineInt2D(fscanner.Text()))
	}
	return codes
}

func parseFileText2D(filename string, s string) [][]string {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	codes := [][]string{}
	for fscanner.Scan() {
		codes = append(codes, strings.Split(fscanner.Text(), ","))
	}
	return codes
}

func parseFileText2DFirstChar(filename string) [][]string {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	codes := [][]string{}
	for fscanner.Scan() {
		s := fscanner.Text()
		codes = append(codes, []string{string(s[0]), s[1:]})
	}
	return codes
}

func getBoolStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func seq(n int, starts ...int) []int {
	s := 0
	if len(starts) > 0 {
		s = starts[0]
	}

	r := make([]int, n-s)
	for i := 0; i < n-s; i++ {
		r[i] = i + s
	}
	return r
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
