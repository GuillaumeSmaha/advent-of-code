package p18

import (
	"testing"
)

func TestDuet(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@t1.txt",
			Result: 4,
		},
		{
			Input:  "@a.txt",
			Result: 1187,
		},
	}

	for i, test := range tests {
		n := Duet(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestTwo(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@t2.txt",
			Result: 3,
		},
		{
			Input:  "@a.txt",
			Result: 5969,
		},
	}

	for i, test := range tests {
		n := Two(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}
