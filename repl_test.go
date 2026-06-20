package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "CharMander eS mi  Fav",
			expected: []string{"charmander", "es", "mi", "fav"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected slice of len: %v and got: %v", len(c.expected), len(actual) )
		}
		for i := range actual {
			word := actual[i]
			if word != c.expected[i] {
				t.Errorf("Expected word: %v and got: %v", c.expected[i], word)
			}
		}
	}
}