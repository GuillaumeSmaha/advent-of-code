package p2

import "testing"

func TestDiffSum(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@ta.txt",
			Result: 18,
		},
		{
			Input:  "@a.txt",
			Result: 43074,
		},
	}

	for i, test := range tests {
		n := DiffSum(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %s: %d (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestEvenlyDivisibleSum(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@tb.txt",
			Result: 9,
		},
		{
			Input:  "@a.txt",
			Result: 280,
		},
	}

	for i, test := range tests {
		n := EvenlyDivisibleSum(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %s: %d (should be %d)", i, test.Input, n, test.Result)
		}
	}
}
