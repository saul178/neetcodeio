package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		n        int
	}{
		{
			name:     "2 steps",
			expected: 2,
			n:        2,
		},
		{
			name:     "3 steps",
			expected: 3,
			n:        3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbingStairs(tt.n)
			assert.Equal(t, tt.expected, got)
		})
	}
}
