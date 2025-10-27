package main

import "math"

func binarySearchRec(nums []int, left int, right int, target int) int {
	if left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			return binarySearchRec(nums, mid+1, right, target)
		} else {
			return binarySearchRec(nums, left, mid-1, target)
		}
	} else {
		return -1
	}
}

func binarSearchIter(nums []int, left int, right int, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchRec(nums []int, target int) int {
	return binarySearchRec(nums, 0, len(nums)-1, target)
}

func searchIter(nums []int, target int) int {
	return binarSearchIter(nums, 0, len(nums)-1, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	totalRows := len(matrix)
	totalCols := len(matrix[0])

	topRow := 0
	bottomRow := totalRows - 1

	// do a binary search to find the row where the target might be
	for topRow <= bottomRow {
		midRow := (topRow + bottomRow) / 2

		if target > matrix[midRow][totalCols-1] {
			topRow = midRow + 1
		} else if target < matrix[midRow][0] {
			bottomRow = midRow - 1
		} else {
			// break out of the loop if we find a row where the target might live
			break
		}
	}

	if topRow > bottomRow {
		return false
	}

	// binary search within the found row that might have the target
	searchRow := (topRow + bottomRow) / 2
	left := 0
	right := totalCols - 1
	for left <= right {
		midCol := (left + right) / 2
		if target > matrix[searchRow][midCol] {
			left = midCol + 1
		} else if target < matrix[searchRow][midCol] {
			right = midCol - 1
		} else {
			return true
		}
	}

	return false
}

func guessNumber(n int) int {
	low := 1
	high := n

	for {
		mid := low + (high-low)/2
		if guess(mid) > 0 {
			high = mid - 1
		} else if guess(mid) < 0 {
			low = mid + 1
		} else {
			return mid
		}
	}
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	res := -1

	for low <= high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			res = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}
