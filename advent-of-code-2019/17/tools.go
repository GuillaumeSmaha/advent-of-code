package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseLine(data string) []int {
	ds := strings.Split(data, ",")
	r := []int{}
	for _, d := range ds {
		i, _ := strconv.Atoi(d)
		r = append(r, i)
	}
	return r
}

func parseFile(filename string) [][]int {
	file, _ := os.Open(filename)
	fscanner := bufio.NewScanner(file)
	codes := [][]int{}
	for fscanner.Scan() {
		codes = append(codes, parseLine(fscanner.Text()))
	}
	return codes
}

func getBoolStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
