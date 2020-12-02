package p9

import (
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "{}",
			Result: 1,
		},
		{
			Input:  "{{{}}}",
			Result: 6,
		},
		{
			Input:  "{{},{}}",
			Result: 5,
		},
		{
			Input:  "{{{},{},{{}}}}",
			Result: 16,
		},
		{
			Input:  "{<a>,<a>,<a>,<a>}",
			Result: 1,
		},
		{
			Input:  "{{<ab>},{<ab>},{<ab>},{<ab>}}",
			Result: 9,
		},
		{
			Input:  "{{<!!>},{<!!>},{<!!>},{<!!>}}",
			Result: 9,
		},
		{
			Input:  "{{<a!>},{<a!>},{<a!>},{<ab>}}",
			Result: 3,
		},
		{
			Input:  "{{{<\"\"!>,<!<'}'ui!!!>!!!>!<{!!!!!>>},{}}}",
			Result: 9,
		},
		{
			Input:  "@a.txt",
			Result: 14204,
		},
	}

	for i, test := range tests {
		n := Score(test.Input)
		if n != test.Result {
			t.Fatalf("a.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}

func TestGarbage(t *testing.T) {
	tests := []struct {
		Input  string
		Result int
	}{
		{
			Input:  "{}",
			Result: 0,
		},
		{
			Input:  "<random characters>",
			Result: 17,
		},
		{
			Input:  "<<<<>",
			Result: 3,
		},
		{
			Input:  "<{!>}>",
			Result: 2,
		},
		{
			Input:  "<!!>",
			Result: 0,
		},
		{
			Input:  "<!!!>>",
			Result: 0,
		},
		{
			Input:  "<{o\"i!a,<{i<a>",
			Result: 10,
		},
		{
			Input:  "@a.txt",
			Result: 6622,
		},
	}

	for i, test := range tests {
		n := Garbage(test.Input)
		if n != test.Result {
			t.Fatalf("b.%d: %s '%d' (should be %d)", i, test.Input, n, test.Result)
		}
	}
}
