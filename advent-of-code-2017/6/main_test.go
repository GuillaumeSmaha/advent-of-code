package p6

import (
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		Banks  []int
		Result int
	}{
		{
			Banks:  []int{0, 2, 7, 0},
			Result: 5,
		},
		{
			Banks:  []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11},
			Result: 4074,
		},
	}

	for i, test := range tests {
		n := Run(test.Banks)
		if n != test.Result {
			t.Fatalf("a.%d: %d (should be %d)", i, n, test.Result)
		}
	}
}

func TestLoop(t *testing.T) {
	tests := []struct {
		Banks  []int
		Result int
	}{
		{
			Banks:  []int{0, 2, 7, 0},
			Result: 4,
		},
		{
			Banks:  []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11},
			Result: 2793,
		},
	}

	for i, test := range tests {
		n := Loop(test.Banks)
		if n != test.Result {
			t.Fatalf("b.%d: %d (should be %d)", i, n, test.Result)
		}
	}
}
