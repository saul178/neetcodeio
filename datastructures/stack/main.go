package main

import (
	"errors"
	"fmt"
	"strings"
)

func removeDuplicates(nums []int) int {
	stack := []int{}

	for i := range nums {
		if len(stack) == 0 || stack[len(stack)-1] != nums[i] {
			stack = append(stack, nums[i])
		}
	}

	return len(stack)
}

func isValidParentheses(s string) bool {
	if len(s) <= 1 {
		return false
	}

	stack := []byte{}
	for i := range s {
		switch s[i] {
		case '(':
			stack = append(stack, s[i])
		case '{':
			stack = append(stack, s[i])
		case '[':
			stack = append(stack, s[i])
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

// answer[i] is how long you have to wait for a day that is warmer
// so i == a day at temp[i] == some temperature;
// count how many days it takes to find a warmer temp where temp[i] > temp[i]++
// we stop as soon as we find a warmer day and the count is added to the array
func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 1 {
		return []int{0}
	}

	answer := make([]int, len(temperatures))
	stack := []int{}

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[index] = i - index
		}
		stack = append(stack, i)
	}

	return answer
}

func simplifyPath(path string) string {
	stack := []string{}
	chunks := strings.Split(path, "/")
	fmt.Println(chunks)

	for _, c := range chunks {
		switch c {
		case ".":
			continue
		case "":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, c)
		}
	}

	return "/" + strings.Join(stack, "/")
}

var ErrorEmptyStack = errors.New("Stack is empty.")

type MinStack struct {
	elements []int
}

func Constructor() *MinStack {
	stack := []int{}
	return &MinStack{elements: stack}
}

func (ms *MinStack) Push(val int) {
	ms.elements = append(ms.elements, val)
}

func (ms *MinStack) Pop() int {
	if len(ms.elements) > 0 {
		val := ms.elements[len(ms.elements)-1]
		ms.elements = ms.elements[:len(ms.elements)-1]
		return val
	} else {
		fmt.Println(ErrorEmptyStack)
		return -1
	}
}

func (ms *MinStack) Top() int {
	if len(ms.elements) > 0 {
		return ms.elements[len(ms.elements)-1]
	} else {
		fmt.Println(ErrorEmptyStack)
		return -1
	}
}

func (ms *MinStack) GetMin() int {
	tmpStack := Constructor()
	minimum := ms.Top()

	for len(ms.elements) > 0 {
		val := ms.Pop()
		minimum = min(minimum, val)
		tmpStack.Push(val)
	}

	for len(tmpStack.elements) > 0 {
		val := tmpStack.Pop()
		ms.Push(val)
	}
	return minimum
}

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			copy(nums[i:], nums[i+1:])
			nums[len(nums)-1] = 0
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
	}
	return len(nums)
}

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	copy(ans, nums)
	ans = append(ans, nums...)
	return ans
}

func main() {
	fmt.Println(getConcatenation([]int{1, 4, 1, 2}))
}

/*
https://leetcode.com/problems/valid-parentheses
https://leetcode.com/problems/simplify-path
https://leetcode.com/problems/daily-temperatures
*/
