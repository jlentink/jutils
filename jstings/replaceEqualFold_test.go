package jstings

import (
	"testing"
)

func TestReplaceEqualFold(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		old      string
		new      string
		n        int
		expected string
	}{
		{
			name:     "Test 1",
			s:        "Hello, World!",
			old:      "world",
			new:      "Universe",
			n:        -1,
			expected: "Hello, Universe!",
		},
		{
			name:     "Test 2",
			s:        "Hello, World!",
			old:      "world",
			new:      "Universe",
			n:        0,
			expected: "Hello, World!",
		},
		{
			name:     "Test 3",
			s:        "Hello, World! World!",
			old:      "world",
			new:      "Universe",
			n:        1,
			expected: "Hello, Universe! World!",
		},
		{
			name:     "Test 4",
			s:        "Hello, World! world",
			old:      "world",
			new:      "Universe",
			n:        -1,
			expected: "Hello, Universe! Universe",
		},
		{
			name:     "Test 5",
			s:        "Hello, World! world",
			old:      "Universe1",
			new:      "Universe2",
			n:        -1,
			expected: "Hello, World! world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceEqualFold(tt.s, tt.old, tt.new, tt.n); got != tt.expected {
				t.Errorf("ReplaceEqualFold() = %v, want %v", got, tt.expected)
			}
		})
	}
}
