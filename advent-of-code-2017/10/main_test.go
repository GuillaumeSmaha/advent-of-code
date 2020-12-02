package p10

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	tests := []struct {
		List   string
		Length string
		Result int
	}{
		{
			List:   "0,1,2,3,4",
			Length: "3,4,1,5",
			Result: 12,
		},
		{
			Length: "225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110",
			Result: 23874,
		},
	}

	for i, test := range tests {
		n := Multiply(test.List, test.Length)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.List, n, test.Result)
		}
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		Length string
		Result string
	}{
		{
			Length: "3,4,1,5",
			Result: "4a19451b02fb05416d73aea0ec8c00c0",
		},
		{
			Length: "225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110",
			Result: "e1a65bfb5a5ce396025fab5528c25a87",
		},
	}

	for i, test := range tests {
		s := Hash(test.Length)
		if s != test.Result {
			t.Fatalf("b.%d: %s '%s' (should be %s)", i, test.Length, s, test.Result)
		}
	}
}
