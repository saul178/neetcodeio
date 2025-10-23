package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountStudents(t *testing.T) {
	tests := []struct {
		name       string
		students   []int
		sandwiches []int
		expected   int
	}{
		{
			name:       "all students eat",
			students:   []int{1, 1, 0, 0},
			sandwiches: []int{0, 1, 0, 1},
			expected:   0,
		},
		{
			name:       "some students eat",
			students:   []int{1, 1, 1, 0, 0, 1},
			sandwiches: []int{1, 0, 0, 0, 1, 1},
			expected:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countStudents(tt.students, tt.sandwiches)
			assert.Equal(t, tt.expected, result)
		})
	}
}
