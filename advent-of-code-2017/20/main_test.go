package p20

import (
	"testing"
)

func TestClosest(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@a.txt",
			Result: 457,
		},
	}

	for i, test := range tests {
		n := Closest(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestCollide(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@a.txt",
			Result: 448,
		},
	}

	for i, test := range tests {
		n := Collide(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}
