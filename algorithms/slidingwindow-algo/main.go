package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(nums []int, target int) int {
	// minimal length of a subarray where the sum is greater than or equal to the target
	minSubArrLen := math.MaxInt
	currSum := 0
	leftPtr := 0

	for rightPtr := 0; rightPtr < len(nums); rightPtr++ {
		currSum += nums[rightPtr]
		for currSum >= target {
			if currSum >= target {
				minSubArrLen = min(minSubArrLen, rightPtr-leftPtr+1)
			}
			currSum -= nums[leftPtr]
			leftPtr++
		}
	}

	if minSubArrLen == math.MaxInt {
		minSubArrLen = 0
	}
	return minSubArrLen
}

func lengthOfLongestSubstring(s string) int {
	// have a set to show the unique chars shown in the string
	if s == "" {
		return 0
	}

	charSet := make(map[byte]bool)
	substringLen := math.MinInt
	leftPtr := 0

	for righPtr := 0; righPtr < len(s); righPtr++ {
		for charSet[s[righPtr]] {
			delete(charSet, s[leftPtr])
			leftPtr++
		}
		charSet[s[righPtr]] = true
		substringLen = max(substringLen, righPtr-leftPtr+1)
	}
	return substringLen
}

func countGoodSubstrings(s string) int {
	charSet := make(map[byte]bool)
	windowSize := 3
	leftPtr := 0
	validCount := 0

	for rightPtr := 0; rightPtr < len(s); rightPtr++ {
		for charSet[s[rightPtr]] {
			delete(charSet, s[leftPtr])
			leftPtr++
		}

		if rightPtr-leftPtr+1 > windowSize {
			delete(charSet, s[leftPtr])
			leftPtr++
		}

		charSet[s[rightPtr]] = true
		if len(charSet) == windowSize {
			validCount++
		}
	}

	return validCount
}

func longestPalindrome(s string) string {
	res := ""
	resLen := 0

	for i := range s {
		// odd length
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right - left + 1) > resLen {
				res = s[left : right+1]
				resLen = right - left + 1
			}
			left--
			right++
		}
		// even length
		l, r := i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > resLen {
				res = s[l : r+1]
				resLen = r - l + 1
			}
			l--
			r++
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("zxyzxyz"))
}
