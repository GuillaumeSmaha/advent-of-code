package p7

import (
	"testing"
)

func TestMaximum(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "@t.txt",
			Result: 1,
		},
		{
			File:   "@a.txt",
			Result: 8022,
		},
	}

	for i, test := range tests {
		n := Maximum(test.File)
		if n != test.Result {
			t.Fatalf("a.%d: '%d' (should be %d)", i, n, test.Result)
		}
	}
}

func TestHighest(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "@t.txt",
			Result: 10,
		},
		{
			File:   "@a.txt",
			Result: 9819,
		},
	}

	for i, test := range tests {
		n := Highest(test.File)
		if n != test.Result {
			t.Fatalf("b.%d: '%d' (should be %d)", i, n, test.Result)
		}
	}
}
