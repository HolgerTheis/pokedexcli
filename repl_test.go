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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " PIKACHU  eevEE  charmander   ",
			expected: []string{"pikachu", "eevee", "charmander"},
		},
		{
			input:    "  BUlbasaur  sQUirtle  ",
			expected: []string{"bulbasaur", "squirtle"},
		},
		{
			input:    "      ",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("FAIL: Difference in output length detected!\nExpected: %v\nActual:   %v", c.expected, actual)
			t.Fail()
		}
		for i := range actual {
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("FAIL: Unexpected word at index %d\nExpected: %s\nActual:   %s", i, expectedWord, word)
				t.Fail()
			}
		}
	}
}
