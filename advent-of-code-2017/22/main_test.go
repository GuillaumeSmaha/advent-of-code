package p22

import (
	"testing"
)

func TestInfect(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@t.txt",
			Result: 5587,
		},
		{
			Input:  "@a.txt",
			Result: 5280,
		},
	}

	for i, test := range tests {
		n := Infect(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestInfect2(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "@t.txt",
			Result: 2511944,
		},
		{
			Input:  "@a.txt",
			Result: 2512261,
		},
	}

	for i, test := range tests {
		n := Infect2(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}
