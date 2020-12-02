package p5

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		Table  []int
		Result int
	}{
		{
			Table:  []int{0, 3, 0, 1, -3},
			Result: 5,
		},
	}

	for i, test := range tests {
		n := Run(test.Table)
		if n != test.Result {
			t.Fatalf("a.%d: %d (should be %d)", i, n, test.Result)
		}
	}
}

func TestRun2(t *testing.T) {
	tests := []struct {
		Table  []int
		Result int
	}{
		{
			Table:  []int{0, 3, 0, 1, -3},
			Result: 10,
		},
	}

	for i, test := range tests {
		n := Run2(test.Table)
		if n != test.Result {
			t.Fatalf("b.%d: %d (should be %d)", i, n, test.Result)
		}
	}
}
