package main

import "os"

func main() {

	e := "1"
	if len(os.Args) > 1 {
		e = os.Args[1]
	}
	switch e {
	case "2":
		// main2()
	default:
		main1()
	}
}
