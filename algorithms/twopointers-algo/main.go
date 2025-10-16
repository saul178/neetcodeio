package main

import (
	"strings"
	"unicode"
)

func sanitizeStr(s string) string {
	s = strings.ToLower(s)
	filteredStr := []rune{}
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			filteredStr = append(filteredStr, c)
		}
	}
	return string(filteredStr)
}

func isPalindrome(s string) bool {
	leftPtr := 0
	rightPtr := len(s) - 1
	newStr := sanitizeStr(s)

	for leftPtr < rightPtr {
		if newStr[leftPtr] != newStr[rightPtr] {
			return false
		}
		leftPtr++
		rightPtr--
	}
	return true
}

/*
https://leetcode.com/problems/3sum
https://leetcode.com/problems/candy-crush
https://leetcode.com/problems/rotate-array
https://leetcode.com/problems/container-with-most-water
*/
