package p17

import (
	"testing"
)

func TestSpin(t *testing.T) {
	tests := []struct {
		Input  int
		Result int
	}{
		{
			Input:  3,
			Result: 638,
		},
		{
			Input:  370,
			Result: 1244,
		},
	}

	for i, test := range tests {
		n := Spin(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %d '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestSpinAngry(t *testing.T) {
	tests := []struct {
		Input  int
		Result int
	}{
		{
			Input:  370,
			Result: 11162912,
		},
	}

	for i, test := range tests {
		n := SpinAngry(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %d '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}
