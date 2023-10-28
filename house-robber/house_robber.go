package main

func main() {
	rob([]int{1, 34, 42, 4, 9})
}

func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// Create a slice to store the maximum amount of money robbed up to the current house
	dp := make([]int, n)

	// Initialize the first two elements in the dp slice
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	// Fill in the dp slice using dynamic programming
	for i := 2; i < n; i++ {
		// The maximum amount of money at the current house is either:
		// 1. The maximum amount from the previous house
		// 2. The maximum amount from two houses back + money from the current house
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	// The last element in the dp slice will contain the maximum amount of money robbed
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
