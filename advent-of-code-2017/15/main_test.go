package p15

import (
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		A      int
		B      int
		Result int
	}{
		{
			A:      65,
			B:      8921,
			Result: 588,
		},
		{
			A:      783,
			B:      325,
			Result: 650,
		},
	}

	for i, test := range tests {
		n := Count(test.A, test.B)
		if n != test.Result {
			t.Fatalf("a.%d: %d,%d '%d' (should be %d)", i, test.A, test.B, n, test.Result)
		}
	}
}

func TestCountSlow(t *testing.T) {
	tests := []struct {
		A      int
		B      int
		Result int
	}{
		{
			A:      65,
			B:      8921,
			Result: 309,
		},
		{
			A:      783,
			B:      325,
			Result: 336,
		},
	}

	for i, test := range tests {
		n := CountSlow(test.A, test.B)
		if n != test.Result {
			t.Fatalf("b.%d: %d,%d '%d' (should be %d)", i, test.A, test.B, n, test.Result)
		}
	}
}
