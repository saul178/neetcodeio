package main

import (
	"fmt"
	"slices"
)

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxFound := 0
	for i := range nums {
		if nums[i] != 0 && nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		maxFound = max(count, maxFound)
	}
	return maxFound
}

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = slices.Delete(nums, i, i+1)
		} else {
			i++
		}
	}
	return len(nums)
}

func replaceElements(arr []int) []int {
	rightMax := -1
	for i := len(arr) - 1; i >= 0; i-- {
		newMax := max(rightMax, arr[i])
		arr[i] = rightMax
		rightMax = newMax
	}
	return arr
}

func main() {
	fmt.Println(replaceElements([]int{1, 0, 1, 1, 0, 1}))
}
