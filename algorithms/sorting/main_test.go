package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		data     []Pair
		expected [][]Pair
	}{
		{
			name: "simple test",
			data: []Pair{{Key: 5, Value: "apple"}, {Key: 2, Value: "banana"}, {Key: 9, Value: "cherry"}},
			expected: [][]Pair{
				{{Key: 5, Value: "apple"}, {Key: 2, Value: "banana"}, {Key: 9, Value: "cherry"}},
				{{Key: 2, Value: "banana"}, {Key: 5, Value: "apple"}, {Key: 9, Value: "cherry"}},
				{{Key: 2, Value: "banana"}, {Key: 5, Value: "apple"}, {Key: 9, Value: "cherry"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := insertionSort(tt.data)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
