package main

// recursive
func climbingStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbingStairs(n-1) + climbingStairs(n-2)
}

// dynamic programming way top-down
var cache map[int]int = make(map[int]int)

func climbingStairsTopDown(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if val, exists := cache[n]; exists {
		return val
	}
	res := climbingStairsTopDown(n-1) + climbingStairsTopDown(n-2)
	cache[n] = res
	return res
}

// dynamic programming way bottom up
func climbStairsBottomUp(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
