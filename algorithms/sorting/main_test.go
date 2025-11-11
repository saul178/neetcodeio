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
			name: "insertion sort test",
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

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		data     []Pair
		expected []Pair
	}{
		{
			name: "merge sort stable test",
			data: []Pair{
				{Key: 5, Value: "apple"},
				{Key: 2, Value: "banana"},
				{Key: 9, Value: "cherry"},
				{Key: 1, Value: "date"},
				{Key: 9, Value: "elderberry"},
			},
			expected: []Pair{
				{Key: 1, Value: "date"},
				{Key: 2, Value: "banana"},
				{Key: 5, Value: "apple"},
				{Key: 9, Value: "cherry"},
				{Key: 9, Value: "elderberry"},
			},
		},
		{
			name: "merge sort stable test 2",
			data: []Pair{
				{Key: 3, Value: "cat"},
				{Key: 2, Value: "dog"},
				{Key: 3, Value: "bird"},
			},
			expected: []Pair{
				{Key: 2, Value: "dog"},
				{Key: 3, Value: "cat"},
				{Key: 3, Value: "bird"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mergeSort(tt.data)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		data     []Pair
		expected []Pair
	}{
		{
			name: "quick sort non stable test",
			data: []Pair{
				{Key: 5, Value: "apple"},
				{Key: 2, Value: "banana"},
				{Key: 9, Value: "cherry"},
				{Key: 1, Value: "date"},
				{Key: 9, Value: "elderberry"},
			},
			expected: []Pair{
				{Key: 1, Value: "date"},
				{Key: 2, Value: "banana"},
				{Key: 5, Value: "apple"},
				{Key: 9, Value: "elderberry"},
				{Key: 9, Value: "cherry"},
			},
		},
		{
			name: "quick sort non stable",
			data: []Pair{
				{Key: 3, Value: "cat"},
				{Key: 2, Value: "dog"},
				{Key: 3, Value: "bird"},
			},
			expected: []Pair{
				{Key: 2, Value: "dog"},
				{Key: 3, Value: "bird"},
				{Key: 3, Value: "cat"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quickSort(tt.data)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestBucketSort(t *testing.T) {
	test := []struct {
		name     string
		data     []int
		expected []int
	}{
		{
			name:     "simple test",
			data:     []int{2, 1, 2, 0, 0, 2},
			expected: []int{0, 0, 1, 2, 2, 2},
		},
	}

	for _, tt := range test {
		got := bucketSort(tt.data)
		assert.Equal(t, tt.expected, got)
	}
}
