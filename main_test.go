package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		expectedLen int
	}{
		{
			name:        "negative size",
			size:        -10,
			expectedLen: 0,
		},
		{
			name:        "zero size",
			size:        0,
			expectedLen: 0,
		},
		{
			name:        "single element",
			size:        1,
			expectedLen: 1,
		},
		{
			name:        "multiple elements",
			size:        100,
			expectedLen: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.expectedLen {
				t.Fatalf("expected slice length %d, got %d", tt.expectedLen, len(got))
			}

			for i, v := range got {
				if v < 0 {
					t.Errorf("expected non-negative element at index %d, got %d", i, v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "nil slice",
			data:     nil,
			expected: 0,
		},
		{
			name:     "empty slice",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "single element",
			data:     []int{42},
			expected: 42,
		},
		{
			name:     "max in the middle",
			data:     []int{3, 15, 99, 24, 8},
			expected: 99,
		},
		{
			name:     "max at the beginning",
			data:     []int{150, 40, 12, 7},
			expected: 150,
		},
		{
			name:     "max at the end",
			data:     []int{5, 18, 33, 210},
			expected: 210,
		},
		{
			name:     "identical elements",
			data:     []int{7, 7, 7, 7},
			expected: 7,
		},
		{
			name:     "negative numbers",
			data:     []int{-50, -12, -3, -89},
			expected: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "empty slice",
			data:     []int{},
			expected: 0,
		},
		{
			name:     "smaller than CHUNKS",
			data:     []int{4, 19, 2},
			expected: 19,
		},
		{
			name:     "exact multiple of CHUNKS",
			data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			expected: 16,
		},
		{
			name:     "not divisible by CHUNKS (remainder in last chunk)",
			data:     []int{10, 2, 8, 4, 1, 9, 3, 7, 5, 12, 100},
			expected: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
