package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// create a slice of test case structs
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "     Take    ME    to    yOUr    leader     ",
			expected: []string{"take", "me", "to", "your", "leader"},
		},
		{
			input:    "exit",
			expected: []string{"exit"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check the length of the actual slice against the expected slice
		// if they dont match, use t.Errorf to print an error message and fail

		if len(actual) != len(c.expected) {
			t.Errorf("FAIL -- cleanInput returned %d words, expected %d words", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				fmt.Println("Error")
				t.Errorf("FAIL -- cleanInput did not match expected test case.")
			}
		}
	}
}
