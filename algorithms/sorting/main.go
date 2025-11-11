package main

import (
	"math"
)

type Pair struct {
	Key   int
	Value string
}

func insertionSort(pairs []Pair) [][]Pair {
	res := [][]Pair{}
	// normally we start i = 1, but because we need to record the initial state of the array we must do i = 0
	// that way the initial state gets recorded.
	for i := 0; i < len(pairs); i++ {
		j := i - 1
		for j >= 0 && pairs[j+1].Key < pairs[j].Key {
			pairs[j], pairs[j+1] = pairs[j+1], pairs[j]
			j--
		}
		// we have to clone the state of the slices after the swap as it's own independent slice.
		// go to remember that slices are always a reference to the underlying array
		clone := append([]Pair(nil), pairs...)
		copy(clone, pairs)
		res = append(res, clone)
	}

	return res
}

func mergeSort(pairs []Pair) []Pair {
	return mergeSortHelper(pairs, 0, len(pairs)-1)
}

func mergeSortHelper(pairs []Pair, start, end int) []Pair {
	if end-start+1 <= 1 {
		return pairs
	}

	mid := (start + end) / 2
	mergeSortHelper(pairs, start, mid)
	mergeSortHelper(pairs, mid+1, end)
	merge(pairs, start, mid, end)

	return pairs
}

func merge(pairs []Pair, start, mid, end int) {
	// create the left and right arrays to length needed for them
	leftArr := make([]Pair, mid-start+1)
	copy(leftArr, pairs[start:mid+1])
	rightArr := make([]Pair, end-mid)
	copy(rightArr, pairs[mid+1:end+1])

	// starting points for left and right array and the orginal array
	i := 0
	j := 0
	k := start

	// merge the two sorted halfs in the og array
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i].Key <= rightArr[j].Key {
			pairs[k] = leftArr[i]
			i++
		} else {
			pairs[k] = rightArr[j]
			j++
		}
		k++
	}

	// check if one of the halfs have any remaining elements to place into the og array
	for i < len(leftArr) {
		pairs[k] = leftArr[i]
		i++
		k++
	}

	for j < len(rightArr) {
		pairs[k] = rightArr[j]
		j++
		k++
	}
}

func quickSort(pairs []Pair) []Pair {
	quickSortHelper(pairs, 0, len(pairs)-1)
	return pairs
}

func quickSortHelper(pairs []Pair, start, end int) {
	if end-start+1 <= 1 {
		return
	}

	pivot := pairs[end] // pivot is the last element always
	leftPtr := start    // pointer for the left part

	// partion the elements smaller than pivot on left side
	for i := start; i < end; i++ {
		if pairs[i].Key < pivot.Key {
			pairs[i], pairs[leftPtr] = pairs[leftPtr], pairs[i]
			leftPtr++
		}
	}
	// move pivot in-between left and right sides
	pairs[end] = pairs[leftPtr]
	pairs[leftPtr] = pivot

	// quicksort the left and right array
	quickSortHelper(pairs, start, leftPtr-1)
	quickSortHelper(pairs, leftPtr+1, end)
}

// only really usable if you know the range of your bucket
// for example if arr=[2,2,1,0,0,2] then bucket size will be 3 where bucketarr=[0,0,0] where idx in bucketarr
// represents arr[i]
func bucketSort(arr []int) []int {
	counts := make([]int, 3)
	for _, n := range arr {
		counts[n]++
	}

	i := 0
	for n, count := range counts {
		for range count {
			arr[i] = n
			i++
		}
	}
	return arr
}

func sortColors(nums []int) {
	counts := make([]int, 3)
	for _, n := range nums {
		counts[n]++
	}

	i := 0
	for n, count := range counts {
		for range count {
			nums[i] = n
			i++
		}
	}
}
