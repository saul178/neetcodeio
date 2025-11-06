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
