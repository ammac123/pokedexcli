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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "what are we doing here???   ",
			expected: []string{"what", "are", "we", "doing", "here???"},
		}, {
			input:    "    ",
			expected: []string{},
		}, {
			input: `Moving    with extra NEWliNe
			 space`,
			expected: []string{"moving", "with", "extra", "newline", "space"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Got length %v, expected %v", len(actual), len(c.expected))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(`cleanInput(%v) == %v, expected %v`, c.input, actual, c.expected)
				return
			}
		}
	}
}
