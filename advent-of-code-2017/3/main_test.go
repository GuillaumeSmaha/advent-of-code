package p3

import "testing"

func TestSteps(t *testing.T) {
	tests := []struct {
		N      int
		Result int
	}{
		{
			N:      1,
			Result: 0,
		},
		{
			N:      12,
			Result: 3,
		},
		{
			N:      23,
			Result: 2,
		},
		{
			N:      1024,
			Result: 31,
		},
		{
			N:      347991,
			Result: 480,
		},
	}

	for i, test := range tests {
		n := Steps(test.N)
		if n != test.Result {
			t.Fatalf("a.%d: %d: %d (should be %d)", i, test.N, n, test.Result)
		}
	}
}

func TestSums(t *testing.T) {
	tests := []struct {
		N      int
		Result int
	}{
		{
			N:      1,
			Result: 1,
		},
		{
			N:      2,
			Result: 1,
		},
		{
			N:      9,
			Result: 25,
		},
		{
			N:      11,
			Result: 54,
		},
		{
			N:      15,
			Result: 133,
		},
		{
			N:      20,
			Result: 351,
		},
		{
			N:      23,
			Result: 806,
		},
		// puzzle input: 347991
		{
			N:      62,
			Result: 330785,
		},
		{
			N:      63,
			Result: 349975,
		},
	}

	for i, test := range tests {
		n := Sums(test.N)
		if n != test.Result {
			t.Fatalf("b.%d: %d: %d (should be %d)", i, test.N, n, test.Result)
		}
	}
}
