package p19

import (
	"testing"
)

func TestWalk(t *testing.T) {
	tests := []struct {
		Input  string
		Result string
		Steps  int
	}{
		{
			Input:  "@t.txt",
			Result: "ABCDEF",
			Steps:  38,
		},
		{
			Input:  "@a.txt",
			Result: "VTWBPYAQFU",
			Steps:  17358,
		},
	}

	for i, test := range tests {
		s, n := Walk(test.Input)
		if s != test.Result || n != test.Steps {
			t.Fatalf("a.%d: %s '%s' %d (should be %s %d)", i, test.Input, s, n, test.Result, test.Steps)
		}
	}
}
