package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head

	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeTwoList(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "empty case",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "list1 is empty case",
			list1:    []int{1, 2, 3, 4},
			list2:    []int{},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "list2 is empty case",
			list1:    []int{},
			list2:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "overlapping values",
			list1:    []int{1, 1, 2, 4, 5},
			list2:    []int{1, 2, 3, 3, 4},
			expected: []int{1, 1, 1, 2, 2, 3, 3, 4, 4, 5},
		},
		{
			name:     "list w/ different lengths",
			list1:    []int{1, 1, 2, 4, 5},
			list2:    []int{1, 2},
			expected: []int{1, 1, 1, 2, 2, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := sliceToList(tt.list1)
			l2 := sliceToList(tt.list2)

			result := mergeTwoLists(l1, l2)
			got := listToSlice(result)

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected []int
	}{
		{
			name:     "list with ordered values",
			list:     []int{0, 1, 2, 3},
			expected: []int{3, 2, 1, 0},
		},
		{
			name:     "empty list",
			list:     []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testList := sliceToList(tt.list)

			result := reverseList(testList)
			got := listToSlice(result)

			assert.Equal(t, tt.expected, got)
		})
	}
}
