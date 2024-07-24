package main

import (
	"testing"
)

func TestClear(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},

		{
			input: "go is awesome",
			expected: []string{
				"go",
				"is",
				"awesome",
			},
		},
	}
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Lenght Mismatch %v vs %v", len(actual), len(cs.expected))
			continue
		}
		for i := range cs.expected {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("Word Mismatch %v vs %v", actualWord, expectedWord)
			}
		}
	}
}
