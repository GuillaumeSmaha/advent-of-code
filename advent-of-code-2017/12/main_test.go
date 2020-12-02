package p12

import (
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "@t.txt",
			Result: 6,
		},
		{
			File:   "@a.txt",
			Result: 175,
		},
	}

	for i, test := range tests {
		n := Count(test.File)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}

func TestTotal(t *testing.T) {
	tests := []struct {
		File   string
		Result int
	}{
		{
			File:   "@t.txt",
			Result: 2,
		},
		{
			File:   "@a.txt",
			Result: 213,
		},
	}

	for i, test := range tests {
		n := Total(test.File)
		if n != test.Result {
			t.Fatalf("b.%d: %s '%d' (should be %d)", i, test.File, n, test.Result)
		}
	}
}
