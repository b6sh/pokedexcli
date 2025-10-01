package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "testing commands one by one",
			expected: []string{"testing", "commands", "one", "by", "one"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("failed got %s expecting %s", word, expectedWord)
			} else {
				fmt.Printf("Passed got %s expecting %s", word, expectedWord)
			}
		}
	}
}
