package p7

import (
	"testing"
)

func TestRoot(t *testing.T) {
	tests := []struct {
		File   string
		Result string
	}{
		{
			File:   "@t.txt",
			Result: "tknk",
		},
		{
			File:   "@a.txt",
			Result: "gmcrj",
		},
	}

	for i, test := range tests {
		s := Root(test.File)
		if s != test.Result {
			t.Fatalf("a.%d: '%s' (should be %s)", i, s, test.Result)
		}
	}
}

func TestWrong(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "@t.txt",
			Result: 60,
		},
		{
			File:   "@a.txt",
			Result: 391,
		},
	}

	for i, test := range tests {
		n := FirstWrong(test.File)
		if n != test.Result {
			t.Fatalf("b.%d: '%d' (should be %d)", i, n, test.Result)
		}
	}
}
