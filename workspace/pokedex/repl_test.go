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
			input:    "hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    bwa bwu wbi",
			expected: []string{"bwa", "bwu", "wbi"},
		},
		{
			input:    "          ",
			expected: []string{},
		},
		{
			input:    "    bwa \n\nbwu\t\n wbi   ",
			expected: []string{"bwa", "bwu", "wbi"},
		},
		{
			input:    "    BWA \n\nbWu\t\n Wbi   ",
			expected: []string{"bwa", "bwu", "wbi"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected length not reached")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words do not match")
			}
		}
	}
}
