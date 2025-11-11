package main

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
