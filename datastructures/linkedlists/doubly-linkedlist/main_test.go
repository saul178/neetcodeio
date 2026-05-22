package main

import "testing"

func TestDoubleLinkedList(t *testing.T) {
	testCases := []struct {
		name        string
		listValues  []int
		searchIdx   int
		expectedVal int
	}{
		{
			name:        "test Get(index)",
			listValues:  []int{2, 3, 5, 6, 1},
			searchIdx:   3,
			expectedVal: 6,
		},
		{
			name:        "test out of bounds",
			listValues:  []int{0, 3},
			searchIdx:   6,
			expectedVal: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ll := Constructor()

			for _, val := range tc.listValues {
				newNode := NewDoublyLinkedNode(val, ll.Tail.Prev, ll.Tail)
				ll.Tail.Prev.Next = newNode
				ll.Tail.Prev = newNode
			}

			actual := ll.Get(tc.searchIdx)
			if actual != tc.expectedVal {
				t.Errorf(
					"Get(%d) failed for case '%s': expected %d got %d",
					tc.searchIdx,
					tc.name,
					tc.expectedVal,
					actual,
				)
			}
		})
	}
}
