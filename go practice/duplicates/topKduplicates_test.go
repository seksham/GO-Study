package main

import (
	"reflect"
	"testing"
)

func TestCheckDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected map[int]int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: map[int]int{},
		},
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: map[int]int{1: 1, 2: 1, 3: 1, 4: 1, 5: 1},
		},
		{
			name:     "With duplicates",
			input:    []int{1, 1, 2, 3, 4, 2, 5, 5},
			expected: map[int]int{1: 2, 2: 2, 3: 1, 4: 1, 5: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkDuplicates(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("checkDuplicates() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTopKduplicates(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		k           int
		expected    []kv
		expectError bool
	}{
		{
			name:        "K greater than slice length",
			input:       []int{1, 2, 3},
			k:           5,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "K is zero",
			input:       []int{1, 1, 2, 2, 3},
			k:           0,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "No duplicates",
			input:       []int{1, 2, 3, 4, 5},
			k:           2,
			expected:    []kv{},
			expectError: false,
		},
		{
			name:        "Normal case",
			input:       []int{1, 1, 2, 3, 4, 2, 5, 5, 5},
			k:           2,
			expected:    []kv{{5, 3}, {1, 2}},
			expectError: false,
		},
		{
			name:        "K greater than number of duplicates",
			input:       []int{1, 1, 2, 2, 3, 3, 4},
			k:           5,
			expected:    []kv{{1, 2}, {2, 2}, {3, 2}},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := topKduplicates(tt.input, tt.k)

			if tt.expectError && err == nil {
				t.Errorf("topKduplicates() expected an error, but got none")
			}

			if !tt.expectError && err != nil {
				t.Errorf("topKduplicates() unexpected error: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("topKduplicates() = %v, want %v", result, tt.expected)
			}
		})
	}
}