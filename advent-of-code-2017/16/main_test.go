package p16

import (
	"testing"
)

func TestDance(t *testing.T) {
	tests := []struct {
		Input  string
		Key    string
		Result string
	}{
		{
			Input:  "@t.txt",
			Key:    "abcde",
			Result: "baedc",
		},
		{
			Input:  "@a.txt",
			Key:    "abcdefghijklmnop",
			Result: "cknmidebghlajpfo",
		},
	}

	for i, test := range tests {
		s := Dance(test.Input, test.Key)
		if s != test.Result {
			t.Fatalf("a.%d: %s '%s' (should be %s)", i, test.Input, s, test.Result)
		}
	}
}

func TestLotsOfDances(t *testing.T) {
	tests := []struct {
		Input  string
		Key    string
		Result string
	}{
		{
			Input:  "@a.txt",
			Key:    "abcdefghijklmnop",
			Result: "cbolhmkgfpenidaj",
		},
	}

	for i, test := range tests {
		s := LotsOfDances(test.Input, test.Key)
		if s != test.Result {
			t.Fatalf("b.%d: %s '%s' (should be %s)", i, test.Input, s, test.Result)
		}
	}
}
