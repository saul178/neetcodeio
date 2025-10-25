package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchRec(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "search the right",
			nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   7,
			expected: 7,
		},
		{
			name:     "search the left",
			nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   2,
			expected: 2,
		},
		{
			name:     "negative to positive",
			nums:     []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   -3,
			expected: 2,
		},
		{
			name:     "target not found",
			nums:     []int{1, 4, 5, 6, 8},
			target:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRec(tt.nums, tt.target)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestBinarySearchIter(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "search the right",
			nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   7,
			expected: 7,
		},
		{
			name:     "search the left",
			nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   2,
			expected: 2,
		},
		{
			name:     "negative to positive",
			nums:     []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:   -3,
			expected: 2,
		},
		{
			name:     "target not found",
			nums:     []int{1, 4, 5, 6, 8},
			target:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchIter(tt.nums, tt.target)
			assert.Equal(t, tt.expected, got)
		})
	}
}
