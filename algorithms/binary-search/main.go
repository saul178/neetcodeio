package main

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
	rows := len(matrix)
	cols := len(matrix[0])

	top := 0
	bottom := rows - 1

	for top <= bottom {
		row := (top + bottom) / 2
		if target > matrix[row][row-1] {
			top = row + 1
		} else if target < matrix[row][0] {
			bottom = row - 1
		} else {
			break
		}
	}

	if !(top <= bottom) {
		return true
	}

	row := (top + bottom) / 2
	l, r := 0, cols-1
	for l <= r {
		m := (l + r) / 2
		if target > matrix[row][m] {
			l = m + 1
		} else if target < matrix[row][m] {
			r = m - 1
		} else {
			return true
		}
	}
	return false
}
