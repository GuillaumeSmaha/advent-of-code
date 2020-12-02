package p24

import (
	"testing"
)

func TestStrongest(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@t.txt",
			Result: 31,
		},
		{
			Input:  "@a.txt",
			Result: 1511,
		},
	}

	for i, test := range tests {
		n := Strongest(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestLongest(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@t.txt",
			Result: 19,
		},
		{
			Input:  "@a.txt",
			Result: 1471,
		},
	}

	for i, test := range tests {
		n := Longest(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}
