package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
		  input: "   hello world   ",
		  expected: []string{"hello", "world"},
		},
		{
		  input: "foo bar",
		  expected: []string{"foo", "bar"},
		},
		{
		  input: "123 456   ",
		  expected: []string{"123", "456"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: expected '%v' (%d), got '%v' (%d)", len(actual), actual, len(c.expected), c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%d's word doesn't match: expected %s, got %s", i, expectedWord, word)
			}
		}
	}
}
