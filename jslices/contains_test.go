package jslices

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		v        int
		expected bool
	}{
		{
			name:     "Test 1",
			list:     []int{1, 2, 3, 4, 5},
			v:        3,
			expected: true,
		},
		{
			name:     "Test 2",
			list:     []int{1, 2, 3, 4, 5},
			v:        6,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.list, tt.v); got != tt.expected {
				t.Errorf("Contains() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestContainsEqualFold(t *testing.T) {
	tests := []struct {
		name     string
		list     []string
		v        string
		expected bool
	}{
		{
			name:     "Test 1",
			list:     []string{"Hello", "World"},
			v:        "hello",
			expected: true,
		},
		{
			name:     "Test 2",
			list:     []string{"Hello", "World"},
			v:        "Universe",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsEqualFold(tt.list, tt.v); got != tt.expected {
				t.Errorf("ContainsEqualFold() = %v, want %v", got, tt.expected)
			}
		})
	}
}
